// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package ppx

type MessageBody = byte
const (
	MessageBodyNONE MessageBody = 0
	MessageBodyHandshake MessageBody = 1
	MessageBodyHandshakeResult MessageBody = 2
	MessageBodyRun MessageBody = 3
	MessageBodyRunResult MessageBody = 4
	MessageBodySample MessageBody = 5
	MessageBodySampleResult MessageBody = 6
	MessageBodyObserve MessageBody = 7
	MessageBodyObserveResult MessageBody = 8
	MessageBodyTag MessageBody = 9
	MessageBodyTagResult MessageBody = 10
	MessageBodyForward MessageBody = 11
	MessageBodyForwardResult MessageBody = 12
	MessageBodyBackward MessageBody = 13
	MessageBodyBackwardResult MessageBody = 14
	MessageBodyReset MessageBody = 15
)

var EnumNamesMessageBody = map[MessageBody]string{
	MessageBodyNONE:"NONE",
	MessageBodyHandshake:"Handshake",
	MessageBodyHandshakeResult:"HandshakeResult",
	MessageBodyRun:"Run",
	MessageBodyRunResult:"RunResult",
	MessageBodySample:"Sample",
	MessageBodySampleResult:"SampleResult",
	MessageBodyObserve:"Observe",
	MessageBodyObserveResult:"ObserveResult",
	MessageBodyTag:"Tag",
	MessageBodyTagResult:"TagResult",
	MessageBodyForward:"Forward",
	MessageBodyForwardResult:"ForwardResult",
	MessageBodyBackward:"Backward",
	MessageBodyBackwardResult:"BackwardResult",
	MessageBodyReset:"Reset",
}

