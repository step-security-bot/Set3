// Code generated by command: go run asm.go -out match.s -stubs match_amd64.go. DO NOT EDIT.

//go:build amd64

package simd

// MatchMetadata performs a 16-way probe of |metadata| using SSE instructions
// nb: |metadata| must be an aligned pointer
func MatchMetadata(metadata *[16]int8, hash int8) uint16

// MatchCRTLhash performs a 16-way probe of the 8 bit value in |hash| to every byte of |ctrl| using SSE2/SSE3 instructions
// nb: |ctrl| must be an aligned pointer
func MatchCRTLhash(ctrl *[16]int8, hash uint64) uint64

// MatchCRTLempty performs a 16-way probe of the 8 bit value for 'empty' (0b1000_0000) to every byte of |ctrl| using SSE2/SSE3 instructions
// nb: |ctrl| must be an aligned pointer
func MatchCRTLempty(ctrl *[16]int8) uint64

// MatchCRTLdeleted performs a 16-way probe of the 8 bit value for 'deleted' (0b1111_1110) to every byte of |ctrl| using SSE2/SSE3 instructions
// nb: |ctrl| must be an aligned pointer
func MatchCRTLdeleted(ctrl *[16]int8) uint64
