package main

import (
	"testing"
)

type test struct {
	data   []int
	answer int
}

func TestSomaTable(t *testing.T) {
	tests := []test{
		test{data: []int{1, 2, 3}, answer: 6},
		test{[]int{10, 10, 10}, 30},
		test{[]int{10, -10, 10}, 10},
	}

	
	for _, v := range tests {
		z:=Soma(v.data...) 
		if z != v.answer {
			t.Errorf("expected:%v , Got:%v ", v.answer, z)
		}
	}

}

func TestSoma(t *testing.T) {
	teste := Soma(1, 2, 3)
	resultado := 6

	if teste != resultado {
		t.Errorf("expected:%v got:%v", resultado, teste)
	}

}
