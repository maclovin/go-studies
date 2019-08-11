package main

import (
  "image"
  "os"
  "fmt"
  "time"
  "image/color"
  "math/rand"
  _ "image/png"

  "github.com/faiface/pixel"
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

func getSpritePosition(spritesheet pixel.Picture, index int) (pixel.Rect) {
  var treesFrames []pixel.Rect
  for x := spritesheet.Bounds().Min.X; x < spritesheet.Bounds().Max.X; x += 200 {
    for y := spritesheet.Bounds().Min.Y; y < spritesheet.Bounds().Max.Y; y += 200 {
      treesFrames = append(treesFrames, pixel.R(x, y, x+200, y+200))
    }
  }

  return treesFrames[index]
}

func run() {
  cfg := pixelgl.WindowConfig{
    Title:  "Midnight Cocktail",
    Bounds: pixel.R(0, 0, 1000, 1000),
    VSync:  true,
  }

  win, err := pixelgl.NewWindow(cfg)
  if err != nil {
    panic(err)
  }

  win.SetSmooth(true)

  tablePic, err := loadPicture("TABLE_GRID.png")
  table := pixel.NewSprite(tablePic, tablePic.Bounds())

  cupsSpritesheet, err := loadPicture("CUPS_SPRITE.png")

  var (
    camPos = pixel.ZV
    camSpeed = 500.0

    cups []*pixel.Sprite
    cupMatrices []pixel.Matrix
  )

  fmt.Println("Running...")

  last := time.Now()

  var (
		frames = 0
		second = time.Tick(time.Second)
	)

  for !win.Closed() {
    deltaTime := time.Since(last).Seconds()
    last = time.Now()

    cam := pixel.IM.Scaled(pixel.ZV, 1).Moved(win.Bounds().Center().Sub(camPos))
    win.SetMatrix(cam)

    if win.JustPressed(pixelgl.MouseButtonLeft) {
      newCup := pixel.NewSprite(cupsSpritesheet, getSpritePosition(cupsSpritesheet, rand.Intn(6)))
      cups = append(cups, newCup)
      mouse := cam.Unproject(win.MousePosition())
      cupMatrices = append(cupMatrices, pixel.IM.Scaled(pixel.ZV, 1).Moved(mouse))
    }

    if win.Pressed(pixelgl.KeyLeft) {
      camPos.X -= camSpeed * deltaTime
    }
    if win.Pressed(pixelgl.KeyRight) {
      camPos.X += camSpeed * deltaTime
    }
    if win.Pressed(pixelgl.KeyDown) {
      camPos.Y -= camSpeed * deltaTime
    }
    if win.Pressed(pixelgl.KeyUp) {
      camPos.Y += camSpeed * deltaTime
    }

    fmt.Println(color.RGBA{0,0,0,255})

    win.Clear(colornames.Black)
    table.Draw(win, pixel.IM.Moved(pixel.Vec{camPos.X, camPos.Y}))

    for i, cup := range cups {
			cup.Draw(win, cupMatrices[i])
		}

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
