// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mysync

import "runtime"

// WaitGroup is like sync.WaitGroup but records the callers of Add and Done.
type WaitGroup struct {
	Callers []uintptr
}

// Add records the callers of Add.
func (wg *WaitGroup) Add(x int) {
	wg.Callers = make([]uintptr, 32)
	n := runtime.Callers(1, wg.Callers)
	wg.Callers = wg.Callers[:n]
}

// Done records the callers of Done.
func (wg *WaitGroup) Done() {
	wg.Add(-1)
}
