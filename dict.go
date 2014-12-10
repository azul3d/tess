// Copyright 2014 The Tess Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tess

// Dict is a list of edges crossing the sweep line, sorted from top to bottom.
//
// The implementation is a doubly-linked list, sorted by the injected edgeLeq
// comparator function. Here it is simple ordering, but see sweep for the list
// of invariants on the edge dictionary this ordering creates.
type Dict struct {
	// head is the head of the doubly-linked DictNode list. At creation time,
	// links back and forward only to itself.
	head *DictNode

	// frame is the tesselator used as the frame for edge/event comparisons.
	frame *GluTesselator

	// leq is the comparison function to maintain the invariants of the Dict.
	// See edgeLeq for source.
	leq func(f *GluTesselator, a, b *ActiveRegion) bool
}

// NewDict returns a new and initialized *Dict using the given frame and leq
// comparator function.
func NewDict(frame *GluTesselator, leq func(f *GluTesselator, a, b *ActiveRegion) bool) *Dict {
	return &Dict{
		head:  NewDictNode(nil, nil, nil),
		frame: frame,
		leq:   leq,
	}
}

// InsertBefore inserts the supplied key into the edge list and returns it's
// new node.
func (d *Dict) InsertBefore(node *DictNode, key *ActiveRegion) *DictNode {
	// TODO(slimsag): fix this.
	return nil

	/*
	  do {
	    node = node.prev;
	  } while (node.key !== null && !this.leq_(this.frame_, node.key, key));

	  // insert the new node and update the surrounding nodes to point to it
	  var newNode = new libtess.DictNode(key, node.next, node);
	  node.next.prev = newNode;
	  node.next = newNode;

	  return newNode;
	*/
}

// Insert inserts the given key into the dict and returns the new node that
// contains it.
func (d *Dict) Insert(key *ActiveRegion) *DictNode {
	return d.InsertBefore(d.head, key)
}

// DeleteNode removes the given node from the list.
func (d *Dict) DeleteNode(node *DictNode) {
	node.Next.Prev = node.Prev
	node.Prev.Next = node.Next
}

// Search returns the node with the smallest key greater than or equal to the
// given key. If there is no such key, it returns a node whose key is nil.
// Similarly, max(d).Next has a nil key, etc.
func (d *Dict) Search(key *ActiveRegion) *DictNode {
	// TODO(slimsag): fix this.
	return nil

	/*
		  var node = this.head_;

		  do {
			node = node.next;
		  } while (node.key !== null && !this.leq_(this.frame_, key, node.key));

		  return node;
	*/
}

// Min returns the node with the smallest key.
func (d *Dict) Min() *DictNode {
	return d.head.Next
}

// NOTE(bckenny): libtess.Dict.getMax isn't called within libtess and isn't part
// of the public API. For now, leaving in but ignoring for coverage.

// Max returns the node with the greatest key.
func (d *Dict) Max() *DictNode {
	return d.head.Prev
}
