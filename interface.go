package quic

import (
	"context"
	"errors"
	"io"
	"net"
	"quic/internal/protocol"
	"time"
)

type StreamID = protocol.StreamID

type VersionNumber = protocol.VersionNumber

const (
	VersionDraft29 = protocol.VersionDraft29
	Version1       = protocol.Version1
)

type Token struct {
	IsRetryToken bool
	RemoteAddr   string
	SentTime     time.Time
}

type ClientToken struct {
	data []byte
}

type TokenStore interface {
	Pop(key StreamID) (token *ClientToken)
	Put(key string, token *ClientToken)
}

var Err0RTTRejected = errors.New("0-RTT rejected")

var SessionTracingKey = sessionTracingCtxKey{}

type sessionTracingCtxKey struct{}

type Stream interface {
	ReceiveStream
	SendStream
	SetDeadline(t time.Time) error
}

type ReceiveStream interface {
	StreamID() StreamID
	io.Reader
	CancelRead(StreamErrorCode)
	SetReadDeadline(t time.Time) error
}

type SendStream interface {
	StreamID() StreamID

	io.Writer
	io.Closer
	CancelWrite(StreamErrorCode)
	Context() context.Context
	SetWriteDeadline(t time.Time) error
}

type Session interface {
	AcceptStream(context.Context) (Stream, error)
	AcceptUniStream(context.Context) (ReceiveStream, error)
	OpenStream() (Stream, error)
	OpenStreamSync(context.Context) (Stream, error)
	OpenUniStream() (SendStream, error)
	OpenUniStreamSync(context.Context) (SendStream, error)
	LocalAddr() net.Addr
	RemoteAddr() net.Addr
	CloseWithError(ApplicationErrorCode, string) error
	Context() context.Context
	ConnectionState() ConnectionState
	SendMessage([]byte) error
	ReceiveMessage() ([]byte, error)
}

type EarlySession interface {
	Session
	HandshakeComplete() context.Context
	NextSession() Session
}

type Config struct {
	Versions             []VersionNumber
	ConnectionIDLength   int
	HandshakeIdleTimeout time.Duration
	MaxIdleTimeout       time.Duration
	AcceptToken          func(clientAddr net.Addr, token *Token) bool

	TokenStore TokenStore

	InitialStreamReceiveWindow       uint64
	MaxStreamReceiveWindow           uint64
	InitialConnectionReceiveWindow   uint64
	MaxConnectionReceiveWindow       uint64
	AllowConnectionWindowIncrease    func(sess Session, delta uint64) bool
	MaxIncomingStreams               int64
	MaxIncomingUniStreams            int64
	StatelessResetKey                []byte
	KeepAlive                        bool
	DisablePathMTUDiscovery          bool
	DisableVersionNegotiationPackets bool
	EnableDatagrams                  bool
	Tracer                           logging.Tracer
}
