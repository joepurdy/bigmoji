package main

import (
	"flag"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

// Version is provided at compile time
var Version = "devel"
var versionFlag = flag.Bool("version", false, "Print the version and exit")

func main() {
	// Parse command line flags and validate there's an input file
	// passed in positional args
	flag.Parse()

	if *versionFlag {
		fmt.Printf("bigmoji %s\n", Version)
		os.Exit(0)
	}

	if len(flag.Args()) < 1 {
		fmt.Println("Usage: bigmoji <input.png>")
		os.Exit(1)
	}

	input := flag.Args()[0]

	if err := validateInput(input); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Get the base filename without the extension
	baseFilename := strings.TrimSuffix(filepath.Base(input), filepath.Ext(input))

	// Open the input file and create the output directory (if it doesn't exist)
	file, err := os.Open(input)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	if _, err := os.Stat("out"); os.IsNotExist(err) {
		os.Mkdir("out", 0755)
	}

	// Decode the input file into an image.Image
	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Pad the image with transparent pixels to make it square
	// and determine the length of the longest side
	paddedImg, side := padImage(img)

	// Slice the image into 16 equal subparts
	subs := sliceImage(paddedImg, side)

	// Encode each subpart as a png and write it to the output directory
	for i, sub := range subs {
		// Increment the index to start at 1 instead of 0 by
		// adding 1 to the index each iteration
		out, err := os.Create(fmt.Sprintf("out/big%s_%d.png", baseFilename, i+1))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer out.Close()

		if err := png.Encode(out, sub); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

}

func validateInput(input string) error {
	// Validate the input file exists
	if _, err := os.Stat(input); os.IsNotExist(err) {
		return fmt.Errorf("Input file `%s` does not exist", input)
	}

	// Validate the input file is a png
	if filepath.Ext(input) != ".png" {
		return fmt.Errorf("Input file `%s` is not a png", input)
	}

	return nil
}

func padImage(img image.Image) (*image.RGBA, int) {
	// Check the image bounds and determine which side is longer
	bounds := img.Bounds()
	side := bounds.Max.X
	if bounds.Max.Y > side {
		side = bounds.Max.Y
	}

	// Create a new image padding the shorter side with transparent pixels
	paddedImg := image.NewRGBA(image.Rect(0, 0, side, side))
	draw.Draw(paddedImg, paddedImg.Bounds(), image.Transparent, image.ZP, draw.Src)
	draw.Draw(paddedImg, bounds, img, bounds.Min, draw.Src)

	return paddedImg, side
}

func sliceImage(img *image.RGBA, side int) []image.Image {
	// Slice the image into 16 equal subparts and return each part
	subs := make([]image.Image, 16)
	idx := 0
	sideLen := side / 4
	for j := 0; j < 4; j++ {
		for i := 0; i < 4; i++ {
			subs[idx] = img.SubImage(image.Rect(i*sideLen, j*sideLen, (i+1)*sideLen, (j+1)*sideLen))
			idx++
		}
	}

	return subs
}
