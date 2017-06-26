/*

	code helper to construct lt packet


	Created: 2017/3/31
	Contact: smileboywtu@gmail.com

 */
package lt

import (
	"os"
	"math"
)

type LTFactor struct {
	K     uint64
	C     float64
	Delta float64
	CDF   []float64
	PRNG  *PRNG
}

type RandomBlockGenerator struct {
	file_name   string `need to supply when init struct`
	file_size   int64
	total_parts uint64
	block_size  uint64 `need to supply when init struct`
	src_bytes   []byte
	src_blocks  map[uint64][]byte
	factor      *LTFactor `need to supply when init struct`
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
	Init src block generator
 */
func (generator *RandomBlockGenerator) InitBlockGenerator() (string, bool) {

	// read all the byte
	f, err := os.Open(generator.file_name)
	defer f.Close()

	if err != nil {
		return string(err.Error()), false
	}
	fi, err := f.Stat()
	if err != nil {
		return string(err.Error()), false
	}

	// read all bytes
	generator.file_size = fi.Size()
	generator.total_parts = uint64(math.Ceil(float64(generator.file_size) / float64(generator.block_size)))
	generator.src_bytes = make([]byte, 0, generator.file_size)
	f.Read(generator.src_bytes)

	block_size := generator.block_size
	limit, offset := uint64(0), uint64(0)
	// save to src blocks
	for i := uint64(0); i < generator.total_parts; i += 1 {
		offset = i * block_size
		generator.src_blocks[i] = make([]byte, block_size)
		if i == generator.total_parts-1 {
			// last part
			limit -= 1
		} else {
			limit = offset + block_size

		}
		copy(generator.src_blocks[i][:], generator.src_bytes[offset:limit])
	}

	return "", true
}

/*
	Get source block random
 */
func (generator *RandomBlockGenerator) GetNextPacket() []byte {
	tmp := make([]byte, 2)
	return tmp
}
