package main

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"github.com/notnil/chess/uci"

	"github.com/notnil/chess"
)

var (
	grid  *fyne.Container
	over  *canvas.Image
	start *canvas.Rectangle
	win   fyne.Window
	eng   *uci.Engine
)

func main() {
	a := app.New()
	win = a.NewWindow("Chess")

	game := chess.NewGame()
	grid = createGrid(game)

	over = canvas.NewImageFromResource(nil)
	over.FillMode = canvas.ImageFillContain
	over.Hide()

	start = canvas.NewRectangle(color.Transparent)
	start.StrokeWidth = 4

	win.SetContent(container.NewMax(grid, container.NewWithoutLayout(start, over)))
	win.Resize(fyne.NewSize(480, 480))

	eng = loadOpponent()
	defer eng.Close()
	win.ShowAndRun()
}

func loadOpponent() *uci.Engine {
	e, err := uci.New("stockfish") // you must have stockfish installed and on $PATH
	if err != nil {
		panic(err)
	}

	if err := e.Run(uci.CmdUCI, uci.CmdIsReady, uci.CmdUCINewGame); err != nil {
		panic(err)
	}
	return e
}

func move(m *chess.Move, game *chess.Game, grid *fyne.Container, over *canvas.Image) {
	off := squareToOffset(m.S1())
	cell := grid.Objects[off].(*fyne.Container)
	img := cell.Objects[2].(*piece)

	over.Resource = resourceForPiece(game.Position().Board().Piece(m.S1()))
	over.Resize(img.Size())
	over.Refresh() // clear our old resource before showing

	over.Show()
	img.Resource = nil
	img.Refresh()

	off = squareToOffset(m.S2())
	cell = grid.Objects[off].(*fyne.Container)
	pos2 := cell.Position()

	a := canvas.NewPositionAnimation(over.Position(), pos2, time.Millisecond*500, func(p fyne.Position) {
		over.Move(p)
		over.Refresh()
	})
	a.Start()
	time.Sleep(time.Millisecond * 550)

	game.Move(m)
	refreshGrid(grid, game.Position().Board())
	over.Hide()

	if game.Outcome() != chess.NoOutcome {
		result := "draw"
		switch game.Outcome().String() {
		case "1-0":
			result = "won"
		case "0-1":
			result = "lost"
		}
		dialog.ShowInformation("Game ended",
			"Game "+result+" because "+game.Method().String(), win)
	}
}

func positionToSquare(pos fyne.Position) chess.Square {
	var offX, offY = -1, -1
	for x := float32(0); x <= pos.X; x += grid.Size().Width / 8 {
		offX++
	}
	for y := float32(0); y <= pos.Y; y += grid.Size().Height / 8 {
		offY++
	}

	return chess.Square((7-offY)*8 + offX)
}

func squareToOffset(sq chess.Square) int {
	x := sq % 8
	y := 7 - ((sq - x) / 8)

	return int(x + y*8)
}
