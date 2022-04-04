package v1

import "testing"

func TestHash(t *testing.T) {
	var hash string = Hash("http://qwertyuiop.com")
	if  len(hash) != 7 {
		t.Error("hash length is not 7")
	}
}
