package utils

import (
	"quic/internal/protocol"
	"time"
)

const (
	rttAlpha          = 0.125
	oneMinusAlpha     = 1 - rttAlpha
	rttBeta           = 0.25
	oneMinusBeta      = 1 - rttBeta
	defaultInitialRTT = 100 * time.Millisecond
)

type RTTStats struct {
	hasMeasurement bool
	minRTT         time.Duration
	latestRTT      time.Duration
	smoothedRTT    time.Duration
	meanDeviation  time.Duration
	maxAckDelay    time.Duration
}

func NewRttStats() *RTTStats {
	return &RTTStats{}
}
func (r *RTTStats) MinRTT() time.Duration { return r.minRTT }

// LatestRTT returns the most recent rtt measurement.
// May return Zero if no valid updates have occurred.
func (r *RTTStats) LatestRTT() time.Duration { return r.latestRTT }

// SmoothedRTT returns the smoothed RTT for the connection.
// May return Zero if no valid updates have occurred.
func (r *RTTStats) SmoothedRTT() time.Duration { return r.smoothedRTT }

// MeanDeviation gets the mean deviation
func (r *RTTStats) MeanDeviation() time.Duration { return r.meanDeviation }

// MaxAckDelay gets the max_ack_delay advertised by the peer
func (r *RTTStats) MaxAckDelay() time.Duration { return r.maxAckDelay }

func (r *RTTStats) PTO(includeMaxAckDelay bool) time.Duration {
	if r.SmoothedRTT() == 0 {
		return 2 * defaultInitialRTT
	}

	pto := r.SmoothedRTT() + Max(4*r.MeanDeviation(), protocol.TimerGranularity)
	if includeMaxAckDelay {
		pto += r.MaxAckDelay()
	}
	return pto
}
func (r *RTTStats) SetMaxAckDelay(mad time.Duration) {
	r.maxAckDelay = mad
}

func (r *RTTStats) UpdateRTT(sendDelta, ackDelay time.Duration, now time.Time) {
	if sendDelta == InfDuration || sendDelta <= 0 {
		return
	}

	if r.minRTT == 0 || r.minRTT > sendDelta {
		r.minRTT = sendDelta
	}

	sample := sendDelta
	if sample-r.minRTT >= ackDelay {
		sample -= ackDelay
	}

	r.latestRTT = sample
	if !r.hasMeasurement {
		r.hasMeasurement = true
		r.smoothedRTT = sample
		r.meanDeviation = sample / 2
	} else {
		r.meanDeviation = time.Duration(oneMinusBeta*float32(r.meanDeviation/time.Microsecond)+rttBeta*float32(Abs(r.smoothedRTT-sample)/time.Microsecond)) * time.Microsecond
		r.smoothedRTT = time.Duration((float32(r.smoothedRTT/time.Microsecond)*oneMinusAlpha)+(float32(sample/time.Microsecond)*rttAlpha)) * time.Microsecond
	}
}
func (r *RTTStats) SetInitialRTT(t time.Duration) {
	if r.hasMeasurement {
		panic("initial RTT set after first measurement")
	}
	r.smoothedRTT = t
	r.latestRTT = t
}

// OnConnectionMigration is called when connection migrates and rtt measurement needs to be reset.
func (r *RTTStats) OnConnectionMigration() {
	r.latestRTT = 0
	r.minRTT = 0
	r.smoothedRTT = 0
	r.meanDeviation = 0
}
func (r *RTTStats) ExpireSmoothedMetrics() {
	r.meanDeviation = Max(r.meanDeviation, Abs(r.smoothedRTT-r.latestRTT))
	r.smoothedRTT = Max(r.smoothedRTT, r.latestRTT)
}
