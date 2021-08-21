package main

import (
	"testing"
)

func TestSoma(t *testing.T) {
	teste := Soma(1, 2, 3)
	resultado := 6

	if teste != resultado {
		t.Errorf("expected:%v got:%v", resultado, teste)
	}

}
