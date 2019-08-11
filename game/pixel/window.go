package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
  "golang.org/x/image/colornames"
)

func run() {
  // VSync: friendly CPU usage and smooth movements
  cfg := pixelgl.WindowConfig{
    Title:  "Pixel Window",
    Bounds: pixel.R(0, 0, 1024, 768),
    VSync: true,
	}

	win, err := pixelgl.NewWindow(cfg)

  if err != nil {
    panic(err)
	}

  win.Clear(colornames.Black)

	for !win.Closed() {
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
