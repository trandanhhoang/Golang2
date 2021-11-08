package test

import (
	"test/la"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodeFillInAnswer(t *testing.T) {
	var tests = []struct {
		input  string
		expect string
	}{
		{"hoang", "**** *"},
		{"danh hoang", "**** *****"},
	}

	for _, tt := range tests {
		t.Run("util-game", func(t *testing.T) {
			got := la.EncodeFillInAnswer(tt.input)
			assert.Equal(t, got, tt.expect, "they should be equal")
			// if got != tt.expect {
			// t.Errorf("expecting %s, got %s", tt.expect, got)
			// }
		})
	}
}
