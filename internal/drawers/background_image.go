package drawers

import (
	"encoding/base64"
	"image"
	"strings"
	_ "image/png"
	_ "image/jpeg"

	"github.com/fogleman/gg"
	"github.com/lroentgenoil/zebrashMod/drawers"
	"github.com/lroentgenoil/zebrashMod/internal/elements"
)

func NewBackgroundImageDrawer() *ElementDrawer {
	return &ElementDrawer{
		Draw: func(gCtx *gg.Context, element any, _ drawers.DrawerOptions, _ *DrawerState) error {
			imgElement, ok := element.(*elements.BackgroundImage)
			if !ok {
				return nil
			}
			
            if imgElement.Image == "" {
				return nil
			}

			// Decodificar el base64
			decoded, err := base64.StdEncoding.DecodeString(strings.TrimSpace(imgElement.Image))
			if err != nil {
				return nil
			}

			// Leer la imagen
			img, _, err := image.Decode(strings.NewReader(string(decoded)))
			if err != nil {
				return nil
			}

			x := imgElement.Position.X
            y := imgElement.Position.Y

			scale := imgElement.Magnification
			if scale <= 0 {
				scale = 1.0
			}

			width := float64(img.Bounds().Dx()) * scale
			height := float64(img.Bounds().Dy()) * scale
			
			tempCtx := gg.NewContext(int(width), int(height))
			tempCtx.Scale(scale, scale)
			tempCtx.DrawImage(img, 0, 0)

			// Dibujar en el contexto principal
			gCtx.DrawImage(tempCtx.Image(), x, y)

			return nil
		},
	}
}
