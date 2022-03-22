package protocol

type StreamType uint8

const (
	StreamTypeUni StreamType = iota
	StreamTypeBidi
)

const InvalidStreamID StreamID = -1

type StreamNum int64

const (
	InvalidStreamNum           = -1
	MaxStreamCount   StreamNum = 1 << 60
)

func (s StreamNum) StreamID(stype StreamType, pers Perspective) StreamID {
	if s == 0 {
		return InvalidStreamID
	}

	var first StreamID
	switch stype {
	case StreamTypeBidi:
		switch pers {
		case PerspectiveClient:
			first = 0
		case PerspectiveServer:
			first = 1
		}
	case StreamTypeUni:
		switch pers {
		case PerspectiveClient:
			first = 2
		case PerspectiveServer:
			first = 3
		}
	}
	return first + 4*StreamID(s-1)
}

type StreamID int64

func (s StreamID) InitiatedBy() Perspective {
	if s%2 == 0 {
		return PerspectiveClient
	}
	return PerspectiveServer
}
func (s StreamID) Type() StreamType {
	if s%4 >= 2 {
		return StreamTypeUni
	}
	return StreamTypeBidi
}
func (s StreamID) StreamNum() StreamNum {
	return StreamNum(s/4) + 1
}
