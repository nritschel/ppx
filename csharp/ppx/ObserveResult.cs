// <auto-generated>
//  automatically generated by the FlatBuffers compiler, do not modify
// </auto-generated>

namespace ppx
{

using global::System;
using global::FlatBuffers;

public struct ObserveResult : IFlatbufferObject
{
  private Table __p;
  public ByteBuffer ByteBuffer { get { return __p.bb; } }
  public static ObserveResult GetRootAsObserveResult(ByteBuffer _bb) { return GetRootAsObserveResult(_bb, new ObserveResult()); }
  public static ObserveResult GetRootAsObserveResult(ByteBuffer _bb, ObserveResult obj) { return (obj.__assign(_bb.GetInt(_bb.Position) + _bb.Position, _bb)); }
  public void __init(int _i, ByteBuffer _bb) { __p.bb_pos = _i; __p.bb = _bb; }
  public ObserveResult __assign(int _i, ByteBuffer _bb) { __init(_i, _bb); return this; }


  public static void StartObserveResult(FlatBufferBuilder builder) { builder.StartObject(0); }
  public static Offset<ObserveResult> EndObserveResult(FlatBufferBuilder builder) {
    int o = builder.EndObject();
    return new Offset<ObserveResult>(o);
  }
};


}
