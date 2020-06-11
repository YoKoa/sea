// +build darwin

package core

import (
	"golang.org/x/sys/unix"
	"sync"
	"sync/atomic"
)

type Poller struct {
	fd       int
	running  atomic.Value
	waitDone chan struct{}
	sockets  sync.Map // [fd]events
}
func Create() (*Poller, error) {
	fd, err := unix.Kqueue()
	if err != nil {
		return nil, err
	}
	_, err = unix.Kevent(fd, []unix.Kevent_t{{
		Ident:  0,
		Filter: unix.EVFILT_USER,
		Flags:  unix.EV_ADD | unix.EV_CLEAR,
	}}, nil, nil)
	if err != nil {
		return nil, err
	}

	return &Poller{
		fd:       fd,
		waitDone: make(chan struct{}),
	}, nil
}


