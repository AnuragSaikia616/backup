package main

import (
	"testing"
	"time"
)

func TestHumanDate(t *testing.T) {
	tests := []struct {
		name string
		t    time.Time
		want string
	}{
		{"UTC", time.Date(2020, 12, 17, 10, 0, 0, 0, time.UTC), "17 Dec 2020 at 10:00"},
		{"Empty", time.Time{}, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hd := humanDate(tt.t)
			if hd != tt.want {
				t.Errorf("\ngot = %q, want %q", hd, tt.want)
			}
		})
	}
}
