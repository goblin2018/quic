package protocol

type KeyPhase uint64

func (p KeyPhase) Bit() KeyPhaseBit {
	if p%2 == 0 {
		return KeyPhaseZero
	}
	return KeyPhaseOne
}

type KeyPhaseBit uint8

const (
	KeyPhaseUndefined KeyPhaseBit = iota
	KeyPhaseZero
	KeyPhaseOne
)

func (p KeyPhaseBit) String() string {
	switch p {
	case KeyPhaseZero:
		return "0"
	case KeyPhaseOne:
		return "1"
	default:
		return "undefined"
	}
}
