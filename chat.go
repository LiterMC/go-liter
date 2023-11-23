package liter

import (
	"encoding/json"
	"fmt"
	"strings"
)

// OptBool is a simplified optional boolean type
// liter.TRUE means true
// liter.FALSE means false
// default value 0 means not set
// other values will cause undefined behaviour
type OptBool int

const (
	TRUE  OptBool = 1
	FALSE OptBool = -1
)

var _ json.Marshaler = (OptBool)(0)
var _ json.Unmarshaler = (*OptBool)(nil)

func (b OptBool) String() string {
	switch b {
	case TRUE:
		return "true"
	case FALSE:
		return "false"
	default:
		return "<undefined bool>"
	}
}

func (b OptBool) IsDefined() bool {
	switch b {
	case TRUE, FALSE:
		return true
	default:
		return false
	}
}

func (b OptBool) Bool() bool {
	switch b {
	case TRUE:
		return true
	case FALSE:
		return false
	default:
		panic("Cannot turn undefined optional bool to bool")
	}
}

func (b OptBool) ToBool(def bool) bool {
	switch b {
	case TRUE:
		return true
	case FALSE:
		return false
	default:
		return def
	}
}

func (b OptBool) MarshalJSON() (buf []byte, err error) {
	return json.Marshal(b.ToBool(false))
}

func (b *OptBool) UnmarshalJSON(buf []byte) (err error) {
	var v bool
	if err = json.Unmarshal(buf, &v); err != nil {
		return
	}
	if v {
		*b = TRUE
	} else {
		*b = FALSE
	}
	return
}

type ClickAction string

const (
	OpenUrl         ClickAction = "open_url"
	RunCommand      ClickAction = "run_command"
	SuggestCommand  ClickAction = "suggest_command"
	ChangePage      ClickAction = "change_page"
	CopyToClipboard ClickAction = "copy_to_clipboard"
)

type HoverAction string

const (
	ShowText   HoverAction = "show_text"
	ShowItem   HoverAction = "show_item"
	ShowEntity HoverAction = "show_entity"
)

type (
	Chat struct {
		Component
		Extra []*Component `json:"extra,omitempty"`
	}
	Component struct {
		// content: parse by order
		Text      string          `json:"text,omitempty"`
		Translate string          `json:"translate,omitempty"`
		With      []*Component    `json:"with,omitempty"` // the values for translate
		Keybind   string          `json:"keybind,omitempty"`
		Score     *ScoreComponent `json:"score,omitempty"`
		// properties
		Bold          OptBool     `json:"bold,omitempty"`
		Italic        OptBool     `json:"italic,omitempty"`
		Underlined    OptBool     `json:"underlined,omitempty"`
		Strikethrough OptBool     `json:"strikethrough,omitempty"`
		Obfuscated    OptBool     `json:"obfuscated,omitempty"`
		Font          string      `json:"font,omitempty"`
		Color         string      `json:"color,omitempty"`
		Insertion     string      `json:"insertion,omitempty"`
		ClickEvent    *ClickEvent `json:"clickEvent,omitempty"`
		HoverEvent    *HoverEvent `json:"hoverEvent,omitempty"`
	}
	ClickEvent struct {
		Action ClickAction `json:"action"`
		Value  any         `json:"value"`
	}
	HoverEvent struct {
		Action  HoverAction `json:"action"`
		Content any         `json:"content"`
	}

	ScoreComponent struct {
		Name      string `json:"name"`
		Objective string `json:"objective"`
		Value     int    `json:"value"`
	}
)

var _ json.Marshaler = (*Chat)(nil)
var _ json.Unmarshaler = (*Chat)(nil)
var _ json.Marshaler = (*Component)(nil)
var _ json.Unmarshaler = (*Component)(nil)

func NewChatFromString(s string) (c *Chat) {
	return &Chat{
		Component: Component{
			Text: s,
		},
	}
}

func (c *Component) toMap() (data map[string]any) {
	data = make(map[string]any, 12)
	if c.Text != "" {
		data["text"] = c.Text
	} else if c.Translate != "" {
		data["translate"] = c.Translate
		if len(c.With) != 0 {
			data["with"] = c.With
		}
	} else if c.Keybind != "" {
		data["keybind"] = c.Keybind
	} else if c.Score != nil {
		data["score"] = c.Score
	} else {
		data["text"] = ""
	}
	if c.Bold != 0 {
		data["bold"] = c.Bold.Bool()
	}
	if c.Italic != 0 {
		data["italic"] = c.Italic.Bool()
	}
	if c.Underlined != 0 {
		data["underlined"] = c.Underlined.Bool()
	}
	if c.Strikethrough != 0 {
		data["strikethrough"] = c.Strikethrough.Bool()
	}
	if c.Obfuscated != 0 {
		data["obfuscated"] = c.Obfuscated.Bool()
	}
	if c.ClickEvent != nil {
		data["clickEvent"] = c.ClickEvent
	}
	if c.HoverEvent != nil {
		data["hoverEvent"] = c.HoverEvent
	}
	return
}

func (c *Component) MarshalJSON() (buf []byte, err error) {
	return json.Marshal(c.toMap())
}

func (c *Component) UnmarshalJSON(buf []byte) (err error) {
	return json.Unmarshal(buf, c)
}

// Plain() will return the plain text value for the component
func (c *Component) Plain() string {
	if c.Text != "" {
		return c.Text
	}
	if c.Translate != "" {
		return c.Translate
	}
	if c.Keybind != "" {
		return c.Keybind
	}
	if c.Score != nil {
		return fmt.Sprintf("<score objective=%s target=%s>", c.Score.Objective, c.Score.Name)
	}
	return ""
}

type ChatType int

const (
	_ ChatType = iota
	TextChat
	TranslateChat
	KeybindChat
	ScoreChat
)

// Type will return the content's type of the chat component.
func (c *Component) Type() ChatType {
	if c.Text != "" {
		return TextChat
	}
	if c.Translate != "" {
		return TranslateChat
	}
	if c.Keybind != "" {
		return KeybindChat
	}
	if c.Score != nil {
		return ScoreChat
	}
	return TextChat
}

func (c *Chat) MarshalJSON() (buf []byte, err error) {
	data := c.toMap()
	if len(c.Extra) != 0 {
		extra := make([]map[string]any, len(c.Extra))
		for i, v := range c.Extra {
			extra[i] = v.toMap()
		}
		data["extra"] = extra
	}
	return json.Marshal(data)
}

func (c *Chat) UnmarshalJSON(buf []byte) (err error) {
	return json.Unmarshal(buf, c)
}

func (c *Chat) Plain() string {
	if len(c.Extra) == 0 {
		return c.Component.Plain()
	}
	var sb strings.Builder
	sb.WriteString(c.Component.Plain())
	for _, v := range c.Extra {
		sb.WriteString(v.Plain())
	}
	return sb.String()
}
