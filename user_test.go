package main

import (
	"fmt"
	"regexp"
	"testing"
)

func TestUserRegexFormat(t *testing.T) {
	testCases := []struct {
		userID string
		isGood bool
	}{
		{"user1369", true},
		{"foo-bar", false},
		{"foo_bar", false},
		{"foobar12345", false},
		{"foobaz1234", true},
		{"123456789", true},
		{" baz123fooz", false},
		{"b_azel", false},
		{"1369user", true},
	}

	rgx := regexp.MustCompile(fmt.Sprintf("^%s$", UserIDFormat))
	for _, tc := range testCases {
		if rgx.MatchString(tc.userID) != tc.isGood {
			t.Errorf("expected %t, got %t for userID %s", tc.isGood, !tc.isGood, tc.userID)
		}
	}
}
