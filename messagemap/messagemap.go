package messagemap

import "sync"

type message struct {
	id           int
	message      string
	isPalindrome bool
}

var (
	mu sync.Mutex
	id int = 0
)

var messagemap = struct {
	sync.RWMutex
	m map[int]*message
}{m: make(map[int]*message)}

func generateID() int {
	mu.Lock()
	id++
	res := id
	mu.Unlock()
	return res
}

func AddMessage(msg string, isPalindrome bool) {
	id := generateID()
}
