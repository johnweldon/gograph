package graph

import (
	"fmt"
	"math"
)

type Vertice struct {
	Name      string
	Data      interface{}
	indegree  []*Vertice
	outdegree []*Vertice
}

type Edge struct {
	Head   *Vertice
	Tail   *Vertice
	Weight float64
	Data   interface{}
}

type Graph struct {
	Directed bool
	vertices []*Vertice
	edges    []Edge
}

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

func (e *Edge) String() string { return fmt.Sprintf("%s%s%g", e.Head.Name, e.Tail.Name, e.Weight) }
func (p *Path) String() string {
	ret := "["
	for _, e := range p.Edges {
		ret = fmt.Sprintf("%s, %s", ret, &e)
	}
	return ret + "]"
}

func (g *Graph) Order() int { return len(g.vertices) }
func (g *Graph) Size() int  { return len(g.edges) }

func (g *Graph) FindAllPaths(from, to string) (ok bool, paths []Path) {
	paths = []Path{}
	return
}

func (g *Graph) FindShortestPath(from, to string) (ok bool, path Path) {
	ok = false
	path = Path{}

	//
	// Dijkstra
	//
	var cur *Vertice

	// initialize distance, visited, current node
	distances := map[string]float64{}
	visited := map[string]bool{}
	for _, v := range g.vertices {
		distances[v.Name] = math.SmallestNonzeroFloat64
		visited[v.Name] = false
	}
	distances[from] = 0

	// set current = from
	if ok, cur = g.FindVertice(from); !ok {
		return
	}

	// begin loop
	for {
		// consider all neighbors and calculate tentative distance
		me := cur.Name
		neighbors := g.getNeighbors(cur)
		smallest := me
		smallestweight := math.MaxFloat64
		for neighbor, weight := range neighbors {
			neighborname := neighbor.Name
			neighborweight := distances[me] + weight
			if neighborweight < smallestweight {
				smallest, smallestweight = neighborname, neighborweight
			}
			if visited[neighborname] {
				continue
			}
			if distances[neighborname] == math.SmallestNonzeroFloat64 {
				distances[neighborname] = neighborweight
				path.Push(Edge{Head: cur, Tail: neighbor, Weight: weight, Data: nil})
			} else if distances[neighborname] > neighborweight {
				distances[neighborname] = neighborweight
				path.Pop()
				path.Push(Edge{Head: cur, Tail: neighbor, Weight: weight, Data: nil})
			}
		}
		// mark current visited and remove from unvisited set
		visited[cur.Name] = true
		// if current = to done
		if cur.Name == to {
			ok = true
			return
		}
		// mark current = smallest neighbor; goto loop
		if ok, cur = g.FindVertice(smallest); !ok {
			return
		}

	}

	return
}

func (g *Graph) getNeighbors(v *Vertice) map[*Vertice]float64 {
	neighbors := map[*Vertice]float64{}
	for _, e := range g.edges {
		if e.Head != v {
			continue
		}
		neighbors[e.Tail] = e.Weight
	}
	return neighbors
}

func (g *Graph) AddVertice(name string, data interface{}) (bool, *Vertice) {
	if ok, existing := g.FindVertice(name); ok {
		return false, existing
	}
	v := &Vertice{Name: name, Data: data}
	g.vertices = append(g.vertices, v)
	return true, v
}

func (g *Graph) AddEdge(from, to string, weight float64, data interface{}) (bool, *Edge) {
	var ok bool
	var head, tail *Vertice
	if ok, head = g.FindVertice(from); !ok {
		_, head = g.AddVertice(from, nil)
	}
	if ok, tail = g.FindVertice(to); !ok {
		_, tail = g.AddVertice(to, nil)
	}
	if ok, _ = findVertice(tail.Name, head.outdegree); !ok {
		head.outdegree = append(head.outdegree, tail)
	}
	if ok, _ = findVertice(head.Name, tail.indegree); !ok {
		tail.indegree = append(tail.indegree, head)
	}
	edge := Edge{head, tail, weight, data}
	g.edges = append(g.edges, edge)
	return true, &edge
}

func (g *Graph) FindVertice(name string) (bool, *Vertice) {
	return findVertice(name, g.vertices)
}

func findVertice(name string, vertices []*Vertice) (bool, *Vertice) {
	for _, v := range vertices {
		if v.Name == name {
			return true, v
		}
	}
	return false, nil
}
