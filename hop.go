package hasOwnProperty

import (
	"bytes"
	"strings"

	jlexer "github.com/mailru/easyjson/jlexer"
)

const delimiter = '.'

func test(in *jlexer.Lexer, path string, start int) (ok bool) {
	if start == len(path)+1 {
		return true
	}
	if in.IsNull() {
		return false
	}

	// We use this somewhat arduous slicing method as it avoids the
	// allocation which is otherwise required by strings.Split
	end := start + strings.IndexByte(path[start:], delimiter)
	if end == start-1 {
		end = len(path)
	}

	needle := path[start:end]
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if key == needle {
			return test(in, path, end+1)
		} else {
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')

	return false
}

// Test checks that the provided property exists in the JSON string. It returns
// a boolean whether it does, and an error if it fails to parse the JSON.
// The "path" may be written in dot notation. For example:
//
//      hasOwnProperty.Test(json, "hello")
//      hasOwnProperty.Test(json, "hello.world")
//
func Test(json []byte, path string) (ok bool) {
	// Quick, cheap smoke test first:
	last := strings.LastIndexByte(path, delimiter) + 1
	if !bytes.Contains(json, []byte(path[last:])) {
		return false
	}

	// Otherwise use the easyjson lexer:
	in := jlexer.Lexer{Data: json}
	return test(&in, path, 0)
}
