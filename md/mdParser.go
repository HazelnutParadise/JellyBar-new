package md

import (
	"github.com/russross/blackfriday/v2"
)

func Parse(input string) string {
	return string(blackfriday.Run([]byte(input)))
}
