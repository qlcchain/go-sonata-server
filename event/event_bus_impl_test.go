/*
 * Copyright (c) 2020. QLC Chain Team
 *
 * This software is released under the MIT License.
 * https://opensource.org/licenses/MIT
 */

package event

import (
	"fmt"
	"sync"
	"testing"
)

func TestDefaultEventBus_Unsubscribe(t *testing.T) {
	eb := New()
	wg := sync.WaitGroup{}

	if id, err := eb.Subscribe("test", func(o *CallbackOption, i int) {
		defer wg.Done()
		fmt.Println(o.ID, "-> ", i)
	}, nil); err != nil {
		t.Fatal(err)
	} else {
		t.Log(id)
	}

	eb.Publish("test", 1)
	wg.Add(1)
	wg.Wait()
}
