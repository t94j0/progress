package progress

import "fmt"

// ProgressTracker holds information about how far the tracker has gotten by
// keeping track of the current count and the maximum. The `Comment` field
// indicates what string should be presented when incrementing
type ProgressTracker struct {
	Comment string
	Count   int
	Maximum int
}

// CreateTracker creates a ProgressTracker object that has a `Maximum` of -1 to
// signify that the maxiumum value is unknown.
func CreateTracker(comment string) *ProgressTracker {
	return &ProgressTracker{comment, 0, -1}
}

// CreateTrackerMax creates a tracker where the `Maximum` field known. This
// changes how the printing is done
func CreateTrackerMax(comment string, max int) *ProgressTracker {
	return &ProgressTracker{comment, 0, max}
}

// Increment simply increments the counter on ProgressTracker
func (p *ProgressTracker) Increment() {
	p.Count += 1
	p.print()
}

// print prints the count according to if the Maximum is known or not
func (p *ProgressTracker) print() {
	if p.Maximum == -1 {
		fmt.Printf("\r%s: %d", p.Comment, p.Count)
	} else {
		fmt.Printf("\r%s: (%d/%d)", p.Comment, p.Count, p.Maximum)
	}
}

// Close is used to create a newline so that the previous string in stdout is
// not written over
func (p *ProgressTracker) Close() {
	fmt.Print("\n")
}
