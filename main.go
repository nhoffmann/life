package main

import (
	"fmt"
	"os"
	"strconv"
	"syscall/js"

	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/markfarnan/go-canvas/canvas"
	"github.com/nhoffmann/life/pattern"
	"github.com/nhoffmann/life/renderer"
	"github.com/nhoffmann/life/rle"
	"github.com/nhoffmann/life/simulator"
)

var done chan struct{}

type App struct {
	s *simulator.Simulator

	cvs *canvas.Canvas2d
	ir  renderer.ImageRenderer
}

func NewApp() App {
	app := App{}

	width := int(js.Global().Get("innerWidth").Float())
	height := int(js.Global().Get("innerHeight").Float())

	app.cvs, _ = canvas.NewCanvas2d(false)
	app.cvs.Create(width, height)
	app.ir = renderer.NewImageRenderer(width, height)

	return app
}

func (app *App) render(gc *draw2dimg.GraphicContext) bool {
	app.s.Evolute()
	return app.ir.Render(gc, app.s.Generation())
}

func (app *App) abort(err error) {
	js.Global().Get("alert").Invoke(err.Error())

	os.Exit(1)
}

func (app *App) generationCount(this js.Value, _ []js.Value) interface{} {
	return app.s.GenerationCount()
}

func (app *App) cellCount(this js.Value, _ []js.Value) interface{} {
	return app.s.CellCount()
}

func (app *App) loadPatternFile(this js.Value, args []js.Value) interface{} {
	rleFile := args[0].String()
	parsedRle, err := rle.Parse(rleFile)
	if err != nil {
		app.abort(err)
	}

	app.s = simulator.NewSimulatorWithPattern(parsedRle.Pattern, parsedRle.Rule)

	return nil
}

func (app *App) setOffset(this js.Value, args []js.Value) interface{} {
	x := args[0].Float()
	y := args[1].Float()

	app.ir.SetOffsetX(app.ir.OffsetX() + x)
	app.ir.SetOffsetY(app.ir.OffsetY() + y)

	return nil
}

func (app *App) center(this js.Value, args []js.Value) interface{} {
	app.ir.SetOffsetX(0)
	app.ir.SetOffsetY(0)

	return nil
}

func (app *App) cellSideLength(this js.Value, args []js.Value) interface{} {
	newCellSideLength, err := strconv.ParseFloat(args[0].String(), 64)
	if err != nil {
		return err
	}

	app.ir.CellSideLength(newCellSideLength)

	return nil
}

func main() {
	app := NewApp()

	patternName := js.Global().Get("pattern").String()
	ruleString := js.Global().Get("rule").String()

	p, ok := pattern.Pattern[patternName]
	if !ok {
		fmt.Errorf("Pattern does not exist: %s", patternName)
	}

	app.s = simulator.NewSimulatorWithPattern(p, ruleString)

	global := js.Global()
	global.Set("generationCount", js.FuncOf(app.generationCount))
	global.Set("cellCount", js.FuncOf(app.cellCount))
	global.Set("loadPatternFile", js.FuncOf(app.loadPatternFile))
	global.Set("setOffset", js.FuncOf(app.setOffset))
	global.Set("center", js.FuncOf(app.center))
	global.Set("cellSideLength", js.FuncOf(app.cellSideLength))

	app.cvs.Start(20, app.render)
	<-done
}
