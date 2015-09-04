package lampartclock

import
(
	"errors"
	"fmt"
)

type Lampart struct {
	nodes map[string]*LampartClock
}

//New provides creating a new object
func New()*Lampart {
	lam := new(Lampart)
	lam.nodes = map[string]*LampartClock{}
	return lam
}

//Add provides store name of node
func (lam *Lampart) Add(title string) {
	lam.nodes[title] = new(LampartClock)
}

//Send provides increment counter of node with title
func (lam *Lampart) Send(title string)error {
	clock, ok := lam.nodes[title]
	if !ok {
		return errors.New(fmt.Sprintf("Node with name %s is not found", title))
	}
	clock.Inc()
	return nil
}

//SenTo provides synchronization between two clocks
func (lam *Lampart) Sync(titlefrom string, titleto string) error {
	clock1, ok := lam.nodes[titlefrom]
	if !ok {
		return errors.New(fmt.Sprintf("Node with name %s is not found", titlefrom))
	}
	clock2, ok2 := lam.nodes[titleto]
	if !ok2 {
		return errors.New(fmt.Sprintf("Node with name %s is not found", titleto))
	}

	clock2.Set(clock1.Get())
	return nil
}

//GetCounter provides getting of counter from node with title
func (lam *Lampart) GetCounter(title string) (uint64, error) {
	clock, ok := lam.nodes[title]
	if !ok {
		return 0, errors.New(fmt.Sprintf("Node with name %s is not found", title))
	}

	return clock.Get(), nil
}
