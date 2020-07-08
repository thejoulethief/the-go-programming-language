package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var palette = []color.Color{color.RGBA{0, 0, 0, 255}, color.RGBA{255, 192, 203, 255}}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

func main() {
	f, err := os.Create("example.gif")
	if err != nil {
		fmt.Printf("%v", err)
	}
	w := bufio.NewWriter(f)
	lissajous(w)
}

func lissajous(out io.Writer) {
	const (
		cycles   = 5     // Number of complex oscillator revsw
		res      = 0.001 // angular res
		size     = 500   // canvas res (size, size)
		nframes  = 512   // number of animation frames
		delay    = 1     // delay between frames in 10ms unit
		freqMult = 2     // change pattern
	)
	freq := rand.Float64() * freqMult
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 //phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
