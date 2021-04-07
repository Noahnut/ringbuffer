package ringbuffer

import (
	"testing"
)

func TestIntCase(t *testing.T) {
	ringb := RingBufferConstruct(5)
	if ringb.IsEmpty() != true {
		t.Error("The ring buffer should be empty")
	}
	ringb.Enqueue(10)
	if ringb.GetFront().(int) != 10 {
		t.Errorf("The result %d not equal %d", ringb.GetFront(), 10)
	}
	ringb.Enqueue(20)
	if ringb.GetRear().(int) != 20 {
		t.Errorf("The result %d not equal %d", ringb.GetRear(), 20)
	}
	ringb.Enqueue(30)
	if ringb.GetRear().(int) != 30 {
		t.Errorf("The result %d not equal %d", ringb.GetRear(), 30)
	}
	ringb.Enqueue(40)
	if ringb.GetRear().(int) != 40 {
		t.Errorf("The result %d not equal %d", ringb.GetRear(), 40)
	}
	ringb.Enqueue(50)
	if ringb.GetRear().(int) != 50 {
		t.Errorf("The result %d not equal %d", ringb.GetRear(), 50)
	}
	if ringb.Enqueue(60) != false {
		t.Error("The ring buffer should be full")
	}
	if ringb.IsFull() != true {
		t.Error("The ring buffer should be full")
	}
	ringb.Dequeue()
	if ringb.GetFront().(int) != 20 {
		t.Errorf("The result %d not equal %d", ringb.GetFront(), 20)
	}

	ringb.Enqueue(60)
	if ringb.GetRear() != 60 {
		t.Errorf("The result %d not equal %d", ringb.GetRear(), 60)
	}
}

func TestLargeSizeInt(t *testing.T) {
	ringb := RingBufferConstruct(1024)
	for i := 0; i < 1024; i++ {
		ringb.Enqueue(i)
	}
	if ringb.GetFront().(int) != 0 {
		t.Errorf("The result %d not equal %d", ringb.GetFront(), 0)
	}

	if ringb.GetRear().(int) != 1023 {
		t.Errorf("The result %d not equal %d", ringb.GetRear(), 1023)
	}

	if ringb.IsFull() != true {
		t.Error("The ring buffer should be full")
	}
}

func TestStringCase(t *testing.T) {
	ringb := RingBufferConstruct(5)
	ringb.Enqueue("testOne")
	if ringb.GetFront().(string) != "testOne" {
		t.Errorf("The result %s not equal %s", ringb.GetFront(), "testOne")
	}

	ringb.Enqueue("testtwo")
	if ringb.GetRear().(string) != "testtwo" {
		t.Errorf("The result %s not equal %s", ringb.GetFront(), "testtwo")
	}

	ringb.Enqueue("testthree")
	if ringb.GetRear().(string) != "testthree" {
		t.Errorf("The result %s not equal %s", ringb.GetFront(), "testthree")
	}

	ringb.Enqueue("testfour")
	if ringb.GetRear().(string) != "testfour" {
		t.Errorf("The result %s not equal %s", ringb.GetFront(), "testfour")
	}

	ringb.Enqueue("testfive")
	if ringb.GetRear().(string) != "testfive" {
		t.Errorf("The result %s not equal %s", ringb.GetFront(), "testfive")
	}

	if ringb.IsFull() != true {
		t.Error("The ring buffer should be full")
	}

	ringb.Dequeue()
	if ringb.GetFront().(string) != "testtwo" {
		t.Errorf("The result %s not equal %s", ringb.GetFront(), "testtwo")
	}

	ringb.Enqueue("testsix")
	if ringb.GetRear().(string) != "testsix" {
		t.Errorf("The result %s not equal %s", ringb.GetFront(), "testsix")
	}
}

type testStruct struct {
	value int
	array []string
}

func TestArrayCase(t *testing.T) {

	ringb := RingBufferConstruct(5)

	testOne := new(testStruct)
	testOne.value = 1
	testOne.array = make([]string, 10)

	ringb.Enqueue(testOne)
	if ringb.GetFront().(*testStruct).value != 1 && len(ringb.GetFront().(*testStruct).array) == 10 {
		t.Errorf("The result %d not equal %d", ringb.GetFront().(*testStruct).value, 1)
	}

	testTwo := new(testStruct)
	testTwo.value = 2
	testTwo.array = make([]string, 15)

	ringb.Enqueue(testOne)
	if ringb.GetFront().(*testStruct).value != 2 && len(ringb.GetFront().(*testStruct).array) == 15 {
		t.Errorf("The result %d not equal %d", ringb.GetFront().(*testStruct).value, 2)
	}

	testThree := new(testStruct)
	testThree.value = 3
	testThree.array = make([]string, 20)

	ringb.Enqueue(testOne)
	if ringb.GetFront().(*testStruct).value != 3 && len(ringb.GetFront().(*testStruct).array) == 20 {
		t.Errorf("The result %d not equal %d", ringb.GetFront().(*testStruct).value, 3)
	}

	testFour := new(testStruct)
	testFour.value = 4
	testFour.array = make([]string, 25)

	ringb.Enqueue(testOne)
	if ringb.GetFront().(*testStruct).value != 4 && len(ringb.GetFront().(*testStruct).array) == 25 {
		t.Errorf("The result %d not equal %d", ringb.GetFront().(*testStruct).value, 4)
	}

	testFive := new(testStruct)
	testFive.value = 5
	testFive.array = make([]string, 30)

	ringb.Enqueue(testOne)
	if ringb.GetFront().(*testStruct).value != 4 && len(ringb.GetFront().(*testStruct).array) == 30 {
		t.Errorf("The result %d not equal %d", ringb.GetFront().(*testStruct).value, 5)
	}

	if ringb.IsFull() != true {
		t.Error("The ring buffer should be full")
	}
}
