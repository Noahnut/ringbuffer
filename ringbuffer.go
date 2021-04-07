package Ringbuffer

type ringbuffer struct {
	ringbufferArray []interface{}
	size            int
	front           int
	rear            int
	len             int
}

func RingBufferConstruct(size int) ringbuffer {
	ring := new(ringbuffer)
	ring.ringbufferArray = make([]interface{}, size)
	ring.rear = -1
	ring.size = size
	return *ring
}

func (this *ringbuffer) enqueue(value interface{}) bool {
	if !this.isFull() {
		this.rear = (this.rear + 1) % this.size
		this.ringbufferArray[this.rear] = value
		this.len++
	} else {
		return false
	}
	return true
}

func (this *ringbuffer) dequeue() bool {
	if !this.isEmpty() {
		this.front = (this.front + 1) % this.size
		this.len--
	} else {
		return false
	}
	return true
}

func (this *ringbuffer) GetFront() interface{} {
	if !this.isEmpty() {
		return this.ringbufferArray[this.front]
	}
	return nil
}

func (this *ringbuffer) GetRear() interface{} {
	if !this.isEmpty() {
		return this.ringbufferArray[this.rear]
	}
	return nil
}

func (this *ringbuffer) isEmpty() bool {
	if this.len == 0 {
		return true
	}
	return false
}

func (this *ringbuffer) isFull() bool {
	if this.len >= this.size {
		return true
	}
	return false
}
