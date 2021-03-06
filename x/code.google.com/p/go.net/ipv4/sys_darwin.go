// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ipv4

import "syscall"

func init() {
	// Seems like kern.osreldate is veiled on latest OS X. We use
	// kern.osrelease instead.
	osver, err := syscall.Sysctl("kern.osrelease")
	if err != nil {
		return
	}
	var i int
	for i = range osver {
		if osver[i] != '.' {
			continue
		}
	}
	// The IP_PKTINFO was introduced in OS X 10.7 (Darwin
	// 11.0.0). See http://support.apple.com/kb/HT1633.
	if i>2 || i==2 && osver[0]>='1' && osver[1]>='1' {
		supportsPacketInfo = true
	}
}
