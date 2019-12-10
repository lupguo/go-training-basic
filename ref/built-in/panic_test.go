package built_in

import "testing"

func TestPanic(t *testing.T) {
	defer func() {
		if rcv := recover(); rcv != nil {
			t.Log("rcv recovery: ",rcv)
		} else {
			t.Log("rcv is nil")
		}
	}()
	M1()
}

func M1() {
	M2()
	panic("M1 Panic")
}

func M2() {
	defer func() {
		if rcv := recover(); rcv != nil {
			//panic("M2 Panic")
		}
	}()
	M3()
}

func M3() {
	panic("M3 Panic")
}
