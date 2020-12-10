package tson

import (
	"fmt"
	"reflect"
	"sync"
	"time"
)

var format = struct {
	layout string
	mutex  *sync.Mutex
}{
	time.RFC3339,
	new(sync.Mutex),
}

// SetLayout - set the layout
func SetLayout(layout string) {
	format.mutex.Lock()
	format.layout = layout
	format.mutex.Unlock()
}

// Time -  alias for time.Time
type Time struct {
	time.Time
}

// UnmarshalJSON - unmarshal JSON
func (t *Time) UnmarshalJSON(data []byte) error {
	if string(data) == "null" || len(data) == 0 {
		t.Time = time.Time{}
		return nil
	}

	layout := fmt.Sprintf(`"%s"`, format.layout)

	tm, err := time.Parse(layout, string(data))
	t.Time = tm
	return err
}

var rtt = reflect.TypeOf(&Time{})
