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
  fmt.Println(matrix[0][0])
  for i := 0; i < len(matrix[0]); i++ {
    for j := 0; j < len(matrix); j++ {
      // fmt.Print()
      // fmt.Print(matrix[i][j])
      fmt.Println(toAscii(matrix[i][j]))
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
  
  // fmt.Println(img)
  // fmt.Println(img.ColorModel())


	bounds := img.Bounds()
  // fmt.Println(bounds)
  bx := int((bounds.Max.X - bounds.Min.X)/CR)
  by := int((bounds.Max.Y - bounds.Min.Y)/CR)
  
  img_matrix := make([][]float64, bx)
  for i := range img_matrix {
    img_matrix[i] = make([]float64, by)
  }
  
  // fmt.Println(bx, by)

  x, y := 0, 0
  for i := bounds.Min.X; i < (bounds.Max.X - (bounds.Max.X % CR)); i += CR {
    for j := bounds.Min.Y; j < (bounds.Max.Y - (bounds.Max.Y % CR)); j += CR {
      avg_value := 0.0
      for k := i; k < i + CR; k++ {
        for l := j; l < j + CR; l++ {
          // fmt.Println(k, l)
          // fmt.Println(img.At(k, l))
          avg_value += convertRGB(img.At(k, l))
          //fmt.Println(img.At(k, l))
        }
      }
      img_matrix[x][y] = avg_value / float64(CR*CR)
      // fmt.Println(avg_value)
      y++
		}
    y = 0
    x++
	}
  /*
  for _, h := range img_matrix {
    for _, g := range h {
      fmt.Println(g)
    }
  }
  */
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
