package main

// MARK: Types

// Attribute is a text style. Attributes can be combined, and not every
// attribute will necessarily be supported by every terminal.
type Attribute struct {
	start uint8
	stop  uint8
}

// MARK: Constants

// style codes
const (
	normal        uint8 = 0
	bold          uint8 = 1
	dim           uint8 = 2
	italic        uint8 = 3
	underlined    uint8 = 4
	blinking      uint8 = 5
	alsoBlink     uint8 = 6
	inverted      uint8 = 7
	hidden        uint8 = 8
	strikethrough uint8 = 9
)

// exported attributes
var (
	Normal        = Attribute{start: normal, stop: 20}
	Bold          = Attribute{start: bold, stop: 22}
	Dim           = Attribute{start: dim, stop: 22}
	Italic        = Attribute{start: italic, stop: 23}
	Underlined    = Attribute{start: underlined, stop: 24}
	Blinking      = Attribute{start: blinking, stop: 25}
	Inverted      = Attribute{start: inverted, stop: 27}
	Hidden        = Attribute{start: hidden, stop: 28}
	Strikethrough = Attribute{start: strikethrough, stop: 29}
)

// MARK: Public Functions

func (a Attribute) Style() Style {
	return &style{
		fg:    0,
		bg:    0,
		attrs: []Attribute{a},
	}
}

// WithForeground applies a foreground color to the Style
func (a Attribute) WithForeground(c Color) Style {
	return a.Style().WithForeground(c)
}

// WithBackground applies a background color to the Style
func (a Attribute) WithBackground(c Color) Style {
	return a.Style().WithBackground(c)
}

// With Attributes applies text attributes to the Style
func (a Attribute) WithAttributes(attr ...Attribute) Style {
	return a.Style().WithAttributes(attr...)
}

// Apply returns a string with the style applied
func (a Attribute) Apply(str string) string {
	return a.Style().Apply(str)
}
