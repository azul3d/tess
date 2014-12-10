// Copyright 2014 The Tess Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tess

// ActiveRegion represents a active sweep region. For each pair of adjacent
// edges crossing the sweep line, there is an ActiveRegion to represent the
// region between them. The active regions are kept in sorted order in a
// dynamic dictionary. As the sweep line crosses each vertex, we update the
// affected regions.
type ActiveRegion struct {
	// TODO(bckenny): I *think* eUp and nodeUp could be passed in as constructor params

	// EUp is the upper edge of the region, directed right to left.
	EUp *GluHalfEdge

	// NodeUp is the dictionary node corresponding to the EUp edge.
	NodeUp *DictNode

	// WindingNumber is used to determine which regions are inside the polygon.
	WindingNumber int

	// Inside tells whether or not this region is inside the polygon.
	Inside bool

	// Sentinel marks fake edges at t = +/-infinity.
	Sentinel bool

	// Dirty marks regions where the upper or lower edge has changed, but we
	// haven't checked whether they intersect yet.
	Dirty bool

	// FixUpperEdge marks temporary edges introduced when we process a "right
	// vertex" (one without any edges leaving to the right).
	FixUpperEdge bool
}

// RegionBelow returns the ActiveRegion below this one.
func (r *ActiveRegion) RegionBelow() *ActiveRegion {
	return r.NodeUp.Prev.Key
}

// RegionAbove returns the ActiveRegion above this one.
func (r *ActiveRegion) RegionAbove() *ActiveRegion {
	return r.NodeUp.Next.Key
}
