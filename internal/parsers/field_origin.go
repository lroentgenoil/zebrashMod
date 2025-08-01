package parsers

import (
	"strconv"

	"github.com/lroentgenoil/zebrashMod/internal/elements"
	"github.com/lroentgenoil/zebrashMod/internal/printers"
)

func NewFieldOriginParser() *CommandParser {
	const code = "^FO"

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

			if len(parts) > 2 {
				if v, err := strconv.Atoi(parts[2]); err == nil {
					switch v {
					case 0:
						val := elements.TextAlignmentLeft
						printer.NextElementAlignment = &val
					case 1:
						val := elements.TextAlignmentRight
						printer.NextElementAlignment = &val
					case 2:
						val := elements.TextAlignmentJustified
						printer.NextElementAlignment = &val
					}
				}
			}

			printer.NextElementPosition = pos.Add(printer.LabelHomePosition)

			return nil, nil
		},
	}
}
