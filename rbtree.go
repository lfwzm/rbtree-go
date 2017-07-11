package rbtree

import (
	"unsafe"
)

const (
	red   = true
	black = false
)

type RbNode struct {
	Parent  uintptr
	RbRight uintptr
	RbLeft  uintptr
}

type RbRoot struct {
	Node *RbNode
}

func NewRbNode() *RbNode {
	return &RbNode{0, 0, 0}
}

func NewRbRoot() *RbRoot {
	return &RbRoot{NewRbNode()}
}

func (n *RbNode) ParentNode() *RbNode {
	l := uintptr(3)
	return (*RbNode)(unsafe.Pointer(n.Parent & ^l))
}

func (n *RbNode) Color() bool {
	return ((n.Parent) & (1)) == 1
}

func (n *RbNode) SetBlack() {
	l := uintptr(1)
	n.Parent = n.Parent & (^l)
}

func (n *RbNode) SetRed() {
	n.Parent = n.Parent | 1
}

func (n *RbNode) SetParent(p *RbNode) {
	n.Parent = uintptr(unsafe.Pointer(p))
}

func (n *RbNode) RbRotateLeft(root *RbRoot) {
	if n == nil || root == nil {
		return
	}

	right := (*RbNode)(unsafe.Pointer(n.RbRight))
	prent := (*RbNode)(unsafe.Pointer(n.Parent))

	if n.RbRight = right.RbLeft; n.RbRight != 0 {
		(*RbNode)(unsafe.Pointer(right.RbLeft)).SetParent(n)
	}
	right.RbLeft = uintptr(unsafe.Pointer(n))

	(*RbNode)(unsafe.Pointer(right)).SetParent(prent)

	if prent != nil {
		if uintptr(unsafe.Pointer(n)) == prent.RbLeft {
			prent.RbLeft = uintptr(unsafe.Pointer(right))
		} else {
			prent.RbRight = uintptr(unsafe.Pointer(right))
		}
	} else {
		root.Node = right

	}
	n.SetParent(right)
}

func (n *RbNode) RbRotateRight(root *RbRoot) {
	if n == nil || root == nil {
		return
	}

	left := (*RbNode)(unsafe.Pointer(n.RbLeft))
	parent := (*RbNode)(unsafe.Pointer(n.Parent))

	if n.RbLeft = left.RbRight; n.RbLeft != 0 {
		(*RbNode)(unsafe.Pointer(left.RbRight)).SetParent(n)
	}
	left.RbRight = uintptr(unsafe.Pointer(n))
	left.SetParent(parent)

	if parent != nil {
		if parent.RbLeft == uintptr(unsafe.Pointer(n)) {
			parent.RbLeft = uintptr(unsafe.Pointer(left))
		} else {
			parent.RbRight = uintptr(unsafe.Pointer(left))
		}
	} else {
		root.Node = left
	}
	n.SetParent(left)
}

func (n *RbNode) RbInsertColor(root *RbRoot) {
	if n == nil || root == nil {
		return
	}

	var parent *RbNode
	var gparent *RbNode
	var optNode *RbNode = n
	parent = (*RbNode)(unsafe.Pointer(optNode.Parent))
	for parent.Color() == true {
		gparent = (*RbNode)(unsafe.Pointer(parent.Parent))
		//父节点为左节点
		if gparent.RbLeft == (uintptr)(unsafe.Pointer(parent)) {
			uncle := (*RbNode)(unsafe.Pointer(gparent.RbRight))
			if uncle != nil && uncle.Color() == true {
				parent.SetBlack()
				uncle.SetBlack()
				optNode = gparent
				parent = (*RbNode)(unsafe.Pointer(optNode.Parent))
				continue
			}
			//左右，左转
			if parent.RbRight == (uintptr)(unsafe.Pointer(optNode)) {
				parent.RbRotateLeft(root)
				parent, optNode = optNode, parent
			}

			parent.SetBlack()
			gparent.SetRed()
			gparent.RbRotateRight(root)

		} else {
			uncle := (*RbNode)(unsafe.Pointer(gparent.RbLeft))
			if uncle != nil && uncle.Color() == true {
				parent.SetBlack()
				uncle.SetBlack()
				optNode = gparent
				parent = (*RbNode)(unsafe.Pointer(optNode.Parent))
				continue
			}

			//左右，左转
			if parent.RbLeft == (uintptr)(unsafe.Pointer(optNode)) {
				parent.RbRotateRight(root)
				parent, optNode = optNode, parent
			}
			parent.SetBlack()
			gparent.SetRed()
			gparent.RbRotateLeft(root)
		}
		parent = (*RbNode)(unsafe.Pointer(optNode.Parent))
	}
	root.Node.SetBlack()
}
