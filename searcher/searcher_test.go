package searcher

import "testing"

// TestSearch Test
func TestSearch(t *testing.T) {
	res := Search("something special", "*pecial")
	if len(res) != 2 {
		t.Error("Incorrect result")
	}
}
