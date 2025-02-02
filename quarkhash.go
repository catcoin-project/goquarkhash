package quarkhash

// +build cgo

/*
#include <stdlib.h>
#include "quarkhash.h"
*/
import "C"

import (
	"unsafe"
)

func QuarkHash(input []byte) []byte {

	cinput := C.CBytes(input)
	defer C.free(cinput)

	coutput := C.CBytes(make([]byte, 32))
	defer C.free(unsafe.Pointer(coutput))

	C.quark_hash(cinput, C.ulonglong(len(input)), coutput)

	return C.GoBytes(coutput, 32)
}
