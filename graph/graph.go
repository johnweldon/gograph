package graph

import (
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

func (g *Graph) Order() int { return len(g.vertices) }
func (g *Graph) Size() int  { return len(g.edges) }

func (g *Graph) FindAllPaths(from, to string) (ok bool, paths []Path) {
	paths = []Path{}
	return
}

func (g *Graph) FindShortestPath(from, to string) (ok bool, path Path) {
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
		for _, e := range g.edges {
			if e.Head != cur {
				continue
			}
			me, neighbor := cur.Name, e.Tail.Name
			relaxed := distances[me] + e.Weight
			if visited[neighbor] {
				continue
			}
			if distances[neighbor] == math.SmallestNonzeroFloat64 {
				distances[neighbor] = relaxed
			} else if distances[neighbor] > relaxed {
				distances[neighbor] = relaxed
			}
		}
		// mark current visited and remove from unvisited set
		visited[cur.Name] = true
		// if current = to done
		// mark current = smallest neighbor; goto loop
	}

	return
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
