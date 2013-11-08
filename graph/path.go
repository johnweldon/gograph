package graph

import (
	"strings"
)

type Path struct {
	Edges []Edge
}

func (p *Path) Pop() (ok bool, e Edge) {
	ok = false
	if len(p.Edges) == 0 {
		return
	}
	e, p.Edges = p.Edges[len(p.Edges)-1], p.Edges[:len(p.Edges)-1]
	ok = true
	return
}

func (p *Path) Push(e Edge) {
	p.Edges = append(p.Edges, e)
}

func (p *Path) Weight() float64 {
	var f float64
	f = 0
	for _, e := range p.Edges {
		f = f + e.Weight
	}
	return f
}

func (p Path) String() string {
	s := []string{}
	for _, e := range p.Edges {
		s = append(s, e.String())
	}
	return "[" + strings.Join(s, ", ") + "]"
}
