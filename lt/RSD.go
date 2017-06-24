/*
	The robust form of distribution is defined by adding an extra set
	of values to the elements of mass function of the ideal soliton
	distribution and then standardising so that the values add up to
	1. The extra set of values, t, are defined in terms of an additional
	real-valued parameter δ (which is interpreted as a failure probability)
	and an integer parameter M (M < N) . Define R as R=N/M. Then the values
	added to p(i), before the final standardisation, are

	t(i) = 1 / (i*M) ------------ (i=1, 2, 3, M-1)
	t(i) = ln(R/δ) / M ---------- (i = M)
	t(i) = 0 -------------------- (i= M+1, ..., N)

	While the ideal soliton distribution has a mode (or spike) at 1, the
	effect of the extra component in the robust distribution is to add an
	additional spike at the value M.

*/
package lt

import "math"

// rho(1) = 1 / K, d=1
// rho(d) = 1 / d*(d-1) d=2, 3, 4, ..., K
// :params K: number of source block
// :return list: rho array list
func GenRho(k uint64) []float64 {

	rho_set := make([]float64, k)

	for i := uint64(1); i <= k; i++ {
		if i == 1 {
			rho_set[i-1] = float64(1) / float64(k)
		} else {
			rho_set[i-1] = float64(1) / float64(i*(i-1))
		}
	}

	return rho_set
}

// :params s: s = c * ln( K / delta ) * sqrt( K )
// :params K: number of source block
// :params delta: delta is a bound on the probability that the decoding fails
// :return list: list of tau
func GenTau(s float64, k uint64, delta float64) []float64 {

	tau_set := make([]float64, k)
	pivot := uint64(math.Floor(float64(k) / s))

	for i := uint64(1); i <= k; i++ {
		if i < pivot {
			tau_set[i-1] = s / float64(k*i)
		} else if i == pivot {
			tau_set[i-1] = s / float64(k) * float64(math.Log(s/delta))
		} else {
			tau_set[i-1] = float64(0)
		}
	}

	return tau_set
}

// calculate the sum of a item in a slice which of the same type
// :params set: slice of values
// :return value: sum of the slice value
func sumSlice(set []float64) float64 {
	var sum float64 = 0
	for _, value := range set {
		sum += value
	}
	return sum
}

// :params k: the number of source block
// :params delta: delta is a bound on the probability that the decoding fails
// :params c: c is a constant of order 1
// :return list: list of mu
func GenMu(k uint64, delta float64, c float64) []float64 {

	mu_set := make([]float64, k)
	var s float64 = c * math.Log(float64(k)/delta) * math.Sqrt(float64(k))

	rho_set := GenRho(k)
	tau_set := GenTau(s, k, delta)

	normalizer := sumSlice(rho_set) + sumSlice(tau_set)

	for index, _ := range rho_set {
		mu_set[index] = (rho_set[index] + tau_set[index]) / normalizer
	}

	return mu_set
}

// :params k: the number of source block
// :params delta: delta is a bound on the probability that the decoding fails
// :params c: c is a constant of order 1
// :return list: list of RSD
func GenRSD(k uint64, delta float64, c float64) []float64 {

	rsd_set := make([]float64, k)
	mu_set := GenMu(k, delta, c)

	for i := uint64(1); i <= k; i++ {
		rsd_set[i-1] = sumSlice(mu_set[:i])
	}

	return rsd_set
}
