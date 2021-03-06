package simulator

import (
	"github.com/nhoffmann/life/grid"
)

type Simulator struct {
	generationCount int
	Rule            Rule
	generation      *grid.Generation
}

func NewSimulator(ruleString string) *Simulator {
	return &Simulator{
		Rule:       ParseRule(ruleString),
		generation: grid.NewGeneration(),
	}
}

func NewSimulatorWithPattern(pattern [][]int, ruleString string) *Simulator {
	s := NewSimulator(ruleString)
	s.generation.LoadPattern(pattern)
	return s
}

func (s *Simulator) Generation() *grid.Generation {
	return s.generation
}

func (s *Simulator) GenerationCount() int {
	return s.generationCount
}

func (s *Simulator) CellCount() int {
	return s.generation.CellCount()
}

// Evolute creates the next state of the grid and then replaces the current one
func (s *Simulator) Evolute() {
	nextGeneration := grid.NewGeneration()

	for cell := range s.generation.CellMap {
		s.evoluteMooreNeighborhood(nextGeneration, cell)
	}

	s.generationCount++
	s.generation = nextGeneration
}

func (s *Simulator) evoluteMooreNeighborhood(nextGeneration *grid.Generation, cell grid.Cell) {
	s.evoluteCell(nextGeneration, cell)

	for _, neighboringCell := range cell.MooreNeighborhood() {
		if !nextGeneration.IsPopulated(neighboringCell) {
			s.evoluteCell(nextGeneration, neighboringCell)
		}
	}
}

func (s *Simulator) evoluteCell(nextGeneration *grid.Generation, cell grid.Cell) {
	if s.cellLives(cell) {
		nextGeneration.Populate(cell)
	}
}

func (s *Simulator) cellLives(cell grid.Cell) bool {
	count := s.generation.PopulatedNeighborsCount(cell)

	if s.generation.IsPopulated(cell) {
		// apply survive rules
		for _, surviveCount := range s.Rule.SurviveCounts {
			if surviveCount == count {
				return true
			}
		}
	} else {
		// apply born rules
		for _, bornCount := range s.Rule.BornCounts {
			if bornCount == count {
				return true
			}
		}

		return false
	}

	return false
}
