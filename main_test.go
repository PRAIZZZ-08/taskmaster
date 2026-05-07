package main

import "testing" // [testing]: The built-in package for unit tests

// [TestCheck]: Must start with 'Test' and take *testing.T
func FuzzCheck(f *testing.F) {
	f.Add("test", 10)
	f.Add("", -1)

	f.Fuzz(func(t *testing.T, title string, cost int) {
		testTask := Task{Title: title, FixedCost: cost}
		_ = testTask.Check()
	})
}
