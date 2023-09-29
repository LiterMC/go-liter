
package liter

import (
	"errors"
	"io"
	"sync"
)

var (
	ErrParserIdNotExists = errors.New("liter.CommandNode: parser id is not exists")
	ErrParserInstanceNotExists = errors.New("liter.CommandNode: command parser instance not exists")
)

type CmdPropEncoder interface {
	Id()(VarInt)
	String()(String)
	Encode(b *PacketBuilder, value any)(err error)
	Decode(r *PacketReader)(value any, err error)
}

var (
	cmdMux sync.RWMutex
	cmdIDMap = make(map[VarInt]CmdPropEncoder, 50)
	cmdSIDMap = make(map[string]CmdPropEncoder, 50)
)

func RegisterCmdEncoder(p CmdPropEncoder)(ok bool){
	id := p.Id()
	sid := p.String()

	cmdMux.Lock()
	defer cmdMux.Unlock()
	if _, ok := cmdIDMap[id]; ok {
		return false
	}
	if _, ok := cmdSIDMap[sid]; ok {
		return false
	}
	cmdIDMap[id] = p
	cmdSIDMap[sid] = p
	return true
}

func UnregisterCmdEncoder(p CmdPropEncoder)(ok bool){
	id := p.Id()
	sid := p.String()

	cmdMux.Lock()
	defer cmdMux.Unlock()
	old1, ok1 := cmdIDMap[id]
	if !ok1 {
		return false
	}
	old2, ok2 := cmdSIDMap[sid]
	if !ok2 || p != old1 || p != old2 { // return false if instances not match
		return false
	}
	delete(cmdIDMap, id)
	delete(cmdSIDMap, sid)
	return true
}


type CmdFlag Byte
const (
	// Values
	CmdTypeRoot    CmdFlag = 0x00
	CmdTypeLiteral CmdFlag = 0x01
	CmdTypeArg     CmdFlag = 0x02
	// Masks
	CmdTypeMask CmdFlag = 0x03 // 0: root, 1: literal, 2: argument. 3 is not used.
	CmdExec     CmdFlag = 0x04 // Set if the node stack to this point constitutes a valid command.
	CmdRedirect CmdFlag = 0x08 // Set if the node redirects to another node.
	CmdSuggest  CmdFlag = 0x10 // Only present for argument nodes.
)

type CommandNode struct {
	Flags           CmdFlag
	Children        []VarInt
	RedirectNode    Optional[VarInt]
	Name            Optional[String]
	ParserId        Optional[VarInt] // 1.19+
	ParserStrId     Optional[String] // 1.7 ~ 1.18
	Properties      any
	SuggestionsType Optional[String]
}

var _ Packet = (*CommandNode)(nil)

func (c *CommandNode)Type()(CmdFlag){
	return c.Flags & CmdTypeMask
}

func (c *CommandNode)HasFlag(f CmdFlag)(bool){
	return c.Flags & f != 0
}

func (c *CommandNode)Vaild()(bool){
	return c.HasFlag(CmdExec)
}

func (c *CommandNode)GetParserId()(VarInt, error){
	if !c.ParserId.Ok {
		if !c.ParserStrId.Ok {
			return -1, ErrParserIdNotExists
		}
		cmdMux.RLock()
		p, ok := cmdSIDMap[c.ParserStrId.V]
		cmdMux.RUnlock()
		if !ok {
			return -1, ErrParserInstanceNotExists
		}
		c.ParserId.Set(p.Id())
	}
	return c.ParserId.V, nil
}

func (c *CommandNode)GetParserStrId()(String, error){
	if !c.ParserStrId.Ok {
		if !c.ParserId.Ok {
			return "", ErrParserIdNotExists
		}
		cmdMux.RLock()
		p, ok := cmdIDMap[c.ParserId.V]
		cmdMux.RUnlock()
		if !ok {
			return "", ErrParserInstanceNotExists
		}
		c.ParserStrId.Set(p.String())
	}
	return c.ParserStrId.V, nil
}

func (c *CommandNode)GetParser()(CmdPropEncoder, error){
	if c.ParserId.Ok {
		cmdMux.RLock()
		p, ok := cmdIDMap[c.ParserId.V]
		cmdMux.RUnlock()
		if !ok {
			return nil, ErrParserInstanceNotExists
		}
		return p, nil
	}
	if c.ParserStrId.Ok {
		cmdMux.RLock()
		p, ok := cmdSIDMap[c.ParserStrId.V]
		cmdMux.RUnlock()
		if !ok {
			return nil, ErrParserInstanceNotExists
		}
		return p, nil
	}
	return nil, ErrParserIdNotExists
}

func (c *CommandNode)Encode(b *PacketBuilder){
	b.VarInt((VarInt)(c.Flags))
	b.VarInt((VarInt)(len(c.Children)))
	for _, child := range c.Children {
		b.VarInt(child)
	}
	if c.HasFlag(CmdRedirect) {
		b.VarInt(c.RedirectNode.Assert())
	}
	t := c.Type()
	if t == CmdTypeLiteral || t == CmdTypeArg {
		b.String(c.Name.Assert())
	}
	if t == CmdTypeArg {
		p, err := c.GetParser()
		if err != nil {
			panic(err)
		}
		b.VarInt(p.Id())
		if err = p.Encode(b, c.Properties); err != nil {
			panic(err)
		}
	}
	if c.HasFlag(CmdSuggest) {
		b.String(c.SuggestionsType.Assert())
	}
}

func (c *CommandNode)DecodeFrom(r *PacketReader)(err error){
	protocol := r.Protocol()
	var ok bool
	var flags VarInt
	if flags, ok = r.VarInt(); !ok {
		return io.EOF
	}
	c.Flags = (CmdFlag)(flags)
	var l VarInt
	if l, ok = r.VarInt(); !ok {
		return io.EOF
	}
	c.Children = make([]VarInt, l)
	for i, _ := range c.Children {
		if c.Children[i], ok = r.VarInt(); !ok {
			return io.EOF
		}
	}
	if c.HasFlag(CmdRedirect) {
		var v VarInt
		if v, ok = r.VarInt(); !ok {
			return io.EOF
		}
		c.RedirectNode.Set(v)
	}
	t := c.Type()
	if t == CmdTypeLiteral || t == CmdTypeArg {
		var v String
		if v, ok = r.String(); !ok {
			return io.EOF
		}
		c.Name.Set(v)
	}
	if t == CmdTypeArg {
		var p CmdPropEncoder
		if protocol < V1_19 {
			var v String
			if v, ok = r.String(); !ok {
				return io.EOF
			}
			c.ParserStrId.Set(v)
			c.ParserId.Ok = false
			if p, err = c.GetParser(); err != nil {
				return
			}
		}else{
			var v VarInt
			if v, ok = r.VarInt(); !ok {
				return io.EOF
			}
			c.ParserId.Set(v)
			c.ParserStrId.Ok = false
			if p, err = c.GetParser(); err != nil {
				return
			}
		}
		if c.Properties, err = p.Decode(r); err != nil {
			return
		}
	}
	if c.HasFlag(CmdSuggest) {
		var v String
		if v, ok = r.String(); !ok {
			return io.EOF
		}
		c.SuggestionsType.Set(v)
	}
	return
}
