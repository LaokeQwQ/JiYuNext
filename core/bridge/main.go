package main

/*
#include <stdlib.h>
*/
import "C"

import (
	"encoding/json"
	"fmt"
	"unsafe"

	core "github.com/LaokeQwQ/JiYuNext/core"
)

//export jy_init
func jy_init(configJSON *C.char) (out *C.char) {
	defer func() {
		if r := recover(); r != nil {
			out = panicResponseToCString("jy_init", r)
		}
	}()
	resp := core.Init(cStringToGo(configJSON))
	return responseToCString(resp)
}

//export jy_execute
func jy_execute(actionJSON *C.char) (out *C.char) {
	defer func() {
		if r := recover(); r != nil {
			out = panicResponseToCString("jy_execute", r)
		}
	}()
	resp := core.Execute(cStringToGo(actionJSON))
	return responseToCString(resp)
}

//export jy_shutdown
func jy_shutdown() (out *C.char) {
	defer func() {
		if r := recover(); r != nil {
			out = panicResponseToCString("jy_shutdown", r)
		}
	}()
	resp := core.Shutdown()
	return responseToCString(resp)
}

//export jy_free
func jy_free(ptr *C.char) {
	if ptr != nil {
		C.free(unsafe.Pointer(ptr))
	}
}

func cStringToGo(v *C.char) string {
	if v == nil {
		return ""
	}
	return C.GoString(v)
}

func responseToCString(resp core.Response) *C.char {
	b, err := json.Marshal(resp)
	if err != nil {
		fallback := core.Response{
			Code:    core.CodeInternal,
			Message: "marshal response failed",
		}
		b, _ = json.Marshal(fallback)
	}
	return C.CString(string(b))
}

func panicResponseToCString(operation string, r any) *C.char {
	msg := fmt.Sprintf("%s panic: %v", operation, r)
	return responseToCString(core.Response{
		Code:    core.CodeInternal,
		Message: msg,
		Result:  nil,
	})
}

func main() {}
