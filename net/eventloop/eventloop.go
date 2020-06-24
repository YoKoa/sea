package eventloop

import (
	"context"
	"github.com/YoKoa/sea/net/core"
	"github.com/YoKoa/sea/net/transport/tcp"
	"golang.org/x/sys/unix"
	"net"
	"sync"
)

var listenConfig = net.ListenConfig{
	Control: tcp.Control,
}

type EventLoop struct {
	sockets sync.Map
	poll    *core.Poller
}

func New() *EventLoop {
	poller, err := core.Create()
	if err != nil {
		panic(err)
	}
	return &EventLoop{
		sockets: sync.Map{},
		poll:    poller,
	}
}

// linux3.9+
func (el *EventLoop) Run(network, address string) {

	// 创建accepter
	listener, err := listenConfig.Listen(context.Background(), network, address)
	if err != nil {
		el.poll.Close()
		panic("reuseport fail")
	}
	// 加入epoll
	l, ok := listener.(*net.TCPListener)
	if !ok {
		panic("could not get file descriptor 1")
	}
	file, err := l.File()
	if err != nil {
		panic(err)
	}
	fd := int(file.Fd())
	if err = unix.SetNonblock(fd, true); err != nil {
		panic(err)
	}
	err = el.AddSocketAndEnableRead(fd, tcp.Connection{
		Fd: fd,
	})
	if err != nil {
		panic(err)
	}
	// 开始wait
	el.poll.Poll(func(fd int, event core.Event) {
		if fd == 0 {
			return
		}

		socket, ok := el.sockets.Load(fd)
		// 接受连接
		if !ok && event&core.EventRead != 0 {
			nfd, sa, err := unix.Accept(fd)
			if err != nil {
				if err == unix.EAGAIN {
					return
				}
				return
			}
			if err := unix.SetNonblock(nfd, true); err != nil {
				return
			}
			el.sockets.Store(fd, &tcp.Connection{
				Fd:         fd,
				SocketAddr: sa,
			})
			if err = el.poll.AddRead(fd); err != nil {
				el.poll.Del(fd)
			}
		}
		if ok && event&core.EventRead != 0 {
			buf := make([]byte, 0xFFFF)
			unix.Read(socket.(*tcp.Connection).Fd, buf)
		}

		if ok && event&core.EventWrite != 0 {

		}

	})
}

// AddSocketAndEnableRead 增加 Socket 到时间循环中，并注册可读事件
func (el *EventLoop) AddSocketAndEnableRead(fd int, s tcp.Connection) error {
	var err error
	el.sockets.Store(fd, s)

	if err = el.poll.AddRead(fd); err != nil {
		el.sockets.Delete(fd)
		return err
	}
	return nil
}
