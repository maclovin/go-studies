package main

import (
	"image"
	_ "image/png"
	"os"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
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

func getSpritePosition(spritesheet pixel.Picture, index int) pixel.Rect {
	var treesFrames []pixel.Rect
	for x := spritesheet.Bounds().Min.X; x < spritesheet.Bounds().Max.X; x += 200 {
		for y := spritesheet.Bounds().Min.Y; y < spritesheet.Bounds().Max.Y; y += 200 {
			treesFrames = append(treesFrames, pixel.R(x, y, x+200, y+200))
		}
	}

	return treesFrames[index]
}

func run() {
	// VSync: friendly CPU usage and smooth movements
	cfg := pixelgl.WindowConfig{
		Title:     "Midnight Cocktail",
		Bounds:    pixel.R(0, 0, 1000, 1000),
		VSync:     true,
		Resizable: true,
	}

	win, err := pixelgl.NewWindow(cfg)

	if err != nil {
		panic(err)
	}

	splashPic, err := loadPicture("SPLASH.png")
	trianglePic1, err := loadPicture("TRIANGLE_1.png")
	trianglePic2, err := loadPicture("TRIANGLE_2.png")
	trianglePic3, err := loadPicture("TRIANGLE_3.png")
	cupsSpritesheet, err := loadPicture("CUPS_SPRITE.png")

	splash := pixel.NewSprite(splashPic, splashPic.Bounds())
	triangle1 := pixel.NewSprite(trianglePic1, trianglePic1.Bounds())
	triangle2 := pixel.NewSprite(trianglePic2, trianglePic2.Bounds())
	triangle3 := pixel.NewSprite(trianglePic3, trianglePic3.Bounds())
	newCup := pixel.NewSprite(cupsSpritesheet, getSpritePosition(cupsSpritesheet, 4))

	rotationAngle := 0.0
	lightMaskValue := 0
	lightMask := colornames.White
	cupMask := colornames.Grey
	last := time.Now()

	imd := imdraw.New(nil)

	imd.Color = pixel.RGB(0.0, 0.0, 0.1)
	imd.Push(pixel.V(0, 150))
	imd.Color = pixel.RGB(0, 0, 0)
	imd.Push(pixel.V(1024, 0))
	imd.Rectangle(0)

	triangleLight := imdraw.New(nil)

	triangleLight.Color = pixel.RGB(0, 0, 0)
	triangleLight.Push(pixel.V(600, 100))
	triangleLight.Color = pixel.RGB(0, 0, 0)
	triangleLight.Push(pixel.V(1000, 100))
	triangleLight.Color = pixel.RGB(0.1, 0.1, 0.1)
	triangleLight.Push(pixel.V(700, 1100))
	triangleLight.Polygon(0)

	for !win.Closed() {
		// Keep animation smoother with whatever FPS we're using
		deltaTime := time.Since(last).Seconds()
		last = time.Now()

		rotationAngle += 0.1 * deltaTime

		switch lightMaskValue {
		case 1:
			lightMask = colornames.Grey
		case 10:
			lightMask = colornames.White
		case 20:
			lightMask = colornames.White
			lightMaskValue = 0
		}

		lightMaskValue += 1
		win.Clear(colornames.Black)
		triangleLight.Draw(win)
		triangle1.DrawColorMask(win, pixel.IM.Scaled(pixel.ZV, 3).Rotated(pixel.ZV, rotationAngle).Moved(pixel.Vec{100, 100}), lightMask)
		triangle2.DrawColorMask(win, pixel.IM.Scaled(pixel.ZV, 3).Rotated(pixel.ZV, rotationAngle).Moved(pixel.Vec{500, 500}), lightMask)
		triangle3.DrawColorMask(win, pixel.IM.Scaled(pixel.ZV, 3).Rotated(pixel.ZV, rotationAngle).Moved(pixel.Vec{800, 800}), lightMask)
		imd.Draw(win)
		newCup.DrawColorMask(win, pixel.IM.Moved(pixel.Vec{824, 150}), cupMask)
		splash.Draw(win, pixel.IM.Scaled(pixel.ZV, 2).Moved(win.Bounds().Center()))
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
