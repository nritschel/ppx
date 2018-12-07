// <auto-generated>
//  automatically generated by the FlatBuffers compiler, do not modify
// </auto-generated>

namespace ppx
{

using global::System;
using global::FlatBuffers;

public struct Backward : IFlatbufferObject
{
  private Table __p;
  public ByteBuffer ByteBuffer { get { return __p.bb; } }
  public static Backward GetRootAsBackward(ByteBuffer _bb) { return GetRootAsBackward(_bb, new Backward()); }
  public static Backward GetRootAsBackward(ByteBuffer _bb, Backward obj) { return (obj.__assign(_bb.GetInt(_bb.Position) + _bb.Position, _bb)); }
  public void __init(int _i, ByteBuffer _bb) { __p.bb_pos = _i; __p.bb = _bb; }
  public Backward __assign(int _i, ByteBuffer _bb) { __init(_i, _bb); return this; }

  public string Name { get { int o = __p.__offset(4); return o != 0 ? __p.__string(o + __p.bb_pos) : null; } }
#if ENABLE_SPAN_T
  public Span<byte> GetNameBytes() { return __p.__vector_as_span(4); }
#else
  public ArraySegment<byte>? GetNameBytes() { return __p.__vector_as_arraysegment(4); }
#endif
  public byte[] GetNameArray() { return __p.__vector_as_array<byte>(4); }
  public Tensor? Input { get { int o = __p.__offset(6); return o != 0 ? (Tensor?)(new Tensor()).__assign(__p.__indirect(o + __p.bb_pos), __p.bb) : null; } }
  public Tensor? GradOutput { get { int o = __p.__offset(8); return o != 0 ? (Tensor?)(new Tensor()).__assign(__p.__indirect(o + __p.bb_pos), __p.bb) : null; } }

  public static Offset<Backward> CreateBackward(FlatBufferBuilder builder,
      StringOffset nameOffset = default(StringOffset),
      Offset<Tensor> inputOffset = default(Offset<Tensor>),
      Offset<Tensor> grad_outputOffset = default(Offset<Tensor>)) {
    builder.StartObject(3);
    Backward.AddGradOutput(builder, grad_outputOffset);
    Backward.AddInput(builder, inputOffset);
    Backward.AddName(builder, nameOffset);
    return Backward.EndBackward(builder);
  }

  public static void StartBackward(FlatBufferBuilder builder) { builder.StartObject(3); }
  public static void AddName(FlatBufferBuilder builder, StringOffset nameOffset) { builder.AddOffset(0, nameOffset.Value, 0); }
  public static void AddInput(FlatBufferBuilder builder, Offset<Tensor> inputOffset) { builder.AddOffset(1, inputOffset.Value, 0); }
  public static void AddGradOutput(FlatBufferBuilder builder, Offset<Tensor> gradOutputOffset) { builder.AddOffset(2, gradOutputOffset.Value, 0); }
  public static Offset<Backward> EndBackward(FlatBufferBuilder builder) {
    int o = builder.EndObject();
    return new Offset<Backward>(o);
  }
};


}
