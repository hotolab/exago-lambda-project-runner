// Do not edit. Bootstrap copy of /home/ceble/Work/Hotolab/go/src/cmd/compile/internal/gc/util.go

//line /home/ceble/Work/Hotolab/go/src/cmd/compile/internal/gc/util.go:1
// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gc

import (
	"os"
	"runtime"
	"runtime/pprof"
)

func (n *Node) Line() string {
	return Ctxt.LineHist.LineString(int(n.Lineno))
}

var atExitFuncs []func()

func atExit(f func()) {
	atExitFuncs = append(atExitFuncs, f)
}

func Exit(code int) {
	for i := len(atExitFuncs) - 1; i >= 0; i-- {
		f := atExitFuncs[i]
		atExitFuncs = atExitFuncs[:i]
		f()
	}
	os.Exit(code)
}

var (
	cpuprofile     string
	memprofile     string
	memprofilerate int64
	traceprofile   string
	traceHandler   func(string)
)

func startProfile() {
	if cpuprofile != "" {
		f, err := os.Create(cpuprofile)
		if err != nil {
			Fatalf("%v", err)
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			Fatalf("%v", err)
		}
		atExit(pprof.StopCPUProfile)
	}
	if memprofile != "" {
		if memprofilerate != 0 {
			runtime.MemProfileRate = int(memprofilerate)
		}
		f, err := os.Create(memprofile)
		if err != nil {
			Fatalf("%v", err)
		}
		atExit(func() {
			runtime.GC() // profile all outstanding allocations
			if err := pprof.WriteHeapProfile(f); err != nil {
				Fatalf("%v", err)
			}
		})
	}
	if traceprofile != "" && traceHandler != nil {
		traceHandler(traceprofile)
	}
}