package main

import (
	"fmt"
	"os"
	"syscall/js"

	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/markfarnan/go-canvas/canvas"
	"github.com/nhoffmann/life/pattern"
	"github.com/nhoffmann/life/simulator"
)

const CELL_WIDTH = 10

var done chan struct{}

var s *simulator.Simulator

var cvs *canvas.Canvas2d
var width float64
var height float64

var alert = js.Global().Get("alert")

func main() {
	width := int(js.Global().Get("innerWidth").Float())
	height := int(js.Global().Get("innerHeight").Float())

	patternName := js.Global().Get("pattern").String()
	ruleString := js.Global().Get("rule").String()

	s = simulator.New(width/CELL_WIDTH, height/CELL_WIDTH, ruleString)

	p, ok := pattern.Pattern[patternName]
	if !ok {
		abort(fmt.Errorf("Pattern does not exist: %s", patternName))
	}

	err := s.LoadPattern(p)
	if err != nil {
		abort(err)
	}

	cvs, _ = canvas.NewCanvas2d(false)
	cvs.Create(width, height)

	cvs.Start(30, Render)
	<-done
}

func Render(gc *draw2dimg.GraphicContext) bool {
	s.Evolute()

	return s.Render(gc)
}

func abort(err error) {
	alert.Invoke(err.Error())

	os.Exit(1)
}
