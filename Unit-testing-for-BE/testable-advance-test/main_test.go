package hehe_test

import (
	"hehe"
	"testing"
)

func TestGetAverageStarsPerRepo(t *testing.T) {
	var tests = []struct {
		username string
		want     float64
	}{
		{"octocat", 1480.375000},
		{"plutov", 15.566667},
	}

	mock := new(hehe.Mock)
	for _, tt := range tests {
		t.Run(tt.username, func(t *testing.T) {
			got, err := hehe.GetAverageStarsPerRepo(mock, tt.username)
			// Don't omit errors even in tests
			if err != nil {
				t.Errorf("expecting nil err, got %v", err)
			}
			if got != tt.want {
				t.Errorf("expecting %f, got %f", tt.want, got)
			}
		})
	}
}
