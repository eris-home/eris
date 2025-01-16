package manager

import "fmt"

type Manager interface {
	Initialize() error
	Process() error
	Shutdown() error
}

type BaseManager struct {
	Name string
}

func NewBaseManager(name string) *BaseManager {
	return &BaseManager{Name: name}
}

func (bm *BaseManager) Initialize() error {
	fmt.Printf("BaseManager [%s]: Initialization complete\n", bm.Name)
	return nil
}

func (bm *BaseManager) Process() error {
	fmt.Printf("BaseManager [%s]: Processing tasks\n", bm.Name)
	return nil
}

func (bm *BaseManager) Shutdown() error {
	fmt.Printf("BaseManager [%s]: Shutdown complete\n", bm.Name)
	return nil
}