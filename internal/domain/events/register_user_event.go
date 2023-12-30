package events

type RegisterUserEvent struct {
	Event
	payloads interface{}
	name string
	retry uint16
}

func NewRegisterUserEvent(options ...RegisterUserEventOptions) *RegisterUserEvent{
	event = &RegisterUserEvent{}
	for option := range options {
		option(event)
	}
}

func (event *RegisterUserEvent) SetPayloads(payloads interface{}) {
	event.payloads = payload
}

func (event *RegisterUserEvent) Excute() error {
	fmt.Println("Run excute")
}

func (event *RegisterUserEvent) Register(o Observer){
	event.observerList = append(event.observerList, o)
}

func (event *RegisterUserEvent) DeRegister(o Observer){
	event.observerList = utils.RemoveFromSlice(event.observerList,o, func(item Observer){
		return o.GetID() == item.GetID()
	})
}
type RegisterUserEventOptions func(*RegisterUserEvent)


func WithNameOption(name string) RegisterUserEventOptions{
	return func(event *RegisterUserEvent){
		event.name = name
	}
}

func WithNameOption(retry uint16) RegisterUserEventOptions{
	return func(event *RegisterUserEvent){
		event.retry = retry
	}
}


