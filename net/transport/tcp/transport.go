package tcp

import "golang.org/x/sys/unix"

type Connection struct {
	Fd       int
	SocketAddr unix.Sockaddr
	//handleC  HandleConnFunc
	//listener net.Listener
	//loop     *eventloop.EventLoop
}





