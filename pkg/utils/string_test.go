package utils_test

import(
	"testing"
)

type RemoveExtraWhitespaceTests struct {
	input string
	expected string
	describe string
}

var removeExtraWhitespaceTests = []removeExtraWhitespaceTest{
  {input: "", expected: "", describe: ""},
  {input: "hello", expected: "hello"},
  {input: "This is a string with spaces.", expected: "This is a string with spaces."},
  {input: "This  text   has  multiple   spaces   between   words.", expected: "This text has multiple spaces between words."},
  {input: "    This text has leading and trailing spaces.    ", expected: "This text has leading and trailing spaces."},
}

func TestRemoveExtraWhitespace(t *testing.T) {
  for _, tt := range removeExtraWhitespaceTests {
    actual := RemoveExtraWhitespace(tt.input)
    if actual != tt.expected {
      t.Errorf("Expected: %s, Got: %s", tt.expected, actual)
    }
  }
}

func BechmarkRemoveExtraWhitespace(t *testing.T) {
  for _, tt := range removeExtraWhitespaceTests {
    RemoveExtraWhitespace(tt.input)
  }
}

type limitTest struct {
  str string
  limit int
  expected string
}

var limitTests = []limitTest{
  {str: "", limit: 10, expected: ""},
  {str: "hello", limit: 5, expected: "hello"},
  {str: "This is a long string.", limit: 10, expected: "This is a"},
  {str: "Hello world!", limit: 0, expected: ""},
  {str: "This is a string.", limit: -1, expected: ""},
}

func TestLimit(t *testing.T) {
  for _, tt := range limitTests {
    actual := Limit(tt.str, tt.limit)
    if actual != tt.expected {
      t.Errorf("Expected: %s, Got: %s for input string: %s and limit: %d", tt.expected, actual, tt.str, tt.limit)
    }
  }
}