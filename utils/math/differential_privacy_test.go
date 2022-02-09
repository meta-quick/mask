package math

import "testing"

func Test_differential_privacy_test(t *testing.T) {
	lap := LaplaceNoise()
	ox := lap.AddNoiseInt64(10,1,1,0.3,0)
	x, _ := lap.ComputeConfidenceIntervalInt64(11,1,1,0.3,0,1)
	println(ox,x.LowerBound,x.UpperBound)
}