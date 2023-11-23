package liter

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"image"
	"image/color"
	"image/png"
	"io"
	"net/http"
	"sync"
)

type Property struct {
	Name  string `json:"name"`
	Value string `json:"value"`
	Sign  string `json:"signature,omitempty"`
}

func (p *Property) UnmarshalJSONValue(ptr any) (err error) {
	buf, err := base64.StdEncoding.DecodeString(p.Value)
	if err != nil {
		return
	}
	return json.Unmarshal(buf, ptr)
}

func (p *Property) Encode(b *PacketBuilder) {
	b.String(p.Name)
	b.String(p.Value)
	hasSign := len(p.Sign) > 0
	b.Bool(hasSign)
	if hasSign {
		b.String(p.Sign)
	}
}

func (p *Property) DecodeFrom(r *PacketReader) (err error) {
	var ok bool
	if p.Name, ok = r.String(); !ok {
		return io.EOF
	}
	if p.Value, ok = r.String(); !ok {
		return io.EOF
	}
	var hasSign bool
	if hasSign, ok = r.Bool(); !ok {
		return io.EOF
	}
	if hasSign {
		if p.Sign, ok = r.String(); !ok {
			return io.EOF
		}
	} else {
		p.Sign = ""
	}
	return
}

type PlayerProfile struct {
	Id             UUID        `json:"id"`
	Name           string      `json:"name"`
	Properties     []*Property `json:"properties`
	ProfileActions []string    `json:"profileActions"`
}

func (p *PlayerProfile) GetProp(name string) (t *Property) {
	for _, t = range p.Properties {
		if t.Name == name {
			return
		}
	}
	return nil
}

const (
	skinSizeX = 64
	skinSizeY = 64
)

// https://minecraft.wiki/w/Skin
type Skin struct {
	// HTTP cache
	Etag string

	img       image.Image
	meta      map[string]any
	headFront image.Image
	initOnce  sync.Once
}

func (s *Skin) init() {
	s.initOnce.Do(s._init)
}

func (s *Skin) _init() {
	const headSize = 8
	s.headFront = s.subImage(8, 8, headSize, headSize)
}

func (s *Skin) subImage(x, y, width, height int) image.Image {
	img := image.NewNRGBA(image.Rect(0, 0, width, height))
	for dx := 0; dx < width; dx++ {
		for dy := 0; dy < height; dy++ {
			x0, y0 := x+dx, y+dy
			c := color.NRGBAModel.Convert(s.img.At(x0, y0)).(color.NRGBA)
			if c.A == 0 {
				c.R, c.G, c.B = 0, 0, 0
			}
			c.A = 0xff
			img.Set(dx, dy, c)
		}
	}
	return img
}

func (s *Skin) Head() image.Image {
	s.init()
	return s.headFront
}

func (p *PlayerProfile) getDefaultSkin() (skin *Skin, err error) {
	return nil, errors.New("Not implemented for default skin yet")
}

func (p *PlayerProfile) GetSkin(cli *http.Client) (skin *Skin, err error) {
	t := p.GetProp("textures")
	if t == nil {
		return p.getDefaultSkin()
	}
	type Textures struct {
		Skin *struct {
			Url  string         `json:"url"`
			Meta map[string]any `json:"metadata"`
		} `json:"SKIN,omitempty"`
		Cape *struct {
			Url string `json:"url"`
		} `json:"CAPE,omitempty"`
	}
	var data struct {
		Timestamp   int64    `json:"timestamp"`
		ProfileId   string   `json:"profileId"`
		ProfileName string   `json:"profileName"`
		Textures    Textures `json:"textures"`
	}
	if err = t.UnmarshalJSONValue(&data); err != nil {
		return
	}
	if data.Textures.Skin == nil {
		return p.getDefaultSkin()
	}
	if cli == nil {
		cli = http.DefaultClient
	}
	res, err := cli.Get(data.Textures.Skin.Url)
	if err != nil {
		return
	}
	defer res.Body.Close()
	skin = &Skin{
		Etag: res.Header.Get("Etag"),
	}
	if skin.img, err = png.Decode(res.Body); err != nil {
		skin = nil
		return
	}
	skin.meta = data.Textures.Skin.Meta
	return
}
