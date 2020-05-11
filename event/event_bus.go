/*
 * Copyright (c) 2020. QLC Chain Team
 *
 * This software is released under the MIT License.
 * https://opensource.org/licenses/MIT
 */

package event

import (
	"io"
)

type CallbackOption struct {
	ID   string
	Data []interface{}
}

// subscriber defines subscription-related bus behavior
type subscriber interface {
	Subscribe(topic string, fn interface{}, option *CallbackOption) (string, error)
	Unsubscribe(topic string, handlerID string) error
}

// publisher defines publishing-related bus behavior
type publisher interface {
	Publish(topic string, args ...interface{})
}

// controller defines bus control behavior (checking handler's presence, synchronization)
type controller interface {
	HasCallback(topic string) bool
	CloseTopic(topic string)
}

type EventBus interface {
	subscriber
	publisher
	controller
	io.Closer
}
