package drawers

import (
	"github.com/fogleman/gg"
	"github.com/lroentgenoil/zebrashMod/drawers"
	"github.com/lroentgenoil/zebrashMod/internal/elements"
)

func NewBackgroundRectangleDrawer() *ElementDrawer {
	return &ElementDrawer{
		Draw: func(gCtx *gg.Context, element any, _ drawers.DrawerOptions, _ *DrawerState) error {
			rect, ok := element.(*elements.BackgroundRectangle)
			if !ok {
				return nil
			}
			
            x := float64(rect.Position.X)
            y := float64(rect.Position.Y)
            w := float64(rect.Width)
            h := float64(rect.Height)

            gCtx.SetRGBA(
			    float64(rect.Color.R) / 255,
			    float64(rect.Color.G) / 255,
			    float64(rect.Color.B) / 255,
			    float64(rect.Color.A) / 255,
			)

            gCtx.DrawRectangle(x, y, w, h)
            gCtx.Fill()

			return nil
		},
	}
}
