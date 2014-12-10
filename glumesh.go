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

	// TODO(slimsag): fix this.
	/*
		var(
			fHead = g.FHead
			vHead = g.VHead
			eHead = g.EHead
			e *GluHalfEdge
		)

		// Faces.
		var(
			f *GluFace
			fPrev = fHead
		)
	*/
	/*
		for fPrev = fHead; (f = fPrev.Next) != fHead; fPrev = f {
			libtess.assert(f.prev === fPrev);
			e = f.anEdge;
			do {
				libtess.assert(e.sym !== e);
				libtess.assert(e.sym.sym === e);
				libtess.assert(e.lNext.oNext.sym === e);
				libtess.assert(e.oNext.sym.lNext === e);
				libtess.assert(e.lFace === f);
				e = e.lNext;
			} while (e !== f.anEdge);
		}
		libtess.assert(f.prev === fPrev && f.anEdge === null && f.data === null);
	*/

	// Vertices.
	/*
		var(
			v *GluVertex
			vPrev = vHead
		)
	*/
	/*
		for (vPrev = vHead; (v = vPrev.next) !== vHead; vPrev = v) {
			libtess.assert(v.prev === vPrev);
			e = v.anEdge;
			do {
				libtess.assert(e.sym !== e);
				libtess.assert(e.sym.sym === e);
				libtess.assert(e.lNext.oNext.sym === e);
				libtess.assert(e.oNext.sym.lNext === e);
				libtess.assert(e.org === v);
				e = e.oNext;
			} while (e !== v.anEdge);
		}
		libtess.assert(v.prev === vPrev && v.anEdge === null && v.data === null);
	*/

	// Edges.
	/*
		ePrev := eHead
		  for (ePrev = eHead; (e = ePrev.next) !== eHead; ePrev = e) {
			libtess.assert(e.sym.next === ePrev.sym);
			libtess.assert(e.sym !== e);
			libtess.assert(e.sym.sym === e);
			libtess.assert(e.org !== null);
			libtess.assert(e.dst() !== null);
			libtess.assert(e.lNext.oNext.sym === e);
			libtess.assert(e.oNext.sym.lNext === e);
		  }
		  libtess.assert(e.sym.next === ePrev.sym &&
			  e.sym === this.eHeadSym &&
			  e.sym.sym === e &&
			  e.org === null && e.dst() === null &&
			  e.lFace === null && e.rFace() === null);
	*/
}
