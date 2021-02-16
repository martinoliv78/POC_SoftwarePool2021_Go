package humanity

import "fmt"

type Pilot struct {
	*Human
}

func (h *Pilot) String() string {
	return fmt.Sprintf("YAHOOOO")
}
