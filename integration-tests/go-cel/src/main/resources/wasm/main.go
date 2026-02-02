package main

import (
	"encoding/json"
	"unsafe"

	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/checker/decls"
)

// Simple bump allocator for host interop, used by "malloc" - which in turn is exported to allow externalized memory
// management
var heap = make([]byte, 10*1024*1024) // 10MB heap
var heapOffset = 0

//go:wasmexport malloc
func malloc(size uint32) uint32 {
	if size == 0 {
		return 0
	}
	ptr := uint32(uintptr(unsafe.Pointer(&heap[heapOffset])))
	heapOffset += int(size)
	return ptr
}

//go:wasmexport free
func free(ptr uint32) {
	// No-op: bump allocator doesn't actually free
}

// evalPolicy evaluates a CEL expression
// Parameters:
//   - policyPtr: pointer to policy string
//   - policyLen: length of policy string
//   - inputPtr: pointer to input JSON bytes
//   - inputLen: length of input JSON bytes
//
// Returns: 1 = policy allows, 0 = policy denies, negative = error
//
//go:wasmexport evalPolicy
func evalPolicy(policyPtr, policyLen, inputPtr, inputLen uint32) int32 {
	// Convert pointers to Go types
	policy := unsafe.String((*byte)(unsafe.Pointer(uintptr(policyPtr))), policyLen)
	inputJSON := unsafe.Slice((*byte)(unsafe.Pointer(uintptr(inputPtr))), inputLen)

	// 1: Parse the JSON input
	var input map[string]any
	if err := json.Unmarshal(inputJSON, &input); err != nil {
		// JSON parse error
		return -1
	}

	// 2: Create CEL environment
	env, err := cel.NewEnv(
		cel.Declarations(
			decls.NewVar("object", decls.NewMapType(decls.String, decls.Dyn)),
		),
	)
	if err != nil {
		// CEL environment creation error
		return -2
	}

	// 3: Compile the CEL expression
	ast, iss := env.Compile(policy)
	if iss.Err() != nil {
		// Compilation error
		return -3
	}

	// 4: Create program
	prg, err := env.Program(ast)
	if err != nil {
	    // Program creation error
		return -4
	}

	// 5: Evaluate the expression
	out, _, err := prg.Eval(map[string]any{
		"object": input,
	})
	if err != nil {
		// CEL runtime error
		return -5
	}

	// 6: Check if result is a boolean true
	if b, ok := out.Value().(bool); ok && b {
		// Policy allows
		return 1
	}
    // Policy denies
	return 0
}

func main() {
	// nada
}
