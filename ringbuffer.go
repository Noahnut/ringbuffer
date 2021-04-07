package ringbuffer

import "sync"

type ringbuff struct {
	ringbufferArray []interface{}
	size            int
	front           int
	rear            int
	len             int
	mux             *sync.RWMutex
}

func RingBufferConstruct(size int) ringbuff {
	ring := new(ringbuff)
	ring.ringbufferArray = make([]interface{}, size)
	ring.rear = -1
	ring.size = size
	ring.mux = &sync.RWMutex{}
	return *ring
}

func (this *ringbuff) enqueue(value interface{}) bool {
	if !this.isFull() {
		this.mux.Lock()
		this.rear = (this.rear + 1) % this.size
		this.ringbufferArray[this.rear] = value
		this.len++
		this.mux.Unlock()
	} else {
		return false
	}
	return true
}

func (this *ringbuff) dequeue() bool {
	if !this.isEmpty() {
		this.mux.Lock()
		this.front = (this.front + 1) % this.size
		this.len--
		this.mux.Unlock()
	} else {
		return false
	}
	return true
}

func (this *ringbuff) GetFront() interface{} {
	this.mux.RLock()
	defer this.mux.RUnlock()
	if !this.isEmpty() {
		return this.ringbufferArray[this.front]
	}
	return nil
}

func (this *ringbuff) GetRear() interface{} {
	this.mux.RLock()
	defer this.mux.RUnlock()
	if !this.isEmpty() {
		return this.ringbufferArray[this.rear]
	}
	return nil
}

func (this *ringbuff) isEmpty() bool {
	this.mux.RLock()
	defer this.mux.RUnlock()
	if this.len == 0 {
		return true
	}

	return false
}

func (this *ringbuff) isFull() bool {
	this.mux.RLock()
	defer this.mux.RUnlock()
	if this.len >= this.size {
		return true
	}
	return false
}
