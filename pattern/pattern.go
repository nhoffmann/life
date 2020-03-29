package pattern

var Pattern = map[string][][]bool{
	"Glider":           Glider,
	"GospersGliderGun": GospersGliderGun,
	"RPentomino":       RPentomino,
	"DieHard":          DieHard,
	"Acorn":            Acorn,
	"Chaos":            Chaos,
	"Chaos2":           Chaos2,
	"Chaos3":           Chaos3,
}

var Glider = [][]bool{
	{false, true, false},
	{false, false, true},
	{true, true, true},
}

var GospersGliderGun = [][]bool{
	{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, true, false, false, false, false, false, false, false, false, false, false, false},
	{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, true, false, true, false, false, false, false, false, false, false, false, false, false, false},
	{false, false, false, false, false, false, false, false, false, false, false, false, true, true, false, false, false, false, false, false, true, true, false, false, false, false, false, false, false, false, false, false, false, false, true, true},
	{false, false, false, false, false, false, false, false, false, false, false, true, false, false, false, true, false, false, false, false, true, true, false, false, false, false, false, false, false, false, false, false, false, false, true, true},
	{true, true, false, false, false, false, false, false, false, false, true, false, false, false, false, false, true, false, false, false, true, true, false, false, false, false, false, false, false, false, false, false, false, false, false, false},
	{true, true, false, false, false, false, false, false, false, false, true, false, false, false, true, false, true, true, false, false, false, false, true, false, true, false, false, false, false, false, false, false, false, false, false, false},
	{false, false, false, false, false, false, false, false, false, false, true, false, false, false, false, false, true, false, false, false, false, false, false, false, true, false, false, false, false, false, false, false, false, false, false, false},
	{false, false, false, false, false, false, false, false, false, false, false, true, false, false, false, true, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false},
	{false, false, false, false, false, false, false, false, false, false, false, false, true, true, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false},
}

var RPentomino = [][]bool{
	{false, true, true},
	{true, true, false},
	{false, true, false},
}

var DieHard = [][]bool{
	{false, false, false, false, false, false, true, false},
	{true, true, false, false, false, false, false, false},
	{false, true, false, false, false, true, true, true},
}

var Acorn = [][]bool{
	{false, true, false, false, false, false, false},
	{false, false, false, true, false, false, false},
	{true, true, false, false, true, true, true},
}

var Chaos = [][]bool{
	{false, false, false, false, false, false, true, false},
	{false, false, false, false, true, false, true, true},
	{false, false, false, false, true, false, true, false},
	{false, false, false, false, true, false, false, false},
	{false, false, true, false, false, false, false, false},
	{true, false, true, false, false, false, false, false},
}

var Chaos2 = [][]bool{
	{true, true, true, false, true},
	{true, false, false, false, false},
	{false, false, false, true, true},
	{false, true, true, false, true},
	{true, false, true, false, true},
}

var Chaos3 = [][]bool{
	{true, true, true, true, true, true, true, true, false, true, true, true, true, true, false, false, false, true, true, true, false, false, false, false, false, false, true, true, true, true, true, true, true, false, true, true, true, true, true},
}
