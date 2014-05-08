package main

import (
	"runtime/noos"
	"sync/atomic"
	"sync/barrier"
	"unsafe"
)

const (
	ready int32 = iota
	closed
	recvWait
	sendWait
)

type Elem int

type ElemS struct {
	e *Elem
}

// ChanS is a synchronous channel.
type ChanS struct {
	event noos.Event
	src   unsafe.Pointer
	dst   unsafe.Pointer
	state int32
}

func (c *ChanS) Close() {
	atomic.StoreInt32(&c.state, closed)
	c.event.Send()
}

func panicIfClosed(state int32) {
	if state == closed {
		panic("send: closed channel")
	}
}

func (c *ChanS) unlockAndSend(e *Elem) {
	// Get c.dst and unlock c.

	dst := c.dst
	// c may be closed - use CAS to avoids reopen.
	atomic.CompareAndSwapInt32(&c.state, recvWait, ready)
	barrier.Compiler()
	atomic.StorePointer(&c.dst, nil)
	atomic.StorePointer(&c.src, nil)
	c.event.Send()

	// Copy data and wake up a receiver.

	**(**Elem)(dst) = *e
	barrier.Memory()
	atomic.StorePointer((*unsafe.Pointer)(dst), nil)
	c.event.Send()
}

func (c *ChanS) unlockAndRecv(e *Elem) {
	// Get c.src and unlock c.

	src := c.src
	// c may be closed - use CAS to avoids reopen.
	atomic.CompareAndSwapInt32(&c.state, sendWait, ready)
	barrier.Compiler()
	atomic.StorePointer(&c.src, nil)
	atomic.StorePointer(&c.dst, nil)
	c.event.Send()

	// Copy the data and wake up a sender.

	*e = **(**Elem)(src)
	barrier.Memory()
	atomic.StorePointer((*unsafe.Pointer)(src), nil)
	c.event.Send()
}

// Send sends to c an element e.
// It blocks if there is no any receiver waiting for communication.
func (c *ChanS) Send(e *Elem) {
	panicIfClosed(atomic.LoadInt32(&c.state))

	for !atomic.CompareAndSwapPointer(&c.src, nil, unsafe.Pointer(&e)) {
		// Other sender locked this channel.
		c.event.Wait()
	}

	if atomic.CompareAndSwapInt32(&c.state, ready, sendWait) {
		// Passive sender.
		p := (*unsafe.Pointer)(unsafe.Pointer(&e))
		for {
			c.event.Wait()
			if atomic.LoadPointer(p) == nil {
				return
			}
			panicIfClosed(atomic.LoadInt32(&c.state))
		}
	}

	// Active sender

	panicIfClosed(atomic.LoadInt32(&c.state))
	c.unlockAndSend(e)
}

// Recv receives an element from c and saves it in a variable pointed by e.
// It blocks if there is no any sender waiting for communication.
func (c *ChanS) Recv(e *Elem) int32 {
	if atomic.LoadInt32(&c.state) == closed {
		return closed
	}

	for !atomic.CompareAndSwapPointer(&c.dst, nil, unsafe.Pointer(&e)) {
		// Other receiver locked this channel.
		c.event.Wait()
	}

	if atomic.CompareAndSwapInt32(&c.state, ready, recvWait) {
		// Passive receiver.
		p := (*unsafe.Pointer)(unsafe.Pointer(&e))
		for {
			c.event.Wait()
			if atomic.LoadPointer(p) == nil {
				return ready
			}
			if atomic.LoadInt32(&c.state) == closed {
				return closed
			}
		}
	}

	// Active receiver

	if atomic.LoadInt32(&c.state) == closed {
		return closed
	}
	c.unlockAndRecv(e)
	return ready
}

// TrySend tries to send to c an element of data pointed by e.
// It returns ready if sending success or sendWait if can't send data
// immediately.
func (c *ChanS) TrySend(e *Elem) int32 {
	panicIfClosed(atomic.LoadInt32(&c.state))

	if !atomic.CompareAndSwapPointer(&c.src, nil, unsafe.Pointer(&e)) {
		// Other sender locked this channel.
		return sendWait
	}

	if state := atomic.LoadInt32(&c.state); state != recvWait {
		// No receivres or c is closed.
		panicIfClosed(state)
		atomic.StorePointer(&c.src, nil)
		c.event.Send()
		return sendWait
	}

	c.unlockAndSend(e)
	return ready
}

// TryRecv tries to receive from c an element of data and save it in variable
// pointed by e.
// It returns ready if receiving success, closed if channel is closed or
// recvWait if can't receive data immediately.
func (c *ChanS) TryRecv(e *Elem) int32 {
	if atomic.LoadInt32(&c.state) == closed {
		return closed
	}

	if !atomic.CompareAndSwapPointer(&c.dst, nil, unsafe.Pointer(&e)) {
		// Other receiver locked this channel.
		return recvWait
	}

	if state := atomic.LoadInt32(&c.state); state != sendWait {
		// No senders or c is closed.
		if state == closed {
			return closed
		}
		atomic.StorePointer(&c.src, nil)
		c.event.Send()
		return recvWait
	}

	c.unlockAndSend(e)
	return ready
}
