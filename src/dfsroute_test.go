package mylib

import (
	"testing"
)

func TestDFSRouteSimple(t *testing.T) {
	nodes := []Node{
		{To: []int{1, 2}},
		{To: []int{0, 3}},
		{To: []int{0, 3}},
		{To: []int{1, 2}},
	}
	route := DFSRoute(nodes, 0, 3)
	if len(route) == 0 {
		t.Fatal("expected a route from 0 to 3")
	}
	if route[0] != 0 || route[len(route)-1] != 3 {
		t.Errorf("route should start at 0 and end at 3, got %v", route)
	}
}

func TestDFSRouteDirect(t *testing.T) {
	nodes := []Node{
		{To: []int{1}},
		{To: []int{0}},
	}
	route := DFSRoute(nodes, 0, 1)
	if len(route) != 2 || route[0] != 0 || route[1] != 1 {
		t.Errorf("expected [0 1], got %v", route)
	}
}

func TestDFSRouteSameNode(t *testing.T) {
	nodes := []Node{
		{To: []int{1}},
		{To: []int{0}},
	}
	route := DFSRoute(nodes, 0, 0)
	if len(route) != 1 || route[0] != 0 {
		t.Errorf("expected [0], got %v", route)
	}
}

func TestDFSRouteUnreachable(t *testing.T) {
	nodes := []Node{
		{To: []int{1}},
		{To: []int{0}},
		{To: []int{}},
	}
	route := DFSRoute(nodes, 0, 2)
	if route != nil {
		t.Errorf("expected nil for unreachable, got %v", route)
	}
}

func TestDFSRouteMultiplePaths(t *testing.T) {
	nodes := []Node{
		{To: []int{1, 2}},
		{To: []int{3}},
		{To: []int{3}},
		{To: []int{}},
	}
	route := DFSRoute(nodes, 0, 3)
	if len(route) == 0 {
		t.Fatal("expected a route")
	}
	if route[0] != 0 || route[len(route)-1] != 3 {
		t.Error("route should start at 0 and end at 3")
	}
}
