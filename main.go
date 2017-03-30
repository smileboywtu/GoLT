package main

import "./lt"
import "time"
import "fmt"
import "reflect"

func main() {

	fmt.Println(time.Now())
	seed := uint32(time.Now().Unix())
	lt.SetSeed(seed)
	fmt.Printf("%d\n", lt.NextInt())

	//fmt.Println(lt.GetByteTable())
	var seq = [...]uint8{
		1, 2, 3, 4, 5, 6,
	}
	fmt.Printf("0x%X\n", lt.GetCRC16(seq[:]))

	x := lt.GenRho(10)
	fmt.Println(x, float32(34*x[0]))

	fmt.Println(lt.GenRSD(10003, 0.5, 0.1))

	a := 0x12345
	fmt.Println(uint8(a))

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

	if z, ok := lt.RestorePackage(c); ok {
		fmt.Println(z.ShowSummery())
		fmt.Println(z.GetBlockData())
		fmt.Println(reflect.DeepEqual(y, z))
	}

	yy := lt.CreateGraphNode([]uint64{1, 2, 3, 4, 5, 2}, []byte{1, 2, 3})
	fmt.Println(yy.GetSummery())

}
