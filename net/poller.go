package net

type channel struct {
}

type activeChannels struct {
}
type Poller interface {
	Poller(*EventLoop)
	// Polls the I/o events
	// Must be called in the loop thread
	poll(timeoutMs int, Channelist *activeChannels)

	// CHanges the interested I/O events
	// Must be called in the loop thread
	updateChannel(*channel)
}
