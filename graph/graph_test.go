package graph_test

import (
	graph "."
	"testing"
)

func TestFindMissingVertice(t *testing.T) {
	g := graph.Graph{}
	if ok, v := g.FindVertice("non-existent"); ok {
		t.Errorf("Should not succeed when searching for non-existent vertice name")
	} else if v != nil {
		t.Errorf("Failed FindVertice call should return nil")
	}
}

func TestAddVertice(t *testing.T) {
	g := graph.Graph{}
	if ok, _ := g.AddVertice("A", 1); !ok {
		t.Errorf("Failed to add basic vertice")
	}
	if ok, _ := g.AddVertice("A", 1); ok {
		t.Errorf("Should not be able to add same vertice twice")
	}
	if ok, _ := g.AddVertice("A", 2); ok {
		t.Errorf("Should not be able to add vertice with same name twice")
	}
}

func TestAddEdge(t *testing.T) {
	g := graph.Graph{}
	if ok, _ := g.AddEdge("A", "B", 1, nil); !ok {
		t.Errorf("Failed to add edge")
	}
	if ok, _ := g.AddEdge("A", "B", 1, nil); !ok {
		t.Errorf("Should be able to add multiple edges with same pair")
	}
}

func TestFindAllPaths(t *testing.T) {
	g := defaultGraph()
	if g.Order() != 5 {
		t.Errorf("Incorrect Order - should be number of Vertices")
	}
	if g.Size() != 9 {
		t.Errorf("Incorrect Size - should be number of Edges")
	}
	if ok, paths := g.FindAllPaths("C", "C"); !ok {
		t.Fatalf("Failed to find paths between 'C' and 'C'")
	} else if len(paths) < 3 {
		t.Errorf("Not enough paths")
	}
}

func TestFindShortestPath(t *testing.T) {
	g := defaultGraph()

	findShortestPath(t, g, "A", "A", -1, 30)
	findShortestPath(t, g, "A", "B", 5, 30)
	findShortestPath(t, g, "A", "C", 9, 30)
	findShortestPath(t, g, "A", "D", -1, 30)
	findShortestPath(t, g, "A", "E", 11, 30)

	findShortestPath(t, g, "B", "A", -1, 30)
	findShortestPath(t, g, "B", "B", -1, 30)
	findShortestPath(t, g, "B", "C", 4, 30)
	findShortestPath(t, g, "B", "D", -1, 30)
	findShortestPath(t, g, "B", "E", 6, 30)

	findShortestPath(t, g, "C", "A", -1, 30)
	findShortestPath(t, g, "C", "B", 5, 30)
	findShortestPath(t, g, "C", "C", -1, 30)
	findShortestPath(t, g, "C", "D", -1, 30)
	findShortestPath(t, g, "C", "E", 2, 30)

	findShortestPath(t, g, "D", "A", -1, 30)
	findShortestPath(t, g, "D", "B", 9, 30)
	findShortestPath(t, g, "D", "C", 13, 30)
	findShortestPath(t, g, "D", "D", -1, 30)
	findShortestPath(t, g, "D", "E", 6, 30)

	findShortestPath(t, g, "E", "A", -1, 30)
	findShortestPath(t, g, "E", "B", 3, 30)
	findShortestPath(t, g, "E", "C", 7, 30)
	findShortestPath(t, g, "E", "D", 15, 30)
	findShortestPath(t, g, "E", "E", -1, 30)
}

func findShortestPath(t *testing.T, g *graph.Graph, from, to string, expected float64, max int) {
	if ok, path := g.FindShortestPath(from, to, max); !ok {
		if expected >= 0 {
			t.Errorf("Failed to find shortest path between %s and %s", from, to)
		}
	} else if path.Weight() != expected {
		t.Errorf("Expected distance %g between %s and %s, got %g \npath: %s", expected, from, to, path.Weight(), &path)
	}
}

func defaultGraph() *graph.Graph {
	//Graph: AB5, BC4, CD8, DC8, DE6, AD5, CE2, EB3, AE7
	/*

	digraph default {
	    A -> B [label=5];
	    B -> C [label=4];
	    C -> D [label=8];
	    D -> E [label=6];
	    A -> D [label=5];
	    C -> E [label=2];
	    E -> B [label=3];
	    A -> E [label=7];
    }

	*/
	g := graph.Graph{}
	g.AddEdge("A", "B", 5, nil)
	g.AddEdge("B", "C", 4, nil)
	g.AddEdge("C", "D", 8, nil)
	g.AddEdge("D", "C", 8, nil)
	g.AddEdge("D", "E", 6, nil)
	g.AddEdge("A", "D", 5, nil)
	g.AddEdge("C", "E", 2, nil)
	g.AddEdge("E", "B", 3, nil)
	g.AddEdge("A", "E", 7, nil)
	return &g
}
