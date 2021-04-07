package Ringbuffer

import (
	"testing"
)

func TestGetFront(t *testing.T) {
	ringb := RingBufferConstruct(10)
	ringb.enqueue(10)
	if ringb.GetFront().(int) != 10 {
		t.Error("not equal")
	}
}

func TestGetRear(t *testing.T) {

}

func TestIsFull(t *testing.T) {
}

func TestIsEmpty(t *testing.T) {

}

func TestEnqueue(t *testing.T) {

}

func TestDequeu(t *testing.T) {

}
