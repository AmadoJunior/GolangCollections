package binary_tree

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func MakeBST(rootVal int) *TreeNode {
	return &TreeNode{val: rootVal}
}

func (t *TreeNode) Insert(value int) {
	if value < t.val {
		if t.left == nil {
			t.left = &TreeNode{val: value}
		} else {
			t.left.Insert(value)
		}
	} else if value > t.val {
		if t.right == nil {
			t.right = &TreeNode{val: value}
		} else {
			t.right.Insert(value)
		}
	}
}

func (t *TreeNode) Search(value int) *TreeNode {
	if t == nil || t.val == value {
		return t
	} else if t.val < value {
		return t.right.Search(value)
	} else {
		return t.left.Search(value)
	}
}

func (t *TreeNode) Min() int {
	if t.left == nil {
		return t.val
	} else {
		return t.left.Min()
	}
}

func (t *TreeNode) Max() int {
	if t.right == nil {
		return t.val
	} else {
		return t.right.Max()
	}
}

func (t *TreeNode) InOrderTraversal(slice *[]int) {
	if t.left != nil {
		t.left.InOrderTraversal(slice)
	}
	*slice = append(*slice, t.val)
	if t.right != nil {
		t.right.InOrderTraversal(slice)
	}
}

func (t *TreeNode) PreOrderTraversal(slice *[]int) {
	*slice = append(*slice, t.val)
	if t.left != nil {
		t.left.PreOrderTraversal(slice)
	}
	if t.right != nil {
		t.right.PreOrderTraversal(slice)
	}
}

func (t *TreeNode) PostOrderTraversal(slice *[]int) {
	if t.left != nil {
		t.left.PostOrderTraversal(slice)
	}
	if t.right != nil {
		t.right.PostOrderTraversal(slice)
	}
	*slice = append(*slice, t.val)
}
