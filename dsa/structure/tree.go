package main

import "fmt"

type TreeNode struct {
	left, right *TreeNode
	value       int
}

type Block interface {
	BlockSize() int
	Encrypt(src, dst []byte)
	Decrypt(src, dst []byte)
}

type Locker interface {
	Lock()
	Unlock()
}

//Mutex
type Mutex struct {
}

func (m *Mutex) Lock() {
	fmt.Println("Try Lock")
}

func (m *Mutex) Unlock() {
	fmt.Println("Try UnLock")
}

//NewMutex
type NewMutex Mutex

//PtrMutex 
type PtrMutex *Mutex

type StMutex struct {
	Mutex
}

func main() {
	mx := new(Mutex)
	mx.Lock()
	defer mx.Unlock()

	//new mutex
	nmx := new(NewMutex)
	_ = nmx

	// new pointer mutex
	pmx := new(PtrMutex)
	_ = pmx

	// new struct mutex
	smx := new(StMutex)
	smx.Lock()
}
