# Life

A Conway's Game Of Life simulator written in Go. It features an infinite board and renders in the browser via web assembly.
It comes with some preloaded patterns to investigate, but also let's you upload patterns stored in the [RLE file format](https://www.conwaylife.com/wiki/Run_Length_Encoded) for life patterns.
It supports all possible rules of Life like cellular automata.


You can see it in action here: https://zentralmaschine.net/life/public

## Roadmap

* [x] ~~**Configure arbitrary rules.** Currently it only supports John Conway's rules, where a cell is born when it has 3 neighbors and stays alive when it has 2 or 3 neighbors (i.e. B3/S23). It would be great to be able to support other rules such as [Highlife](https://en.wikipedia.org/wiki/Highlife_(cellular_automaton)) or whatever comes to mind.~~
* [x] ~~**Add [RLE](https://www.conwaylife.com/wiki/Run_Length_Encoded) support.** Currently there a re some patterns provided, but there is currently no way to define patterns yourself. Support for loading RLE would alleviate this limitation.~~
* [x] ~~**Add web interface for loading RLE files.** RLE file support is present, but currently not offered in the web interface.~~
* [x] ~~**Infinite board.** Currently the board wraps around at the edges.~~
* [x] ~~**Add ability to change board configuration in real time.** Cell size, center of the board, line width etc. should all be configurable.~~
* [ ] **Javascript native renderer** Currently the game state gets rendered via [draw2d](https://github.com/llgcode/draw2d) in go. The resulting image is then rendered in a canvas in the browser through [go-canvas](https://github.com/markfarnan/go-canvas). While this keeps the rendering implementation in Go code, it is not efficient and it would be preferable to do the rendering in a more native environment directly.
* [ ] **3D Life.** A three dimensional version of Life :)

## Development

Clone the repo and get the dependencies.

In order to compile the application run:

    $ GOOS=js GOARCH=wasm go build -o ./public/main.wasm ./main.go

The public folder contains a [caddy](https://caddyserver.com/) configuration so you can run the webserver like so:

    $ cd public && caddy
