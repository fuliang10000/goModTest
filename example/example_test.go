package example

import "testing"

func TestAdd(t *testing.T) {
	r := Add(2, 3)
	if r != 5 {
		t.Fatalf("Add(2, 3) error, expect:%d, actual:%d", r, 5)
	}
	t.Logf("test Add success.")
}
func TestSub(t *testing.T) {
	r := Sub(3, 2)
	if r != 1 {
		t.Fatalf("Sub(3, 2) error, expect:%d, actual:%d", r, 1)
	}
	t.Logf("test Sub success.")
}
