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

  splashPic, err := loadPicture("SPLASH.png")
  trianglePic1, err := loadPicture("TRIANGLE_1.png")
  trianglePic2, err := loadPicture("TRIANGLE_2.png")
  trianglePic3, err := loadPicture("TRIANGLE_3.png")

  splash := pixel.NewSprite(splashPic, splashPic.Bounds())
  triangle1 := pixel.NewSprite(trianglePic1, trianglePic1.Bounds())
  triangle2 := pixel.NewSprite(trianglePic2, trianglePic2.Bounds())
  triangle3 := pixel.NewSprite(trianglePic3, trianglePic3.Bounds())

  rotationAngle := 0.0
  lightMaskValue := 0
  lightMask := colornames.White

	for !win.Closed() {
    rotationAngle += 0.001

    switch lightMaskValue {
    case 1:
      lightMask = colornames.Grey
    case 10:
      lightMask = colornames.White
    case 30:
      lightMask = colornames.White
      lightMaskValue = 0
    }

    lightMaskValue += 1
    win.Clear(colornames.Black)

    triangle1.DrawColorMask(win, pixel.IM.Scaled(pixel.ZV, 3).Rotated(pixel.ZV, rotationAngle).Moved(pixel.Vec{100, 100}), lightMask)
    triangle2.DrawColorMask(win, pixel.IM.Scaled(pixel.ZV, 3).Rotated(pixel.ZV, rotationAngle).Moved(pixel.Vec{500, 500}), lightMask)
    triangle3.DrawColorMask(win, pixel.IM.Scaled(pixel.ZV, 3).Rotated(pixel.ZV, rotationAngle).Moved(pixel.Vec{800, 800}), lightMask)
    splash.Draw(win, pixel.IM.Scaled(pixel.ZV, 2).Moved(win.Bounds().Center()))
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
