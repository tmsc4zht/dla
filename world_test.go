package main

import (
	"testing"
)

func TestOutSide(t *testing.T) {
	w := newWorld(3, 3)

	cases := []struct {
		x   int
		y   int
		b   int
		out bool
	}{
		{x: -1, y: -1, b: 0, out: true},
		{x: 0, y: 0, b: 0, out: false},
		{x: 2, y: 2, b: 0, out: false},
		{x: 3, y: 3, b: 0, out: true},
		{x: 0, y: 0, b: 1, out: true},
		{x: 1, y: 1, b: 1, out: false},
		{x: 2, y: 2, b: 1, out: true},
	}

	for _, c := range cases {
		got := w.outside(c.x, c.y, c.b)
		if got != c.out {
			t.Fatalf("%v faild", c)
		}
	}
}

func TestIsTouch(t *testing.T) {
	w := newWorld(5, 5)

	err := w.set(2, 2, true)
	if err != nil {
		t.Fatal(err)
	}

	cases := []struct {
		x int
		y int
		t bool
	}{
		{x: 1, y: 1, t: false},
		{x: 2, y: 1, t: true},
		{x: 3, y: 1, t: false},
		{x: 1, y: 2, t: true},
		{x: 2, y: 2, t: false},
		{x: 3, y: 2, t: true},
		{x: 1, y: 3, t: false},
		{x: 2, y: 3, t: true},
		{x: 3, y: 3, t: false},
	}

	for _, c := range cases {
		got, err := w.isTouch(c.x, c.y)
		if err != nil {
			t.Fatal(err)
		}
		if got != c.t {
			t.Fatalf("%v faild", c)
		}
	}

}
