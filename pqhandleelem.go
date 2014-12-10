// Copyright 2014 The Tess Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tess

// TODO(bckenny): more specific typing on key

// TODO(slimsag): This is far from idiomatic Go.

type PQHandleElem struct {
	// TODO(bckenny): if key could instead be an indexed into another store, makes heap storage a lot easier

	Key *PQKey

	// TODO(slimsag): was set to "0" in JS version, correct or not?

	Node *PQHandle
}

// Realloc allocates a PQHandleElem array of the given size. If oldArray is not
// nil, it's contents are copied to the beginning of the new array. The rest of
// the array is filled with new PQHandleElem.
func (n *PQHandleElem) Realloc(oldArray []*PQHandleElem, size int) []*PQHandleElem {
	newArray := make([]*PQHandleElem, size)
	copy(newArray, oldArray)

	for i := len(oldArray); i < len(newArray); i++ {
		newArray[i] = new(PQHandleElem)
	}
	return newArray
}
