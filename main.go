package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

var (
	title        string = "Tic Tac Toe Game"
	screenWidth  int32  = 600
	screenHeight int32  = 600
	gridX        int32  = screenWidth / 4
	gridY        int32  = screenHeight / 4
	gridWidth    int32  = screenHeight / 2
	gridHeight   int32  = screenHeight / 2
	Grid         string = "........."
	crossTex     *sdl.Texture
	circleTex    *sdl.Texture
	startTex     *sdl.Texture
	titleTex     *sdl.Texture
	window       *sdl.Window
	renderer     *sdl.Renderer
)

const (
	imgWidth    = 100
	imgHeight   = 100
	startWidth  = 100
	startHeight = 79
	titlewidth  = 600
	titleheight = 338
)

func initializeGrid() {
	renderer.SetDrawColor(255, 255, 255, 0)
	gridX = screenWidth / 4
	gridY = screenHeight / 4
	gridWidth = screenWidth / 2
	gridHeight = screenHeight / 2
	renderer.DrawRect(&sdl.Rect{X: gridX, Y: gridY, W: gridWidth, H: gridHeight})
	renderer.SetDrawColor(0, 255, 0, 0)
	// Draw Grid
	var i int32 = 0
	for ; i < 2; i++ {
		renderer.DrawLine(gridX, gridY+(i+1)*gridHeight/3, gridX+gridWidth, gridY+(i+1)*gridHeight/3)
		renderer.DrawLine(gridX+(i+1)*gridWidth/3, gridY, gridX+(i+1)*gridWidth/3, gridY+gridHeight)
	}
	i = 0
	for ; i < 9; i++ {
		if Grid[i] == UsrMove {
			px, py := calcCoord(int8(i+1), true)
			renderer.Copy(circleTex,
				&sdl.Rect{X: 0, Y: 0, W: imgWidth, H: imgHeight},
				&sdl.Rect{X: px, Y: py, W: imgWidth, H: imgHeight})
		} else if Grid[i] == OppMove {
			px, py := calcCoord(int8(i+1), true)
			renderer.Copy(crossTex,
				&sdl.Rect{X: 0, Y: 0, W: imgWidth, H: imgHeight},
				&sdl.Rect{X: px, Y: py, W: imgWidth, H: imgHeight})
		}
	}
	renderer.SetDrawColor(0, 0, 0, 0)
}

func startScreen() {
	renderer.Copy(titleTex,
		&sdl.Rect{X: 0, Y: 0, W: titlewidth, H: titleheight},
		&sdl.Rect{X: 0, Y: 0, W: screenWidth, H: gridY})
	renderer.Copy(startTex,
		&sdl.Rect{X: 0, Y: 0, W: startWidth, H: startHeight},
		&sdl.Rect{X: gridX, Y: screenHeight - startHeight - 100, W: gridWidth, H: gridHeight / 2})

}
func hidestartScreen() {
	renderer.FillRect(&sdl.Rect{X: gridX, Y: gridY + gridHeight, W: gridWidth, H: gridHeight / 2})
	renderer.FillRect(&sdl.Rect{X: 0, Y: 0, W: screenWidth, H: gridY})
}

func calcCoord(pos int8, img bool) (x, y int32) {
	pd := int32((pos - 1) / 3)
	pm := int32((pos - 1) % 3)
	colWidth := gridWidth / 3
	rowWidth := gridHeight / 3

	x = gridX + colWidth*(pm) + colWidth/2
	y = gridY + rowWidth*(pd) + rowWidth/2
	if img {
		x -= imgWidth / 2
		y -= imgHeight / 2
	}
	return x, y
}

func UpdateGrid(pos int8, player rune) {
	if pos != -1 {
		px, py := calcCoord(pos, true)
		if player == UsrMove {
			renderer.Copy(circleTex,
				&sdl.Rect{X: 0, Y: 0, W: imgWidth, H: imgHeight},
				&sdl.Rect{X: px, Y: py, W: imgWidth, H: imgHeight})
		} else {
			renderer.Copy(crossTex,
				&sdl.Rect{X: 0, Y: 0, W: imgWidth, H: imgHeight},
				&sdl.Rect{X: px, Y: py, W: imgWidth, H: imgHeight})
		}
	}
}

func _checkInsideGrid(x, y int32) bool {
	return (x >= gridX && x <= gridX+gridWidth && y >= gridY && y <= gridY+gridHeight)
}

func getPos(x, y int32) int8 {
	y -= gridY
	x -= gridX
	colWidth := gridWidth / 3
	rowWidth := gridHeight / 3
	px := x / colWidth
	py := y / rowWidth
	return int8((1 + 3*py) + px)
}

