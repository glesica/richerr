package main

import (
	"errors"
	"fmt"
	"github.com/glesica/richerr"
	"os"
)

func main() {
	var w, h, d int

	fmt.Println("Input dimensions (w,h,d):")
	_, err := fmt.Scanf("%d,%d,%d\n", &w, &h, &d)
	if err != nil {
		os.Exit(1)
	}

	v, err := Volume(w, h, d)
	if err != nil {
		fmt.Printf("something went terribly wrong: %s\n", err.Error())
		var richErr richerr.Error
		if errors.As(err, &richErr) {
			fmt.Printf("%+v\n", richErr.Fields())
		}
		os.Exit(1)
	}

	fmt.Printf("Volume: %d\n", v)
}

// Volume represents a function within our control, so
// we have a chance to introduce richer errors here.
func Volume(width, height, depth int) (int, error) {
	a, err := Area(width, height)
	if err != nil {
		return 0, richerr.Wrap(err, "area calculation failed").
			WithFields(richerr.Fields{
				{"width", width},
				{"height", height},
			})
	}

	if depth <= 0 {
		return 0, richerr.New("depth must be greater than 0").
			WithField("depth", depth)
	}

	return a * depth, nil
}

// Area is pretending to be an external library function,
// or other code that is not under our control.
func Area(width, height int) (int, error) {
	if width <= 0 || height <= 0 {
		return 0, errors.New("negative dimension")
	}

	return width * height, nil
}
