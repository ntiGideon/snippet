package main

import (
	"snippetbox.ogidi.net/internal/assert"
	"testing"
	"time"
)

func TestHumanDate(t *testing.T) {
	tm := time.Date(2025, 03, 20, 12, 23, 0, 0, time.UTC)
	hd := humanDate(tm)
	if hd != "20 Mar 2025 at 12:23" {
		t.Errorf("got %s, expected %s", hd, "20 Mar 2025 at 12:23")
	}

}

func TestHumanTime(t *testing.T) {
	tests := []struct {
		name string
		tm   time.Time
		want string
	}{
		{
			name: "UTC",
			tm:   time.Date(2024, 3, 17, 10, 15, 0, 0, time.UTC),
			want: "17 Mar 2024 at 10:15",
		},
		{
			name: "empty",
			tm:   time.Time{},
			want: "",
		},
		{
			name: "CET",
			tm:   time.Date(2024, 3, 17, 9, 15, 0, 0, time.FixedZone("CET", 1*60*60)),
			want: "17 Mar 2024 at 09:15",
		},
	}
	// Use the t.Run() function to run a sub-test for each test case. The
	// first parameter to this is the name of the test (which is used to
	// identify the sub-test in any log output) and the second parameter is
	// and anonymous function containing the actual test for each case.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hd := humanDate(tt.tm)
			assert.Equal(t, hd, tt.want)
		})
	}
}

func TestInvalidHumanDate(t *testing.T) {
	tests := []struct {
		name string
		tm   time.Time
		want string
	}{
		{
			name: "UTC not valid",
			tm:   time.Date(2024, 3, 17, 10, 15, 0, 0, time.UTC),
			want: "17 Mar 2024 at 11:15",
		},
		{
			name: "empty not valid",
			tm:   time.Time{},
			want: "not a valid date",
		},
		{
			name: "CET not valid",
			tm:   time.Date(2024, 3, 17, 9, 15, 0, 0, time.FixedZone("CET", 1*60*60)),
			want: "17 Mar 2024 at 09:5",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hd := humanDate(tt.tm)
			assert.NotEqual(t, hd, tt.want)
		})
	}
}
