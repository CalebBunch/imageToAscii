package main

import (
	"fmt"
	"image/png"
	"log"
	"os"
  "image/color"
)

const PATH string = "image.png"
const ASCIIMAP string = ".:!i><~+_-?][}{1)(/0*#&8%@$"
const CR int = 8

func main() {
  matrix := readImage(PATH)
  for _, row := range matrix {
    for _, val := range row {
      fmt.Print(toAscii(val))
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
  bx := int((bounds.Max.X - bounds.Min.X)/CR)
  by := int((bounds.Max.Y - bounds.Min.Y)/CR)
  
  img_matrix := make([][]float64, bx)
  for i := range img_matrix {
    img_matrix[i] = make([]float64, by)
  }
  
  x, y := 0, 0
  for i := bounds.Min.X; i < (bounds.Max.X - (bounds.Max.X % CR)); i += CR {
    for j := bounds.Min.Y; j < (bounds.Max.Y - (bounds.Max.Y % CR)); j += CR {
      avg_value := 0.0
      for k := i; k < i + CR; k++ {
        for l := j; l < j + CR; l++ {
          avg_value += convertRGB(img.At(k, l))
        }
      }
      img_matrix[x][y] = avg_value / float64(CR*CR)
      y++
		}
    y = 0
    x++
	}
  return img_matrix
}

func toAscii(shade float64) string {
  index := int(shade/(255.0/float64(len(ASCIIMAP))))
  return string(ASCIIMAP[index])
}

func convertRGB(c color.Color) float64 {
  // [0, 65535]
  r, g, b, _ := c.RGBA()
  shade := 0.299 * float64(r) + 0.587 * float64(g) + 0.114 * float64(b)
  shade *= (255.0 / 65535.0)
  return shade
}
