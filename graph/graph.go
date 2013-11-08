package graph

import (
	"math"
)

type Graph struct {
	Directed bool
	vertices []*Vertice
	edges    []Edge
}

func (g *Graph) Order() int { return len(g.vertices) }
func (g *Graph) Size() int  { return len(g.edges) }

func (g *Graph) FindAllPaths(from, to string) (ok bool, paths []Path) {
	paths = []Path{}
	return
}

func (g *Graph) FindShortestPath(from, to string, max int) (ok bool, path Path) {
	ok = false
	path = Path{}

	//
	// Dijkstra
	//

	distances := map[string]float64{}
	visited := map[string]bool{}
	for _, v := range g.vertices {
		distances[v.Name] = math.SmallestNonzeroFloat64
		visited[v.Name] = false
	}
	distances[from] = 0

	var cur, closest *Vertice
	if ok, cur = g.FindVertice(from); !ok {
		return
	}

	for {
		me := cur.Name
		neighbors := g.getNeighbors(cur)

		smallest := ""
		smallestweight := math.MaxFloat64

		for neighbor, weight := range neighbors {
			nname := neighbor.Name
			nweight := distances[me] + weight
			if visited[nname] {
				continue
			}

			if nweight < smallestweight {
				smallest, smallestweight = nname, nweight
			}

			if distances[nname] == math.SmallestNonzeroFloat64 {
				distances[nname] = nweight
			} else if distances[nname] > nweight {
				distances[nname] = nweight
			}
		}

		// mark current visited and remove from unvisited set
		visited[cur.Name] = true
		// if current = to done
		if cur.Name == to && len(path.Edges) > 0 {
			ok = true
			return
		}

		// mark current = smallest neighbor; goto loop
		if ok, closest = g.FindVertice(smallest); !ok {
			return
		}
		path.Push(Edge{Head: cur, Tail: closest, Weight: smallestweight - distances[cur.Name], Data: nil})

		cur = closest

		// short circuit
		if max > 0 && len(path.Edges) >= max {
			ok = false
			return
		}
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
