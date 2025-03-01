// Copyright 2019 the u-root Authors. All rights reserved
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rtc

import (
	"time"

	"golang.org/x/sys/unix"
)

// Read implements Read for the Linux RTC
func (r *RTC) Read() (time.Time, error) {
	rt, err := unix.IoctlGetRTCTime(int(r.file.Fd()))
	if err != nil {
		return time.Time{}, err
	}

	return time.Date(int(rt.Year)+1900,
		time.Month(rt.Mon+1),
		int(rt.Mday),
		int(rt.Hour),
		int(rt.Min),
		int(rt.Sec),
		0,
		time.UTC), nil
}

// Set implements Set for the Linux RTC
func (r *RTC) Set(tu time.Time) error {
	rt := unix.RTCTime{Sec: int32(tu.Second()),
		Min:   int32(tu.Minute()),
		Hour:  int32(tu.Hour()),
		Mday:  int32(tu.Day()),
		Mon:   int32(tu.Month() - 1),
		Year:  int32(tu.Year() - 1900),
		Wday:  int32(0),
		Yday:  int32(0),
		Isdst: int32(0)}

	return unix.IoctlSetRTCTime(int(r.file.Fd()), &rt)
}
