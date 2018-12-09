// <auto-generated>
//  automatically generated by the FlatBuffers compiler, do not modify
// </auto-generated>

namespace ppx
{

using global::System;
using global::FlatBuffers;

public struct BatchOperationResult : IFlatbufferObject
{
  private Table __p;
  public ByteBuffer ByteBuffer { get { return __p.bb; } }
  public static BatchOperationResult GetRootAsBatchOperationResult(ByteBuffer _bb) { return GetRootAsBatchOperationResult(_bb, new BatchOperationResult()); }
  public static BatchOperationResult GetRootAsBatchOperationResult(ByteBuffer _bb, BatchOperationResult obj) { return (obj.__assign(_bb.GetInt(_bb.Position) + _bb.Position, _bb)); }
  public void __init(int _i, ByteBuffer _bb) { __p.bb_pos = _i; __p.bb = _bb; }
  public BatchOperationResult __assign(int _i, ByteBuffer _bb) { __init(_i, _bb); return this; }

  public Message? Results(int j) { int o = __p.__offset(4); return o != 0 ? (Message?)(new Message()).__assign(__p.__indirect(__p.__vector(o) + j * 4), __p.bb) : null; }
  public int ResultsLength { get { int o = __p.__offset(4); return o != 0 ? __p.__vector_len(o) : 0; } }

  public static Offset<BatchOperationResult> CreateBatchOperationResult(FlatBufferBuilder builder,
      VectorOffset resultsOffset = default(VectorOffset)) {
    builder.StartObject(1);
    BatchOperationResult.AddResults(builder, resultsOffset);
    return BatchOperationResult.EndBatchOperationResult(builder);
  }

  public static void StartBatchOperationResult(FlatBufferBuilder builder) { builder.StartObject(1); }
  public static void AddResults(FlatBufferBuilder builder, VectorOffset resultsOffset) { builder.AddOffset(0, resultsOffset.Value, 0); }
  public static VectorOffset CreateResultsVector(FlatBufferBuilder builder, Offset<Message>[] data) { builder.StartVector(4, data.Length, 4); for (int i = data.Length - 1; i >= 0; i--) builder.AddOffset(data[i].Value); return builder.EndVector(); }
  public static VectorOffset CreateResultsVectorBlock(FlatBufferBuilder builder, Offset<Message>[] data) { builder.StartVector(4, data.Length, 4); builder.Add(data); return builder.EndVector(); }
  public static void StartResultsVector(FlatBufferBuilder builder, int numElems) { builder.StartVector(4, numElems, 4); }
  public static Offset<BatchOperationResult> EndBatchOperationResult(FlatBufferBuilder builder) {
    int o = builder.EndObject();
    return new Offset<BatchOperationResult>(o);
  }
};


}
