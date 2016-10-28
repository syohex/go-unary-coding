package unary

import (
	"reflect"
	"testing"
)

func TestUnaryEecode(t *testing.T) {
	expecteds := []struct {
		input    []int
		expected []byte
	}{
		{[]int{1}, []byte{128}},
		{[]int{1, 2}, []byte{160}},
		{[]int{1, 2, 3}, []byte{164}},
		{[]int{1, 2, 3, 4}, []byte{164, 64}},
		{[]int{1, 2, 3, 4, 7}, []byte{164, 64, 128}},
		{[]int{1, 2, 3, 4, 7, 8}, []byte{164, 64, 128, 128}},
		{[]int{1, 2, 3, 4, 7, 16}, []byte{164, 64, 128, 0, 128}},
	}

	for _, e := range expecteds {
		ns := Encode(e.input)
		if !reflect.DeepEqual(ns, e.expected) {
			t.Errorf("unexpected encode: got %v, expected %v",
				ns, e.expected)
		}
	}
}

func TestUnaryDecode(t *testing.T) {
	expecteds := []struct {
		input    []byte
		expected []int
	}{
		{[]byte{1}, []int{1}},
		{[]byte{255}, []int{1, 1, 1, 1, 1, 1, 1, 1}},
		{[]byte{37}, []int{1, 2, 3}},
		{[]byte{0,0,0,1}, []int{25}},
	}

	for _, e := range expecteds {
		ns := Decode(e.input)
		if !reflect.DeepEqual(ns, e.expected) {
			t.Errorf("unexpected decode: got %v, expected %v",
				ns, e.expected)
		}
	}
}
