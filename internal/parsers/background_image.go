package parsers

import (
	"strconv"

	"github.com/ingridhq/zebrash/internal/elements"
	"github.com/ingridhq/zebrash/internal/printers"
)

func NewBackgroundImageParser() *CommandParser {
	const code = "~BI"

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

			result := &elements.BackgroundImage{
				Position: printer.NextElementPosition,
			}
			
			if len(parts) > 2 {
				if magnification, err :=  strconv.ParseFloat(parts[2], 64); err == nil {
					result.Magnification = magnification
				}
			}

			if len(parts) > 3 {
				result.Image = parts[3]
				// if imgData, err := base64.StdEncoding.DecodeString(parts[3]); err == nil {
				// }
			}

			return result, nil
		},
	}
}