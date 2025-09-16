// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Malloc small size classes.
//
// See malloc.go for overview.
// See also mksizeclasses.go for how we decide what size classes to use.

package runtime

// Returns size of the memory block that mallocgc will allocate if you ask for the size.
// 内存对齐函数
func roundupsize(size uintptr) uintptr {
	// Small sizes are rounded up to a multiple of 8. 小的内存块按8字节对齐
	if size < _MaxSmallSize {
		if size <= smallSizeMax-8 {
			return uintptr(class_to_size[size_to_class8[divRoundUp(size, smallSizeDiv)]])
		} else {
			return uintptr(class_to_size[size_to_class128[divRoundUp(size-smallSizeMax, largeSizeDiv)]])
		}
	}
	// Large sizes are rounded up to a multiple of _PageSize. 大的内存块按页大小对齐
	if size+_PageSize < size {
		return size
	}
	return alignUp(size, _PageSize)
}
