package DataStructures

// Node struct left smaller right larger nunber.
type BinaryTree struct {
	Key   int
	Left  *BinaryTree
	Right *BinaryTree
}

// Insert key to Binary Search Tree
// if the key argument is greater than current key go right
// if the key agrument is less than current key go left
// if the current key node is not null then make recursive call
func (n *BinaryTree) Insert(key int) {
	if n.Key < key {
		if n.isRightAvailable() {
			n.Right = &BinaryTree{Key: key}
			return
		}

		n.Right.Insert(key)
	} else if n.Key > key {
		if n.isLeftAvailable() {
			n.Left = &BinaryTree{Key: key}
			return
		}

		n.Left.Insert(key)
	}
}

func (n *BinaryTree) isLeftAvailable() bool {
	return n.Left == nil
}

func (n *BinaryTree) isRightAvailable() bool {
	return n.Right == nil
}

// Search Binary tree for a value
// recursively call method and based on current node move down left or right
// until we reach our value or the end of the tree
func (n *BinaryTree) Search(key int) bool {

	if n == nil {
		return false
	}

	if n.Key < key {
		return n.Right.Search(key)
	} else if n.Key > key {
		return n.Left.Search(key)
	}

	return true
}

// Initialize
func NewBinarySearchTree() *BinaryTree {
	return &BinaryTree{}
}
