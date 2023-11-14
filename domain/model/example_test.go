package model

import (
	"testing"
)

func TestExample_EmptyExample(t *testing.T) {
	tests := []struct {
		name     string
		example  *Example
		expected bool
	}{
		{
			name:     "正常系: モデルが空の場合",
			example:  EmptyExample(),
			expected: true,
		},
		{
			name: "正常系: モデルが空でない場合",
			example: &Example{
				ID:          1,
				ExampleKey:  "test_key",
				ExampleName: "test_name",
			},
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if empty := test.example.IsEmpty(); empty != test.expected {
				t.Errorf("Expected IsEmpty() to return %v, but got %v", test.expected, empty)
			}
		})
	}
}
