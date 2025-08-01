package parsers

import (
	"github.com/lroentgenoil/zebrashMod/internal/printers"
)

func NewFieldValueParser() *CommandParser {
	const code = "^FV"

	return &CommandParser{
		CommandCode: code,
		Parse: func(command string, printer *printers.VirtualPrinter) (any, error) {
			printer.NextElementFieldData = commandText(command, code)

			return nil, nil
		},
	}
}
