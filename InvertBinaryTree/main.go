package main

import "fmt"

//         15
//       /    \
//     10      20
//    /  \    /  \
//   8   12  17   25
//
// this is an example of a binary search tree
// the top node (15) is the root node and
// the bottom nodes (8,12,17,25) are all leaf nodes
// the right node is always bigger and the left node is always smaller than the parent node
// the advantage of a binary search tree is that it is fast O(log n)

// Node: represents the components of a binary tree
type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

// Insert: Will add a node to a tree
// The key to add should not already be in the tree
func (n *Node) Insert(k int) {
	if n.Key < k {
		//move right
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			//recursive
			n.Right.Insert(k)
		}
	} else if n.Key > k {
		//move left
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			//recursive
			n.Left.Insert(k)
		}
	}
}

// Search will take in a key value
// and return true if there is a node for that value
func (n *Node) Search(k int) bool {

	// Base Case
	if n == nil {
		return false
	}

	if n.Key < k {
		//move right
		return n.Right.Search(k)
	} else if n.Key > k {
		//move left
		return n.Left.Search(k)
	}

	return true
}

// in-order traversal
func (n *Node) InOrderTraversal() {
	if n == nil {
		return
	}
	n.Left.InOrderTraversal()
	fmt.Println(n.Key)
	n.Right.InOrderTraversal()
}

// pre-order traversal
func (n *Node) PreOrderTraversal() {
	if n == nil {
		return
	}
	fmt.Println(n.Key)
	n.Left.PreOrderTraversal()
	n.Right.PreOrderTraversal()
}

// post-order traversal
func (n *Node) PostOrderTraversal() {
	if n == nil {
		return
	}
	n.Left.PostOrderTraversal()
	n.Right.PostOrderTraversal()
	fmt.Println(n.Key)
}

//Inverts the Binary Tree
func InvertTree(n *Node) *Node {
	if n == nil {
		return n
	}

	invertedLeft := InvertTree(n.Left)
	invertedRight := InvertTree(n.Right)

	n.Left = invertedRight
	n.Right = invertedLeft

	return n
}

func main() {
	tree := &Node{Key: 100}
	tree.Insert(10)
	tree.Insert(50)
	tree.Insert(45)
	tree.Insert(500)
	fmt.Println(tree)
	fmt.Println()
	tree.PreOrderTraversal()
}