func DrawLine() {
	var checkHorizontal bool = false
	var checkVertical bool = false
	var x1, y1, x2, y2 int32
	for i := 0; i < 3; i++ {
		h := 3 * i
		checkHorizontal = checkHorizontal || (Grid[h] == Grid[h+1] && Grid[h+1] == Grid[h+2] && Grid[h+1] != '.')
		checkVertical = checkVertical || (Grid[i] == Grid[i+3] && Grid[i+3] == Grid[i+6] && Grid[i+3] != '.')
		if checkHorizontal {
			renderer.SetDrawColor(0, 0, 255, 0)
			x1, y1 = calcCoord(int8(h+1), false)
			x2, y2 = calcCoord(int8(h+3), false)
			renderer.DrawLine(x1, y1, x2, y2)
			renderer.SetDrawColor(0, 0, 0, 0)
			return
		}
		if checkVertical {
			renderer.SetDrawColor(0, 0, 255, 0)
			x1, y1 = calcCoord(int8(i+1), false)
			x2, y2 = calcCoord(int8(i+6+1), false)
			renderer.DrawLine(x1, y1, x2, y2)
			renderer.SetDrawColor(0, 0, 0, 0)
			return
		}
	}
	var checkDiagonal bool = (Grid[0] == Grid[4] && Grid[4] == Grid[8] && Grid[4] != '.')
	if checkDiagonal {
		renderer.SetDrawColor(0, 0, 255, 0)
		x1, y1 = calcCoord(int8(1), false)
		x2, y2 = calcCoord(int8(9), false)
		renderer.DrawLine(x1, y1, x2, y2)
		renderer.SetDrawColor(0, 0, 0, 0)
		return
	}
	checkDiagonal = (Grid[2] == Grid[4] && Grid[4] == Grid[6] && Grid[4] != '.')
	if checkDiagonal {
		renderer.SetDrawColor(0, 0, 255, 0)
		x1, y1 = calcCoord(int8(3), false)
		x2, y2 = calcCoord(int8(7), false)
		renderer.DrawLine(x1, y1, x2, y2)
		renderer.SetDrawColor(0, 0, 0, 0)
	}
}

func playGame() (rcode int32) {
	hidestartScreen()
	var pos int8
	rcode = 0
	for {
		pos = -1
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch Etype := event.(type) {
			case *sdl.QuitEvent:
				rcode = 1
				return
			case *sdl.MouseButtonEvent:
				mousex, mousey, state := sdl.GetMouseState()
				if _checkInsideGrid(mousex, mousey) && state == 1 {
					pos = getPos(mousex, mousey)
					if Grid[pos-1] != '.' {
						fmt.Printf("Already inserted\n")
						pos = -1
					} else {
						Grid = _replace(Grid, int(pos-1), UsrMove)
					}
				}
			case *sdl.WindowEvent:
				if Etype.Event == sdl.WINDOWEVENT_SIZE_CHANGED {
					screenWidth, screenHeight = window.GetSize()
					initializeGrid()
				}
			}
		}
		UpdateGrid(pos, UsrMove)
		if checkFinish(Grid) {
			fmt.Printf("Congratulations! You won\n")
			DrawLine()
			Grid = "........."
			return
		}
		if pos != -1 {
			pos = GetOptMove(Grid)
			if pos == -1 {
				fmt.Printf("Draw\n")
				Grid = "........."
				return
			} else {
				UpdateGrid(pos, OppMove)
				Grid = _replace(Grid, int(pos-1), OppMove)
			}
		}
		if checkFinish(Grid) {
			fmt.Printf("Computer won\n")
			DrawLine()
			Grid = "........."
			return
		}
		renderer.Present()
	}
}

func main() {
	var err error
	if err = sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Printf("Initialization Error: %s", err)
		return
	}

	window, err = sdl.CreateWindow(title, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, screenWidth, screenHeight, sdl.WINDOW_OPENGL)
	if err != nil {
		fmt.Printf("Window Initialization Error: %s", err)
		return
	}
	defer window.Destroy()

	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Printf("Renderer Initialization Error: %s", err)
		return
	}
	defer renderer.Destroy()

	img, err := sdl.LoadBMP("sprites/cross.bmp")
	if err != nil {
		fmt.Printf("Loading Sprite Error: %s", err)
		return
	}
	defer img.Free()
	crossTex, err = renderer.CreateTextureFromSurface(img)
	if err != nil {
		fmt.Printf("Texture Creation Error: %s\n", err)
	}
	defer crossTex.Destroy()
	img, err = sdl.LoadBMP("sprites/circle.bmp")
	if err != nil {
		fmt.Printf("Loading Sprite Error: %s", err)
		return
	}
	circleTex, err = renderer.CreateTextureFromSurface(img)
	if err != nil {
		fmt.Printf("Texture Creation Error: %s", err)
		return
	}
	defer circleTex.Destroy()
	img, err = sdl.LoadBMP("sprites/start.bmp")
	if err != nil {
		fmt.Printf("Loading Sprite Error: %s", err)
		return
	}
	startTex, err = renderer.CreateTextureFromSurface(img)
	if err != nil {
		fmt.Printf("Texture Creation Error: %s", err)
		return
	}
	img, err = sdl.LoadBMP("sprites/title.bmp")
	if err != nil {
		fmt.Printf("Loading Sprite Error: %s", err)
		return
	}
	titleTex, err = renderer.CreateTextureFromSurface(img)
	if err != nil {
		fmt.Printf("Texture Creation Error: %s", err)
		return
	}
	initializeGrid()
	defer startTex.Destroy()
	for {
		startScreen()
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch Etype := event.(type) {
			case *sdl.QuitEvent:
				return
			case *sdl.WindowEvent:
				if Etype.Event == sdl.WINDOWEVENT_SIZE_CHANGED {
					screenWidth, screenHeight = window.GetSize()
					startScreen()
				}
			case *sdl.KeyboardEvent:
				if Etype.Keysym.Scancode == sdl.SCANCODE_SPACE {
					renderer.Clear()
					initializeGrid()
					if playGame() == 1 {
						return
					}
				}
			}
		}
		renderer.Present()
	}
}
