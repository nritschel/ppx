# automatically generated by the FlatBuffers compiler, do not modify

# namespace: ppx

import flatbuffers

class ForwardResult(object):
    __slots__ = ['_tab']

    @classmethod
    def GetRootAsForwardResult(cls, buf, offset):
        n = flatbuffers.encode.Get(flatbuffers.packer.uoffset, buf, offset)
        x = ForwardResult()
        x.Init(buf, n + offset)
        return x

    # ForwardResult
    def Init(self, buf, pos):
        self._tab = flatbuffers.table.Table(buf, pos)

    # ForwardResult
    def Values(self, j):
        o = flatbuffers.number_types.UOffsetTFlags.py_type(self._tab.Offset(4))
        if o != 0:
            x = self._tab.Vector(o)
            x += flatbuffers.number_types.UOffsetTFlags.py_type(j) * 4
            x = self._tab.Indirect(x)
            from .Tensor import Tensor
            obj = Tensor()
            obj.Init(self._tab.Bytes, x)
            return obj
        return None

    # ForwardResult
    def ValuesLength(self):
        o = flatbuffers.number_types.UOffsetTFlags.py_type(self._tab.Offset(4))
        if o != 0:
            return self._tab.VectorLen(o)
        return 0

def ForwardResultStart(builder): builder.StartObject(1)
def ForwardResultAddValues(builder, values): builder.PrependUOffsetTRelativeSlot(0, flatbuffers.number_types.UOffsetTFlags.py_type(values), 0)
def ForwardResultStartValuesVector(builder, numElems): return builder.StartVector(4, numElems, 4)
def ForwardResultEnd(builder): return builder.EndObject()
