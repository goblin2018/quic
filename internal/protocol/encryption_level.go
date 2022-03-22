package protocol

type EncryptionLevel uint8

const (
	EncryptionInitial EncryptionLevel = 1 + iota
	EncryptionHandshake
	Encryption0RTT
	Encryption1RTT
)

func (e EncryptionLevel) String() string {
	switch e {
	case EncryptionInitial:
		return "Initial"
	case EncryptionHandshake:
		return "Handshake"
	case Encryption0RTT:
		return "0-RTT"
	case Encryption1RTT:
		return "1-RTT"
	}
	return "unknown"
}
