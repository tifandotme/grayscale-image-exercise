package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"strings"
)

type Pixel struct {
	R, G, B, A uint8
}

func main() {
	// uncomment line below to see the result of your code from task1
	runTask1()

	// uncomment line below to see the result of your code from task2
	runTask2()
}

// Given an two-dimensional slice of pixel, convert it to grayscale
func task1(pixels [][]Pixel) [][]Pixel {
	// Write your code here
	// ----------------------------------------

	for x, v := range pixels {
		for y, pixel := range v {
			gray := uint8((int(pixel.R) + int(pixel.G) + int(pixel.B)) / 3)

			pixels[x][y].R = gray
			pixels[x][y].G = gray
			pixels[x][y].B = gray
		}
	}

	return pixels
	// ----------------------------------------
}

// Given an two-dimensional slice of pixels from an actual image, convert it to grayscale
func task2(pixels [][]Pixel) [][]Pixel {
	// Write your code here
	// ----------------------------------------

	for x, v := range pixels {
		for y, pixel := range v {
			gray := uint8((int(pixel.R) + int(pixel.G) + int(pixel.B)) / 3)

			pixels[x][y].R = gray
			pixels[x][y].G = gray
			pixels[x][y].B = gray
		}
	}

	return pixels
	// ----------------------------------------
}

// Code To Support Task 2, Please don't modify them unless necessary

const originalImageFile = "./img/img.png"
const greyscaleImageFile = "./result/greyscale-img.png"

func runTask1() {
	pixels := [][]Pixel{
		{Pixel{R: 10, G: 20, B: 30, A: 1}, Pixel{R: 40, G: 50, B: 60, A: 1}},
		{Pixel{R: 70, G: 80, B: 90, A: 1}, Pixel{R: 100, G: 110, B: 120, A: 1}},
	}

	fmt.Printf("original: \n")
	printArrayOfPixels(pixels)
	grayPixels := task1(pixels)

	fmt.Printf("\ngreyscale: \n")
	printArrayOfPixels(grayPixels)
}

func printArrayOfPixels(pixels [][]Pixel) {
	if len(pixels) <= 0 {
		fmt.Println("no pixels to print")
		return
	}

	fmt.Printf("-------------\n")
	for i := 0; i < len(pixels); i++ {
		fmt.Printf("|")
		for j := 0; j < len(pixels[0]); j++ {
			fmt.Printf("(R: %3d G: %3d B: %d)|", pixels[i][j].R, pixels[i][j].G, pixels[i][j].B)
		}
		fmt.Printf("\n")
	}
	fmt.Printf("-------------\n")
}

// Below are codes to make this exercise more realistic
// The process can be more straightforward
// However, since we want to balance the complexity
// We mainly divide the steps into 3:
// 1. Extract the pixels
// 2. Modify the pixels
// 3. Create a new image from the modified pixels
func runTask2() {
	// Open the image file
	file, err := os.Open(originalImageFile)
	if err != nil {
		fmt.Println("Failed to open image:", err)
		return
	}
	defer file.Close()

	// Decode the image
	img, format, err := image.Decode(file)
	if err != nil {
		fmt.Println("Failed to decode image:", err)
		return
	}

	// Manipulate the image (invert colors)
	rotImg := greyScaleImage(img)

	// Save the manipulated image
	outputFile, err := os.Create(greyscaleImageFile)
	if err != nil {
		fmt.Println("Failed to create output file:", err)
		return
	}
	defer outputFile.Close()

	switch strings.ToLower(format) {
	case "png":
		png.Encode(outputFile, rotImg)
	default:
		fmt.Println("Unsupported format:", format)
	}
}

// Extracting the pixels
func extractPixels(src image.Image) [][]Pixel {
	bounds := src.Bounds()
	pixels := make([][]Pixel, bounds.Dy())
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		row := make([]Pixel, bounds.Dx())
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := src.At(x, y).RGBA()
			row[x-bounds.Min.X] = Pixel{
				R: uint8(r >> 8),
				G: uint8(g >> 8),
				B: uint8(b >> 8),
				A: uint8(a >> 8),
			}
		}
		pixels[y-bounds.Min.Y] = row
	}
	return pixels
}

// From the modifed pixels, we create a new image
func createImageFromPixels(pixels [][]Pixel) *image.RGBA {
	height, width := len(pixels), len(pixels[0])
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			p := pixels[y][x]
			img.Set(x, y, color.RGBA{p.R, p.G, p.B, p.A})
		}
	}
	return img
}

func greyScaleImage(src image.Image) *image.RGBA {
	pixels := extractPixels(src)
	modifiedPixels := task2(pixels)
	return createImageFromPixels(modifiedPixels)
}
