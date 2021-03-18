package str

import (
	"strings"
	"unicode/utf8"

	uuid "github.com/iris-contrib/go.uuid"
)

func UUID() string {
	u, _ := uuid.NewV4()
	return strings.ReplaceAll(u.String(), "-", "")
}

func Len(str string) int {
	return utf8.RuneCountInString(str)
}
