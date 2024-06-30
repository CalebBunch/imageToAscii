package main

import (
	"flag"
	"fmt"
	"image/color"
	"image/png"
	"log"
	"os"
)

const ASCIIMAP string = " .:!i><~+_-?][}{1)(/0*#&8%@$"

var PATH string
var CRH int
var CRV int

func main() {
	cx := flag.Int("cx", 20, "X Compression")
	cy := flag.Int("cy", 30, "Y Compression")
	path := flag.String("path", "image.png", "Path to Image")
	flag.Parse()
	CRH = *cx
	CRV = *cy
	PATH = *path

	matrix := readImage(PATH)
	for i := 0; i < len(matrix[0]); i++ {
		for j := 0; j < len(matrix); j++ {
			fmt.Print(toAscii(matrix[j][i]))
		}
		fmt.Println()
	}

}

func readImage(path string) [][]float64 {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	img, err := png.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	bounds := img.Bounds()
	bx := int((bounds.Max.X - bounds.Min.X) / CRH)
	by := int((bounds.Max.Y - bounds.Min.Y) / CRV)

	img_matrix := make([][]float64, bx)
	for i := range img_matrix {
		img_matrix[i] = make([]float64, by)
	}

	x, y := 0, 0
	for i := bounds.Min.X; i < (bounds.Max.X - (bounds.Max.X % CRH)); i += CRH {
		for j := bounds.Min.Y; j < (bounds.Max.Y - (bounds.Max.Y % CRV)); j += CRV {
			avg_value := 0.0
			for k := i; k < i+CRH; k++ {
				for l := j; l < j+CRV; l++ {
					avg_value += convertRGB(img.At(k, l))
				}
			}
			img_matrix[x][y] = avg_value / float64(CRH*CRV)
			y++
		}
		y = 0
		x++
	}
	return img_matrix
}

func toAscii(shade float64) string {
	index := int(shade / (255.0 / float64(len(ASCIIMAP))))
	return string(ASCIIMAP[index])
}

func convertRGB(c color.Color) float64 {
	r, g, b, _ := c.RGBA()
	shade := 0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)
	shade *= (255.0 / 65535.0)
	return shade
}
