package socks4

import (
	"context"
	"fmt"
	"net"
	"net/url"

	"github.com/xjasonlyu/tun2socks/v2/dialer"
	M "github.com/xjasonlyu/tun2socks/v2/metadata"
	"github.com/xjasonlyu/tun2socks/v2/proxy"
	"github.com/xjasonlyu/tun2socks/v2/proxy/base"
	"github.com/xjasonlyu/tun2socks/v2/transport/socks4"
)

const Proto = "socks4"

type Socks4 struct {
	*base.Base

	userID string
}

func NewSocks4(addr, userID string) (*Socks4, error) {
	return &Socks4{
		Base:   base.New(addr, Proto),
		userID: userID,
	}, nil
}

func (ss *Socks4) DialContext(ctx context.Context, metadata *M.Metadata) (c net.Conn, err error) {
	c, err = dialer.DialContext(ctx, "tcp", ss.Base.Addr())
	if err != nil {
		return nil, fmt.Errorf("connect to %s: %w", ss.Base.Addr(), err)
	}
	base.SetKeepAlive(c)

	defer func(c net.Conn) {
		base.SafeConnClose(c, err)
	}(c)

	err = socks4.ClientHandshake(c, metadata.DestinationAddress(), socks4.CmdConnect, ss.userID)
	return
}

func parseSocks4(u *url.URL) (base.Proxy, error) {
	address, userID := u.Host, u.User.Username()
	return NewSocks4(address, userID)
}

func init() {
	proxy.RegisterProtocol(Proto, parseSocks4)
}
