// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package ppx

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type BatchOperation struct {
	_tab flatbuffers.Table
}

func GetRootAsBatchOperation(buf []byte, offset flatbuffers.UOffsetT) *BatchOperation {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &BatchOperation{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *BatchOperation) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *BatchOperation) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *BatchOperation) Operations(obj *Message, j int) bool {
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

func (rcv *BatchOperation) OperationsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func BatchOperationStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func BatchOperationAddOperations(builder *flatbuffers.Builder, operations flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(operations), 0)
}
func BatchOperationStartOperationsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func BatchOperationEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
