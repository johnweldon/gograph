package graph

import (
	"fmt"
)

type Edge struct {
	Head   *Vertice
	Tail   *Vertice
	Weight float64
	Data   interface{}
}

func (e *Edge) String() string { return fmt.Sprintf("%s%s%g", e.Head.Name, e.Tail.Name, e.Weight) }
