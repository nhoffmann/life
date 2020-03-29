package renderer

import (
	"image/color"

	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/llgcode/draw2d/draw2dkit"
	"github.com/nhoffmann/life/grid"
)

type ImageRenderer struct {
	cellSideLength float64
	imageWidth     float64
	imageHeight    float64
	lineWidth      float64

	rowCount          int
	columnCount       int
	centerRowIndex    int
	centerColumnIndex int

	cellColor  color.RGBA
	boardColor color.RGBA
}

func NewImageRenderer(width, height int) ImageRenderer {
	ir := ImageRenderer{}

	ir.cellSideLength = 10
	ir.lineWidth = 1
	ir.imageWidth = float64(width)
	ir.imageHeight = float64(height)

	ir.rowCount = int(ir.imageWidth / ir.cellSideLength)
	ir.columnCount = int(ir.imageHeight / ir.cellSideLength)
	ir.centerRowIndex = ir.rowCount / 2
	ir.centerColumnIndex = ir.columnCount / 2

	ir.cellColor = color.RGBA{0x00, 0x00, 0x00, 0xFF}
	ir.boardColor = color.RGBA{0xff, 0xff, 0xff, 0xff}

	return ir
}

func (ir *ImageRenderer) Render(gc *draw2dimg.GraphicContext, g *grid.Generation) bool {
	gc.SetFillColor(ir.boardColor)
	gc.Clear()

	gc.SetLineWidth(ir.lineWidth)

	gc.SetStrokeColor(ir.boardColor)
	gc.SetFillColor(ir.cellColor)
	gc.BeginPath()
	for cell := range g.CellMap {
		if ir.cellInViewport(cell) {
			ir.drawCell(gc, cell)
		}
	}
	gc.FillStroke()
	gc.Close()

	return true
}

func (ir *ImageRenderer) cellInViewport(cell grid.Cell) bool {
	x, y := ir.coordinates(cell)

	if x < 0 || y < 0 || x > ir.imageWidth || y > ir.imageHeight {
		return false
	}

	return true
}

func (ir *ImageRenderer) drawCell(gc *draw2dimg.GraphicContext, cell grid.Cell) {
	startX, startY := ir.coordinates(cell)
	endX := startX + ir.cellSideLength
	endY := startY + ir.cellSideLength

	draw2dkit.Rectangle(gc, startX, startY, endX, endY)
}

func (ir *ImageRenderer) coordinates(cell grid.Cell) (x, y float64) {
	renderRow := ir.centerRowIndex + cell.X
	renderColumn := ir.centerColumnIndex + cell.Y

	return float64(renderRow) * ir.cellSideLength, float64(renderColumn) * ir.cellSideLength
}