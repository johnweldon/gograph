package graph

type Vertice struct {
	Name      string
	Data      interface{}
	indegree  []*Vertice
	outdegree []*Vertice
}

func findVertice(name string, vertices []*Vertice) (bool, *Vertice) {
	for _, v := range vertices {
		if v.Name == name {
			return true, v
		}
	}
	return false, nil
}
