package lt

// const params
const (
	PRNG_A        = 16807
	PRNG_M        = (1 << 31) - 1
	PRNG_MAX_RAND = PRNG_M - 1
)

// member
var state uint32

// member method
func init() {
	state = 2067261
}

func SetSeed(seed uint32) {
	state = seed
}

func GetStat() uint32 {
	return state
}

func NextInt() uint32 {
	state = state * PRNG_A % PRNG_M
	return uint32(state)
}

func GetProbability() float32 {
	return float32(NextInt()) / float32(PRNG_MAX_RAND)
}
