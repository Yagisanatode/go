package validate

import (
	"testing"
)

func TestValidate(t *testing.T) {
	t.Run("path contains 'clasp'", func(t *testing.T) {
		got := hasClasp(claspName)
		want := true

		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})
	t.Run("does not contain 'clasp'", func(t *testing.T) {
		got := hasClasp("BADNAME")
		want := false

		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})

}

// type MockPathFinder struct{}

// func (d MockPathFinder) ExecPath() string {
// 	path := "C:\\Users\\Admin\\Documents\\Yagisanatode\\Study\\Go\\go-claspall\\go\\claspall.exe"
// 	// err := new(error)
// 	return path
// }
