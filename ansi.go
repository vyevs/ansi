package ansi

import (
	"fmt"
)

const FGBlack = "\033[38;2;0;0;0m"
const FGWhite = "\033[38;2;255;255;255m"
const FGRed = "\033[38;2;255;0;0m"
const FGBlue = "\033[38;2;0;0;255m"
const FGGreen = "\033[38;2;0;255;0m"

const BGBlack = "\033[48;2;0;0;0m"
const BGWhite = "\033[48;2;255;255;255m"
const BGRed = "\033[48;2;255;0;0m"
const BGBlue = "\033[48;2;0;0;255m"
const BGGreen = "\033[48;2;0;255;0m"

const Bold = "\033[1m"
const Faint = "\033[2m"
const Italic = "\033[3m"
const Underline = "\033[4m"
const SlowBlink = "\033[5m"
const FastBlink = "\033[6m"
const StopBlink = "\033[25m"
const Invert = "\033[7m"
const Gothic = "\033[20m"
const DoubleUnderline = "\033[21m"
const Framed = "\033[51m"
const Encircled = "\033[52m"
const Overlined = "\033[53m"

const Clear = "\033[m"

func FGColor(r, g, b byte) string {
	return fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
}

func BGColor(r, g, b byte) string {
	return fmt.Sprintf("\033[48;2;%d;%d;%dm", r, g, b)
}

