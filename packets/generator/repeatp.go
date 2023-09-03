
package main

import (
	"sort"
	"sync"
)

type repeatPacketStat struct {
	Protocol int
	State string
	Bound string
	Id int
}

type repeatPacketInfo struct {
	I int
	Matched bool
	Stats []repeatPacketStat
	Name string
	Fields []packetField
	Table *Table

	firstPkt *packetInfo
}

var repeatPacketMux sync.RWMutex
var repeatPackets = make(map[string][]*repeatPacketInfo)

func (rpi *repeatPacketInfo)addPkt(protocol int, info *packetInfo){
	if len(rpi.Stats) == 1 {
		rpi.firstPkt.RepeatInfo = rpi
	}
	info.RepeatInfo = rpi
	i := sort.Search(len(rpi.Stats), func(i int)(bool){ return rpi.Stats[i].Protocol < protocol })
	rpi.Stats = insert(rpi.Stats, i, repeatPacketStat{protocol, info.State, info.Bound, info.Id})
}

func (rpi *repeatPacketInfo)LowestProtocol()(int){
	return rpi.Stats[0].Protocol
}

func (rpi *repeatPacketInfo)HighestProtocol()(int){
	return rpi.Stats[len(rpi.Stats) - 1].Protocol
}

func pktFieldsEqual(a, b []packetField)(bool){
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		w := b[i]
		if v.Name != w.Name || v.Type != w.Type {
			return false
		}
	}
	return true
}

func registerPacketInfo(info *packetInfo, protocol int){
	repeatPacketMux.RLock()
	lst := repeatPackets[info.Name]
	repeatPacketMux.RUnlock()
	i := sort.Search(len(lst), func(i int)(bool){ return lst[i].LowestProtocol() < protocol })
	ok := false
	if i < len(lst) {
		rp := lst[i]
		if info.Matched {
			ok = rp.Matched && pktFieldsEqual(rp.Fields, info.Fields)
		}else if !rp.Matched {
			ok = tableEqualIgnoreNote(rp.Table, info.Table)
		}
		if ok {
			rp.addPkt(protocol, info)
			loger.Tracef("Found repeated packet %s", info.Name)
			goto DONE
		}
	}
	if i > 0 {
		rp := lst[i - 1]
		if info.Matched {
			ok = rp.Matched && pktFieldsEqual(rp.Fields, info.Fields)
		}else if !rp.Matched {
			ok = tableEqualIgnoreNote(rp.Table, info.Table)
		}
		if ok {
			rp.addPkt(protocol, info)
			loger.Tracef("Found repeated packet %s", info.Name)
			goto DONE
		}
	}
	lst = insert(lst, i, &repeatPacketInfo{
		I: i,
		Matched: info.Matched,
		Stats: []repeatPacketStat{repeatPacketStat{protocol, info.State, info.Bound, info.Id}},
		Name: info.Name,
		Fields: info.Fields,
		Table: info.Table,
		firstPkt: info,
	})
	for _, v := range lst[i + 1:] {
		v.I++
	}
	DONE:
	repeatPacketMux.Lock()
	repeatPackets[info.Name] = lst
	repeatPacketMux.Unlock()
}

