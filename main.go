package main

import (
	"fmt"
	"log"
)

type State interface {
	Run() (State, error)
}

type StartState struct{}

func (s *StartState) Run() (State, error) {
	fmt.Println("Hello from StartState!")
	return &OpenBookState{title: "1984"}, nil
}

type DoneState struct{}

func (s *DoneState) Run() (State, error) {
	fmt.Println("Hello from DoneState!")
	return nil, nil
}

type OpenBookState struct {
	title string
}

func (o *OpenBookState) Run() (State, error) {
	fmt.Println("Hello from OpenBookState!")
	fmt.Println("Opened a book named", o.title)
	return &ReadBookState{title: o.title}, nil
}

type ReadBookState struct {
	title string
}

func (r *ReadBookState) Run() (State, error) {
	fmt.Println("Hello from ReadBookState!")
	fmt.Println("Reading a book named", r.title)
	return &FlipPageState{title: r.title, page: 1}, nil
}

type FlipPageState struct {
	title string
	page  int
}

func (f *FlipPageState) Run() (State, error) {
	fmt.Println("Hello from FlipPageState!")
	if f.page <= 5 {
		fmt.Printf("On page %d of book %s\n", f.page, f.title)
		return &FlipPageState{title: f.title, page: f.page + 1}, nil
	}
	return &FinishBookState{title: f.title}, nil
}

type FinishBookState struct {
	title string
}

func (f *FinishBookState) Run() (State, error) {
	fmt.Println("Hello from FinishBookState!")
	fmt.Println("Finished reading a book named", f.title)
	return new(DoneState), nil
}

type StateMachine struct {
	state State
}

func (s *StateMachine) Run() error {
	for {
		state, err := s.state.Run()
		if err != nil {
			return err
		}
		if state == nil {
			fmt.Println("StateMachine is done!")
			return nil
		}
		s.state = state
	}
}

func NewStateMachine(startState State) *StateMachine {
	return &StateMachine{state: startState}
}

func main() {
	stateMachine := NewStateMachine(new(StartState))
	err := stateMachine.Run()
	if err != nil {
		log.Fatal(err)
	}
}
