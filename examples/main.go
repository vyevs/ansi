package main

import (
	"fmt"
	"math/rand"

	"github.com/vyevs/ansi"
)

func main() {
	fmt.Println(ansi.FGRed + "red fg" + ansi.Clear)
	fmt.Println(ansi.FGGreen + "green fg" + ansi.Clear)
	fmt.Println(ansi.FGBlue + "blue fg" + ansi.Clear)
	fmt.Println(ansi.FGBlack + "black fg" + ansi.Clear)
	fmt.Println(ansi.FGWhite + "white fg" + ansi.Clear)

	fmt.Println(ansi.BGRed + "red bg" + ansi.Clear)
	fmt.Println(ansi.BGGreen + "green bg" + ansi.Clear)
	fmt.Println(ansi.BGBlue + "blue bg" + ansi.Clear)
	fmt.Println(ansi.BGBlack + "black bg" + ansi.Clear)
	fmt.Println(ansi.BGWhite + "white gb" + ansi.Clear)

	fmt.Println(ansi.Bold + "bold" + ansi.Clear)
	fmt.Println(ansi.Faint + "faint" + ansi.Clear)
	fmt.Println(ansi.Italic + "italic" + ansi.Clear)
	fmt.Println(ansi.Underline + "underline" + ansi.Clear)
	fmt.Println(ansi.SlowBlink + "slow blink" + ansi.Clear)
	fmt.Println(ansi.FastBlink + "fast blink" + ansi.Clear)
	fmt.Println(ansi.Invert + "invert" + ansi.Clear)
	fmt.Println(ansi.Gothic + "gothic" + ansi.Clear)
	fmt.Println(ansi.DoubleUnderline + "double underline" + ansi.Clear)
	fmt.Println(ansi.Framed + "framed" + ansi.Clear)
	fmt.Println(ansi.Encircled + "encircled" + ansi.Clear)
	fmt.Println(ansi.Overlined + "overlined" + ansi.Clear)

	{
		src := rand.Uint64()
		r := byte(src)
		g := byte(src >> 8)
		b := byte(src >> 16)
		fmt.Println(ansi.FGColor(r, g, b) + "random color fg" + ansi.Clear)
	}
	{
		src := rand.Uint64()
		r := byte(src)
		g := byte(src >> 8)
		b := byte(src >> 16)
		fmt.Println(ansi.FGColor(r, g, b) + "random color bg" + ansi.Clear)
	}
	{

		src := rand.Uint64()
		fgR := byte(src)
		fgG := byte(src >> 8)
		fgB := byte(src >> 16)
		bgR := byte(src >> 24)
		bgG := byte(src >> 32)
		bgB := byte(src >> 40)
		fmt.Println(ansi.FGColor(fgR, fgG, fgB) + ansi.FGColor(bgR, bgG, bgB) + "random color fg and bg" + ansi.Clear)
	}
	{

		src := rand.Uint64()
		fgR := byte(src)
		fgG := byte(src >> 8)
		fgB := byte(src >> 16)
		bgR := byte(src >> 24)
		bgG := byte(src >> 32)
		bgB := byte(src >> 40)
		fmt.Println(ansi.FGColor(fgR, fgG, fgB) + ansi.FGColor(bgR, bgG, bgB) + ansi.SlowBlink + ansi.Underline + "random color fg and bg blinking and underlined" + ansi.Clear)
	}

	for _, c := range ansi.Colors {
		fmt.Println(ansi.FGColorName(c.Name) + c.Name)
	}

	fmt.Print(ansi.Clear)
}
