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
		{[]int{1}, []byte{1}},
		{[]int{91}, []byte{0,0,0,0,0,0,0,0,0,0,0,4}},
		{[]int{1, 2}, []byte{5}},
		{[]int{1, 2, 3}, []byte{37}},
		{[]int{1, 2, 3, 4}, []byte{37, 2}},
		{[]int{1, 2, 3, 4, 7}, []byte{37, 2, 1}},
		{[]int{1, 2, 3, 4, 7, 8}, []byte{37, 2, 1, 1}},
		{[]int{1, 2, 3, 4, 7, 16}, []byte{37, 2, 1, 0, 1}},
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
		{[]byte{0, 0, 0, 1}, []int{25}},
	}

	for _, e := range expecteds {
		ns := Decode(e.input)
		if !reflect.DeepEqual(ns, e.expected) {
			t.Errorf("unexpected decode: got %v, expected %v",
				ns, e.expected)
		}
	}
}

func TestUnaryEncodeDecode(t *testing.T) {
	inputs := [][]int{
		{1,2,3,4},
		{91, 93, 23903, 293},
	}

	for _, i := range inputs {
		bs := Encode(i)
		gs := Decode(bs)
		if !reflect.DeepEqual(gs, i) {
			t.Errorf("unexpected decode: got %v, expected %v",
				gs, i)
		}
	}
}
