// Copyright 2014 The Tess Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tess

const PriorityQInitSize = 32

type PQKey int

type PriorityQ struct {
	keys []*PQKey

	// Array of indexes into keys.
	order []int

	size        int
	max         int
	initialized bool

	// TODO(bckenny): leq was inlined by define in original, but appears to
	// just be vertLeq, as passed. keep an eye on this as to why its not used.
	leq func(a, b *PQKey) bool

	heap *PriorityQHeap
}

func NewPriorityQ(leq func(a, b *PQKey) bool) *PriorityQ {
	return &PriorityQ{
		keys: PQKeyRealloc(nil, PriorityQInitSize),
		max:  PriorityQInitSize,
		heap: NewPriorityQHeap(leq),
	}
}

func (p *PriorityQ) DeleteQ() {
	// TODO(bckenny): unnecessary, I think.
	p.heap.DeleteHeap()
	p.heap = nil
	p.order = nil
	p.keys = nil
	// NOTE(bckenny): nulled at callsite (sweep.donePriorityQ_)
}

func (p *PriorityQ) Init() {
	// TODO(bckenny): reuse. in theory, we don't have to empty this, as access is
	// dictated by this.size_, but array.sort doesn't know that
	p.order = nil

	// Create an array of indirect pointers to the keys, so that
	// the handles we have returned are still valid.
	//
	// TODO(bckenny): valid for when? it appears we can just store indexes into
	// keys, but what did this mean?
	for i := 0; i < p.size; i++ {
		p.order[i] = i
	}

	// sort the indirect pointers in descending order of the keys themselves
	// TODO(bckenny): make sure it's ok that keys[a] === keys[b] returns 1
	// TODO(bckenny): unstable sort means we may get slightly different polys in
	// different browsers, but only when passing in equal points
	// TODO(bckenny): make less awkward closure?

	// TODO(slimsag): fix this.
	/*
		  var comparator = (function(keys, leq) {
			return function(a, b) {
			  return leq(keys[a], keys[b]) ? 1 : -1;
			};
		  })(this.keys_, this.leq_);
		  this.order_.sort(comparator);
	*/

	p.max = p.size
	p.initialized = true
	p.heap.Init()

	if DEBUG {
		var r = p.size - 1
		for i := 0; i < r; i++ {
			assert(p.leq(p.keys[p.order[i+1]], p.keys[p.order[i]]), "p.leq(p.keys[p.order[i + 1]], p.keys[p.order[i]])")
		}
	}
}

func (p *PriorityQ) Insert(keyNew *PQKey) PQHandle {
	// NOTE(bckenny): originally returned LONG_MAX as alloc failure signal. no
	// longer does.
	if p.initialized {
		return p.heap.Insert(keyNew)
	}

	var curr = p.size
	/*
		// TODO(slimsag): fix this.
		if (++p.size >= p.max) {
			// If the heap overflows, double its size.
			p.max *= 2
			p.keys = PQKeyRealloc(p.keys, p.max);
		}
	*/

	p.keys[curr] = keyNew

	// Negative handles index the sorted array.
	return PQHandle(-(curr + 1))
}

// PQKeyRealloc allocates a PQKey array of the given size. If oldArray is not
// nil, its contents are copied to the beginning of the new array. The rest of
// the array is filled with nil.
func PQKeyRealloc(oldArray []*PQKey, size int) []*PQKey {
	newArray := make([]*PQKey, size)
	copy(newArray, oldArray)

	for i := len(oldArray); i < len(newArray); i++ {
		newArray[i] = new(PQKey)
	}
	return newArray

}

func (p *PriorityQ) ExtractMin() *PQKey {
	if p.size == 0 {
		return p.heap.ExtractMin()
	}

	// TODO(slimsag): fix this.
	return nil

	/*
		var sortMin = this.keys_[this.order_[this.size_ - 1]];
		if (!this.heap_.isEmpty()) {
			var heapMin = this.heap_.minimum();
			if (this.leq_(heapMin, sortMin)) {
				return this.heap_.extractMin();
			}
		}

		do {
			--this.size_;
		} while (this.size_ > 0 && this.keys_[this.order_[this.size_ - 1]] === null);

		return sortMin;
	*/
}

/*
 * [minimum description]
 * @return {libtess.PQKey} [description].

libtess.PriorityQ.prototype.minimum = function() {
  if (this.size_ === 0) {
    return this.heap_.minimum();
  }

  var sortMin = this.keys_[this.order_[this.size_ - 1]];
  if (!this.heap_.isEmpty()) {
    var heapMin = this.heap_.minimum();
    if (this.leq_(heapMin, sortMin)) {
      return heapMin;
    }
  }

  return sortMin;
};


 * [remove description]
 * @param {libtess.PQHandle} curr [description].

libtess.PriorityQ.prototype.remove = function(curr) {
  if (curr >= 0) {
    this.heap_.remove(curr);
    return;
  }
  curr = -(curr + 1);

  libtess.assert(curr < this.max_ && this.keys_[curr] !== null);

  this.keys_[curr] = null;
  while (this.size_ > 0 && this.keys_[this.order_[this.size_ - 1]] === null) {
    --this.size_;
  }
};
*/
