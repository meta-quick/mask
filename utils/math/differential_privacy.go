package math

import (
	"github.com/google/differential-privacy/go/noise"
)

func LaplaceNoise() noise.Noise {
	return  noise.Laplace()
}

