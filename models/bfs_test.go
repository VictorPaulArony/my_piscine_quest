package models

import (
	"testing"
)

// TestBfs_SinglePath tests the Bfs function with a simple single path.
func TestBfs_SinglePath(t *testing.T) {
	// Create a colony with a single path
	colony := &Colony{
		NoOfAnts: 3,
		Rooms: []*Room{
			{Name: "start", RoomAndConnectedLinks: map[string][]string{"start": {"room1"}}},
			{Name: "room1", RoomAndConnectedLinks: map[string][]string{"room1": {"end"}}},
			{Name: "end", RoomAndConnectedLinks: map[string][]string{}},
		},
		StartRoom: Room{Name: "start"},
		EndRoom:   Room{Name: "end"},
	}

	// Run BFS
	paths := colony.Bfs()

	// Validate the result
	expectedPaths := [][]string{{"start", "room1", "end"}}
	if len(paths) != len(expectedPaths) {
		t.Fatalf("Expected %d paths, got %d", len(expectedPaths), len(paths))
	}
	for i, path := range paths {
		if len(path) != len(expectedPaths[i]) {
			t.Errorf("Expected path length %d, got %d", len(expectedPaths[i]), len(path))
		}
		for j, room := range path {
			if room != expectedPaths[i][j] {
				t.Errorf("Expected room %s, got %s", expectedPaths[i][j], room)
			}
		}
	}
}

// TestBfs_MultiplePaths tests the Bfs function with multiple valid paths.
func TestBfs_MultiplePaths(t *testing.T) {
	// Create a colony with multiple paths
	colony := &Colony{
		NoOfAnts: 3,
		Rooms: []*Room{
			{Name: "start", RoomAndConnectedLinks: map[string][]string{"start": {"room1", "room2"}}},
			{Name: "room1", RoomAndConnectedLinks: map[string][]string{"room1": {"end"}}},
			{Name: "room2", RoomAndConnectedLinks: map[string][]string{"room2": {"end"}}},
			{Name: "end", RoomAndConnectedLinks: map[string][]string{}},
		},
		StartRoom: Room{Name: "start"},
		EndRoom:   Room{Name: "end"},
	}

	// Run BFS
	paths := colony.Bfs()

	// Validate the result
	expectedPaths := [][]string{
		{"start", "room1", "end"},
		{"start", "room2", "end"},
	}
	if len(paths) != len(expectedPaths) {
		t.Fatalf("Expected %d paths, got %d", len(expectedPaths), len(paths))
	}
	for i, path := range paths {
		if len(path) != len(expectedPaths[i]) {
			t.Errorf("Expected path length %d, got %d", len(expectedPaths[i]), len(path))
		}
		for j, room := range path {
			if room != expectedPaths[i][j] {
				t.Errorf("Expected room %s, got %s", expectedPaths[i][j], room)
			}
		}
	}
}

// TestBfs_NoPath tests the Bfs function when no valid path exists.
func TestBfs_NoPath(t *testing.T) {
	// Create a colony with no valid path
	colony := &Colony{
		NoOfAnts: 3,
		Rooms: []*Room{
			{Name: "start", RoomAndConnectedLinks: map[string][]string{"start": {"room1"}}},
			{Name: "room1", RoomAndConnectedLinks: map[string][]string{"room1": {"room2"}}},
			{Name: "room2", RoomAndConnectedLinks: map[string][]string{"room2": {"room3"}}},
			{Name: "room3", RoomAndConnectedLinks: map[string][]string{}},
			{Name: "end", RoomAndConnectedLinks: map[string][]string{}},
		},
		StartRoom: Room{Name: "start"},
		EndRoom:   Room{Name: "end"},
	}

	// Run BFS
	paths := colony.Bfs()

	// Validate the result
	if len(paths) != 0 {
		t.Errorf("Expected no paths, got %d paths", len(paths))
	}
}

// TestBfs_ComplexPath tests the Bfs function with a more complex graph.
func TestBfs_ComplexPath(t *testing.T) {
	// Create a colony with a complex graph
	colony := &Colony{
		NoOfAnts: 3,
		Rooms: []*Room{
			{Name: "start", RoomAndConnectedLinks: map[string][]string{"start": {"room1", "room2"}}},
			{Name: "room1", RoomAndConnectedLinks: map[string][]string{"room1": {"room3", "end"}}},
			{Name: "room2", RoomAndConnectedLinks: map[string][]string{"room2": {"room3"}}},
			{Name: "room3", RoomAndConnectedLinks: map[string][]string{"room3": {"end"}}},
			{Name: "end", RoomAndConnectedLinks: map[string][]string{}},
		},
		StartRoom: Room{Name: "start"},
		EndRoom:   Room{Name: "end"},
	}

	// Run BFS
	paths := colony.Bfs()

	// Validate the result
	expectedPaths := [][]string{
		{"start", "room1", "end"},
		{"start", "room1", "room3", "end"},
		{"start", "room2", "room3", "end"},
	}
	if len(paths) != len(expectedPaths) {
		t.Fatalf("Expected %d paths, got %d", len(expectedPaths), len(paths))
	}
	for i, path := range paths {
		if len(path) != len(expectedPaths[i]) {
			t.Errorf("Expected path length %d, got %d", len(expectedPaths[i]), len(path))
		}
		for j, room := range path {
			if room != expectedPaths[i][j] {
				t.Errorf("Expected room %s, got %s", expectedPaths[i][j], room)
			}
		}
	}
}
