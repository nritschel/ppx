// automatically generated by the FlatBuffers compiler, do not modify

package ppx

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Normal struct {
	_tab flatbuffers.Table
}

func GetRootAsNormal(buf []byte, offset flatbuffers.UOffsetT) *Normal {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Normal{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Normal) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Normal) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Normal) Mean(obj *Tensor) *Tensor {
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

func (rcv *Normal) Stddev(obj *Tensor) *Tensor {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
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

func NormalStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func NormalAddMean(builder *flatbuffers.Builder, mean flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(mean), 0)
}
func NormalAddStddev(builder *flatbuffers.Builder, stddev flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(stddev), 0)
}
func NormalEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
