// Copyright 2014 The Tess Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tess

// GluFace is a circularly linked-list face.
//
// Each face has a pointer to the next and previous faces in the circular list,
// and a pointer to a half-edge with this face as the left face (null if this
// is the dummy header). There is also a field "data" for client data.
type GluFace struct {
	// Pointers to the next and previous faces.
	Next, Prev *GluFace

	// AnEdge is a half-edge with this left face.
	AnEdge *GluHalfEdge

	// Data is the client's own data.
	//
	// TODO(slimsag): can we eliminate this?
	Data interface{}

	// Inside tells whether or not this face is in the polygon interior.
	Inside bool
}

// NewGluFace returns a new and initialized *GluFace.
//
// If either next or prev faces are nil, they are set to the returned face
// itself.
func NewGluFace(next, prev *GluFace) *GluFace {
	n := &GluFace{
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
