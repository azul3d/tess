// Copyright 2014 The Tess Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tess

// TODO(slimsag): cleanup the HalfEdge documentation and ensure symbol names
// match their Go version.

// HalfEdge is the fundamental data structure. Two half-edges go together to
// make an edge, but they point in opposite directions. Each half-edge has a
// pointer to its mate (the "symmetric" half-edge sym), its origin vertex
// (org), the face on its left side (lFace), and the adjacent half-edges in the
// CCW direction around the origin vertex (oNext) and around the left face
// (lNext). There is also a "next" pointer for the global edge list (see
// below).
//
// The notation used for mesh navigation:
//
//  sym   = the mate of a half-edge (same edge, but opposite direction)
//  oNext = edge CCW around origin vertex (keep same origin)
//  dNext = edge CCW around destination vertex (keep same dest)
//  lNext = edge CCW around left face (dest becomes new origin)
//  rNext = edge CCW around right face (origin becomes new dest)
//
// "prev" means to substitute CW for CCW in the definitions above.
//
// The circular edge list is special; since half-edges always occur in pairs (e
// and e.sym), each half-edge stores a pointer in only one direction. Starting
// at eHead and following the e.next pointers will visit each *edge* once (ie.
// e or e.sym, but not both). e.sym stores a pointer in the opposite direction,
// thus it is always true that e.sym.next.sym.next === e.
type GluHalfEdge struct {
	// TODO(bckenny): are these the right defaults? (from gl_meshNewMesh requirements)

	// Next is a pointer to the next edge.
	//
	//  Prev == Sym.Next
	//
	Next *GluHalfEdge

	// TODO(bckenny): how can these be required if created in pairs? move to factory
	// creation only?

	// Sym is the same edge but in the opposite direction.
	Sym *GluHalfEdge

	// ONext is the next edge CCW around origin.
	ONext *GluHalfEdge

	// LNext is the next edge CCW around the left face.
	LNext *GluHalfEdge

	// Org is the origin vertex (OVertex too long).
	Org *GluVertex

	// LFace is the left face.
	LFace *GluFace

	// activeRegion is a region with this upper edge (see sweep code).
	activeRegion *ActiveRegion

	// winding is the change in winding number when crossing from the right
	// face to the left face.
	winding int
}

// NewGluHalfEdge returns a new and initialized *GluHalfEdge.
//
// If the next half-edge is nil, it is set to the returned half-edge itself.
func NewGluHalfEdge(next *GluHalfEdge) *GluHalfEdge {
	e := &GluHalfEdge{
		Next: next,
	}
	if e.Next == nil {
		e.Next = e
	}
	return e
}

func (e *GluHalfEdge) RFace() *GluFace {
	return e.Sym.LFace
}

func (e *GluHalfEdge) Dst() *GluVertex {
	return e.Sym.Org
}

func (e *GluHalfEdge) OPrev() *GluHalfEdge {
	return e.Sym.LNext
}

func (e *GluHalfEdge) LPrev() *GluHalfEdge {
	return e.ONext.Sym
}

// NOTE(bckenny): GluHalfEdge.DPrev is called nowhere in libtess and isn't part
// of the current public API. It could be useful for mesh traversal and
// manipulation if made public, however.

// DPrev returns the edge clockwise around destination vertex (keep same dest).
func (e *GluHalfEdge) DPrev() *GluHalfEdge {
	return e.LNext.Sym
}

func (e *GluHalfEdge) RPrev() *GluHalfEdge {
	return e.Sym.ONext
}

func (e *GluHalfEdge) DNext() *GluHalfEdge {
	return e.RPrev().Sym
}

// NOTE(bckenny): GluHalfEdge.RNext is called nowhere in libtess and isn't part
// of the current public API. It could be useful for mesh traversal and
// manipulation if made public, however.

func (e *GluHalfEdge) RNext() *GluHalfEdge {
	return e.OPrev().Sym
}
