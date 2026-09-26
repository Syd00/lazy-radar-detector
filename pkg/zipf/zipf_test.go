package zipf_test

import (
	"errors"
	"sync"
	"testing"

	"lazy-radar-detector/pkg/zipf"
)

// Verify that invalid parameters return expected errors
func TestNewValidation(t *testing.T) {
	tests := []struct {
		name        string
		size        int
		alpha       float64
		expectedErr error
	}{
		{name: "size zero", size: 0, alpha: 2.0, expectedErr: zipf.ErrInvalidSize},
		{name: "size negative", size: -3, alpha: 2.0, expectedErr: zipf.ErrInvalidSize},
		{name: "alpha zero", size: 2, alpha: 0.0, expectedErr: zipf.ErrInvalidAlpha},
		{name: "alpha negative", size: 2, alpha: -2.0, expectedErr: zipf.ErrInvalidAlpha},
		{name: "valid parameters", size: 2, alpha: 3.0, expectedErr: nil},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := zipf.New(tc.size, tc.alpha)
			if !errors.Is(err, tc.expectedErr) {
				t.Fatalf("Expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}

// Verify that the distribution is correct
func TestZipfDistribution(t *testing.T) {

	classes := map[int]int{1: 0, 2: 0}
	gen, err := zipf.New(2, 3)

	if err != nil {
		t.Fatalf("Error during generator creation: %v", err)
	}

	for range 1000 {
		class := gen.NextInt()
		if class < 1 || class > 2 {
			t.Fatalf("Expected class between 1 and 2, got %d", class)
		}
		classes[class]++
	}

	if classes[1] <= classes[2] {
		t.Errorf("Expected class 1 > class 2")
	}
}

func TestConcurrentNextInt(t *testing.T) {
	gen, err := zipf.New(10, 2.0)

	if err != nil {
		t.Fatalf("Error during generator creation: %v", err)
	}

	var wg sync.WaitGroup

	for range 10 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range 200 {
				gen.NextInt()
			}
		}()
	}
	wg.Wait()
}
