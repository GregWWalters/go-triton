package main

// MARK: Types

// Color is a basic text color available in the shell
type Color uint8

// MARK: Constants

// ANSI foreground colors
const (
	Black   Color = 30
	Red     Color = 31
	Green   Color = 32
	Yellow  Color = 33
	Blue    Color = 34
	Magenta Color = 35
	Cyan    Color = 36
	White   Color = 37
	Reset   Color = 39
)

// MARK: Public Functions

func (c Color) Style() Style {
	return &style{
		fg:    c,
		bg:    0,
		attrs: nil,
	}
}

// WithForeground applies a foreground color to the Style
func (c Color) WithForeground(color Color) Style {
	return c.Style().WithForeground(color)
}

// WithBackground applies a background color to the Style
func (c Color) WithBackground(color Color) Style {
	return c.Style().WithBackground(color)
}

// With Attributes applies text attributes to the Style
func (c Color) WithAttributes(a ...Attribute) Style {
	return c.Style().WithAttributes(a...)
}

// Apply returns a string with the style applied
func (c Color) Apply(str string) string {
	return c.Style().Apply(str)
}
