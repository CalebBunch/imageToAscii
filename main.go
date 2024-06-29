package main

import (
	"fmt"
	"image/png"
	"log"
	"os"
  "image/color"
)

const PATH string = "image.png"

func main() {
  matrix := readImage(PATH)
  fmt.Println(matrix[0][0])
  for i := 0; i < len(matrix); i++ {
    for j := 0; j < len(matrix[0]); j++ {
      fmt.Println(matrix[i][j])
    }
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
  fmt.Println(bounds)
  bx := int((bounds.Max.X - bounds.Min.X)/4)
  by := int((bounds.Max.Y - bounds.Min.Y)/4)
  
  img_matrix := make([][]float64, bx)
  for i := range img_matrix {
    img_matrix[i] = make([]float64, by)
  }

  for i := bounds.Min.X; i < ((bounds.Max.X/4) - ((bounds.Max.X/4) % 4)); i += 4 {
    for j := bounds.Min.Y; j < ((bounds.Max.Y/4) - ((bounds.Max.Y/4)%4)); j += 4 {
      avg_value := 0.0
      for k := i; k < i + 4; k++ {
        for l := j; l < j + 4; l++ {
          avg_value += convertRGB(img.At(i, j))
        }
      }
      img_matrix[i][j] = avg_value / 16.0
		}
	}
  return img_matrix
}

func convertRGB(c color.Color) float64 {
  // [0, 65535]
  r, g, b, _ := c.RGBA()
  shade := 0.299 * float64(r) + 0.587 * float64(g) + 0.114 * float64(b)
  shade *= (255.0 / 65535.0)
  return shade
}
