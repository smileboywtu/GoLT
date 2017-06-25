/*

	code helper to construct lt packet


	Created: 2017/3/31
	Contact: smileboywtu@gmail.com

 */
package lt

import "os"

type LTFactor struct {
	K     uint64
	C     float64
	Delta float64
	CDF   []float64
	PRNG  *PRNG
}

type RandomBlockGenerator struct {
	filename   string
	block_size uint64
	src_bytes  []byte
}

/*

	init the lt factor with given parameters

 */
func InitLtFactor(k uint64, c float64, delta float64, state uint64, prng_a uint64, prng_m uint64) (*LTFactor) {

	factor := new(LTFactor)

	factor.K = k
	factor.C = c
	factor.Delta = delta

	factor.PRNG = new(PRNG)
	factor.PRNG.State = state
	factor.PRNG.PRNG_A = prng_a
	factor.PRNG.PRNG_M = prng_m

	factor.CDF = GenRSD(k, delta, c)

	return factor
}

/*
	set the seed of the lt factor prng
 */
func (factor *LTFactor) SetSeed(seed uint64) {
	factor.PRNG.State = seed
}

/*
	get the state of lt factor state
 */
func (factor *LTFactor) GetState() uint64 {
	return factor.PRNG.State
}

/*
	generate sample src blocks for send
 */
func (factor *LTFactor) GetSrcBlocks() []uint64 {

	degree := factor.GetSampleDegree()

	i := uint64(0)
	blocks := map[uint64]bool{}
	for i < degree {
		number := factor.PRNG.NextInt() % factor.K
		if _, ok := blocks[number]; !ok {
			blocks[number] = true
			i++
		}
	}

	_blocks := make([]uint64, len(blocks))
	for key := range blocks {
		_blocks = append(_blocks, key)
	}
	return _blocks
}

func (factor *LTFactor) GetSampleDegree() uint64 {

	// get probability
	p := factor.PRNG.GetProbability()

	i := uint64(0)
	size := uint64(len(factor.CDF))
	for i < size {
		if factor.CDF[i] > p {
			return i + 1
		}
		i++
	}

	return i + 1
}

/*
	Get source block random
 */
func (generator *RandomBlockGenerator) NewBlockGenerator() (string, bool) {

	// read all the byte

	f, err := os.Open(generator.filename)
	if err {
		return string(err), false
	}
	fi, err := f.Stat()
	if err != nil {
		return string(err), false
	}

	// read all bytes
	generator.src_bytes = make([]byte, 0, fi.Size())
	f.Read(generator.src_bytes)

	// save to src blocks

	return "", true
}
