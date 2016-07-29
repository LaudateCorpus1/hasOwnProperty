package hasOwnProperty

import "testing"

func TestHasOwnProperty(t *testing.T) {
	tt := []struct {
		json, path string
		expected   bool
	}{
		{`null`, "foo", false},
		{`{"foo":42}`, "foo", true},
		{`{"f`, "foo", false},
		{`{"foo":{"bar":true}}`, "foo", true},
		{`{"foo":{"bar":true}}`, "foo.bar", true},
		{`{"foo":{"baz":true}}`, "foo.bar", false},
		{`{"foo":null}`, "foo.bar", false},
	}

	for _, tc := range tt {
		if Test([]byte(tc.json), tc.path) != tc.expected {
			t.Errorf("Fail: %v != Test(%s, %s)", tc.expected, tc.json, tc.path)
		}
	}
}

func BenchmarkHasOwnProperty(b *testing.B) {
	json := []byte(`{"foo":{"bar":true}}`)
	for i := 0; i < b.N; i++ {
		_ = Test(json, "foo.bar")
	}
}

func BenchmarkHasOwnPropertyWhenNonexistent(b *testing.B) {
	json := []byte(`{"foo":{}}`)
	for i := 0; i < b.N; i++ {
		_ = Test(json, "foo.bar")
	}
}
