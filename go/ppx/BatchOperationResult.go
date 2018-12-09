// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package ppx

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type BatchOperationResult struct {
	_tab flatbuffers.Table
}

func GetRootAsBatchOperationResult(buf []byte, offset flatbuffers.UOffsetT) *BatchOperationResult {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &BatchOperationResult{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *BatchOperationResult) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *BatchOperationResult) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *BatchOperationResult) Results(obj *Message, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *BatchOperationResult) ResultsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func BatchOperationResultStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func BatchOperationResultAddResults(builder *flatbuffers.Builder, results flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(results), 0)
}
func BatchOperationResultStartResultsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func BatchOperationResultEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
