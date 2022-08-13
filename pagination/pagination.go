package pagination

import (
	"github.com/fajarardiyanto/flt-go-utils/parser"
	"math"
)

func GetPage(limit, offset string) int {
	return (parser.StrToInt(offset) - 1) * parser.StrToInt(limit)
}

func TotalPage(total int64, limit string) string {
	return parser.IntToStr(int64(math.Ceil(float64(total) / float64(parser.StrToInt(limit)))))
}
