package base

import (
	"context"
	"errors"
	"net"

	M "github.com/xjasonlyu/tun2socks/v2/metadata"
	"github.com/xjasonlyu/tun2socks/v2/proxy/proto"
)

type Base struct {
	address  string
	protocol string
}

func (b *Base) Addr() string {
	return b.address
}

func (b *Base) Proto() proto.Proto {
	return proto.Proto(b.protocol)
}

func (b *Base) DialContext(context.Context, *M.Metadata) (net.Conn, error) {
	return nil, errors.ErrUnsupported
}

func (b *Base) DialUDP(*M.Metadata) (net.PacketConn, error) {
	return nil, errors.ErrUnsupported
}

func New(address, protocol string) *Base {
	return &Base{
		address:  address,
		protocol: protocol,
	}
}
