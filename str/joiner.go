package str

import (
	"fmt"
	"strings"
)

type joiner struct {
	delimiter string
	prefix    string
	suffix    string
	data      []string
}

func (joiner *joiner) Append(data string) *joiner {
	if data != "" {
		joiner.data = append(joiner.data, data)
	}
	return joiner
}

func (joiner *joiner) Size() int {
	return len(joiner.data)
}

func (joiner *joiner) String() string {
	return fmt.Sprintf("%s%s%s", joiner.prefix, strings.Join(joiner.data, joiner.delimiter), joiner.suffix)
}

func NewJoiner(delimiter, prefix, suffix string) *joiner {
	return &joiner{
		delimiter: delimiter,
		prefix:    prefix,
		suffix:    suffix,
		data:      []string{},
	}
}
