package lt

/*

	LT Packet Method helper for build LT Packet

	Created: 2017/3/29
	Contact: smileboywtu@gmail.com

 */

import "fmt"

const DEFAULT_HEADER_SIZE = 20
const DEFAULT_BLOCK_SIZE = 80

const CRC_DATA_OFFSET = 0
const FILE_SIZE_OFFSET = 2
const BLOCK_SIZE_OFFSET = 10
const BLOCK_SEED_OFFSET = 12
const BLOCK_DATA_OFFSET = 20

// members
type LTPacket struct {
	file_size  uint64
	block_size uint16
	block_seed uint64
	block_data []byte
}

// version
var version string = "1.0.0"

// constructor
func init() {

}

// member function
func GetVersion() string {
	return version
}

// create new LTPacket
func InitPacket(filesize uint64, blocksize uint16, blockseed uint64, blockdata []byte) LTPacket {
	return LTPacket{
		file_size:  filesize,
		block_size: blocksize,
		block_seed: blockseed,
		block_data: blockdata,
	}
}

/*
	Getters and Setters for packet
 */
func (packet *LTPacket) GetFileSize() uint64 {
	return packet.file_size
}

func (packet *LTPacket) SetFileSize(fs uint64) {
	packet.file_size = fs
}

func (packet *LTPacket) GetBlockSize() uint16 {
	return packet.block_size
}

func (packet *LTPacket) SetBlockSize(bs uint16) {
	packet.block_size = bs
}

func (packet *LTPacket) GetBlockSeed() uint64 {
	return packet.block_seed
}

func (packet *LTPacket) SetBlockSeed(bs uint64) {
	packet.block_seed = bs
}

func (packet *LTPacket) GetBlockData() []byte {
	return packet.block_data
}

func (packet *LTPacket) SetBlockData(bd []byte) {
	tmp := make([]byte, len(bd))
	copy(tmp, bd)
	packet.block_data = tmp
}

func (packet *LTPacket) ShowSummery() string {
	return fmt.Sprintf(
		"packet summery:"+
			"\n|%-10s|%10d|\n"+
			"|%-10s|%10d|\n"+
			"|%-10s|%10d|\n",
		"file size", packet.file_size,
		"block size", packet.block_size,
		"block seed", packet.block_seed,
	)
}

/*
	set data in byte array with data offset with start, with length
 */
func setDataInByteArray(bytes []byte, data uint64, start int8, len int8) {

	for shift := len - 1; shift >= 0; shift-- {
		var getter uint64 = 0xFF
		getter = (getter << (uint8(shift) * 8) & data) >> (uint8(shift) * 8)
		bytes[start] = uint8(getter)
		start++
	}
}

/*
	get data in byte array with data offset with start, with length
 */
func getDataInByteArray(bytes []byte, start int8, len int8) uint64 {
	var getter uint64 = 0x00
	for shift := len - 1; shift >= 0; shift-- {
		current := uint64(bytes[start] & 0xFF)
		getter |= current << (uint8(shift) * 8)
		start++
	}
	return getter
}

/*
	build the final packet to byte array
 */
func (packet *LTPacket) BuildPacket() []byte {

	final_packet := make([]byte, DEFAULT_HEADER_SIZE+packet.block_size)

	setDataInByteArray(
		final_packet, packet.file_size,
		FILE_SIZE_OFFSET, BLOCK_SIZE_OFFSET-FILE_SIZE_OFFSET,
	)
	setDataInByteArray(
		final_packet, uint64(packet.block_size),
		BLOCK_SIZE_OFFSET, BLOCK_SEED_OFFSET-BLOCK_SIZE_OFFSET,
	)
	setDataInByteArray(
		final_packet, packet.block_seed,
		BLOCK_SEED_OFFSET, BLOCK_DATA_OFFSET-BLOCK_SEED_OFFSET,
	)
	copy(final_packet[BLOCK_DATA_OFFSET:], packet.block_data[:])

	// set crc 16 code
	crc16 := GetCRC16(final_packet[FILE_SIZE_OFFSET:DEFAULT_HEADER_SIZE-FILE_SIZE_OFFSET])
	setDataInByteArray(
		final_packet, uint64(crc16),
		CRC_DATA_OFFSET, FILE_SIZE_OFFSET-CRC_DATA_OFFSET,
	)

	return final_packet
}

/*
	restore packet from byte array
 */
//func (packet *LTPacket) RestorePackage(bytes []byte) LTPacket {
//
//}
