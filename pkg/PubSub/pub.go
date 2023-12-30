package PubSub

type Pub interface {
	Notify()
	Register(o Observer)
	Deregister(o Observer)
}

type Observer interface {
	Excute(event Pub)
}