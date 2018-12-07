# automatically generated by the FlatBuffers compiler, do not modify

# namespace: ppx

import flatbuffers

class Backward(object):
    __slots__ = ['_tab']

    @classmethod
    def GetRootAsBackward(cls, buf, offset):
        n = flatbuffers.encode.Get(flatbuffers.packer.uoffset, buf, offset)
        x = Backward()
        x.Init(buf, n + offset)
        return x

    # Backward
    def Init(self, buf, pos):
        self._tab = flatbuffers.table.Table(buf, pos)

    # Backward
    def Name(self):
        o = flatbuffers.number_types.UOffsetTFlags.py_type(self._tab.Offset(4))
        if o != 0:
            return self._tab.String(o + self._tab.Pos)
        return None

    # Backward
    def Input(self):
        o = flatbuffers.number_types.UOffsetTFlags.py_type(self._tab.Offset(6))
        if o != 0:
            x = self._tab.Indirect(o + self._tab.Pos)
            from .Tensor import Tensor
            obj = Tensor()
            obj.Init(self._tab.Bytes, x)
            return obj
        return None

    # Backward
    def GradOutput(self):
        o = flatbuffers.number_types.UOffsetTFlags.py_type(self._tab.Offset(8))
        if o != 0:
            x = self._tab.Indirect(o + self._tab.Pos)
            from .Tensor import Tensor
            obj = Tensor()
            obj.Init(self._tab.Bytes, x)
            return obj
        return None

def BackwardStart(builder): builder.StartObject(3)
def BackwardAddName(builder, name): builder.PrependUOffsetTRelativeSlot(0, flatbuffers.number_types.UOffsetTFlags.py_type(name), 0)
def BackwardAddInput(builder, input): builder.PrependUOffsetTRelativeSlot(1, flatbuffers.number_types.UOffsetTFlags.py_type(input), 0)
def BackwardAddGradOutput(builder, gradOutput): builder.PrependUOffsetTRelativeSlot(2, flatbuffers.number_types.UOffsetTFlags.py_type(gradOutput), 0)
def BackwardEnd(builder): return builder.EndObject()
