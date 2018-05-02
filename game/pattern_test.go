package game

import (
	"regexp"
	"testing"
)

func TestPatternMatching(t *testing.T) {
	// Rows
	assertMatch(t, "^xxx......$", "xxx------")
	assertMatch(t, "^...xxx...$", "---xxx---")
	assertMatch(t, "^......xxx$", "------xxx")
	
	// Columns
	assertMatch(t, "^x..x..x..$", "x--x--x--")
	assertMatch(t, "^.x..x..x.$", "-x--x--x-")
	assertMatch(t, "^..x..x..x$", "--x--x--x")

	// Diagonals
	assertMatch(t, "^x...x...x$", "x---x---x")
	assertMatch(t, "^..x.x.x..$", "--x-x-x--")
}

func assertMatch(t *testing.T, pattern, text string) {
	matched, err := regexp.MatchString("^xxx......$", "xxx------")
	if err != nil {
		t.Fatalf("Error matching %q to %q. %s", pattern, text, err.Error())
	}
	if !matched {
		t.Fatalf("Expected %q to match %q", pattern, text)
	}
}
