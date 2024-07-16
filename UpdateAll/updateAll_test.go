package updateall

import "testing"

func TestGetClaspAllJsonData(t *testing.T) {
	t.Run("path contains 'clasp'", func(t *testing.T) {
		got := getClaspAllJsonData()
		t.Log(got)

	})
	// t.Run("does not contain 'clasp'", func(t *testing.T) {
	// 	got := hasClasp("BADNAME")
	// 	want := false

	// 	if got != want {
	// 		t.Errorf("got %t want %t", got, want)
	// 	}
	// })

}
