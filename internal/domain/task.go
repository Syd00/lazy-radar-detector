package domain

// TaskType è il nostro tipo personalizzato basato su int
type TaskType int

const (
	_                  TaskType = iota
	TaskEmptySpace              // Empty space
	TaskTargetDetected          // Target
)

func (t TaskType) String() string {
	switch t {
	case TaskEmptySpace:
		return "EmptySpace"
	case TaskTargetDetected:
		return "TargetDetected"
	default:
		return "Unknown"
	}
}
