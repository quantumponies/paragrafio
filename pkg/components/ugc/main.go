package ugc

import (
	"github.com/microcosm-cc/bluemonday"
)

func GetUGCPolicy() *bluemonday.Policy {
	return bluemonday.UGCPolicy().
		AllowAttrs("class", "data-id").OnElements("div", "p", "pre", "code", "span", "blockquote").
		AllowAttrs("href", "target").OnElements("a")
}