package chained

import (
	"fmt"
	"regexp"
)

// Peer represents the path to the root address of
// another blockchain.
type Peer struct {
	Host string
	Port uint
}

// NewPeer will create and validate a peer with the
// given parameters. An error will be returned if the
// peer is invalid.
func NewPeer(h string, p uint) (Peer, error) {
	peer := Peer{h, p}

	if !peer.IsValid() {
		return Peer{}, ErrInvalidData
	}

	return peer, nil
}

// Address will return the full address of the peer
//
// Example output:
//   http://localhost:3000/
//   http://192.168.1.1:6060/
func (p Peer) Address(https bool) string {
	prot := "http:"
	if https {
		prot = "https:"
	}

	return fmt.Sprintf("%v//%v:%v/", prot, p.Host, p.Port)
}

// JSONString converts the received peer to a
// JSON string and returns it.
func (p Peer) JSONString() string {
	s, _ := toJSONString(p)
	return s
}

// IsValid inspects the fields of the received
// peer and determines if its contents are
// valid.
//
// A peer must have a host with the proper format
// (ex. localhost or 192.168.1.1), and must have a
// port value above 1000.
func (p Peer) IsValid() bool {
	// Host must be in a proper format
	reg := regexp.MustCompile(`^localhost|(\d{0,3}\.){3}\d{0,3}$`)
	if !reg.MatchString(p.Host) {
		return false
	}

	// Port must be at least 1000
	if p.Port < 1000 {
		return false
	}

	return true
}
