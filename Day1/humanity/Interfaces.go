package humanity

import "fmt"

type Preparer interface {
	Prepare() error
}

func (h *Human) Prepare() error {
	h.Ready = true
	fmt.Printf("Human: %s, %d years old from %s is ready !\n", h.Name, h.Age, h.Country)
	return nil
}

func PrepareMissionPart(objs ...Preparer) error {

	for _, act := range objs {
		act.Prepare()
	}
	return nil
}

type Checker interface {
	Check() bool
}

func (h *Human) Check() bool {
	if h.Ready == false {
		fmt.Printf("Human: %s, %d years old from %s is not ready !\n", h.Name, h.Age, h.Country)
	}
	return h.Ready
}

func CheckMissionPart(objs ...Checker) bool {
	var curr bool = true
	for _, act := range objs {
		if act.Check() == false {
			curr = false
		}
	}
	return curr
}

type PrepareChecker interface {
	Preparer
	Checker
}

func PrepareAndCheckMissionPart(objs ...PrepareChecker) bool {
	var curr bool = true
	for _, act := range objs {
		act.Prepare()
		if act.Check() == false {
			curr = false
		}
	}
	return curr
}
