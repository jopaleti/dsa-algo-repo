// BFS --- Breadth First Search
/*
Breadth First Search (BFS) is a fundamental graph traversal algorithm .
It involves visiting all the connected nodes of a graph in a level-by-level manner.
*/
package main

import "fmt"

type Node struct {
	Value int
	Neighbors []*Node
}

// BFS returns the length of the shortest path between the root and the target node
func BFS(root, target *Node) int {
	if root == nil || target == nil {
		return -1
	}

	queue := []*Node{root} // Store all node which are waiting to be proccessed
	step := 0	// Number of steps needed from root to current node

	// Now, iterating through the elements in the queue
	for len(queue) > 0 {
		size := len(queue)
		
		for i:=0; i<size; i++ {
			cur := queue[0]
			
			queue = queue[1:]

			// Return step if cur is target
			if cur == target {
				return step
			}

			// Add neighbors to the queue
			for _, next := range cur.Neighbors {
				queue = append(queue, next)
			}
		}
		step++
	}
	return -1	// there is no path from root to target
}

func main() {
	// Create a sample graph for demonstration
    node1 := &Node{Value: 1}
    node2 := &Node{Value: 2}
    node3 := &Node{Value: 3}
    node4 := &Node{Value: 4}
    node5 := &Node{Value: 5}

    node1.Neighbors = []*Node{node2, node3}
    node2.Neighbors = []*Node{node4}
    node3.Neighbors = []*Node{node4, node5}
    node4.Neighbors = []*Node{node5}

    // Test BFS function
    fmt.Println(BFS(node1, node5)) // Output: 2
}