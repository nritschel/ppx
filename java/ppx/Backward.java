// automatically generated by the FlatBuffers compiler, do not modify

package ppx;

import java.nio.*;
import java.lang.*;
import java.util.*;
import com.google.flatbuffers.*;

@SuppressWarnings("unused")public final class Backward extends Table {
  public static Backward getRootAsBackward(ByteBuffer _bb) { return getRootAsBackward(_bb, new Backward()); }
  public static Backward getRootAsBackward(ByteBuffer _bb, Backward obj) { _bb.order(ByteOrder.LITTLE_ENDIAN); return (obj.__assign(_bb.getInt(_bb.position()) + _bb.position(), _bb)); }
  public void __init(int _i, ByteBuffer _bb) { bb_pos = _i; bb = _bb; }
  public Backward __assign(int _i, ByteBuffer _bb) { __init(_i, _bb); return this; }

  public String name() { int o = __offset(4); return o != 0 ? __string(o + bb_pos) : null; }
  public ByteBuffer nameAsByteBuffer() { return __vector_as_bytebuffer(4, 1); }
  public ByteBuffer nameInByteBuffer(ByteBuffer _bb) { return __vector_in_bytebuffer(_bb, 4, 1); }
  public Tensor arguments(int j) { return arguments(new Tensor(), j); }
  public Tensor arguments(Tensor obj, int j) { int o = __offset(6); return o != 0 ? obj.__assign(__indirect(__vector(o) + j * 4), bb) : null; }
  public int argumentsLength() { int o = __offset(6); return o != 0 ? __vector_len(o) : 0; }

  public static int createBackward(FlatBufferBuilder builder,
      int nameOffset,
      int argumentsOffset) {
    builder.startObject(2);
    Backward.addArguments(builder, argumentsOffset);
    Backward.addName(builder, nameOffset);
    return Backward.endBackward(builder);
  }

  public static void startBackward(FlatBufferBuilder builder) { builder.startObject(2); }
  public static void addName(FlatBufferBuilder builder, int nameOffset) { builder.addOffset(0, nameOffset, 0); }
  public static void addArguments(FlatBufferBuilder builder, int argumentsOffset) { builder.addOffset(1, argumentsOffset, 0); }
  public static int createArgumentsVector(FlatBufferBuilder builder, int[] data) { builder.startVector(4, data.length, 4); for (int i = data.length - 1; i >= 0; i--) builder.addOffset(data[i]); return builder.endVector(); }
  public static void startArgumentsVector(FlatBufferBuilder builder, int numElems) { builder.startVector(4, numElems, 4); }
  public static int endBackward(FlatBufferBuilder builder) {
    int o = builder.endObject();
    return o;
  }
}

