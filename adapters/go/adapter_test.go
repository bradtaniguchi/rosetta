package main

import (
	"reflect"
	"testing"
)

func TestGetToFormatReturnsCamelIfGivenCamel(t *testing.T) {
	type Test struct {
		RequestOptions
		expected string
	}

	toFormat, err := GetToFormat(RequestOptions{
		To: "camel",
	})

	if err != nil {
		t.Errorf("GetToFormat({To: \"camel\"}) returned err: %v", err)
	}
	if toFormat != "camel" {
		t.Errorf("GetToFormat({To: \"camel\"}) did not return camel, it returned %s", toFormat)
	}
}

func TestGetToFormatReturnsCamelIfGivenCamelDirectly(t *testing.T) {
	type Test struct {
		RequestOptions
		expected string
	}

	toFormat, err := GetToFormat(RequestOptions{
		Camel: true,
	})

	if err != nil {
		t.Errorf("GetToFormat(Camel: true}) returned err: %v", err)
	}
	if toFormat != "camel" {
		t.Errorf("GetToFormat({Camel: true }) did not return camel, it returned %s", toFormat)
	}
}

func TestGetToFormatReturnsKebabIfGivenKebab(t *testing.T) {
	toFormat, err := GetToFormat(RequestOptions{
		To: "kebab",
	})

	if err != nil {
		t.Errorf("GetToFormat({To: \"kebab\"}) returned err: %v", err)
	}
	if toFormat != "kebab" {
		t.Errorf("GetToFormat({To: \"kebab\"}) did not return kebab, it returned %s", toFormat)
	}
}

func TestGetToFormatReturnsKebabIfGivenKebabDirectly(t *testing.T) {
	toFormat, err := GetToFormat(RequestOptions{
		Kebab: true,
	})

	if err != nil {
		t.Errorf("GetToFormat({Kebab: true}) returned err: %v", err)
	}
	if toFormat != "kebab" {
		t.Errorf("GetToFormat({Kebab: true}) did not return kebab, it returned %s", toFormat)
	}
}

func TestGetToFormatReturnsPascalIfGivenPascal(t *testing.T) {
	toFormat, err := GetToFormat(RequestOptions{
		To: "pascal",
	})

	if err != nil {
		t.Errorf("GetToFormat({To: \"pascal\"}) returned err: %v", err)
	}
	if toFormat != "pascal" {
		t.Errorf("GetToFormat({To: \"pascal\"}) did not return pascal, it returned %s", toFormat)
	}
}

func TestGetToFormatReturnsPascalIfGivenPascalDirectly(t *testing.T) {
	toFormat, err := GetToFormat(RequestOptions{
		Pascal: true,
	})

	if err != nil {
		t.Errorf("GetToFormat({Pascal: true}) returned err: %v", err)
	}
	if toFormat != "pascal" {
		t.Errorf("GetToFormat({Pascal: true}) did not return pascal, it returned %s", toFormat)
	}
}

func TestGetToFormatReturnsSnakeIfGivenSnake(t *testing.T) {
	toFormat, err := GetToFormat(RequestOptions{
		To: "snake",
	})

	if err != nil {
		t.Errorf("GetToFormat({To: \"snake\"}) returned err: %v", err)
	}
	if toFormat != "snake" {
		t.Errorf("GetToFormat({To: \"snake\"}) did not return snake, it returned %s", toFormat)
	}
}

func TestGetToFormatReturnsSnakeIfGivenSnakeDirectly(t *testing.T) {
	toFormat, err := GetToFormat(RequestOptions{
		Snake: true,
	})

	if err != nil {
		t.Errorf("GetToFormat({Snake: true}) returned err: %v", err)
	}
	if toFormat != "snake" {
		t.Errorf("GetToFormat({Snake: true}) did not return snake, it returned %s", toFormat)
	}
}

func TestCapitalizeWorks(t *testing.T) {
	w := Capitalize("hello")

	if w != "Hello" {
		t.Error("Capitalize(\"hello\") did not return \"Hello\"")
	}
}

func TestCapitalizeHandlesEmptyString(t *testing.T) {
	w := Capitalize("")

	if w != "" {
		t.Error("Capitalize(\"\") did not return \"\"")
	}
}

func TestTokenize(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "snake_case splits on underscores",
			input:    "hello_world_foo",
			expected: []string{"hello", "world", "foo"},
		},
		{
			name:     "kebab-case splits on hyphens",
			input:    "hello-world-foo",
			expected: []string{"hello", "world", "foo"},
		},
		{
			name:     "camelCase splits on uppercase boundaries",
			input:    "helloWorldFoo",
			expected: []string{"hello", "world", "foo"},
		},
		{
			name:     "camelCase handles acronyms correctly",
			input:    "HTTPServer",
			expected: []string{"http", "server"},
		},
		{
			name:     "PascalCase handles acronyms correctly",
			input:    "HTTPServer",
			expected: []string{"http", "server"},
		},
		{
			name:     "camelCase handles acronyms in middle",
			input:    "getHTTPResponse",
			expected: []string{"get", "http", "response"},
		},
		{
			name:     "camelCase handles acronyms at end",
			input:    "serveHTTP",
			expected: []string{"serve", "http"},
		},
		{
			name:     "all-caps acronym returns single token",
			input:    "HTTP",
			expected: []string{"http"},
		},
		{
			name:     "two-letter acronyms handled correctly",
			input:    "IPAddress",
			expected: []string{"ip", "address"},
		},
		{
			name:     "PascalCase splits on uppercase boundaries",
			input:    "HelloWorldFoo",
			expected: []string{"hello", "world", "foo"},
		},
		{
			name:     "handles multi-byte unicode characters",
			input:    "caféAuLait",
			expected: []string{"café", "au", "lait"},
		},
		{
			name:     "handles multi-byte unicode characters with acronyms",
			input:    "HTTPÜberServer",
			expected: []string{"http", "über", "server"},
		},
		{
			name:     "single word returns single token",
			input:    "hello",
			expected: []string{"hello"},
		},
		{
			name:     "leading and trailing whitespace is trimmed",
			input:    "  hello_world  ",
			expected: []string{"hello", "world"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Tokenize(tt.input)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("Tokenize(%q) = %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}
