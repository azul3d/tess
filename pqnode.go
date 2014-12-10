// Copyright 2014 The Tess Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tess

// TODO(bckenny): maybe just have these created inline as literals (or unboxed
// directly - PQHandle is just an array index number).

// TODO(slimsag): This is far from idiomatic Go.

// PQNode represents a priority queue node.
type PQNode struct {
	Handle PQHandle
}

// Realloc allocates a PQNode array of the given size. If oldArray is not nil,
// it's contents are copied to the beginning of the new array. The rest of the
// array is filled with new PQNodes.
func (n *PQNode) Realloc(oldArray []*PQNode, size int) []*PQNode {
	newArray := make([]*PQNode, size)
	copy(newArray, oldArray)

	for i := len(oldArray); i < len(newArray); i++ {
		newArray[i] = new(PQNode)
	}
	return newArray
}
