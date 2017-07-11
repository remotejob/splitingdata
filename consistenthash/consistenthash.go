package consistenthash

import (
	"hash/crc32"
	"sort"
)

// Ring is a network of distributed nodes.
type Ring struct {
	Nodes Nodes
}

// Nodes is an array of nodes.
type Nodes []Node

// Node is a single entity in a ring.
type Node struct {
	Id     string
	HashId uint32
}

func NewRing() *Ring {
	return &Ring{Nodes: Nodes{}}
}
func NewNode(id string) *Node {
	crc32q := crc32.MakeTable(0xD5828281)
	return &Node{
		Id:     id,
		HashId: crc32.Checksum([]byte(id), crc32q),
	}
}

func (r *Ring) AddNode(id string) {
	node := NewNode(id)
	r.Nodes = append(r.Nodes, *node)
	sort.Sort(r.Nodes)
}

func (n Nodes) Len() int           { return len(n) }
func (n Nodes) Less(i, j int) bool { return n[i].HashId < n[j].HashId }
func (n Nodes) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }

func (r *Ring) Get(key string) string {
	crc32q := crc32.MakeTable(0xD5828281)
	searchfn := func(i int) bool {
		return r.Nodes[i].HashId >= crc32.Checksum([]byte(key), crc32q)
	}
	i := sort.Search(r.Nodes.Len(), searchfn)
	if i >= r.Nodes.Len() {
		i = 0
	}
	return r.Nodes[i].Id
}
