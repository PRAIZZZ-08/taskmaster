package todo

import "testing"

func TestSaveAndLoad(t *testing.T) {
	testFile := "test_tasks.json"
	original := []Task{
		{ID: 1, Description: "Test Task", IsDone: false},
	}

	_ = SaveTasks(testFile, original)
	loaded, _ := LoadTasks(testFile)

	if len(loaded) != len(original) {
		t.Errorf("Expected %d tasks, got %d", len(original), len(loaded))
	}

	// Index into the slice.
	if loaded[0].Description != original[0].Description {
		t.Errorf("Data mismatch! Expected %s, got %s", original[0].Description, loaded[0].Description)
	}
}

func TestEmptyLoad(t *testing.T) {
	_, err := LoadTasks("non_existent_file.json")

	if err == nil {
		t.Errorf("Expected an error for missing file, but got nil")
	}
}
