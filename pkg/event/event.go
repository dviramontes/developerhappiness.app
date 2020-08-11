package event

type SlackEvent struct {}

type Handler interface {

}

func NewSlackEventHandler() *SlackEvent{
	return &SlackEvent{}
}
