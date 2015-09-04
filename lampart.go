package lampartclock

import
(
	"errors"
	"fmt"
)

type Lampart struct {
	nodes map[string]*LampartClock
}

func New()*Lampart {
	lam := new(Lampart)
	lam.nodes = map[string]*LampartClock{}
	return lam
}

func (lam *Lampart) Add(title string) {
	lam.nodes[title] = new(LampartClock)
}

func (lam *Lampart) Send(title string)(error) {
	clock, ok := lam.nodes[title]
	if !ok {
		return errors.New(fmt.Sprintf("Node with name %s is not found", title))
	}
	clock.Inc()
	return nil
}
