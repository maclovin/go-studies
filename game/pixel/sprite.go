package main

import (
  "image"
	"os"
	_ "image/png"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
  "golang.org/x/image/colornames"
)

func loadPicture(path string) (pixel.Picture, error) {
  file, err := os.Open(path)

  if err!= nil {
    return nil, err
  }

  defer file.Close()

  img, _, err := image.Decode(file)

  if err != nil {
    return nil, err
  }

  return pixel.PictureDataFromImage(img), nil
}

func run() {
  // VSync: friendly CPU usage and smooth movements
  cfg := pixelgl.WindowConfig{
    Title:  "Midnight Cocktail",
    Bounds: pixel.R(0, 0, 1000, 1000),
    VSync: true,
	}

	win, err := pixelgl.NewWindow(cfg)

  if err != nil {
    panic(err)
	}

  pic, err := loadPicture("SPLASH.png")
  if err != nil {
    panic(err)
  }

  sprite := pixel.NewSprite(pic, pic.Bounds())

  win.Clear(colornames.Black)
  sprite.Draw(win, pixel.IM.Scaled(pixel.ZV, 2).Moved(win.Bounds().Center()))

	for !win.Closed() {
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
