package main

import "testing" // [testing]: The built-in package for unit tests [4]

// [TestCheck]: Must start with 'Test' and take *testing.T
func TestCheck(t *testing.T) {
<<<<<<< HEAD
	task := Task{Title: "Test Task"}
	expected := "Audited: Test Task with cost 0"

	// [Check]: Running the actual task logic
	result := task.Check()
=======
	project := Project{Title: "Test Project"}
	expected := "Audited: Test Project with cost 0"

	// [Check]: Running the actual project logic
	result := project.Check()
>>>>>>> 8ccf087 (test: updated test logic)

	if result != expected {
		// [Errorf]: Reports a failure and what the expected vs actual was
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}
