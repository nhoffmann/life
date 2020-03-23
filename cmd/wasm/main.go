package main

import (
	"syscall/js"

	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/markfarnan/go-canvas/canvas"
	"github.com/nhoffmann/life/pattern"
	"github.com/nhoffmann/life/simulator"
)

var done chan struct{}

var s *simulator.Simulator

var cvs *canvas.Canvas2d
var width float64
var height float64

func main() {
	width := int(js.Global().Get("innerWidth").Float())
	height := int(js.Global().Get("innerHeight").Float())

	s = simulator.New(width/10, height/10)

	s.LoadPattern(pattern.Acorn)

	cvs, _ = canvas.NewCanvas2d(false)
	cvs.Create(width, height)

	cvs.Start(30, Render)
	<-done
}

func Render(gc *draw2dimg.GraphicContext) bool {
	s.Evolute()

	return s.Render(gc)
}
