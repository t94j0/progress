package progress

import "fmt"

type ProgressTracker struct {
	Comment string
	Count   int
	Maximum int
}

func CreateProgressTracker(comment string) *ProgressTracker {
	return &ProgressTracker{comment, 0, -1}
}

func CreateProgressTrackerMax(comment string, max int) *ProgressTracker {
	return &ProgressTracker{comment, 0, max}
}

func (p *ProgressTracker) Increment() {
	p.Count += 1
	if p.Maximum == -1 {
		fmt.Printf("\r%s: %d", p.Comment, p.Count)
	} else {
		fmt.Printf("\r%s: (%d/%d)", p.Comment, p.Count, p.Maximum)
	}

}

func (p *ProgressTracker) Close() {
	fmt.Print("\n")
}
