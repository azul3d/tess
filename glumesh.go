// Copyright 2014 The Tess Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tess

// GluMesh is a mesh type. Create one using NewGluMesh.
type GluMesh struct {
	// VHead is a dummy header for the vertex list.
	VHead *GluVertex

	// FHead is a dummy header for the face list.
	FHead *GluFace

	// EHead is a dummy header for the edge list.
	EHead *GluHalfEdge

	// EHeadSym is EHead's symmetric counterpart.
	EHeadSym *GluHalfEdge
}

// NewGluMesh returns a new and initialized *GluMesh structure.
//
// It has no edges, no vertices, and no loops (what we usually call a "face").
func NewGluMesh() *GluMesh {
	m := &GluMesh{
		VHead:    &GluVertex{},
		FHead:    &GluFace{},
		EHead:    &GluHalfEdge{},
		EHeadSym: &GluHalfEdge{},
	}

	// Pair the half edge head and it's symmetrical counterpart together.
	m.EHead.Sym = m.EHeadSym
	m.EHeadSym.Sym = m.EHead
	return m
}

// assert is a panic-causing assertion:
//
//  assert(a != b, "a != b")
//
func assert(cond bool, val string) {
	if !cond {
		panic(val)
	}
}

// Check checks this mesh for self-consistency.
func (g *GluMesh) Check() {
	if !DEBUG {
		return
	}

	var (
		fHead = g.FHead
		vHead = g.VHead
		eHead = g.EHead
		e     *GluHalfEdge
	)

	// Faces.
	var (
		f     *GluFace
		fPrev = fHead
	)
	for {
		f = fPrev.Next
		if !(f != fHead) {
			break
		}

		assert(f.Prev == fPrev, "f.Prev == fPrev")
		e = f.AnEdge
		for {
			assert(e.Sym != e, "e.Sym != e")
			assert(e.Sym.Sym == e, "e.Sym.Sym == e")
			assert(e.LNext.ONext.Sym == e, "e.LNext.ONext.Sym == e")
			assert(e.ONext.Sym.LNext == e, "e.ONext.Sym.LNext == e")
			assert(e.LFace == f, "e.LFace == f")
			e = e.LNext
			if !(e != f.AnEdge) {
				break
			}
		}

		fPrev = f
	}
	assert(f.Prev == fPrev && f.AnEdge == nil && f.Data == nil, "f.Prev == fPrev && f.AnEdge == nil && f.Data == nil")

	// Vertices.
	var (
		v     *GluVertex
		vPrev = vHead
	)
	for {
		v = vPrev.Next
		if !(v != vHead) {
			break
		}

		assert(v.Prev == vPrev, "v.Prev == vPrev")
		e = v.AnEdge
		for {
			assert(e.Sym != e, "e.Sym != e")
			assert(e.Sym.Sym == e, "e.Sym.Sym == e")
			assert(e.LNext.ONext.Sym == e, "e.LNext.ONext.Sym == e")
			assert(e.ONext.Sym.LNext == e, "e.ONext.Sym.LNext == e")
			assert(e.Org == v, "e.Org == v")
			e = e.ONext
			if !(e != v.AnEdge) {
				break
			}
		}

		vPrev = v
	}
	assert(v.Prev == vPrev && v.AnEdge == nil && v.Data == nil, "v.Prev == vPrev && v.AnEdge == nil && v.Data == nil")

	// Edges.
	ePrev := eHead
	for {
		e = ePrev.Next
		if !(e != eHead) {
			break
		}

		assert(e.Sym.Next == ePrev.Sym, "e.Sym.Next == ePrev.Sym")
		assert(e.Sym != e, "e.Sym != e")
		assert(e.Sym.Sym == e, "e.Sym.Sym == e")
		assert(e.Org != nil, "e.Org != nil")
		assert(e.Dst() != nil, "e.Dst() != nil")
		assert(e.LNext.ONext.Sym == e, "e.LNext.ONext.Sym == e")
		assert(e.ONext.Sym.LNext == e, "e.ONext.Sym.LNext == e")

		ePrev = e
	}
	assert(e.Sym.Next == ePrev.Sym && e.Sym == g.EHeadSym && e.Sym.Sym == e && e.Org == nil && e.Dst() == nil && e.LFace == nil && e.RFace() == nil, "")
}
