package utils

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"time"

	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/multiformats/go-multibase"
)

const (
	DefaultDialTimeout = 10 * time.Second
	DefaultConnectTimeout = 30 * time.Second
)

func DialContext(ctx context.Context, addr string) (net.Conn, error) {
	dialer := &net.Dialer{
		Timeout: DefaultDialTimeout,
	}
	return dialer.DialContext(ctx, "tcp", addr)
}

func ConnectContext(ctx context.Context, addr string) (net.Conn, error) {
	conn, err := DialContext(ctx, addr)
	if err!= nil {
		return nil, err
	}
	err = conn.SetDeadline(time.Now().Add(DefaultConnectTimeout))
	if err!= nil {
		return nil, err
	}
	return conn, nil
}

func NewTLSConfig(certFile, keyFile string) (*tls.Config, error) {
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err!= nil {
		return nil, err
	}
	return &tls.Config{
		Certificates: []tls.Certificate{cert},
	}, nil
}

func NewLibp2pHost(ctx context.Context, privateKey *ecdsa.PrivateKey, listenAddr string) (*network.Host, error) {
	id, err := peer.IDFromPrivateKey(privateKey)
	if err!= nil {
		return nil, err
	}
	host, err := network.NewHost(ctx, id, listenAddr)
	if err!= nil {
		return nil, err
	}
	return host, nil
}

func GetFreePort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err!= nil {
		return 0, err
	}
	l, err := net.ListenTCP("tcp", addr)
	if err!= nil {
		return 0, err
	}
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port, nil
}

func GetExternalIP() (net.IP, error) {
	ifaces, err := net.Interfaces()
	if err!= nil {
		return nil, err
	}
	for _, iface := range ifaces {
		addrs, err := iface.Addrs()
		if err!= nil {
			return nil, err
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip!= nil &&!ip.IsLoopback() && ip.To4()!= nil {
				return ip, nil
			}
		}
	}
	return nil, errors.New("no external IP found")
}