var Colors = []struct {
	Name string
	RGB  [3]byte
}{
	{"black", [3]byte{0, 0, 0}},
	{"white", [3]byte{255, 255, 255}},
	{"red", [3]byte{255, 0, 0}},
	{"lime", [3]byte{0, 255, 0}},
	{"blue", [3]byte{0, 0, 255}},
	{"yellow", [3]byte{255, 255, 0}},
	{"cyan", [3]byte{0, 255, 255}},
	{"aqua", [3]byte{0, 255, 255}},
	{"magenta", [3]byte{255, 0, 255}},
	{"fuchsia", [3]byte{255, 0, 255}},
	{"silver", [3]byte{192, 192, 192}},
	{"gray", [3]byte{128, 128, 128}},
	{"maroon", [3]byte{128, 0, 0}},
	{"olive", [3]byte{128, 128, 0}},
	{"green", [3]byte{0, 128, 0}},
	{"purple", [3]byte{128, 0, 128}},
	{"violet", [3]byte{238, 130, 238}},
	{"teal", [3]byte{0, 128, 128}},
	{"navy", [3]byte{0, 0, 128}},
	{"brown", [3]byte{165, 42, 42}},
	{"orange", [3]byte{255, 165, 0}},
	{"gold", [3]byte{255, 215, 0}},
	{"pink", [3]byte{255, 192, 203}},

	{"dark red", [3]byte{139, 0, 0}},
	{"firebrick", [3]byte{178, 34, 34}},
	{"crimson", [3]byte{220, 20, 60}},
	{"tomato", [3]byte{255, 99, 71}},
	{"coral", [3]byte{255, 127, 80}},
	{"indian red", [3]byte{205, 92, 92}},
	{"light coral", [3]byte{240, 128, 128}},
	{"dark salmon", [3]byte{233, 150, 122}},
	{"salmon", [3]byte{250, 128, 114}},
	{"light salmon", [3]byte{255, 160, 122}},
	{"orange red", [3]byte{255, 69, 0}},
	{"dark orange", [3]byte{255, 140, 0}},
	{"dark golden rod", [3]byte{184, 134, 11}},
	{"golden rod", [3]byte{218, 165, 32}},
	{"pale golden rod", [3]byte{238, 232, 170}},
	{"dark khaki", [3]byte{189, 183, 107}},
	{"khaki", [3]byte{240, 230, 140}},
	{"yellow green", [3]byte{154, 205, 50}},
	{"olive green", [3]byte{85, 107, 47}},
	{"olive drab", [3]byte{107, 142, 35}},
	{"lawn green", [3]byte{124, 252, 0}},
	{"chartreuse", [3]byte{127, 255, 0}},
	{"green yellow", [3]byte{173, 255, 47}},
	{"dark green", [3]byte{0, 100, 0}},
	{"forest green", [3]byte{34, 139, 34}},
	{"lime green", [3]byte{50, 205, 50}},
	{"light green", [3]byte{144, 238, 144}},
	{"pale green", [3]byte{152, 251, 152}},
	{"dark sea green", [3]byte{143, 188, 143}},
	{"medium spring green", [3]byte{0, 250, 154}},
	{"spring green", [3]byte{0, 255, 127}},
	{"sea green", [3]byte{46, 139, 87}},
	{"medium aqua marine", [3]byte{102, 205, 170}},
	{"medium sea green", [3]byte{60, 179, 113}},
	{"light sea green", [3]byte{32, 178, 170}},
	{"dark slate gray", [3]byte{47, 79, 79}},
	{"dark cyan", [3]byte{0, 139, 139}},
	{"light cyan", [3]byte{224, 255, 255}},
	{"dark turquoise", [3]byte{0, 206, 209}},
	{"turquoise", [3]byte{64, 224, 208}},
	{"medium turquoise", [3]byte{72, 209, 204}},
	{"pale turquoise", [3]byte{175, 238, 238}},
	{"aqua marine", [3]byte{127, 255, 212}},
	{"powder blue", [3]byte{176, 224, 230}},
	{"cadet blue", [3]byte{95, 158, 160}},
	{"steel blue", [3]byte{70, 130, 180}},
	{"corn flower blue", [3]byte{100, 149, 237}},
	{"deep sky blue", [3]byte{0, 191, 255}},
	{"dodger blue", [3]byte{30, 144, 255}},
	{"light blue", [3]byte{173, 216, 230}},
	{"sky blue", [3]byte{135, 206, 235}},
	{"light sky blue", [3]byte{135, 206, 250}},
	{"midnight blue", [3]byte{25, 25, 112}},
	{"dark blue", [3]byte{0, 0, 139}},
	{"medium blue", [3]byte{0, 0, 205}},
	{"royal blue", [3]byte{65, 105, 225}},
	{"blue violet", [3]byte{138, 43, 226}},
	{"indigo", [3]byte{75, 0, 130}},
	{"dark slate blue", [3]byte{72, 61, 139}},
	{"slate blue", [3]byte{106, 90, 205}},
	{"medium slate blue", [3]byte{123, 104, 238}},
	{"medium purple", [3]byte{147, 112, 219}},
	{"dark magenta", [3]byte{139, 0, 139}},
	{"dark violet", [3]byte{148, 0, 211}},
	{"dark orchid", [3]byte{153, 50, 204}},
	{"medium orchid", [3]byte{186, 85, 211}},
	{"thistle", [3]byte{216, 191, 216}},
	{"plum", [3]byte{221, 160, 221}},
	{"orchid", [3]byte{218, 112, 214}},
	{"medium violet red", [3]byte{199, 21, 133}},
	{"pale violet red", [3]byte{219, 112, 147}},
	{"deep pink", [3]byte{255, 20, 147}},
	{"hot pink", [3]byte{255, 105, 180}},
	{"light pink", [3]byte{255, 182, 193}},
	{"antique white", [3]byte{250, 235, 215}},
	{"beige", [3]byte{245, 245, 220}},
	{"bisque", [3]byte{255, 228, 196}},
	{"blanched almond", [3]byte{255, 235, 205}},
	{"wheat", [3]byte{245, 222, 179}},
	{"corn silk", [3]byte{255, 248, 220}},
	{"lemon chiffon", [3]byte{255, 250, 205}},
	{"light golden rod yellow", [3]byte{250, 250, 210}},
	{"light yellow", [3]byte{255, 255, 224}},
	{"saddle brown", [3]byte{139, 69, 19}},
	{"sienna", [3]byte{160, 82, 45}},
	{"chocolate", [3]byte{210, 105, 30}},
	{"peru", [3]byte{205, 133, 63}},
	{"sandy brown", [3]byte{244, 164, 96}},
	{"burly wood", [3]byte{222, 184, 135}},
	{"tan", [3]byte{210, 180, 140}},
	{"rosy brown", [3]byte{188, 143, 143}},
	{"moccasin", [3]byte{255, 228, 181}},
	{"navajo white", [3]byte{255, 222, 173}},
	{"peach puff", [3]byte{255, 218, 185}},
	{"misty rose", [3]byte{255, 228, 225}},
	{"lavender blush", [3]byte{255, 240, 245}},
	{"linen", [3]byte{250, 240, 230}},
	{"old lace", [3]byte{253, 245, 230}},
	{"papaya whip", [3]byte{255, 239, 213}},
	{"sea shell", [3]byte{255, 245, 238}},
	{"mint cream", [3]byte{245, 255, 250}},
	{"slate gray", [3]byte{112, 128, 144}},
	{"lavender", [3]byte{230, 230, 250}},
	{"floral white", [3]byte{255, 250, 240}},
	{"alice blue", [3]byte{240, 248, 255}},
	{"ghost white", [3]byte{248, 248, 255}},
	{"honeydew", [3]byte{240, 255, 240}},
	{"ivory", [3]byte{255, 255, 240}},
	{"azure", [3]byte{240, 255, 255}},
	{"snow", [3]byte{255, 250, 250}},
	{"dim gray", [3]byte{105, 105, 105}},
	{"dark gray", [3]byte{169, 169, 169}},
	{"light gray", [3]byte{211, 211, 211}},
	{"gainsboro", [3]byte{220, 220, 220}},
	{"white smoke", [3]byte{245, 245, 245}},
}

var colorNameToRGB = func() map[string][3]byte {
	colorNameToRGB := make(map[string][3]byte, len(Colors))
	for _, color := range Colors {
		colorNameToRGB[color.Name] = color.RGB
	}
	return colorNameToRGB
}()

func FGColorName(name string) string {
	rgb := colorNameToRGB[name]
	return FGColor(rgb[0], rgb[1], rgb[2])
}

func BGColorName(name string) string {
	rgb := colorNameToRGB[name]
	return BGColor(rgb[0], rgb[1], rgb[2])
}
