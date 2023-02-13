package event


type Fast int

const (
	Normal Fast = iota
	FastCheese
    FastFish
    FastWine
    FastFull
)

func (f Fast) String() string {
	switch f {
		case Normal:
			return "Normal"
		case FastCheese:
			return "FastCheest"
		case FastFish:
			return "FastFish"
		case FastWine:
			return "FastWine"
		case FastFull:
			return "FastFull"
	}
	return "unknown"
}


type Event struct {
	saint string
	reading string
	fast Fast
}

func NewEvent(saint string, reading string, fast Fast) Event {
	event := Event{saint:saint, reading:reading, fast:fast}
	return event
}
