package service_test

import (
	"testing"

	"github.com/AlexLitovchenko/softweather/task/internal/app/service"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"abcabcbb", "abc"},
		{"bbbb", "b"},
		{"pwwkew", "wke"},
		{"", ""},
		{"a", "a"},
	}
	s := service.New()
	for _, test := range tests {
		result := s.LongestSubstring(test.input)
		if result != test.expected {
			t.Errorf("Input string- %s, wait- %s, have- %s", test.input, test.expected, result)
		}
	}
}
