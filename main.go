package main

import (
	"flag"
	"fmt"
	"image"
	"image/png"
	"math/rand"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run() error {
	w := flag.Int("w", 300, "width")
	h := flag.Int("h", 300, "height")
	n := flag.Int("n", 10000, "particle")
	flag.Parse()

	randInit()

	f, err := os.Create("out.png")
	if err != nil {
		return fmt.Errorf("could not create out.png :%v", err)
	}

	img, err := createImage(*w, *h, *n)
	if err != nil {
		return fmt.Errorf("could not create Image :%v", err)
	}

	if err = png.Encode(f, img); err != nil {
		return fmt.Errorf("could not encode image to png :%v", err)
	}
	return nil
}

func createImage(w, h, nmax int) (image.Image, error) {
	a := newWorld(w, h)

	cx, cy := w/2, h/2

	if err := a.set(cx, cy, true); err != nil {
		return nil, err
	}

	rMax := 0.0
	if w > h {
		rMax = float64(h) / 2
	} else {
		rMax = float64(w) / 2
	}

	a.r = 2

bigloop:
	for n := 0; n < nmax; n++ {
		x, y := randCircle(float64(a.r), cx, cy)

		for {
			t, err := a.get(x, y)
			if err != nil {
				fmt.Fprintf(os.Stderr, "sudeni kuro check: %v\n", err)
				x, y = randCircle(float64(a.r), cx, cy)
				continue
			}
			if t {
				fmt.Fprintf(os.Stderr, "sudeni kuro")
				x, y = randCircle(float64(a.r), cx, cy)
				continue
			}

			t, err = a.isTouch(x, y)
			if err != nil {
				fmt.Fprintf(os.Stderr, "touch check err: %v\n", err)
				x, y = randCircle(float64(a.r), cx, cy)
				continue
			}

			if t {
				if err := a.set(x, y, true); err != nil {
					fmt.Fprintf(os.Stderr, "could not set point :%v\n", err)
				}
				r := distance(cx, cy, x, y) + 1
				if r+1 > a.r {
					a.r = r
				}
				fmt.Printf("%5d %5d %5d %f\n", n, x, y, a.r)
				if a.r > rMax {
					break bigloop
				}
				break
			}

			switch rand.Intn(4) {
			case 0:
				x++
			case 1:
				x--
			case 2:
				y++
			case 3:
				y--
			}
		}
	}

	return a, nil
}
