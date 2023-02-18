//go:build windows
// +build windows

package proxy

import (
	"go-server/lib/conn"
)

func HandleTrans(c *conn.Conn, s *TunnelModeServer) error {
	return nil
}
