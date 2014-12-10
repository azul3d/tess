// Copyright 2014 The Tess Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tess

// TODO(bckenny): keys appear to always be GluVertex in this case?

const PriorityQHeapInitSize = 32

type PriorityQHeap struct {
	// Nodes is the heap itself. Active nodes are stored in the range 1..size.
	// Each node stores only an index into handles.
	nodes []*PQNode

	// Handles is a list of PQHandleElem, each handle stores a key plus a
	// pointer back to the node which currently represents that key:
	//
	//  nodes[handles[i].node].handle == i
	//
	handles []*PQHandleElem

	// TODO(bckenny): size and max should probably be libtess.PQHandle for
	// correct typing (see PriorityQ.js)

	// size is the size of the queue.
	size int

	// Max is the queue's current allocated space.
	max int

	// FreeList is the index of the next free hold in the handles array. Handle
	// in that slot has next item in freeList in it's node propert. If there
	// are no holes, FreeList == 0 and one at the end of handles must be used.
	FreeList PQHandle

	// Initialized indicates that the heap has been initialized via Init. If
	// false, inserts are fast insertions at the end of a list. If true, all
	// inserts will now be correctly ordered in the queue before rendering.
	Initialized bool

	// TODO(bckenny): leq was inlined by define in original, but appears to
	// be vertLeq, as passed. Using injected version, but is it better just to
	// manually inline?
	leq func(a, b *PQKey) bool
}

func NewPriorityQHeap(leq func(a, b *PQKey) bool) *PriorityQHeap {
	h := &PriorityQHeap{
		nodes:   PQNodeRealloc(nil, PriorityQHeapInitSize+1),
		handles: PQHandleElemRealloc(nil, PriorityQHeapInitSize+1),
		max:     PriorityQHeapInitSize,
	}

	// So that minimum returns nil.
	h.nodes[1].Handle = 1
	return h
}

// Initializing ordering of the heap. Must be called before any method other
// than insert is called to ensure correctness when removing or querying.
func (h *PriorityQHeap) Init() {
	// This method of building a heap is O(n), rather than O(n lg n).
	for i := h.size; i >= 1; i-- {
		h.floatDown(PQHandle(i))
	}
	h.Initialized = true
}

func (h *PriorityQHeap) DeleteHeap() {
	// TODO(bckenny): unnecessary, I think.
	h.handles = nil
	h.nodes = nil
	// NOTE(bckenny): nulled at callsite in PriorityQ.deleteQ
}

// Insert inserts a new key into the heap. It returns a handle that can be used
// to remove the key.
func (h *PriorityQHeap) Insert(keyNew *PQKey) PQHandle {
	return PQHandle(0)

	// TODO(slimsag): fix this.
	/*
	  var curr = ++this.size_;

	  // if the heap overflows, double its size.
	  if ((curr * 2) > this.max_) {
	    this.max_ *= 2;
	    this.nodes_ = libtess.PQNode.realloc(this.nodes_, this.max_ + 1);
	    this.handles_ = libtess.PQHandleElem.realloc(this.handles_, this.max_ + 1);
	  }

	  var free;
	  if (this.freeList_ === 0) {
	    free = curr;
	  } else {
	    free = this.freeList_;
	    this.freeList_ = this.handles_[free].node;
	  }

	  this.nodes_[curr].handle = free;
	  this.handles_[free].node = curr;
	  this.handles_[free].key = keyNew;

	  if (this.initialized_) {
	    this.floatUp_(curr);
	  }

	  return free;
	*/
}

// IsEmpty tells whether the heap is empty.
func (h *PriorityQHeap) IsEmpty() bool {
	return h.size == 0
}

// Minimum returns the minimum key in the heap. if the heap is empty, nil will
// be returned.
func (h *PriorityQHeap) Minimum() *PQKey {
	return h.handles[h.nodes[1].Handle].Key
}

// ExtractMin removes the minimum key from the heap and returns it. If the heap
// is empty, nil will be returned.
func (h *PriorityQHeap) ExtractMin() *PQKey {
	var (
		n    = h.nodes
		h2   = h.handles
		hMin = n[1].Handle
		min  = h2[hMin].Key
	)

	if h.size > 0 {
		n[1].Handle = n[h.size].Handle
		h2[n[1].Handle].Node = PQHandle(1)

		h2[hMin].Key = nil
		h2[hMin].Node = h.FreeList
		h.FreeList = hMin

		// TODO(slimsag): fix this.
		/*
			if (--h.size > 0) {
			  h.floatDown(1);
			}
		*/
	}

	return min
}

// Remove removes the key associated with handle hCurr (returned from Insert)
// from heap.
func (h *PriorityQHeap) Remove(hCurr PQHandle) {
	var (
		//n = h.nodes
		h2 = h.handles
	)
	assert(hCurr >= 1 && int(hCurr) <= h.max && h2[hCurr].Key != nil, "hCurr >= 1 && int(hCurr) <= h.max && h2[hCurr].Key != nil")

	// TODO(slimsag): fix this.
	/*
	  var curr = h[hCurr].node;
	  n[curr].handle = n[this.size_].handle;
	  h[n[curr].handle].node = curr;

	  if (curr <= --this.size_) {
	    if (curr <= 1 ||
	        this.leq_(h[n[curr >> 1].handle].key, h[n[curr].handle].key)) {

	      this.floatDown_(curr);
	    } else {
	      this.floatUp_(curr);
	    }
	  }

	  h[hCurr].key = null;
	  h[hCurr].node = this.freeList_;
	  this.freeList_ = hCurr;
	*/
}

func (h *PriorityQHeap) floatDown(curr PQHandle) {
	// TODO(slimsag): fix this.
	/*
	  var n = this.nodes_;
	  var h = this.handles_;

	  var hCurr = n[curr].handle;
	  for (;;) {
	    // The children of node i are nodes 2i and 2i+1.
	    // set child to the index of the child with the minimum key
	    var child = curr << 1;
	    if (child < this.size_ &&
	        this.leq_(h[n[child + 1].handle].key, h[n[child].handle].key)) {

	      ++child;
	    }

	    libtess.assert(child <= this.max_);

	    var hChild = n[child].handle;
	    if (child > this.size_ || this.leq_(h[hCurr].key, h[hChild].key)) {
	      n[curr].handle = hCurr;
	      h[hCurr].node = curr;
	      break;
	    }
	    n[curr].handle = hChild;
	    h[hChild].node = curr;
	    curr = child;
	  }
	*/
}

func (h *PriorityQHeap) floatUp(curr PQHandle) {
	// TODO(slimsag): fix this.
	/*
	  var n = this.nodes_;
	  var h = this.handles_;

	  var hCurr = n[curr].handle;
	  for (;;) {
	    var parent = curr >> 1;
	    var hParent = n[parent].handle;
	    if (parent === 0 || this.leq_(h[hParent].key, h[hCurr].key)) {
	      n[curr].handle = hCurr;
	      h[hCurr].node = curr;
	      break;
	    }

	    n[curr].handle = hParent;
	    h[hParent].node = curr;
	    curr = parent;
	  }
	*/
}
