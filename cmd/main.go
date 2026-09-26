package main

import (
	"fmt"
	"lazy-radar-detector/internal/domain"
	"math"
)

func main() {
	sig := domain.Signal{
		Range: 25.5,
		Theta: 30.2,
	}

	task := domain.TaskTargetDetected

	rad := sig.Theta * (math.Pi / 180.0)
	scan := domain.Scan{
		RadarID: 1,
		Range:   sig.Range,
		Theta:   sig.Theta,
		X:       math.Round(sig.Range*math.Sin(rad)*100) / 100,
		Y:       math.Round(sig.Range*math.Cos(rad)*100) / 100,
		RCS:     10.0,
		SNR:     8.9,
		Class:   2,
	}

	fmt.Printf("Segnale: Range=%.2f, Theta=%.2f\n", sig.Range, sig.Theta)
	fmt.Printf("Task: %s", task)
	fmt.Printf("Scan: %#v, class: %s", scan, scan.Class)
}
