package merkletree

import (
	"bytes"
	"crypto/sha256"
	"hash"
)

type Content interface {
	CalculateHash() ([]byte, error)
	Equals(other Content) (bool, error)
}

//
type Node struct {
	Tree   *MerkleTree
	Parent *Node
	Left   *Node
	Right  *Node
	leaf   bool
	dup    bool
	Hash   []byte
	C      Content
}

type MerkleTree struct {
	Root         *Node
	merkleRoot   []byte
	Leafs        []*Node
	hashStrategy func() hash.Hash
}

// 递归计算 对应节点hash
// 返回hash值， error
func (n *Node) verifyNode() ([]byte, error) {
	// 如果是叶子节点， 直接计算返回
	if n.leaf {
		return n.C.CalculateHash()
	}
	rightBytes, err := n.Right.verifyNode()
	if err != nil {
		return nil, err
	}

	leftBytes, err := n.Left.verifyNode()
	if err != nil {
		return nil, err
	}

	h := n.Tree.hashStrategy()
	if _, err := h.Write(append(leftBytes, rightBytes...)); err != nil {
		return nil, err
	}

	return h.Sum(nil), nil
}

func (n *Node) calculateNodeHash() ([]byte, error) {
	if n.leaf {
		return n.C.CalculateHash()
	}

	h := n.Tree.hashStrategy()
	if _, err := h.Write(append(n.Left.Hash, n.Right.Hash...)); err != nil {
		return nil, err
	}

	return h.Sum(nil), nil
}

func NewTree(cs []Content) (*MerkleTree, error) {
	var defaultHashStrategy = sha256.New
	return NewTreeWithHashStrategy(cs, defaultHashStrategy)
}

func NewTreeWithHashStrategy(cs []Content, hashStrategy func() hash.Hash) (*MerkleTree, error) {
	t := &MerkleTree{
		hashStrategy: hashStrategy,
	}
	root, leafs, err := buildWithContent(cs, t)
	if err != nil {
		return nil, err
	}
	t.Root = root
	t.Leafs = leafs
	t.merkleRoot = root.Hash
	return t, nil
}

func (m *MerkleTree) GetMerklePath(content Content) ([][]byte, []int64, error) {
	for _, current := range m.Leafs {
		ok, err := current.C.Equals(content)
		if err != nil {
			return nil, nil, err
		}
		if ok {
			currentParent := current.Parent
			var merklePath [][]byte
			var index []int64

			for currentParent != nil {
				if bytes.Equal(currentParent.Left.Hash, current.Hash) {
					merklePath = append(merklePath, currentParent.Right.Hash)
					index = append(index, 1)
				} else {
					merklePath = append(merklePath, currentParent.Left.Hash)
					index = append(index, 0)
				}

				current = currentParent
				currentParent = currentParent.Parent
			}
			return merklePath, index, nil
		}
	}

	return nil, nil, nil
}

func buildWithContent(cs []Content, t *MerkleTree) (*Node, []*Node, error) {

	return nil, nil, nil
}
