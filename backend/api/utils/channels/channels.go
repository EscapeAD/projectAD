package channels

// OK channel ok
func OK(done <-chan bool) bool {
	select {
	case ok := <-done:
		if ok {
			return true
		}
	}
	return false
}
