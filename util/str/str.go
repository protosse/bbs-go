package str

import (
	"regexp"
	"strings"

	uuid "github.com/iris-contrib/go.uuid"
)

func UUID() string {
	u, _ := uuid.NewV4()
	return strings.ReplaceAll(u.String(), "-", "")
}

func RemoveUnmarshalerDecoder(str string) string {
	re := regexp.MustCompile(`unmarshalerDecoder: ([\s\S]+), error found`)
	return re.FindAllStringSubmatch(str, -1)[0][1]
}
