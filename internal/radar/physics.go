package radar

import "lazy-radar-detector/internal/domain"

// ComputeIdealScan esegue la trasformazione fisica deterministica senza rumore
func ComputeIdealScan(sig domain.Signal) domain.Scan {
	// 1. Calcola X e Y da Range e Theta (trigonometria)
	// 2. Assegna RCS ed SNR nominali
	// 3. Classifica il Task (TaskEmptySpace o TaskTargetDetected)
	// 4. Ritorna il RadarScan "perfetto"
}
