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

type Positions struct {
	TopLeft, TopRight, BottomLeft, BottomRight, CenterTop, CenterRight, CenterBottom, CenterLeft pixel.Vec
}

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

// Anchor something with the declared bounds
func anchorTo(bounds pixel.Rect, margin pixel.Vec) Positions {
	var (
		TopLeft      = pixel.Vec{bounds.Min.X + margin.X, bounds.Max.Y - margin.Y}
		TopRight     = pixel.Vec{bounds.Max.X - margin.X, bounds.Max.Y - margin.Y}
		BottomLeft   = pixel.Vec{bounds.Min.X + margin.X, bounds.Min.Y + margin.Y}
		BottomRight  = pixel.Vec{bounds.Max.X - margin.X, bounds.Min.Y + margin.Y}
		CenterTop    = pixel.Vec{bounds.Center().X, bounds.Max.Y - margin.Y}
		CenterRight  = pixel.Vec{bounds.Max.X - margin.X, bounds.Center().Y}
		CenterBottom = pixel.Vec{bounds.Center().X, bounds.Min.Y + margin.Y}
		CenterLeft   = pixel.Vec{bounds.Min.X + margin.X, bounds.Center().Y}
	)

	return Positions{TopLeft, TopRight, BottomLeft, BottomRight, CenterTop, CenterRight, CenterBottom, CenterLeft}
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:     "Resize me",
		Bounds:    pixel.R(0, 0, 148, 148),
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
		point.Draw(win, pixel.IM.Moved(anchorTo(win.Bounds(), pixel.Vec{25, 25}).TopLeft))
		point.Draw(win, pixel.IM.Moved(anchorTo(win.Bounds(), pixel.Vec{25, 25}).TopRight))
		point.Draw(win, pixel.IM.Moved(anchorTo(win.Bounds(), pixel.Vec{25, 25}).BottomRight))
		point.Draw(win, pixel.IM.Moved(anchorTo(win.Bounds(), pixel.Vec{25, 25}).BottomLeft))
		point.Draw(win, pixel.IM.Moved(anchorTo(win.Bounds(), pixel.Vec{25, 25}).CenterTop))
		point.Draw(win, pixel.IM.Moved(anchorTo(win.Bounds(), pixel.Vec{25, 25}).CenterRight))
		point.Draw(win, pixel.IM.Moved(anchorTo(win.Bounds(), pixel.Vec{25, 25}).CenterBottom))
		point.Draw(win, pixel.IM.Moved(anchorTo(win.Bounds(), pixel.Vec{25, 25}).CenterLeft))
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
