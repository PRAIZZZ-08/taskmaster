package main

import "testing" // [testing]: The built-in package for unit tests [4]

// [TestCheck]: Must start with 'Test' and take *testing.T
func TestCheck(t *testing.T) {
	project := Project{Title: "Test Project"}
	expected := "Audited: Test Project with cost 0"

	// [Check]: Running the actual project logic
	result := project.Check()

	if result != expected {
		// [Errorf]: Reports a failure and what the expected vs actual was
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}
