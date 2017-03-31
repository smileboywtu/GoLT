package lt

/*

	Pseudo Random Number Generator

 */

const (
	PRNG_A = 16807
	PRNG_M = (1 << 31) - 1
)

type PRNG struct {
	State  uint64
	PRNG_A uint64
	PRNG_M uint64
}

func (prng *PRNG) SetSeed(seed uint64) {
	prng.State = seed
}

func (prng *PRNG) GetStat() uint64 {
	return prng.State
}

func (prng *PRNG) NextInt() uint64 {
	prng.State = prng.State * prng.PRNG_A % prng.PRNG_M
	return prng.State
}

func (prng *PRNG) GetProbability() float64 {
	return float64(prng.NextInt()) / float64(prng.PRNG_M-1)
}
