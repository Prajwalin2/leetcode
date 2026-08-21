package main

type Node struct {
	Val       int
	Neighbors []*Node
}

func main() {
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	m := map[int]*Node{}

	dfs(node, m)

	return m[node.Val]
}

func dfs(node *Node, copies map[int]*Node) {
	newNode := new(Node)
	newNode.Val = node.Val

	copies[node.Val] = newNode

	for _, neighbor := range node.Neighbors {
		if copies[neighbor.Val] == nil {
			dfs(neighbor, copies)
		}

		newNode.Neighbors = append(newNode.Neighbors, copies[neighbor.Val])
	}
}
