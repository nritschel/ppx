<?php
// automatically generated by the FlatBuffers compiler, do not modify

namespace ppx;

class MessageBody
{
    const NONE = 0;
    const Handshake = 1;
    const HandshakeResult = 2;
    const Run = 3;
    const RunResult = 4;
    const Sample = 5;
    const SampleResult = 6;
    const Observe = 7;
    const ObserveResult = 8;
    const Tag = 9;
    const TagResult = 10;
    const Forward = 11;
    const ForwardResult = 12;
    const Backward = 13;
    const BackwardResult = 14;
    const BatchOperation = 15;
    const BatchOperationResult = 16;
    const Reset = 17;

    private static $names = array(
        "NONE",
        "Handshake",
        "HandshakeResult",
        "Run",
        "RunResult",
        "Sample",
        "SampleResult",
        "Observe",
        "ObserveResult",
        "Tag",
        "TagResult",
        "Forward",
        "ForwardResult",
        "Backward",
        "BackwardResult",
        "BatchOperation",
        "BatchOperationResult",
        "Reset",
    );

    public static function Name($e)
    {
        if (!isset(self::$names[$e])) {
            throw new \Exception();
        }
        return self::$names[$e];
    }
}
