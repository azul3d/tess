// Copyright 2014 The Tess Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tess

// DictNode is a doubly-linked-list node with a *ActiveRegion payload.
//
// The key for this node and the next and previous nodes in the parent Dict
// list can be provided to insert it into an existing list (or all can be
// omitted if this is to be the founding node of the list).
type DictNode struct {
	// Key is the *ActiveRegion key for this node, or nil if the head of the
	// list.
	Key *ActiveRegion

	// Pointers to the next and previous DictNode's in parent list or to self
	// if this is the first node.
	Next, Prev *DictNode
}

// NewDictNode returns a new and initialized *DictNode.
//
// If either next or prev nodes are nil, they are set to the returned node
// itself.
func NewDictNode(key *ActiveRegion, next, prev *DictNode) *DictNode {
	n := &DictNode{
		Key:  key,
		Next: next,
		Prev: prev,
	}
	if n.Next == nil {
		n.Next = n
	}
	if n.Prev == nil {
		n.Prev = n
	}
	return n
}
