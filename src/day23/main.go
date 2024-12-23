package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

type Edge [2]string
type V map[string]bool
type E map[Edge]bool

func parse() (v V, e E) {
	data, _ := os.ReadFile("data/day23.txt")
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	v = make(V)
	e = make(E)

	for _, line := range lines {
		edge := strings.Split(strings.TrimSpace(string(line)), "-")
		v[edge[0]] = true
		v[edge[1]] = true
		e[Edge{edge[0], edge[1]}] = true
		e[Edge{edge[1], edge[0]}] = true
	}
	return v, e
}

func solve1(v V, e E) (score int) {
	for edge := range e {
		for w := range v {
			if e[Edge{edge[0], w}] && e[Edge{edge[1], w}] &&
				(w[0] == 't' || edge[0][0] == 't' || edge[1][0] == 't') {
				score++
			}
		}
	}
	return score / 6
}

func N(v string, e E) (neigh V) {
	neigh = make(V)
	for edge := range e {
		if edge[0] == v {
			neigh[edge[1]] = true
		} else if edge[1] == v {
			neigh[edge[0]] = true
		}
	}
	return neigh
}

func union(self V, other V) (result V) {
	result = make(V)
	for k, v := range self {
		result[k] = v
	}
	for k, v := range other {
		result[k] = v
	}
	return result
}

func intersection(self V, other V) (result V) {
	result = make(V)
	for k := range self {
		if other[k] {
			result[k] = true
		}
	}
	return result
}

func difference(self V, other V) (result V) {
	result = make(V)
	for k := range self {
		if !other[k] {
			result[k] = true
		}
	}
	return result
}

func bk(r V, p V, x V, e E) V {
	if len(p) == 0 && len(x) == 0 {
		return copy(r)
	}
	mclq := make(V)
	for v := range p {
		n := N(v, e)
		ss := union(r, V{v: true})
		mc := bk(union(r, V{v: true}), intersection(p, n), intersection(x, n), e)
		if len(mc) > len(mclq) {
			mclq = mc
		}
		p = difference(p, ss)
		x = union(x, ss)
	}
	return mclq
}

func copy(originalMap V) (newMap V) {
	newMap = make(V)
	for k, v := range originalMap {
		newMap[k] = v
	}
	return newMap
}

func solve2(v V, e E) string {
	cq := bk(make(V), copy(v), make(V), e)

	var verts []string
	for q := range cq {
		verts = append(verts, q)
	}
	sort.Slice(verts, func(i, j int) bool { return verts[i] < verts[j] })
	return strings.Join(verts, ",")
}

func main() {
	v, e := parse()
	println(len(v))
	fmt.Println("ANSWER 1: ", solve1(v, e))
	fmt.Println("ANSWER 2: ", solve2(v, e))
}
