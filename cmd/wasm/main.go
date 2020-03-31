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
	"github.com/nhoffmann/life/simulator"
)

var done chan struct{}

var s *simulator.Simulator

var cvs *canvas.Canvas2d
var ir renderer.ImageRenderer
var width float64
var height float64

var alert = js.Global().Get("alert")

type Dom struct {
	document js.Value
}

func NewDom() Dom {
	return Dom{
		document: js.Global().Get("document"),
	}
}

func (d *Dom) getElementById(id string) js.Value {
	return d.document.Call("getElementById", id)
}

func main() {
	width := int(js.Global().Get("innerWidth").Float())
	height := int(js.Global().Get("innerHeight").Float())

	patternName := js.Global().Get("pattern").String()
	ruleString := js.Global().Get("rule").String()

	dom := NewDom()

	cellSideLengthCb := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		newCellSideLength, err := strconv.ParseFloat(args[0].Get("target").Get("value").String(), 64)
		if err != nil {
			return err
		}
		ir.CellSideLength(newCellSideLength)
		return nil
	})

	rangeSlider := dom.getElementById("cellSideLength-range")
	rangeSlider.Call("addEventListener", "change", cellSideLengthCb)

	moveUp := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		ir.SetOffsetY(ir.OffsetY() + 20)
		return nil
	})

	moveDown := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		ir.SetOffsetY(ir.OffsetY() - 20)
		return nil
	})

	moveLeft := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		ir.SetOffsetX(ir.OffsetX() + 20)
		return nil
	})

	moveRight := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		ir.SetOffsetX(ir.OffsetX() - 20)
		return nil
	})

	moveCenter := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		ir.SetOffsetX(0)
		ir.SetOffsetY(0)
		return nil
	})

	upButton := dom.getElementById("up")
	upButton.Call("addEventListener", "click", moveUp)

	downButton := dom.getElementById("down")
	downButton.Call("addEventListener", "click", moveDown)

	leftButton := dom.getElementById("left")
	leftButton.Call("addEventListener", "click", moveLeft)

	rightButton := dom.getElementById("right")
	rightButton.Call("addEventListener", "click", moveRight)

	centerButton := dom.getElementById("center")
	centerButton.Call("addEventListener", "click", moveCenter)

	p, ok := pattern.Pattern[patternName]
	if !ok {
		abort(fmt.Errorf("Pattern does not exist: %s", patternName))
	}

	s = simulator.NewSimulatorWithPattern(p, ruleString)

	js.Global().Set("generationCount", js.FuncOf(generationCount))
	js.Global().Set("cellCount", js.FuncOf(cellCount))

	cvs, _ = canvas.NewCanvas2d(false)
	cvs.Create(width, height)
	ir = renderer.NewImageRenderer(width, height)

	cvs.Start(20, Render)
	<-done
}

func Render(gc *draw2dimg.GraphicContext) bool {
	s.Evolute()
	return ir.Render(gc, s.Generation())
}

func abort(err error) {
	alert.Invoke(err.Error())

	os.Exit(1)
}

func generationCount(this js.Value, _ []js.Value) interface{} {
	return s.GenerationCount()
}

func cellCount(this js.Value, _ []js.Value) interface{} {
	return s.CellCount()
}
