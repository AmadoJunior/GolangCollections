package graph

type Vertex struct {
	key      int
	adjacent []*Vertex
}

type Graph struct {
	vertices []*Vertex
}

func NewGraph() *Graph {
	return &Graph{vertices: make([]*Vertex, 0)}
}

func (g *Graph) AddVertex(k int) {
	if !g.contains(g.vertices, k) {
		g.vertices = append(g.vertices, &Vertex{key: k})
	}
}

func (g *Graph) AddEdge(from int, to int) {
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)

	if fromVertex != nil && toVertex != nil {
		if !g.contains(fromVertex.adjacent, toVertex.key) {
			fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
		}
	}
}

func (g *Graph) getVertex(k int) *Vertex {
	for i, v := range g.vertices {
		if k == v.key {
			return g.vertices[i]
		}
	}
	return nil
}

func (g *Graph) contains(s []*Vertex, k int) bool {
	for _, v := range s {
		if k == v.key {
			return true
		}
	}
	return false
}

func (g *Graph) DFS(v *Vertex, slice *[]int, seen map[int]bool) {
	*slice = append(*slice, v.key)
	seen[v.key] = true
	for _, v := range v.adjacent {
		if !seen[v.key] {
			g.DFS(v, slice, seen)
		}
	}
}

func (g *Graph) BFS(v *Vertex, slice *[]int, seen map[int]bool) {
	// Create Queue
	queue := []*Vertex{v}

	//While Queue Length > 0
	for len(queue) > 0 {
		//Pop Queue
		v = queue[0]
		queue = queue[1:]
		if !seen[v.key] {
			//Append Value To Result
			*slice = append(*slice, v.key)
			//Append To Queue
			queue = append(queue, v.adjacent...)
			//Update Seen
			seen[v.key] = true
		}
	}
}
