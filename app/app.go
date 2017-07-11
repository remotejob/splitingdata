package main

import (
	"fmt"
	"hash/fnv"

	"github.com/gholt/ring"
)

// "stathat.com/c/consistent"

// "stathat.com/c/consistent"

// var castagnoliTable = crc32.MakeTable(crc32.Castagnoli)

// const shards int = 40

var (
	node1id = "node1"
	node2id = "node2"
	node3id = "node3"
	node4id = "node4"
)

func main() {
	builder := ring.NewBuilder(64)
	// (active, capacity, no tiers, no addresses, meta, no config)
	builder.AddNode(true, 1, nil, nil, "NodeA", nil)
	builder.AddNode(true, 1, nil, nil, "NodeB", nil)
	builder.AddNode(true, 1, nil, nil, "NodeC", nil)
	builder.AddNode(true, 1, nil, nil, "NodeD", nil)
	builder.AddNode(true, 1, nil, nil, "NodeE", nil)
	builder.AddNode(true, 1, nil, nil, "NodeF", nil)
	// This rebalances if necessary and provides a usable Ring instance.
	ring := builder.Ring()
	// This value indicates how many bits are in use for determining ring
	// partitions.

	users := []string{"user0", "user1", "user2", "user3", "user4", "user5", "user6", "user7", "user8", "user9", "user10"}
	partitionBitCount := ring.PartitionBitCount()
	for _, item := range users {
		// We're using fnv hashing here, but you can use whatever you like.
		// We don't actually recommend fnv, but it's useful for this example.
		hasher := fnv.New64a()
		hasher.Write([]byte(item))
		partition := uint32(hasher.Sum64() >> (64 - partitionBitCount))
		// We can just grab the first node since this example just uses one
		// replica. See Builder.SetReplicaCount for more information.
		node := ring.ResponsibleNodes(partition)[0]
		fmt.Printf("%s is handled by %v\n", item, node.Meta())
	}
	// memcacheServers := []string{"0", "1", "2", "3"}

	// ring := hashring.New(memcacheServers)
	// // server, _ := ring.GetNode("my_key")

	// log.Println(server)
	// ring := ketama.NewRing([]*ketama.Node{
	// 	ketama.NewNode("000", "binding data0", 1),
	// 	ketama.NewNode("001", "binding data1", 1),
	// 	ketama.NewNode("002", "binding data2", 1),
	// 	ketama.NewNode("003", "binding data3", 1),
	// 	ketama.NewNode("004", "binding data4", 1),
	// 	ketama.NewNode("005", "binding data5", 1),
	// })
	// fmt.Printf("Get a server by key \"key1\": %v\n", ring.Get("key1"))
	// fmt.Printf("Get a server by key \"key2\": %v\n", ring.Get("key2"))
	// fmt.Printf("Get a server by key \"key3\": %v\n", ring.Get("key3"))
	// fmt.Printf("Get a server by key \"key4\": %v\n", ring.Get("key4"))
	// fmt.Printf("Get a server by key \"key5\": %v\n", ring.Get("key5"))
	// fmt.Printf("Get a server by key \"key1\" again: %v\n", ring.Get("key1"))
	// c := consistenthash.NewRing()

	// c.AddNode(node1id)
	// c.AddNode(node2id)
	// c.AddNode(node3id)
	// // c.AddNode(node4id)
	// // c.AddNode("node5")

	// users := []string{"user0", "user1", "user2", "user3", "user4", "user5", "user6", "user7", "user8", "user9", "user10"}
	// for _, u := range users {
	// 	server, _ := ring.GetNode(u)
	// 	// if err != nil {
	// 	// 	log.Fatal(err)
	// 	// }
	// 	fmt.Printf("%s => %s\n", u, server)
	// }
	// node := r.Get("user0")

	// log.Println(node)

	// 64 is fine unless you have a specific use case.
	// builder := ring.NewBuilder(64)
	// // (active, capacity, no tiers, no addresses, meta, no conf)
	// builder.AddNode(true, 1, nil, nil, "NodeA", nil)
	// builder.AddNode(true, 1, nil, nil, "NodeB", nil)
	// builder.AddNode(true, 1, nil, nil, "NodeC", nil)
	// builder.AddNode(true, 1, nil, nil, "NodeD", nil)
	// // This rebalances if necessary and provides a usable Ring instance.
	// ring := builder.Ring()
	// // This value indicates how many bits are in use for determining ring
	// // partitions.
	// partitionBitCount := ring.PartitionBitCount()
	// for _, item := range []string{"First", "Second", "Third"} {
	// 	// We're using fnv hashing here, but you can use whatever you like.
	// 	// We don't actually recommend fnv, but it's useful for this example.
	// 	hasher := fnv.New64a()
	// 	hasher.Write([]byte(item))
	// 	partition := uint32(hasher.Sum64() >> (64 - partitionBitCount))
	// 	// We can just grab the first node since this example just uses one
	// 	// replica. See Builder.SetReplicaCount for more information.
	// 	node := ring.ResponsibleNodes(partition)[0]
	// 	fmt.Printf("%s is handled by %v\n", item, node.Meta())
	// }
	// crc := crc32.New(castagnoliTable)
	// crc.Write([]byte("abcd"))

	// // k	fmt.Printf("Sum32 : %x \n", crc)

	// log.Println(crc)

	// h := crc32.NewIEEE()
	// h.Write([]byte("value"))
	// hash := float64(h.Sum32())

	// shard_no := math.Mod(hash, float64(shards))

	// // res := hash % 2

	// log.Println(shard_no)

	// h.Write([]byte("value2"))
	// hash1 := float64(h.Sum32())
	// shard_no1 := math.Mod(hash1, float64(shards))
	// log.Println(shard_no1)
	// hasher := jump.New(10, jump.CRC64)
	// h0 := hasher.Hash("serv0")

	// log.Println(h0)
	// h := jump.HashString("127.0.0.1", 8, jump.CRC64)
	// // log.Println(crc32([]byte("foobar")))
	// c := consistent.New()
	// c.Add("cacheA")
	// c.Add("cacheB")
	// c.Add("cacheC")
	// c.Add("cacheD")

	// users := []string{"user0", "user1", "user2", "user3", "user4", "user5", "user6", "user7", "user8", "user9", "user10"}
	// for _, u := range users {
	// 	server, err := c.Get(u)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Printf("%s => %s\n", u, server)
	// }

	// c := consistent.New()

	// // adds the hosts to the ring
	// c.Add("serv0")
	// c.Add("serv1")
	// c.Add("serv2")
	// c.Add("serv3")
	// // c.Add("serv4")

	// users := []string{"user_mcnulty", "user_bunk", "user_omar", "user_bunny", "user_stringer", "test1", "test2"}
	// for _, u := range users {
	// 	server, err := c.GetLeast(u)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	c.Inc(server)
	// 	fmt.Printf("%s => %s\n", u, server)
	// 	c.Done(server)

	// }
	// Returns the host that owns `key`.
	//
	// As described in https://en.wikipedia.org/wiki/Consistent_hashing
	//
	// It returns ErrNoHosts if the ring has no hosts in it.
	// host, err := c.Get("/app.html")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Println(host)
}
