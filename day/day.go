package day

import (
	"orthocal/event"
)

type Day struct {
    day int 
    event event.Event
}

func (d Day) GetDay() int {
	return d.day
}

func NewDay (day_num int, saint string, reading string, fast event.Fast ) Day {
	event := event.NewEvent(saint, reading, fast)
	day_object := Day{day:day_num, event:event}
	return day_object
}
