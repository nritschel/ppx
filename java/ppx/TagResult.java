// automatically generated by the FlatBuffers compiler, do not modify

package ppx;

import java.nio.*;
import java.lang.*;
import java.util.*;
import com.google.flatbuffers.*;

@SuppressWarnings("unused")public final class TagResult extends Table {
  public static TagResult getRootAsTagResult(ByteBuffer _bb) { return getRootAsTagResult(_bb, new TagResult()); }
  public static TagResult getRootAsTagResult(ByteBuffer _bb, TagResult obj) { _bb.order(ByteOrder.LITTLE_ENDIAN); return (obj.__assign(_bb.getInt(_bb.position()) + _bb.position(), _bb)); }
  public void __init(int _i, ByteBuffer _bb) { bb_pos = _i; bb = _bb; }
  public TagResult __assign(int _i, ByteBuffer _bb) { __init(_i, _bb); return this; }


  public static void startTagResult(FlatBufferBuilder builder) { builder.startObject(0); }
  public static int endTagResult(FlatBufferBuilder builder) {
    int o = builder.endObject();
    return o;
  }
}

