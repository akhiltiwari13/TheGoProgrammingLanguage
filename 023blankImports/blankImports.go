// alternates the input image between png and jpeg image formats.This needs to be changed.
package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"os"
)

func main() {
	if err := convert(os.Stdin, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "jpeg: %v\n", err)
		os.Exit(1)
	}
}
func convert(in io.Reader, out io.Writer) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stdout, "Input format =", kind)
	if kind == "jpeg" {
		return png.Encode(out, img)
	} else {
		return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
	}
}
