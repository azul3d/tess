// Copyright 2014 The Tess Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tess

// GluVertex is a circularly linked-list vertex in 3D space.
//
// Each vertex has a pointer to next and previous vertices in the circular
// list, and a pointer to a half-edge with this vertex as the origin (null if
// this is the dummy header). There is also a field "data" for client data.
type GluVertex struct {
	// Next and previous vertex pointers.
	Next, Prev *GluVertex

	// AnEdge is a half-edge with this origin.
	AnEdge *GluHalfEdge

	// Data is the client's own data.
	//
	// TODO(slimsag): can we eliminate this?
	Data interface{}

	// The vertex location in 3D.
	Coords [3]float32

	// Components of projection onto the sweep plane.
	S, T float32

	// To allow deletion from priority queue.
	PQHandle PQHandle
}

// NewGluVertex returns a new and initialized *GluVertex.
//
// If either next or prev vertex are nil, they are set to the returned vertex
// itself.
func NewGluVertex(next, prev *GluVertex) *GluVertex {
	v := &GluVertex{
		Next: next,
		Prev: prev,
	}
	if v.Next == nil {
		v.Next = v
	}
	if v.Prev == nil {
		v.Prev = v
	}
	return v
}
