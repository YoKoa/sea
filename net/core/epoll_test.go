// +build linux

package core

import (
	"golang.org/x/sys/unix"
	"testing"
)

func TestEpoll(t *testing.T) {
	fd, _, _ := unix.Syscall(unix.SYS_GETPID, 0, 0, 0) // 用不到的就补上 0
	r0, _, _ := unix.Syscall(unix.SYS_EVENTFD2, 0, 0, 0)

	eventFd := int(r0)

	_ = unix.EpollCtl(int(fd), unix.EPOLL_CTL_ADD, eventFd, &unix.EpollEvent{
		Events: unix.EPOLLET | unix.EPOLLIN,
		Fd:     int32(eventFd),
	})




}
