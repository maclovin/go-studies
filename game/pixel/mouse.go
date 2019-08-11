package main

import (
  "image"
  "os"
  "fmt"
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

  cup := pixel.NewSprite(cupsSpritesheet, getSpritePosition(cupsSpritesheet, 5))

  fmt.Println("Running...")
  for !win.Closed() {
    win.Clear(colornames.Black)
    table.Draw(win, pixel.IM.Moved(win.Bounds().Center()))
    cup.Draw(win, pixel.IM.Moved(win.Bounds().Center()))
    win.Update()
  }
}

func main() {
  pixelgl.Run(run)
}
