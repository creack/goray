package main

import (
	"encoding/hex"
	"image/color"
	"math"
)

var KeyListStr = map[string]int{
	"  ":      ' ',
	"<esc>":   65307,
	"\n":      65293,
	"<left>":  65361,
	"<up>":    65362,
	"<right>": 65363,
	"<down>":  65364,
}
var KeyListInt = map[int]string{}

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

var colorMap = map[string]color.Color{
	"black":   color.Black,
	"white":   color.White,
	"red":     RgbIntToColor(0xFF0000),
	"lime":    RgbIntToColor(0x00FF00),
	"yellow":  RgbIntToColor(0xFFFF00),
	"blue":    RgbIntToColor(0x0000FF),
	"aqua":    RgbIntToColor(0x00FFFF),
	"magenta": RgbIntToColor(0xFF00FF),
	"purple":  RgbIntToColor(0x800080),
	"silver":  RgbIntToColor(0xC0C0C0),
	"gray":    RgbIntToColor(0x808080),
	"green":   RgbIntToColor(0x008000),
	"teal":    RgbIntToColor(0x008080),
	"olive":   RgbIntToColor(0x808000),
	"navy":    RgbIntToColor(0x000080),
	"maroon":  RgbIntToColor(0x800000),
}

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
	return RgbBytesToColor(colorBytes), nil
}

func RgbBytesToColor(rgb []byte) color.Color {
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

func RgbIntToColor(rgb uint32) color.Color {
	return &color.RGBA{
		B: uint8(rgb & 255),
		G: uint8((rgb >> 8) & 255),
		R: uint8((rgb >> 16) & 255),
		A: uint8((rgb >> 24) & 255),
	}
}

func Delta(a, b, c float64) float64 {
	// b^2 - 4ac
	return b*b - 4*a*c
}

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
