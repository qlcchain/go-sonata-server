/*
 * Copyright (c) 2019 QLC Chain Team
 *
 * This software is released under the MIT License.
 * https://opensource.org/licenses/MIT
 */

package event

import (
	"fmt"
	"reflect"
	"sync"
	"time"

	"github.com/rs/xid"

	"github.com/gammazero/workerpool"
)

const (
	defaultQueueSize         = 100
	eventBusWaitingQueueSize = 1024
)

// DefaultEventBus
type DefaultEventBus struct {
	handlers  sync.Map
	queueSize int
}

func (eb *DefaultEventBus) Close() error {
	eb.handlers.Range(func(key, value interface{}) bool {
		eb.CloseTopic(key.(string))
		return true
	})
	return nil
}

type handlerOption struct {
	isSync bool
}

type eventHandler struct {
	id       string
	callBack reflect.Value
	option   *handlerOption
	extra    *CallbackOption
	pool     *workerpool.WorkerPool
}

// New returns new DefaultEventBus with empty handlers.
func New() EventBus {
	b := &DefaultEventBus{
		queueSize: defaultQueueSize,
		handlers:  sync.Map{},
	}
	return EventBus(b)
}

func NewEventBus(queueSize int) EventBus {
	b := &DefaultEventBus{
		queueSize: queueSize,
	}
	return EventBus(b)
}

var (
	once  sync.Once
	eb    EventBus
	cache = sync.Map{}
)

func SimpleEventBus() EventBus {
	once.Do(func() {
		eb = New()
	})

	return eb
}

func GetEventBus(id string) EventBus {
	if id == "" {
		return SimpleEventBus()
	}

	if v, ok := cache.Load(id); ok {
		return v.(EventBus)
	}

	eb := New()
	cache.Store(id, eb)
	return eb
}

// doSubscribe handles the subscription logic and is utilized by the public Subscribe functions
func (eb *DefaultEventBus) doSubscribe(topic string, fn interface{}, callbackOption *CallbackOption, option *handlerOption) (string, error) {
	kind := reflect.TypeOf(fn).Kind()
	if kind != reflect.Func {
		return "", fmt.Errorf("%s is not of type reflect.Func", kind)
	}
	id := xid.New().String()
	if callbackOption == nil {
		callbackOption = &CallbackOption{}
	}
	callbackOption.ID = id
	handler := &eventHandler{
		id:       id,
		callBack: reflect.ValueOf(fn),
		option:   option,
		extra:    callbackOption,
		pool:     workerpool.New(eb.queueSize),
	}

	if value, ok := eb.handlers.Load(topic); ok {
		list := value.(*eventHandlers)
		list.Add(handler)
	} else {
		handlers := newEventHandler()
		handlers.Add(handler)
		eb.handlers.Store(topic, handlers)
	}
	return handler.id, nil
}

// Subscribe subscribes to a topic.
// Returns error if `fn` is not a function.
func (eb *DefaultEventBus) Subscribe(topic string, fn interface{}, option *CallbackOption) (string, error) {
	return eb.doSubscribe(topic, fn, option, &handlerOption{isSync: false})
}

// HasCallback returns true if exists any callback subscribed to the topic.
func (eb *DefaultEventBus) HasCallback(topic string) bool {
	if v, ok := eb.handlers.Load(topic); ok {
		handlers := v.(*eventHandlers)
		return handlers.Size() > 0
	}
	return false
}

// Close unsubscribe all handlers from given topic
func (eb *DefaultEventBus) CloseTopic(topic string) {
	if value, ok := eb.handlers.Load(topic); ok {
		value.(*eventHandlers).Clear()
		eb.handlers.Delete(topic)
	}
}

// Unsubscribe removes callback defined for a topic.
// Returns error if there are no callbacks subscribed to the topic.
func (eb *DefaultEventBus) Unsubscribe(topic string, handler string) error {
	if value, ok := eb.handlers.Load(topic); ok {
		if err := value.(*eventHandlers).RemoveCallback(handler); err != nil {
			return err
		}

		return nil
	}

	return fmt.Errorf("topic %s doesn't exist", topic)
}

// Publish executes callback defined for a topic. Any additional argument will be transferred to the callback.
func (eb *DefaultEventBus) Publish(topic string, args ...interface{}) {
	eb.handlers.Range(func(key, value interface{}) bool {
		topicPattern := key.(string)
		handlers := value.(*eventHandlers)
		if handlers.Size() > 0 && MatchSimple(topicPattern, topic) {
			all := handlers.All()
			for _, handler := range all {
				h := handler
				rArgs := eb.setUpPublish(h.extra, args...)

				// waiting until the queue is ready
				if h.pool.WaitingQueueSize() >= eventBusWaitingQueueSize {
					checkInterval := time.NewTicker(10 * time.Millisecond)

				checkWaitingQueueOut:
					for {
						select {
						case <-checkInterval.C:
							if h.pool.WaitingQueueSize() < eventBusWaitingQueueSize {
								checkInterval.Stop()
								break checkWaitingQueueOut
							}
						}
					}
				}

				if h.option.isSync {
					h.pool.SubmitWait(func() {
						h.callBack.Call(rArgs)
					})
				} else {
					h.pool.Submit(func() {
						h.callBack.Call(rArgs)
					})
				}
			}
		}

		return true
	})
}

func (eb *DefaultEventBus) setUpPublish(option *CallbackOption, args ...interface{}) []reflect.Value {
	passedArguments := make([]reflect.Value, 0)
	passedArguments = append(passedArguments, reflect.ValueOf(option))
	for _, arg := range args {
		passedArguments = append(passedArguments, reflect.ValueOf(arg))
	}
	return passedArguments
}
