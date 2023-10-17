package blockchain

import "testing"

func TestNewProof(t *testing.T) {
	bl := CreateBlock("test", []byte{})
	target := NewProof(bl)
	if target == nil {
		t.Error("NewProof failed")
	}
	// assert t == true
	t.Log("NewProof passed")

}
