package utils

import (
	"os"
	"testing"
)

// TestFileReader_ValidInput tests the FileReader function with a valid input file.
func TestFileReader_ValidInput(t *testing.T) {
	// Create a temporary valid input file
	content := `3
##start
room1 1 1
##end
room2 2 2
room3 3 3
room1-room2
room2-room3
`
	tmpFile, err := os.CreateTemp("", "valid_input_*.txt")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.WriteString(content); err != nil {
		t.Fatalf("Failed to write to temporary file: %v", err)
	}
	tmpFile.Close()

	// Test FileReader
	colony := FileReader(tmpFile.Name())
	if colony == nil {
		t.Fatal("Expected a valid colony, got nil")
	}

	// Validate the number of ants
	if colony.NoOfAnts != 3 {
		t.Errorf("Expected 3 ants, got %d", colony.NoOfAnts)
	}

	// Validate the start room
	if colony.StartRoom.Name != "room1" {
		t.Errorf("Expected start room 'room1', got '%s'", colony.StartRoom.Name)
	}

	// Validate the end room
	if colony.EndRoom.Name != "room2" {
		t.Errorf("Expected end room 'room2', got '%s'", colony.EndRoom.Name)
	}

	// Validate the number of rooms
	if len(colony.Rooms) != 3 {
		t.Errorf("Expected 3 rooms, got %d", len(colony.Rooms))
	}

	// Validate the links
	expectedLinks := [][]string{{"room1", "room2"}, {"room2", "room3"}}
	for i, link := range colony.Rooms[0].Link {
		if link[0] != expectedLinks[i][0] || link[1] != expectedLinks[i][1] {
			t.Errorf("Expected link %v, got %v", expectedLinks[i], link)
		}
	}
}

// TestFileReader_InvalidAnts tests the FileReader function with an invalid number of ants.
func TestFileReader_InvalidAnts(t *testing.T) {
	// Create a temporary file with invalid ants
	content := `0
##start
room1 1 1
##end
room2 2 2
room1-room2
`
	tmpFile, err := os.CreateTemp("", "invalid_ants_*.txt")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.WriteString(content); err != nil {
		t.Fatalf("Failed to write to temporary file: %v", err)
	}
	tmpFile.Close()

	// Test FileReader
	colony := FileReader(tmpFile.Name())
	if colony != nil {
		t.Error("Expected nil colony due to invalid ants, got a valid colony")
	}
}

// TestFileReader_MissingStartRoom tests the FileReader function with a missing start room.
func TestFileReader_MissingStartRoom(t *testing.T) {
	// Create a temporary file with no start room
	content := `3
##end
room2 2 2
room1-room2
`
	tmpFile, err := os.CreateTemp("", "missing_start_*.txt")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.WriteString(content); err != nil {
		t.Fatalf("Failed to write to temporary file: %v", err)
	}
	tmpFile.Close()

	// Test FileReader
	colony := FileReader(tmpFile.Name())
	if colony != nil {
		t.Error("Expected nil colony due to missing start room, got a valid colony")
	}
}

// TestFileReader_MissingEndRoom tests the FileReader function with a missing end room.
func TestFileReader_MissingEndRoom(t *testing.T) {
	// Create a temporary file with no end room
	content := `3
##start
room1 1 1
room1-room2
`
	tmpFile, err := os.CreateTemp("", "missing_end_*.txt")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.WriteString(content); err != nil {
		t.Fatalf("Failed to write to temporary file: %v", err)
	}
	tmpFile.Close()

	// Test FileReader
	colony := FileReader(tmpFile.Name())
	if colony != nil {
		t.Error("Expected nil colony due to missing end room, got a valid colony")
	}
}

// TestFileReader_DuplicateRooms tests the FileReader function with duplicate rooms.
func TestFileReader_DuplicateRooms(t *testing.T) {
	// Create a temporary file with duplicate rooms
	content := `3
##start
room1 1 1
##end
room2 2 2
room1 1 1
room1-room2
`
	tmpFile, err := os.CreateTemp("", "duplicate_rooms_*.txt")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.WriteString(content); err != nil {
		t.Fatalf("Failed to write to temporary file: %v", err)
	}
	tmpFile.Close()

	// Test FileReader
	colony := FileReader(tmpFile.Name())
	if colony != nil {
		t.Error("Expected nil colony due to duplicate rooms, got a valid colony")
	}
}

// TestFileReader_InvalidLinks tests the FileReader function with invalid links.
func TestFileReader_InvalidLinks(t *testing.T) {
	// Create a temporary file with invalid links
	content := `3
##start
room1 1 1
##end
room2 2 2
room1-room2-room3
`
	tmpFile, err := os.CreateTemp("", "invalid_links_*.txt")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.WriteString(content); err != nil {
		t.Fatalf("Failed to write to temporary file: %v", err)
	}
	tmpFile.Close()

	// Test FileReader
	colony := FileReader(tmpFile.Name())
	if colony != nil {
		t.Error("Expected nil colony due to invalid links, got a valid colony")
	}
}

// TestFileReader_SelfLink tests the FileReader function with a self-link.
func TestFileReader_SelfLink(t *testing.T) {
	// Create a temporary file with a self-link
	content := `3
##start
room1 1 1
##end
room2 2 2
room1-room1
`
	tmpFile, err := os.CreateTemp("", "self_link_*.txt")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.WriteString(content); err != nil {
		t.Fatalf("Failed to write to temporary file: %v", err)
	}
	tmpFile.Close()

	// Test FileReader
	colony := FileReader(tmpFile.Name())
	if colony != nil {
		t.Error("Expected nil colony due to self-link, got a valid colony")
	}
}
