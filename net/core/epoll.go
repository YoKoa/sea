// +build linux

package core

import (
	"github.com/YoKoa/sea/sync"
	"golang.org/x/sys/unix"
)

var wakeBytes = []byte{1, 0, 0, 0, 0, 0, 0, 0}


type Poller struct {
	fd       int
	eventFd  int
	running  sync.Bool
	waitDone chan struct{}
}

func Create() (*Poller, error) {
	fd, err := unix.EpollCreate1(0)
	if err != nil {
		return nil, err
	}

	r0, _, errno := unix.Syscall(unix.SYS_EVENTFD2, 0, 0, 0)
	if errno != 0 {
		return nil, errno
	}
	eventFd := int(r0)

	// stub for wake loop
	err = unix.EpollCtl(fd, unix.EPOLL_CTL_ADD, eventFd, &unix.EpollEvent{
		Events: unix.EPOLLIN,
		Fd:     int32(eventFd),
	})
	if err != nil {
		_ = unix.Close(fd)
		_ = unix.Close(eventFd)
		return nil, err
	}

	return &Poller{
		fd:       fd,
		eventFd:  eventFd,
		waitDone: make(chan struct{}),
	}, nil
}

func (ep *Poller) Wake() error {
	_, err := unix.Write(ep.eventFd, wakeBytes)
	return err
}
