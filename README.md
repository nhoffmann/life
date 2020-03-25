# Life

A Conway's Game Of Life simulator written in Go. It includes rendering on the command line and on a webpage via web assembly.


You can see it in action here: http://nhoffmann.github.com/life/cmd/wasm

## Roadmap

* [ ] **More configuration options.**
* [x] ~~**Configure arbitrary rules.** Currently it only supports John Conway's rules, where a cell is born when it has 3 neighbors and stays alive when it has 2 or 3 neighbors (i.e. B3/S23). It would be great to be able to support other rules such as [Highlife](https://en.wikipedia.org/wiki/Highlife_(cellular_automaton)) or whatever comes to mind.~~
* [ ] **Add [RLE](https://www.conwaylife.com/wiki/Run_Length_Encoded) support.** Currently there a re some patterns provided, but there is currently no way to define patterns yourself. Support for loading RLE would alleviate this limitation.
* [ ] **Infinite board.** Currently the board wraps around at the edges.
* [ ] **3D Life.** A three dimensional version of Life :)
