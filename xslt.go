package xslt

// #include <stdlib.h>
// #include <string.h>
// #include "xslt.h"
// #cgo CFLAGS: -I ./  -I ./libs
// #cgo CXXFLAGS: -I ./ -I ./libs
// #cgo linux LDFLAGS:  -L ./libs -Wl,--start-group -lcxslt -lstdc++ -lm -pthread -ldl -lxslt  -lxml2 -lzlib -licudata -licudt -licui18n -licutools -licuuc -Wl,--end-group
// #cgo windows LDFLAGS: -L ./libs -lcxslt -lxslt -lxml2 -lzlib -licudata -licudt -licui18n -licutools -licuuc
// #cgo darwin LDFLAGS: -L　./libs -lcxslt -lxslt -lxml2 -lzlib -licudata -licudt -licui18n -licutools -licuuc
import "C"

import (
	"errors"
	"runtime"
	"unsafe"
)

var (
	ErrXSLTFailure     = errors.New("XSL transformation failed")
	ErrXSLParseFailure = errors.New("Failed to parse XSL")
)

type Stylesheet struct {
	ptr C.xslt_stylesheet_ptr
}

func (xs *Stylesheet) free() {
	if xs.ptr != nil {
		C.xslt_free_style(&xs.ptr)
		xs.ptr = nil
	}
}

func (xs *Stylesheet) Transform(xml []byte) ([]byte, error) {
	var (
		cxml *C.char
		cout *C.char
		ret  C.int
		size C.size_t
	)

	cxml = C.CString(string(xml))
	defer C.free(unsafe.Pointer(cxml))

	ret = C.xslt_apply_style(xs.ptr, cxml, &cout, &size)
	if ret != 0 {
		defer C.free(unsafe.Pointer(cout))
		return nil, ErrXSLTFailure
	}

	ptr := unsafe.Pointer(cout)
	defer C.free(ptr)

	return C.GoBytes(ptr, C.int(size)), nil
}

func NewStylesheet(xsl []byte) (*Stylesheet, error) {
	var (
		cxsl *C.char
		cssp C.xslt_stylesheet_ptr
		ret  C.int
	)

	cxsl = C.CString(string(xsl))
	defer C.free(unsafe.Pointer(cxsl))

	ret = C.xslt_make_style(cxsl, &cssp)
	if ret != 0 {
		return nil, ErrXSLParseFailure
	}

	st := &Stylesheet{ptr: cssp}
	runtime.SetFinalizer(st, (*Stylesheet).free)
	return st, nil
}
