package main

import (
	"fmt"
	"image"
	_ "image/png"
	"os"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}

// Normalize the pixel way to handle vectors.
func toPixelVector(win *pixelgl.Window, x float64, y float64) pixel.Vec {
	var (
		newY = win.Bounds().Max.Y - y
	)

	return pixel.Vec{x, newY}
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:     "Coordinates",
		Bounds:    pixel.R(0, 0, 500, 500),
		VSync:     true,
		Resizable: true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	win.SetSmooth(false)

	pointPic, err := loadPicture("POINT.png")
	point := pixel.NewSprite(pointPic, pointPic.Bounds())

	var (
		frames = 0
		second = time.Tick(time.Second)
	)

	fmt.Println("Running...")

	for !win.Closed() {
		win.Clear(pixel.RGB(1, 1, 1))
		point.Draw(win, pixel.IM.Moved(toPixelVector(win, win.Bounds().Min.X+200, win.Bounds().Min.Y+200)))

		win.Update()

		frames++
		select {
		case <-second:
			win.SetTitle(fmt.Sprintf("%s | FPS: %d", cfg.Title, frames))
			frames = 0
		default:
		}
	}
}

func main() {
	pixelgl.Run(run)
}
