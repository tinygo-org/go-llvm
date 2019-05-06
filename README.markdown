# Go bindings to system LLVM

This library provides bindings to a system-installed LLVM.

Currently supported:

  * LLVM 8 from [apt.llvm.org](http://apt.llvm.org/) on Debian.

## Usage

If you have a supported LLVM installation, you should be able to do a simple `go get`:

    go get github.com/tinygo-org/go-llvm

After downloading, you may have to update the configuration for your system:

    make config

If you built your own LLVM, you can also use that:

    # update LLVM files
    make update SRCDIR=<dir>
    
    # configure this LLVM build
    make config BUILDDIR=<builddir>

Note that you may have to comment out parts of `backports.go` if you use a
newer version of LLVM.

## License

Most of the files are extracted from the LLVM source tree, specifically all
\*.go, \*.cpp, and \*.h files come directly from
[bindings/go/llvm](https://github.com/llvm-mirror/llvm/tree/release_80/bindings/go/llvm)
in the LLVM source tree. They are all released under the [University of
Illinois/NCSA Open Source License](http://releases.llvm.org/8.0.0/LICENSE.TXT)
which is a permissive BSD-like license. Check upstream LLVM for detailed
copyright information.

This README, the backports\* files, and the Makefile are separate from LLVM but
are licensed under the same license.
