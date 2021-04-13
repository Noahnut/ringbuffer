[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT) [![Build Status](https://www.travis-ci.com/Noahnut/ringbuffer.svg?branch=main)](https://www.travis-ci.com/Noahnut/ringbuffer)
# Simple ringbuffer
Simple thread-free ringbuffer implement by golang

## install
----
    go get github.com/Noahnut/ringbuffer
 
### Ringbuffer struct

```go
type ringbuff struct {
	ringbufferArray []interface{}
	size            int
	front           int
	rear            int
	len             int
	mux             *sync.RWMutex
}
```

### Enqueu
add element to the ringbuffer
if ringbuffer is full return false
```go
func (this *ringbuff) Enqueue(value interface{}) bool
```

### Dequeue
delete element from the ringbuffer
if ringbuffer is empty return false
```go
func (this *ringbuff) Dequeue() bool
```

### GetFront
get the top element on the queue
if the queue is empty return nil
```go
func (this *ringbuff) GetFront() interface{}
```

### GetRear
get the last element on the queue
if the queue is empty return nil
```go
func (this *ringbuff) GetRear() interface{}
```
### IsEmpty
check the ringbuffer is empty or not
if is empty return true
```go
func (this *ringbuff) IsEmpty() bool
```

### IsFull
check the ringbuffer is full or not
if is empty return true
```go
func (this *ringbuff) IsFull() bool
```
