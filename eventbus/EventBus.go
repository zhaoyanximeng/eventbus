package eventbus

import "sync"

type EventBus struct {
	subscribes map[string]EventDataChannel
	lock sync.Mutex
}

func NewEventBus() *EventBus {
	return &EventBus{subscribes: make(map[string]EventDataChannel)}
}

// 订阅对应的事件
func (this *EventBus) Sub(topic string) EventDataChannel {
	this.lock.Lock()
	defer this.lock.Unlock()
	if e,ok := this.subscribes[topic] ; ok {
		return e
	} else {
		ch := make(EventDataChannel)
		this.subscribes[topic] = ch

		return ch
	}
}

func (this *EventBus) Pub(topic string,data interface{}) {
	this.lock.Lock()
	defer this.lock.Unlock()
	if e,ok := this.subscribes[topic]; ok {
		go func() {
			e <- &EventData{Data: data}
		}()
	}
}


