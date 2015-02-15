package utils

import (
	"encoding/hex"
	"image/color"
	"math"
)

var (
	// KeyListStr maps all keys with underlying
	// X11 value.
	// See `init()`
	KeyListStr = map[string]int{
		" ":       ' ',
		"<esc>":   65307,
		"\n":      65293,
		"<left>":  65361,
		"<up>":    65362,
		"<right>": 65363,
		"<down>":  65364,
	}
	// KeyListInt is a reverse map from KeyListStr
	KeyListInt = map[int]string{}
)

func init() {
	// Populate the KeyMap
	for i := byte('a'); i <= 'z'; i++ {
		KeyListStr[string([]byte{i})] = int(i)
	}
	for i := byte('A'); i <= 'Z'; i++ {
		KeyListStr[string([]byte{i})] = int(i)
	}
	// Populate the inverted KeyMap
	for k, v := range KeyListStr {
		KeyListInt[v] = k
	}
}

// colorMap maps common colors names with actual value.
// TODO(creack): Remove this. Do we still use it?
var colorMap = map[string]color.Color{
	"black":   color.Black,
	"white":   color.White,
	"red":     RGBIntToColor(0xFF0000),
	"lime":    RGBIntToColor(0x00FF00),
	"yellow":  RGBIntToColor(0xFFFF00),
	"blue":    RGBIntToColor(0x0000FF),
	"aqua":    RGBIntToColor(0x00FFFF),
	"magenta": RGBIntToColor(0xFF00FF),
	"purple":  RGBIntToColor(0x800080),
	"silver":  RGBIntToColor(0xC0C0C0),
	"gray":    RGBIntToColor(0x808080),
	"green":   RGBIntToColor(0x008000),
	"teal":    RGBIntToColor(0x008080),
	"olive":   RGBIntToColor(0x808000),
	"navy":    RGBIntToColor(0x000080),
	"maroon":  RGBIntToColor(0x800000),
}

// DecodeColor decodes a color string to an actual
// Color object. The string can be a color name (see `colorMap`
// or an arbitrary hexadecial RGB value.
func DecodeColor(val string) (color.Color, error) {
	if color, exists := colorMap[val]; exists {
		return color, nil
	}
	if val[:2] == "0x" || val[:2] == "0X" {
		val = val[2:]
	}
	colorBytes, err := hex.DecodeString(val)
	if err != nil {
		return nil, err
	}
	return RGBBytesToColor(colorBytes), nil
}

// RGBBytesToColor creates a Color object
// from its RGBA binary representation.
func RGBBytesToColor(rgb []byte) color.Color {
	for len(rgb) < 4 {
		rgb = append(rgb, 0)
	}
	return &color.RGBA{
		R: rgb[0],
		G: rgb[1],
		B: rgb[2],
		A: rgb[3],
	}
}

//RGBIntToColor creates a Color object from
// its RGBA integer representation.
func RGBIntToColor(rgb uint32) color.Color {
	return &color.RGBA{
		B: uint8(rgb & 255),
		G: uint8((rgb >> 8) & 255),
		R: uint8((rgb >> 16) & 255),
		A: uint8((rgb >> 24) & 255),
	}
}

// Delta is a small helper to determine the delta
// of the given 2nd degree equation.
// (a^2 * x + b * x + c = 0
func Delta(a, b, c float64) float64 {
	// b^2 - 4ac
	return b*b - 4*a*c
}

// SecondDegree is a small helper to solve
// the given 2nd degree equation.
// (a * x^2 + b * x + c = 0
// Returns the smallest positive solution.
// Returns 0 when no solution.
func SecondDegree(a, b, c float64) float64 {
	delta := Delta(a, b, c)
	// delta negative: no solution
	if delta < 0 {
		return 0
	}
	// Two solution: (-b + sqrt(delta)) / 2a and (-b - sqrl(delta)) / 2a
	var (
		k1 = (-b + math.Sqrt(delta)) / (2 * a)
		k2 = (-b - math.Sqrt(delta)) / (2 * a)
	)
	if k1 > 0 && k1 < k2 {
		return k1
	}
	return k2
}
