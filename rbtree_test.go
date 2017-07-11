package rbtree

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestParentNode(t *testing.T) {
	parent := NewRbNode()
	son := NewRbNode()
	son.Parent = uintptr(unsafe.Pointer(parent))
	sonp := son.ParentNode()
	if parent != sonp {
		t.Fail()
	}
}

func TestSetBlack(t *testing.T) {
	parent := NewRbNode()
	son := NewRbNode()
	son.Parent = uintptr(unsafe.Pointer(parent))
	son.SetBlack()
	sonp := son.ParentNode()
	if parent != sonp {
		t.Fail()
	}
}

func TestSetRed(t *testing.T) {
	parent := NewRbNode()
	son := NewRbNode()
	son.Parent = uintptr(unsafe.Pointer(parent))
	son.SetRed()
	fmt.Println(son.Parent)
	sonp := son.ParentNode()
	fmt.Println(uintptr(unsafe.Pointer(sonp)))
	if parent != sonp {
		t.Fail()
	}

	if son.Color() != red {
		t.Fail()
	}
}

func TestRbRotateLeft(t *testing.T) {
	parent := NewRbNode()
	node := NewRbNode()
	right := NewRbNode()
	left := NewRbNode()
	rightson := NewRbNode()
	leftson := NewRbNode()

	parent.RbLeft = uintptr(unsafe.Pointer(node))
	node.Parent = uintptr(unsafe.Pointer(parent))
	node.RbLeft = uintptr(unsafe.Pointer(left))
	node.RbRight = uintptr(unsafe.Pointer(right))
	left.Parent = uintptr(unsafe.Pointer(node))
	right.Parent = uintptr(unsafe.Pointer(node))
	rightson.SetParent(right)
	leftson.SetParent(right)
	right.RbLeft = uintptr(unsafe.Pointer(leftson))
	right.RbRight = uintptr(unsafe.Pointer(rightson))
	root := RbRoot{parent}
	node.RbRotateLeft(&root)

	if parent.RbLeft != uintptr(unsafe.Pointer(right)) {
		t.Fatal("right is not parent son")
	}

	if uintptr(unsafe.Pointer(right)) != node.Parent {
		t.Fatal("right is not node's parent")
	}

	if uintptr(unsafe.Pointer(leftson)) != node.RbRight {
		t.Fatal("right's right is nod node's left")
	}
}

func TestRbRotateLeftRoot(t *testing.T) {
	rootNode := NewRbNode()
	right := NewRbNode()
	rootNode.RbRight = uintptr(unsafe.Pointer(right))

	root := RbRoot{rootNode}

	rootNode.RbRotateLeft(&root)

	if root.Node.RbLeft != uintptr(unsafe.Pointer(rootNode)) {
		t.Fatal("root node error")
	}

	if rootNode.Parent != uintptr(unsafe.Pointer(root.Node)) {
		t.Fatal("node 's parent is no root")
	}
}

//发现个问题，转换的时候想当然的将子节点当做是已经是一个OK的指针了。
func TestRbRotateRightRoot(t *testing.T) {
	rootNode := NewRbNode()
	left := NewRbNode()
	rootNode.RbLeft = uintptr(unsafe.Pointer(left))

	root := RbRoot{rootNode}

	rootNode.RbRotateRight(&root)

	if root.Node.RbRight != uintptr(unsafe.Pointer(rootNode)) {
		t.Fatal("root node error")
	}

	if rootNode.Parent != uintptr(unsafe.Pointer(root.Node)) {
		t.Fatal("node 's parent is no root")
	}
}

func TestRbRotateRight(t *testing.T) {
	parent := NewRbNode()
	node := NewRbNode()
	right := NewRbNode()
	left := NewRbNode()
	rightson := NewRbNode()
	leftson := NewRbNode()

	parent.RbLeft = uintptr(unsafe.Pointer(node))
	node.Parent = uintptr(unsafe.Pointer(parent))
	node.RbLeft = uintptr(unsafe.Pointer(left))
	node.RbRight = uintptr(unsafe.Pointer(right))
	left.Parent = uintptr(unsafe.Pointer(node))
	right.Parent = uintptr(unsafe.Pointer(node))
	rightson.SetParent(left)
	leftson.SetParent(left)
	left.RbLeft = uintptr(unsafe.Pointer(leftson))
	left.RbRight = uintptr(unsafe.Pointer(rightson))
	root := RbRoot{parent}
	node.RbRotateRight(&root)

	if parent.RbLeft != uintptr(unsafe.Pointer(left)) {
		t.Fatal("right is not parent son")
	}

	if uintptr(unsafe.Pointer(left)) != node.Parent {
		t.Fatal("right is not node's parent")
	}

	if uintptr(unsafe.Pointer(rightson)) != node.RbLeft {
		t.Fatal("right's right is nod node's left")
	}
}
