package domain

type Scan struct {
	RadarID int
	Range   float64
	Theta   float64
	X       float64
	Y       float64
	RCS     float64
	SNR     float64
	Class   TaskType
}
