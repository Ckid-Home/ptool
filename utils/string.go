package utils

import (
	"fmt"
	"strconv"
	"strings"

	runewidth "github.com/mattn/go-runewidth"
)

func Capitalize(str string) string {
	return strings.ToUpper(str[:1]) + str[1:]
}

func ContainsI(str string, substr string) bool {
	return strings.Contains(
		strings.ToLower(str),
		strings.ToLower(substr),
	)
}

func ParseInt(str string) int64 {
	str = strings.ReplaceAll(str, ",", "")
	v, _ := strconv.ParseInt(str, 10, 0)
	return v
}

func PrintStringInWidth(str string, width int64, padRight bool) {
	strWidth := int64(0)
	pstr := ""
	for _, char := range str {
		runeWidth := int64(runewidth.RuneWidth(char))
		if strWidth+runeWidth > width {
			break
		}
		pstr += string(char)
		strWidth += runeWidth
	}
	if padRight {
		pstr += strings.Repeat(" ", int(width-strWidth))
	} else {
		pstr = strings.Repeat(" ", int(width-strWidth)) + pstr
	}
	fmt.Print(pstr)
}

func SanitizeText(text string) string {
	text = strings.ReplaceAll(text, "\u00ad", "") // &shy;  invisible Soft hyphen
	text = strings.TrimSpace(text)
	return text
}