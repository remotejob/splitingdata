package main

import (
	"hash/crc32"
	"log"
	"math"
)

var castagnoliTable = crc32.MakeTable(crc32.Castagnoli)

func main() {
	crc := crc32.New(castagnoliTable)
	crc.Write([]byte("abcd"))

	// k	fmt.Printf("Sum32 : %x \n", crc)

	log.Println(crc)

	h := crc32.NewIEEE()
	h.Write([]byte("value"))
	hash := float64(h.Sum32())

	shard_no := math.Mod(hash, float64(10))

	// res := hash % 2

	log.Println(shard_no)

	h.Write([]byte("value2"))
	hash1 := float64(h.Sum32())
	shard_no1 := math.Mod(hash1, float64(10))
	log.Println(shard_no1)

	// log.Println(crc32([]byte("foobar")))

}
