package numeric

import (
	"golang.org/x/text/language"
    "golang.org/x/text/message"
)

func IntToRupiah(value int64) string {
	printer := message.NewPrinter(language.Indonesian)
	return "Rp" + printer.Sprintf("%d", value)
}