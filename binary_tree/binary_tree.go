package binary_tree

type Comparable interface {
	Less(other Comparable) bool
	More(other Comparable) bool
	Equal(other Comparable) bool
}

type TreeNode[T Comparable] struct {
	val   T
	left  *TreeNode[T]
	right *TreeNode[T]
}

func MakeBST[T Comparable](rootVal T) *TreeNode[T] {
	return &TreeNode[T]{val: rootVal}
}

func (t *TreeNode[T]) Insert(value T) {
	if value.Less(t.val) {
		if t.left == nil {
			t.left = &TreeNode[T]{val: value}
		} else {
			t.left.Insert(value)
		}
	} else if value.More(t.val) {
		if t.right == nil {
			t.right = &TreeNode[T]{val: value}
		} else {
			t.right.Insert(value)
		}
	}
}

func (t *TreeNode[T]) Search(value T) *TreeNode[T] {
	if t == nil || value.Equal(t.val) {
		return t
	} else if value.More(t.val) {
		return t.right.Search(value)
	} else {
		return t.left.Search(value)
	}
}

func (t *TreeNode[T]) Min() T {
	if t.left == nil {
		return t.val
	} else {
		return t.left.Min()
	}
}

func (t *TreeNode[T]) Max() T {
	if t.right == nil {
		return t.val
	} else {
		return t.right.Max()
	}
}

func (t *TreeNode[T]) InOrderTraversal(slice *[]T) {
	if t.left != nil {
		t.left.InOrderTraversal(slice)
	}
	*slice = append(*slice, t.val)
	if t.right != nil {
		t.right.InOrderTraversal(slice)
	}
}

func (t *TreeNode[T]) PreOrderTraversal(slice *[]T) {
	*slice = append(*slice, t.val)
	if t.left != nil {
		t.left.PreOrderTraversal(slice)
	}
	if t.right != nil {
		t.right.PreOrderTraversal(slice)
	}
}

func (t *TreeNode[T]) PostOrderTraversal(slice *[]T) {
	if t.left != nil {
		t.left.PostOrderTraversal(slice)
	}
	if t.right != nil {
		t.right.PostOrderTraversal(slice)
	}
	*slice = append(*slice, t.val)
}
