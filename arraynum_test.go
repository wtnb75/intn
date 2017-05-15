package intn

import "testing"

func TestArrayNumMin(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Logf("error: %v", err)
		}
	}()
	ar := NewArrayNum(1)
	t.Error("no error:", ar)
}
