package registry

import (
	"fmt"
	"sync"
	"time"
)

type CallBack func(event AsyncResultEvent) error

type Listener interface {
	OnEvent(event Event) error
	OnEventAsync(event Event) CallBack
}

type Subject interface {
	AddObserver(observer Listener)
	RemoveObserver(observer Listener)
	NotifyObservers(event Event)
	GetListeners() []Listener
}

type ConcreteSubject struct {
	observers []Listener
	mutex     sync.Mutex
}

func (s *ConcreteSubject) GetListeners() []Listener {
	return s.observers
}

func (s *ConcreteSubject) AddObserver(observer Listener) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.observers = append(s.observers, observer)
}

func (s *ConcreteSubject) RemoveObserver(observer Listener) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for i, o := range s.observers {
		if o == observer {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		}
	}
}

func (s *ConcreteSubject) NotifyObservers(event Event) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for _, observer := range s.observers {
		if event.Async {
			observer.OnEvent(event)
		} else {
			result := make(chan AsyncResultEvent)
			var back CallBack
			go func() {
				// 注意panic和error
				back = observer.OnEventAsync(event)
				result <- AsyncResultEvent{
					Code: Success,
					Data: nil,
				}
			}()
			select {
			case res := <-result:
				fmt.Println("received result:", res)
				back(res)
			case <-time.After(5000 * time.Millisecond):
				fmt.Println("timeout")
				res := AsyncResultEvent{
					Code: TimeOut,
					Data: nil,
				}
				back(res)
			}
		}
	}

}

type ConcreteObserver struct {
	subject Subject
}

func (o *ConcreteObserver) OnEvent(event Event) error {
	return nil
}

func (o *ConcreteObserver) OnEventAsync(event Event) CallBack {

	// init callback
	cb := CallBack(myCallback)
	return cb
}

func myCallback(event AsyncResultEvent) error {
	fmt.Println("Result:", event.Code)
	fmt.Println("Error:", event.Data)
	return nil
}

// 初始化所有的listener
func init() {

}
