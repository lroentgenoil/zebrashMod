package elements

import(
	"image/color"
)

type BackgroundRectangle struct {
	Position	LabelPosition
	Width  		int
	Height 		int
	Color  		color.RGBA
}
