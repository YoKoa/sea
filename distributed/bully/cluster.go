package bully

import "sync"

type PeerSet interface {
	Add(id, address string)
	Remove(id string)
	Send(id string, msg []byte)
}

type Peer struct {
	Id   string
	Addr string
}

type Cluster struct {
	mu    *sync.RWMutex
	peers map[string]*Peer
}

func (c *Cluster) Add(id, address string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.peers[id] = &Peer{
		Id:   id,
		Addr: address,
	}
}

func (c *Cluster) Remove(id string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.peers, id)
}

func (c *Cluster) Send(id string, msg []byte) {

}
