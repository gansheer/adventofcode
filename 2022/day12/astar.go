package astar

import (
	"math"
	"sort"
)

type GraphA struct {
	vertices []VertexA
	edges    []EdgeA
}

type VertexA struct {
	x           int
	y           int
	value       int
	heuristique int
}

type EdgeA struct {
	from VertexA
	to   VertexA
}

// want to do go 18

func voisins(g GraphA, u VertexA) []VertexA {
	voisins := make([]VertexA, 0)
	for _, v := range g.vertices {
		if containsEdge(g.edges, EdgeA{from: u, to: v}) {
			if v.x == u.x+1 && v.y == u.y {
				voisins = append(voisins, v)
			} else if v.x == u.x-1 && v.y == u.y {
				voisins = append(voisins, v)
			} else if v.x == u.x && v.y == u.y+1 {
				voisins = append(voisins, v)
			} else if v.x == u.x && v.y == u.y-1 {
				voisins = append(voisins, v)
			}
		}
	}
	return voisins
}

func contains(s []VertexA, e VertexA) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// bug : value and heuristique of vertex might be different
func containsEdge(s []EdgeA, e EdgeA) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func containsInferiorValue(s []VertexA, e VertexA) bool {
	for _, a := range s {
		if a == e && a.value < e.value {
			return true
		}
	}
	return false
}

func distance(startx int, starty int, endx int, endy int) int {
	one := math.Abs(float64((startx - endx) * (startx - endx)))
	two := math.Abs(float64((starty - endy) * (starty - endy)))
	distance := math.Sqrt(one + two)
	return int(distance)
}

func compareParHeuristique(n1, n2 VertexA) bool {
	if n1.heuristique < n2.heuristique {
		return true
	} else if n1.heuristique == n2.heuristique {
		return true
	} else {
		return false
	}
}

func cheminPlusCourt(g GraphA, end VertexA, start VertexA) int {
	// closed list
	var closedLists []VertexA
	// open list par comparateur heuristique
	var openList []VertexA = make([]VertexA, 0)
	openList = append(openList, start)
	// tant que openList n'est pas vide
	for len(openList) > 0 {
		u := openList[0]
		if u.x == end.x && u.y == end.y {
			//reconstituerChemin(u)
			return len(openList)
		} else {
			neighboors := voisins(g, u)

			//pour chaque voisin v de u dans g
			for _, v := range neighboors {
				if !(contains(closedLists, v) || containsInferiorValue(openList, v)) {
					v.value = v.value + 1
					v.heuristique = v.value + distance(v.x, v.y, end.x, end.y)
					openList = append(openList, v)
					sort.Slice(openList, func(p, q int) bool {
						return openList[p].heuristique < openList[q].heuristique
					})

				}

			}
		}
		openList = append(closedLists, u)

	}

	return -1
}
