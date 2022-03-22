package protocol

type PacketNumber int64

const InvalidPacketNumber PacketNumber = -1

type PacketNumberLen uint8

const (
	PacketNumberLen1 PacketNumberLen = 1
	PacketNumberLen2 PacketNumberLen = 2
	PacketNumberLen3 PacketNumberLen = 3
	PacketNumberLen4 PacketNumberLen = 4
)

func DecodePacketNumber(
	packetNumberLength PacketNumberLen,
	lastPacketNumber PacketNumber,
	wirePacketNumber PacketNumber,
) PacketNumber {
	var epochDelta PacketNumber
	switch packetNumberLength {
	case PacketNumberLen1:
		epochDelta = PacketNumber(1) << 8
	case PacketNumberLen2:
		epochDelta = PacketNumber(1) << 16
	case PacketNumberLen3:
		epochDelta = PacketNumber(1) << 24
	case PacketNumberLen4:
		epochDelta = PacketNumber(1) << 32
	}
	epoch := lastPacketNumber & ^(epochDelta - 1)
	var prevEpochBegin PacketNumber
	if epoch > epochDelta {
		prevEpochBegin = epoch - epochDelta
	}
	nextEpochBegin := epoch + epochDelta
	return closestTo(
		lastPacketNumber+1,
		epoch+wirePacketNumber,
		closestTo(lastPacketNumber+1, prevEpochBegin+wirePacketNumber, nextEpochBegin+wirePacketNumber),
	)
}
func closestTo(target, a, b PacketNumber) PacketNumber {
	if delta(target, a) < delta(target, b) {
		return a
	}
	return b
}

func delta(a, b PacketNumber) PacketNumber {
	if a < b {
		return b - a
	}
	return a - b
}
func GetPacketNumberLengthForHeader(packetNumber, leastUnacked PacketNumber) PacketNumberLen {
	diff := uint64(packetNumber - leastUnacked)
	if diff < (1 << (16 - 1)) {
		return PacketNumberLen2
	}
	if diff < (1 << (24 - 1)) {
		return PacketNumberLen3
	}
	return PacketNumberLen4
}
