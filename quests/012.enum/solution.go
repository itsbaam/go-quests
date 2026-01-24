package enum

type OrderState int

const (
	StateCreated OrderState = iota
	StatePaid
	StatePacked
	StateShipped
	StateDelivered
	StateCancelled
	StateReturned
)

var stateStatus = map[OrderState]string{
	StateCreated:   "created",
	StatePaid:      "paid",
	StatePacked:    "packed",
	StateShipped:   "shipped",
	StateDelivered: "delivered",
	StateCancelled: "cancelled",
	StateReturned:  "returned",
}

func (s OrderState) String() string {
	return stateStatus[s]
}

func NextState(current OrderState, action string) OrderState {
	if _, ok := stateStatus[current]; !ok {
		panic("unknow state")
	}

	if current == StateCancelled || current == StateReturned {
		return current
	}

	switch current {
	case StateCreated:
		switch action {
		case "pay":
			return StatePaid
		case "cancel":
			return StateCancelled
		}
	case StatePaid:
		switch action {
		case "pack":
			return StatePacked
		case "cancel":
			return StateCancelled
		}
	case StatePacked:
		switch action {
		case "ship":
			return StateShipped
		}
	case StateShipped:
		switch action {
		case "deliver":
			return StateDelivered
		case "return":
			return StateReturned
		}
	case StateDelivered:
		switch action {
		case "return":
			return StateReturned
		}
	}

	return current
}
