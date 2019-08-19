package main

import (
	"fmt"
	"image"
	_ "image/png"
	"os"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

type Gridsystem struct {
	parentBounds pixel.Rect
	margin       pixel.Vec
	columnCount  float64
}

// TODO
func (g *Gridsystem) getColumn(index float64) pixel.Rect {
	var (
		pieces = g.columnCount
		Min    = pixel.Vec{index * (g.parentBounds.Max.X / pieces), g.parentBounds.Min.Y}
		Max    = pixel.Vec{index*(g.parentBounds.Max.X/pieces) + g.parentBounds.Max.X/pieces, g.parentBounds.Max.Y}
	)

	return pixel.Rect{Min, Max}
}

func (g *Gridsystem) setColumns(columns float64) {
	g.columnCount = columns
}

func (g *Gridsystem) setMargin(margin pixel.Vec) {
	g.margin = margin
}

// Positions returns some values useful to anchor any component with the bound
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

// Anchor something with declared bounds and margins
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

func resizedWindow(window *pixelgl.Window, boundsStore *pixel.Rect) bool {
	currentBounds := window.Bounds()

	if currentBounds.Min.X != boundsStore.Min.X || currentBounds.Min.Y != boundsStore.Min.Y || currentBounds.Max.X != boundsStore.Max.X || currentBounds.Max.Y != boundsStore.Max.Y {
		*boundsStore = currentBounds

		return true
	}

	return false
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:     "Resize me",
		Bounds:    pixel.R(0, 0, 1000, 500),
		VSync:     true,
		Resizable: true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	imd := imdraw.New(nil)
	pointPic, err := loadPicture("POINT.png")
	point := pixel.NewSprite(pointPic, pointPic.Bounds())
	gridIndPic, err := loadPicture("GRID_INDICATOR.png")

	if err != nil {
		panic(err)
	}
	gridInd := pixel.NewSprite(gridIndPic, gridIndPic.Bounds())

	var (
		frames       = 0
		second       = time.Tick(time.Second)
		windowBounds = win.Bounds()
	)

	fmt.Println("Running...")
	fmt.Println(win.Bounds())
	imd.Color = pixel.RGB(1, 0, 1)
	imd.Push(anchorTo(win.Bounds(), pixel.Vec{50, 50}).TopLeft, anchorTo(win.Bounds(), pixel.Vec{50, 50}).BottomRight)
	imd.Rectangle(0)

	grid := Gridsystem{win.Bounds(), pixel.Vec{0, 0}, 8}
	fmt.Println(grid.getColumn(1))

	for !win.Closed() {
		win.Clear(pixel.RGB(0, 0, 0))
		point.Draw(win, pixel.IM.Moved(anchorTo(win.Bounds(), pixel.Vec{25, 25}).TopLeft))
		point.Draw(win, pixel.IM.Moved(anchorTo(win.Bounds(), pixel.Vec{25, 25}).TopRight))
		point.Draw(win, pixel.IM.Moved(anchorTo(win.Bounds(), pixel.Vec{25, 25}).BottomRight))
		point.Draw(win, pixel.IM.Moved(anchorTo(win.Bounds(), pixel.Vec{25, 25}).BottomLeft))
		point.Draw(win, pixel.IM.Moved(anchorTo(win.Bounds(), pixel.Vec{25, 25}).CenterTop))
		point.Draw(win, pixel.IM.Moved(anchorTo(win.Bounds(), pixel.Vec{25, 25}).CenterRight))
		point.Draw(win, pixel.IM.Moved(anchorTo(win.Bounds(), pixel.Vec{25, 25}).CenterBottom))
		point.Draw(win, pixel.IM.Moved(anchorTo(win.Bounds(), pixel.Vec{25, 25}).CenterLeft))

		if resizedWindow(win, &windowBounds) {
			fmt.Println("Window Resized", windowBounds)

			imd = imdraw.New(nil)
			imd.Color = pixel.RGB(1, 0, 1)
			imd.Push(anchorTo(win.Bounds(), pixel.Vec{50, 50}).TopLeft, anchorTo(win.Bounds(), pixel.Vec{50, 50}).BottomRight)
			imd.Rectangle(0)
			grid = Gridsystem{win.Bounds(), pixel.Vec{0, 0}, 8}
		}

		imd.Draw(win)
		gridInd.Draw(win, pixel.IM.Moved(anchorTo(grid.getColumn(0), pixel.Vec{13, 25}).TopLeft))
		gridInd.Draw(win, pixel.IM.Moved(anchorTo(grid.getColumn(0), pixel.Vec{13, 25}).CenterLeft))
		gridInd.Draw(win, pixel.IM.Moved(anchorTo(grid.getColumn(1), pixel.Vec{13, 25}).TopLeft))
		gridInd.Draw(win, pixel.IM.Moved(anchorTo(grid.getColumn(2), pixel.Vec{13, 25}).TopLeft))
		gridInd.Draw(win, pixel.IM.Moved(anchorTo(grid.getColumn(3), pixel.Vec{13, 25}).TopLeft))
		gridInd.Draw(win, pixel.IM.Moved(anchorTo(grid.getColumn(4), pixel.Vec{13, 25}).CenterLeft))
		gridInd.Draw(win, pixel.IM.Moved(anchorTo(grid.getColumn(4), pixel.Vec{13, 25}).TopLeft))
		gridInd.Draw(win, pixel.IM.Moved(anchorTo(grid.getColumn(5), pixel.Vec{13, 25}).TopLeft))
		gridInd.Draw(win, pixel.IM.Moved(anchorTo(grid.getColumn(6), pixel.Vec{13, 25}).TopLeft))
		gridInd.Draw(win, pixel.IM.Moved(anchorTo(grid.getColumn(7), pixel.Vec{13, 25}).TopLeft))
		gridInd.Draw(win, pixel.IM.Moved(anchorTo(grid.getColumn(7), pixel.Vec{13, 25}).CenterLeft))
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
