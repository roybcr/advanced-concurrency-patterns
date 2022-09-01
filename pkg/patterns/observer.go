package patterns

import (
	"fmt"
	"time"
)

// The Observer Pattern defines one-to-many dependency between a Notifier and Observers.
// When Notifier changes its state, all Observers are notified with Events.

type Event struct {
	Data interface{}
}

type Observer interface {
	OnNotify(Event)
}

type Notifier interface {
	Subscribe(Observer)
	Unsubscribe(Observer)
	Notify(Event)
}

type observer struct {
	id string
}

type notifier struct {
	observers map[Observer]struct{}
}

func (o *observer) OnNotify(e Event) {
	fmt.Printf("Observer %s recieved event %v\n", (*o).id, e.Data)
}

func (n *notifier) Subscribe(o Observer) {
	(*n).observers[o] = struct{}{}
}

func (n *notifier) Unsubscribe(o Observer) {
	delete((*n).observers, o)
}

func (n *notifier) Notify(e Event) {
	for o := range (*n).observers {
		o.OnNotify(e)
	}
}

func Observe() {

	n := notifier{
		observers: map[Observer]struct{}{},
	}

	o1 := observer{id: "001"}
	o2 := observer{id: "002"}

	defer n.Unsubscribe(&o1)
	defer n.Unsubscribe(&o2)

	n.Subscribe(&o1)
	n.Subscribe(&o2)

	n.Notify(Event{Data: "Mouse Click"})
	time.Sleep(time.Second * 2)

	n.Notify(Event{Data: 23})
	time.Sleep(time.Second * 4)

	n.Notify(Event{Data: "New User Registered!"})

}
