package client

import (
	"net"
	"time"
)

// 一个对net.Conn的包装类型，每次读/写都修改其deadline为当前时间+读/写超时时长
type timeoutConn struct {
	// the actual network connection
	net.Conn

	// timeout for every read from the connection
	ReadTimeout time.Duration

	// timeout for every write to the connection
	WriteTimeout time.Duration
}

func NewTimeoutConn(conn net.Conn, readTimeout, writeTimeout time.Duration) net.Conn {
	return &timeoutConn{
		Conn:         conn,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}
}

func (tr *timeoutConn) Read(p []byte) (n int, err error) {
	var zero time.Duration
	if tr.ReadTimeout != zero {
		tr.Conn.SetReadDeadline(time.Now().Add(tr.ReadTimeout))
	}
	return tr.Conn.Read(p)
}

func (tr *timeoutConn) Write(p []byte) (n int, err error) {
	var zero time.Duration
	if tr.ReadTimeout != zero {
		tr.Conn.SetReadDeadline(time.Now().Add(tr.ReadTimeout))
	}
	return tr.Conn.Write(p)
}