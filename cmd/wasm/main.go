package main

import (
	"fmt"
	"os"
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

func main() {
	width := int(js.Global().Get("innerWidth").Float())
	height := int(js.Global().Get("innerHeight").Float())

	patternName := js.Global().Get("pattern").String()
	ruleString := js.Global().Get("rule").String()

	p, ok := pattern.Pattern[patternName]
	if !ok {
		abort(fmt.Errorf("Pattern does not exist: %s", patternName))
	}

	s = simulator.NewSimulatorWithPattern(p, ruleString)

	cvs, _ = canvas.NewCanvas2d(false)
	cvs.Create(width, height)
	s.Renderer = renderer.NewImageRenderer(width, height)

	cvs.Start(20, Render)
	<-done
}

func Render(gc *draw2dimg.GraphicContext) bool {
	s.Evolute()
	return s.Renderer.Render(gc, s.Generation())
}

func abort(err error) {
	alert.Invoke(err.Error())

	os.Exit(1)
}
