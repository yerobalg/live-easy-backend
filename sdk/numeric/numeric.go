package numeric

import (
	"regexp"
	"fmt"
)

func IntToRupiah(value int64) string {
	str := fmt.Sprintf("%d", value)
	result := regexp.MustCompile(`/\B(?=(\d{3})+(?!\d))/g`)
	return "Rp" + result.ReplaceAllString(str, ".")
}