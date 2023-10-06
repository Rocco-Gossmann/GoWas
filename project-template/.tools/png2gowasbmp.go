package main

import (
	"fmt"
	"image"
	_ "image/png"
	"log"
	"os"
)

func printHelp() {
	fmt.Printf(`
usage: go run png2compressedbmp.go source.png  out.go  PackageName   BitmapVarName
	source.png    => path+filename of the png image you want to convert
	out.go        => path+filename of the go file containing the Bitmap definition
	PackageName   => name of generated files package
	BitmapVarName => name under which the Bitmap will be accessable in Go
`)
}

func main() {

	// Make sure all the required intputs exist
	//------------------------------------------------------------------------------
	if len(os.Args) != 5 {
		printHelp()
		return
	}

	// Read Input File
	//------------------------------------------------------------------------------
	filename := os.Args[1]
	outfilename := os.Args[2]

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)

	xs, xe, ys, ye := img.Bounds().Min.X, img.Bounds().Max.X, img.Bounds().Min.Y, img.Bounds().Max.Y

	// Generate Output File
	//------------------------------------------------------------------------------
	outfile, err := os.Create(outfilename)
	if err != nil {
		log.Fatalln(err)
	}
	defer outfile.Close()

	// Write Headers
	//------------------------------------------------------------------------------
	outfile.WriteString(fmt.Sprintf(`package %v

import "github.com/rocco-gossmann/GoWas/core"

var mem_%v = []uint32 {
`, os.Args[3], os.Args[4]))

	// Write Body
	//------------------------------------------------------------------------------

	var totalPixels = 0
	var lastPX uint32 = 0x80000000
	var repeatCount uint32 = 0

	for y := ys; y < ye; y++ {
		for x := xs; x < xe; x++ {
			totalPixels++
			r, g, b, a := img.At(x, y).RGBA()

			var out uint32 = 0
			nr := byte(float64(r) / 0xffff * 0xff)
			ng := byte(float64(g) / 0xffff * 0xff)
			nb := byte(float64(b) / 0xffff * 0xff)
			na := byte(float64(a) / 0xffff * 0xff)

			out = uint32(nr)<<16 + uint32(ng)<<8 + uint32(nb)

			if na == 255 {
				out |= 0x01000000
			}

			if out != lastPX || repeatCount == 127 {
				if lastPX != 0x80000000 {
					outfile.WriteString(fmt.Sprintf("0x%08x, ", lastPX+(repeatCount<<25)))
				}
				lastPX = out
				repeatCount = 0
			} else {
				repeatCount++
			}

		}
	}
	outfile.WriteString(fmt.Sprintf("0x%08x, \n", lastPX+(repeatCount<<25)))

	// Write Footer
	//------------------------------------------------------------------------------
	outfile.WriteString(fmt.Sprintf(`}

var BMP%v = core.CreateBitmapFromCompressed(%v, %v, &mem_%v)`, os.Args[4], xe-xs, totalPixels, os.Args[4]))
}
