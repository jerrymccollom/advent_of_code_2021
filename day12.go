package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func readLinesFromFile(fname string) (result []string, err error) {
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(b), "\n")

	// filter out empty lines
	result = make([]string, 0, len(lines))
	for _, l := range lines {
		if len(l) == 0 {
			continue
		}
		result = append(result, l)
	}
	return result, nil
}

type Vertex struct {
	Key      string
	Vertices map[string]*Vertex
}

type Graph struct {
	Vertices   map[string]*Vertex
	Start, End *Vertex
}

func newGraph() *Graph {
	return &Graph{
		Vertices: map[string]*Vertex{},
	}
}

func (g *Graph) AddVertex(key string) {
	v := g.Vertices[key]
	if v == nil {
		v := &Vertex{
			Key:      key,
			Vertices: map[string]*Vertex{},
		}
		g.Vertices[key] = v
	}
}

// The AddEdge method adds an edge between two vertices in the graph
func (g *Graph) AddEdge(k1, k2 string) {
	v1 := g.Vertices[k1]
	v2 := g.Vertices[k2]
	v1.Vertices[v2.Key] = v2
	v2.Vertices[v1.Key] = v1
	g.Vertices[v1.Key] = v1
	g.Vertices[v2.Key] = v2
}

var paths []string

func (g *Graph) isSmall(v *Vertex) bool {
	return strings.ToLower(v.Key) == v.Key && v != g.Start
}

func (g *Graph) findPaths(start, end *Vertex, path string, visited map[string]int, lastSmall *Vertex) {
	if g.isSmall(start) {
		if start == lastSmall {
			if visited[start.Key] > 1 {
				return
			}
		} else if visited[start.Key] > 0 {
			return
		}

	}

	visited[start.Key] += 1
	for _, edge := range start.Vertices {
		if edge.Key == g.End.Key {
			paths = append(paths, path+",end")
		} else if edge != g.Start {
			if lastSmall == nil && g.isSmall(start) {
				g.findPaths(edge, end, path+","+edge.Key, visited, start)
			}
			g.findPaths(edge, end, path+","+edge.Key, visited, lastSmall)
		}
	}
	visited[start.Key] -= 1
}

func createGraph(data []string) *Graph {
	g := newGraph()
	for _, d := range data {
		sf := strings.Split(d, "-")
		g.AddVertex(sf[0])
		g.AddVertex(sf[1])
		g.AddEdge(sf[0], sf[1])
	}
	g.Start = g.Vertices["start"]
	g.End = g.Vertices["end"]
	return g
}

const DEBUG = false

func day12part1(data []string) {
	g := createGraph(data)
	g.findPaths(g.Start, g.End, "start", make(map[string]int), g.Start)
	sort.Strings(paths)
	if DEBUG {
		for _, p := range paths {
			fmt.Println(p)
		}
	}
	fmt.Println(len(paths))
}

func unique(slice []string) (u []string) {
	keys := make(map[string]bool)
	u = []string{}
	for _, s := range slice {
		if _, v := keys[s]; !v {
			keys[s] = true
			u = append(u, s)
		}
	}
	return u
}

func day12part2(data []string) {
	g := createGraph(data)
	paths = []string{}
	g.findPaths(g.Start, g.End, "start", make(map[string]int), nil)
	paths = unique(paths)
	if DEBUG {
		for _, p := range paths {
			fmt.Println(p)
		}
	}
	sort.Strings(paths)
	fmt.Println(len(paths))
}

func main() {
	data, _ := readLinesFromFile("day12.txt")
	day12part1(data)
	fmt.Println()
	day12part2(data)
}
