package lt

import "fmt"

/*

	Resolve Graph

	Created: 2017/3/30
	Copyright: (c) chenbo<smileboywtu@gmail.com>

 */

type GraphNode struct {
	blocks []uint64
	bytes  []byte
}

type Tuple struct {
	block uint64
	bytes []byte
}

type GraphResovler struct {
	number_of_block uint64
	graph           map[uint64][]GraphNode
	resolved_block  map[uint64][]byte
}

// constructor
func init() {

}

/*
    create graph node
*/
func CreateGraphNode(blocks []uint64, data []byte) GraphNode {

	encontered := map[uint64]bool{}
	result := []uint64{}

	for _, block := range blocks {
		if _, ok := encontered[block]; ok {
			continue
		}
		result = append(result, block)
		encontered[block] = true
	}

	return GraphNode{
		blocks: result,
		bytes:  data,
	}
}

/*
	show graph node summery
 */
func (node *GraphNode) GetSummery() string {
	return fmt.Sprintf(
		"|%-10s|%v\n"+
			"|%-10s|%v\n",
		"blocks", node.blocks,
		"bytes", node.bytes,
	)
}

/*
	member func for graph resolver

	add blocks to resolver
 */
func (resolver *GraphResovler) AddBlock(blocks []uint64, data []byte) bool {
	if len(blocks) == 1 {
		eliminate := resolver.Resolve(blocks[0], data)
		for len(eliminate) > 0 {
			current := eliminate[0]
			eliminate = append(
				eliminate,
				resolver.Resolve(current.block, current.bytes)...,
			)
			eliminate = eliminate[1:]
		}
	} else {
		reserve := make([]uint64, len(blocks))
		copy(reserve, blocks)
		for index, block := range blocks {
			// remove block
			if _, ok := resolver.resolved_block[block]; ok {
				for index := range data {
					data[index] ^= resolver.resolved_block[block][index]
				}
				reserve = append(reserve[:index], reserve[index+1:]...)
			}
		}

		if len(reserve) == 1 {
			return resolver.AddBlock(reserve, data)
		} else {

			node := new(GraphNode)
			node.blocks = reserve
			node.bytes = data

			for _, block := range reserve {
				if _, ok := resolver.graph[block]; ok {
					resolver.graph[block] = append(
						resolver.graph[block],
						*node,
					)
				} else {
					resolver.graph[block] = []GraphNode{*node}
				}
			}
		}

	}

	return uint64(len(resolver.resolved_block)) >= resolver.number_of_block
}

/*
	add block to graph and resolve
 */
func (resolver *GraphResovler) Resolve(block uint64, bytes []byte) ([]Tuple) {

	// add to resolved blocks
	tmp := make([]byte, len(bytes))
	copy(tmp, bytes)
	resolver.resolved_block[block] = tmp

	eliminated := []Tuple{}
	if nodes, ok := resolver.graph[block]; ok {
		delete(resolver.graph, block)
		for _, item := range nodes {
			// resolve
			for index, byte := range bytes {
				item.bytes[index] ^= byte
			}
			// remove from graph node blocks
			for index, value := range item.blocks {
				if value == block {
					item.blocks = append(
						item.blocks[:index],
						item.blocks[index+1:]...,
					)
					break
				}
			}
			// check reserve
			if len(item.blocks) == 1 {
				eliminated = append(
					eliminated,
					Tuple{
						block: item.blocks[0],
						bytes: item.bytes,
					},
				)
			}
		}
	}

	return eliminated
}
