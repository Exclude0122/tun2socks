package proxy

import (
	"net/url"
	"sync"

	"github.com/xjasonlyu/tun2socks/v2/proxy/base"
)

type ErrProtocol struct {
	Protocol string
}

func (e *ErrProtocol) Error() string {
	return "unknown protocol: " + e.Protocol
}

var (
	protocols   = make(map[string]func(u *url.URL) (base.Proxy, error))
	prototolsMu = sync.Mutex{}
)

func RegisterProtocol(name string, factory func(u *url.URL) (base.Proxy, error)) {
	prototolsMu.Lock()
	defer prototolsMu.Unlock()

	if _, ok := protocols[name]; ok {
		panic("protocol already registered: " + name)
	}
	protocols[name] = factory
}

func Parse(urlString string) (base.Proxy, error) {
	u, err := url.Parse(urlString)
	if err != nil {
		return nil, err
	}

	prototolsMu.Lock()
	defer prototolsMu.Unlock()

	factory, ok := protocols[u.Scheme]
	if !ok {
		return nil, &ErrProtocol{Protocol: u.Scheme}
	}

	return factory(u)
}
