// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package ppx

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type ForwardResult struct {
	_tab flatbuffers.Table
}

func GetRootAsForwardResult(buf []byte, offset flatbuffers.UOffsetT) *ForwardResult {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ForwardResult{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *ForwardResult) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ForwardResult) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *ForwardResult) Output(obj *Tensor) *Tensor {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Tensor)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func ForwardResultStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func ForwardResultAddOutput(builder *flatbuffers.Builder, output flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(output), 0)
}
func ForwardResultEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
