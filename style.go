package main

import (
	"strconv"
	"strings"
)

// MARK: Types

// Style is a text style combining foreground and background text colors with
// text attributes
type Style interface {
	// WithForeground applies a foreground color to the Style
	WithForeground(c Color) Style

	// WithBackground applies a background color to the Style
	WithBackground(c Color) Style

	// With Attributes applies text attributes to the Style
	WithAttributes(a ...Attribute) Style

	// Apply returns a string with the style applied
	Apply(str string) string
}

// style is a text style combining foreground and background text colors with
// text attributes
type style struct {
	fg    Color
	bg    Color
	attrs []Attribute
}

// MARK: Public Methods

// NewStyle returns a new, empty Style
func NewStyle() Style {
	return new(style)
}

// WithForeground applies a foreground color to the Style
func (s *style) WithForeground(c Color) Style {
	if c != 0 {
		s.fg = c
	}
	return s
}

// WithBackground applies a background color to the Style
func (s *style) WithBackground(c Color) Style {
	if c != 0 {
		s.bg = c + 10
	}
	return s
}

// With Attributes applies text attributes to the Style
func (s *style) WithAttributes(a ...Attribute) Style {
	s.attrs = append(s.attrs, a...)
	return s
}

// Apply returns a string with the style applied
func (s *style) Apply(str string) string {
	var start, stop, styled string
	var startCodes, stopCodes []string
	if s.bg != 0 {
		startCodes = append(startCodes, strconv.Itoa(int(s.bg)))
		stopCodes = append(stopCodes, strconv.Itoa(int(Reset+10)))
	}
	if s.fg != 0 {
		startCodes = append(startCodes, strconv.Itoa(int(s.fg)))
		stopCodes = append(stopCodes, strconv.Itoa(int(Reset)))
	}
	for _, c := range s.attrs {
		startCodes = append(startCodes, strconv.Itoa(int(c.start)))
		stopCodes = append(stopCodes, strconv.Itoa(int(c.stop)))
	}

	if len(startCodes) > 0 {
		start = "\u001b[" + strings.Join(startCodes, ";") + "m"
		stop = "\u001b[" + strings.Join(stopCodes, ";") + "m"
	}

	styled = start + str + stop

	return styled
}
