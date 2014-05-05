package main

import (
	"image/color"
)

var KeyList = map[string]int{
	" ": ' ',
	"": 65307,
	"\\e": 65307,
	"esc": 65307,
	"\n":  65293,
}

func rgbToColor(rgb uint32) color.Color {
	return &color.RGBA{
		B: uint8(rgb & 255),
		G: uint8((rgb >> 8) & 255),
		R: uint8((rgb >> 16) & 255),
		A: uint8((rgb >> 24) & 255),
	}
}
