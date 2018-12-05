// <auto-generated>
//  automatically generated by the FlatBuffers compiler, do not modify
// </auto-generated>

namespace ppx
{

using global::System;
using global::FlatBuffers;

public struct BackwardResult : IFlatbufferObject
{
  private Table __p;
  public ByteBuffer ByteBuffer { get { return __p.bb; } }
  public static BackwardResult GetRootAsBackwardResult(ByteBuffer _bb) { return GetRootAsBackwardResult(_bb, new BackwardResult()); }
  public static BackwardResult GetRootAsBackwardResult(ByteBuffer _bb, BackwardResult obj) { return (obj.__assign(_bb.GetInt(_bb.Position) + _bb.Position, _bb)); }
  public void __init(int _i, ByteBuffer _bb) { __p.bb_pos = _i; __p.bb = _bb; }
  public BackwardResult __assign(int _i, ByteBuffer _bb) { __init(_i, _bb); return this; }

  public Tensor? Values(int j) { int o = __p.__offset(4); return o != 0 ? (Tensor?)(new Tensor()).__assign(__p.__indirect(__p.__vector(o) + j * 4), __p.bb) : null; }
  public int ValuesLength { get { int o = __p.__offset(4); return o != 0 ? __p.__vector_len(o) : 0; } }

  public static Offset<BackwardResult> CreateBackwardResult(FlatBufferBuilder builder,
      VectorOffset valuesOffset = default(VectorOffset)) {
    builder.StartObject(1);
    BackwardResult.AddValues(builder, valuesOffset);
    return BackwardResult.EndBackwardResult(builder);
  }

  public static void StartBackwardResult(FlatBufferBuilder builder) { builder.StartObject(1); }
  public static void AddValues(FlatBufferBuilder builder, VectorOffset valuesOffset) { builder.AddOffset(0, valuesOffset.Value, 0); }
  public static VectorOffset CreateValuesVector(FlatBufferBuilder builder, Offset<Tensor>[] data) { builder.StartVector(4, data.Length, 4); for (int i = data.Length - 1; i >= 0; i--) builder.AddOffset(data[i].Value); return builder.EndVector(); }
  public static VectorOffset CreateValuesVectorBlock(FlatBufferBuilder builder, Offset<Tensor>[] data) { builder.StartVector(4, data.Length, 4); builder.Add(data); return builder.EndVector(); }
  public static void StartValuesVector(FlatBufferBuilder builder, int numElems) { builder.StartVector(4, numElems, 4); }
  public static Offset<BackwardResult> EndBackwardResult(FlatBufferBuilder builder) {
    int o = builder.EndObject();
    return new Offset<BackwardResult>(o);
  }
};


}
