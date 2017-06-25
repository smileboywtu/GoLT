package main

import "./lt"
import "time"
import "fmt"
import (
	"reflect"
	"strings"
)

func main() {

	// Test PRNG
	fmt.Println(strings.Repeat("=", 80))
	fmt.Println("测试随机数：")
	seed := uint64(time.Now().Unix())
	prng := lt.PRNG{0, lt.PRNG_A, lt.PRNG_M}
	prng.SetSeed(seed)
	fmt.Printf("%d\n", prng.NextInt())


	// GET CRC16
	fmt.Println(strings.Repeat("=", 80))
	fmt.Println("测试CRC16校验码：")
	var seq = [...]uint8{
		1, 2, 3, 4, 5, 6,
	}
	fmt.Printf("0x%X\n", lt.GetCRC16(seq[:]))

	// Test RHO
	fmt.Println(strings.Repeat("=", 80))
	fmt.Println("测试RHO随机码：")
	x := lt.GenRho(10)
	fmt.Println(x, float32(34*x[0]))
	fmt.Println(lt.GenRSD(100, 0.5, 0.1))

	// Pack Data
	fmt.Println(strings.Repeat("=", 80))
	fmt.Println("测试LT编解码：")
	y := lt.InitPacket(500, 3, 55679813123, []byte{1, 2, 3})
	fmt.Println(y.ShowSummery())
	fmt.Println(y.GetFileSize())
	y.SetFileSize(5)
	fmt.Println(y.GetFileSize())

	y.SetBlockData([]byte{1, 2, 3})
	fmt.Println(y.GetBlockData())

	y.SetBlockData([]byte{9, 6, 7})
	fmt.Println(y.GetBlockData())

	fmt.Println(y.ShowSummery())
	y.SetBlockSize(3)
	c := y.BuildPacket()
	fmt.Println(c)

	// Restore Data
	if z, ok := lt.RestorePackage(c); ok {
		fmt.Println(z.ShowSummery())
		fmt.Println(z.GetBlockData())
		fmt.Println(reflect.DeepEqual(y, z))
	}

	yy := lt.CreateGraphNode([]uint64{1, 2, 3, 4, 5, 2}, []byte{1, 2, 3})
	fmt.Println(yy.GetSummery())

}
