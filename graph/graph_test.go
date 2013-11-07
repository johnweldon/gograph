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
	//Graph: AB5, BC4, CD8, DC8, DE6, AD5, CE2, EB3, AE7
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
func TestFindShortestPath(t *testing.T) { t.Errorf("Not-implemented") }
