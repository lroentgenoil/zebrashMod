package parsers

import (
	"strconv"

	"github.com/ingridhq/zebrash/internal/elements"
	"github.com/ingridhq/zebrash/internal/printers"
)

func NewBackgroundRectangleParser() *CommandParser {
	const code = "~BR"

	return &CommandParser{
		CommandCode: code,
		Parse: func(command string, printer *printers.VirtualPrinter) (any, error) {
			pos := elements.LabelPosition{
				CalculateFromBottom: false,
			}

			parts := splitCommand(command, code, 0)

			if len(parts) > 0 {
				if v, err := toPositiveIntField(parts[0]); err == nil {
					pos.X = v
				}
			}

			if len(parts) > 1 {
				if v, err := toPositiveIntField(parts[1]); err == nil {
					pos.Y = v
				}
			}

			printer.NextElementPosition = pos.Add(printer.LabelHomePosition)

			result := &elements.BackgroundRectangle{
				Position: printer.NextElementPosition,
			}
			
			if len(parts) > 2 {
				if width, err := strconv.Atoi(parts[2]); err == nil {
					result.Width = width
				}
			}

			if len(parts) > 2 {
				if height, err := strconv.Atoi(parts[3]); err == nil {
					result.Height = height
				}
			}

			if len(parts) > 3 {
				if r, err := strconv.Atoi(parts[4]); err == nil {
					result.Color.R = uint8(r)
				}
			}

			if len(parts) > 4 {
				if g, err := strconv.Atoi(parts[5]); err == nil {
					result.Color.G = uint8(g)
				}
			}

			if len(parts) > 5 {
				if b, err := strconv.Atoi(parts[6]); err == nil {
					result.Color.B = uint8(b)
				}
			}

			result.Color.A = 255

			return result, nil
		},
	}
}