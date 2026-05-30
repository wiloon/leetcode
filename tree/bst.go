package tree

// TreeNode is a node in a binary search tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// BST is a binary search tree rooted at root.
type BST struct {
	root *TreeNode
}

// NewBST builds a BST from the given values (inserted in order).
func NewBST(vals ...int) *BST {
	b := &BST{}
	for _, v := range vals {
		b.Insert(v)
	}
	return b
}

// Insert adds val into the tree. Average time: O(h), worst O(n).
func (b *BST) Insert(val int) {
	b.root = insert(b.root, val)
}

func insert(node *TreeNode, val int) *TreeNode {
	if node == nil {
		return &TreeNode{Val: val}
	}
	if val < node.Val {
		node.Left = insert(node.Left, val)
	} else if val > node.Val {
		node.Right = insert(node.Right, val)
	}
	return node
}

// Search returns true if val exists in the tree.
func (b *BST) Search(val int) bool {
	return search(b.root, val)
}

func search(node *TreeNode, val int) bool {
	for node != nil {
		switch {
		case val < node.Val:
			node = node.Left
		case val > node.Val:
			node = node.Right
		default:
			return true
		}
	}
	return false
}

// Delete removes val from the tree. Returns false if val was not found.
func (b *BST) Delete(val int) bool {
	var found bool
	b.root, found = deleteNode(b.root, val)
	return found
}

func deleteNode(node *TreeNode, val int) (*TreeNode, bool) {
	if node == nil {
		return nil, false
	}
	if val < node.Val {
		var found bool
		node.Left, found = deleteNode(node.Left, val)
		return node, found
	}
	if val > node.Val {
		var found bool
		node.Right, found = deleteNode(node.Right, val)
		return node, found
	}

	// node matches val
	if node.Left == nil {
		return node.Right, true
	}
	if node.Right == nil {
		return node.Left, true
	}
	// two children: replace with in-order successor (min of right subtree)
	succ := node.Right
	for succ.Left != nil {
		succ = succ.Left
	}
	node.Val = succ.Val
	node.Right, _ = deleteNode(node.Right, succ.Val)
	return node, true
}

// InOrder returns values in ascending order (in-order traversal).
func (b *BST) InOrder() []int {
	out := make([]int, 0)
	inOrder(b.root, &out)
	return out
}

func inOrder(node *TreeNode, out *[]int) {
	if node == nil {
		return
	}
	inOrder(node.Left, out)
	*out = append(*out, node.Val)
	inOrder(node.Right, out)
}
