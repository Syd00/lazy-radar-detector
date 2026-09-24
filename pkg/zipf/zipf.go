package zipf

import (
	"errors"
	"math"
	"math/rand"
	"sync"
	"time"
)

var (
	ErrInvalidSize  = errors.New("size must be greater than 0")
	ErrInvalidAlpha = errors.New("alpha must be greater than 0")
)

type Zipf struct {
	size       int
	alpha      float64
	normFactor float64
	rnd        *rand.Rand
	mu         sync.Mutex
}

// Initialize zipf struct
func New(size int, alpha float64) (*Zipf, error) {
	if size <= 0 {
		return nil, ErrInvalidSize
	}

	if alpha <= 0 {
		return nil, ErrInvalidAlpha
	}

	normFactor := 0.0
	for i := 1; i <= size; i++ {
		normFactor += 1.0 / math.Pow(float64(i), alpha)
	}

	return &Zipf{
		size:       size,
		alpha:      alpha,
		normFactor: normFactor,
		rnd:        rand.New(rand.NewSource(time.Now().UnixNano())),
	}, nil
}

// NextInt restituisce un numero estratto secondo la distribuzione di Zipf
func (z *Zipf) NextInt() int {
	z.mu.Lock()
	defer z.mu.Unlock()
	for {
		// Estrai un rango casuale tra 1 e size
		rank := z.rnd.Intn(z.size) + 1

		// Calcola la probabilità (frequenza) per quel rango
		frequency := (1.0 / math.Pow(float64(rank), z.alpha)) / z.normFactor

		// Accetta il valore se un numero casuale è < probabilità
		if z.rnd.Float64() < frequency {
			return rank
		}
	}
}
