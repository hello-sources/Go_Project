package main

import "testiing"

func testad(t *testing.T) {
	if ans := add(1, 2); ans != 3 {
		t.Error("add(1, 2) answer should be 3")
	}
}