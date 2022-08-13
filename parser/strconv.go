package parser

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func IntToStr(i int64) string {
	return strconv.Itoa(int(i))
}

func StrToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func StrToInt64(s string) int64 {
	i, _ := strconv.ParseInt(s, 10, 64)
	return i
}

func FloatToStr(f float64) string {
	return fmt.Sprintf("%f", f)
}

func StringLink(s string) string {
	space := regexp.MustCompile(`\s+`)
	reg, err := regexp.Compile(`[^\w]`)
	if err != nil {
		return err.Error()
	}

	processedString := reg.ReplaceAllString(strings.TrimSpace(strings.ToLower(s)), " ")
	str := space.ReplaceAllString(processedString, " ")
	result := strings.ReplaceAll(str, " ", "-")
	return result
}
