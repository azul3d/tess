// Copyright 2014 The Tess Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tess

type PriorityQHeap struct{}

func NewPriorityQHeap(leq func(a, b *PQKey) bool) *PriorityQHeap {
	return nil
}

func (p *PriorityQHeap) Init() {
}

func (p *PriorityQHeap) DeleteHeap() {
}

func (p *PriorityQHeap) Insert(keyNew *PQKey) PQHandle {
	return PQHandle(0)
}

func (p *PriorityQHeap) ExtractMin() *PQKey {
	return nil
}
