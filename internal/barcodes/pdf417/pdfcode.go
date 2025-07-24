package pdf417

import (
	"image"
	"image/color"

	"github.com/lroentgenoil/zebrashMod/internal/barcodes/utils"
	"github.com/lroentgenoil/zebrashMod/internal/images"
)

type pdfBarcode struct {
	width  int
	height int
	code   *utils.BitList
}

func (c *pdfBarcode) ColorModel() color.Model {
	return color.RGBAModel
}

func (c *pdfBarcode) Bounds() image.Rectangle {
	return image.Rect(0, 0, c.width, c.height)
}

func (c *pdfBarcode) At(x, y int) color.Color {
	if c.code.GetBit(y*c.width + x) {
		return images.ColorBlack
	}
	return images.ColorTransparent
}
