package messagemap

import "sync"

// Message struct
type Message struct {
	ID           int    `json:"id"`
	Msg          string `json:"msg"`
	IsPalindrome bool   `json:"isPalindrome"`
}

var (
	mu sync.Mutex
	id int = 0
)

var messagemap = struct {
	sync.RWMutex
	m map[int]*Message
}{m: make(map[int]*Message)}

func generateID() int {
	mu.Lock()
	id++
	res := id
	mu.Unlock()
	return res
}

// AddMessage is a func that
func AddMessage(msg string, isPalindrome bool) (*Message, error) {
	id := generateID()

	message := &Message{
		ID:           id,
		Msg:          msg,
		IsPalindrome: isPalindrome,
	}

	messagemap.Lock()
	messagemap.m[id] = message
	messagemap.Unlock()

	return message, nil
}

/*
func UpdateMessage(msg string, id string, isPalindrome bool) (*Message, error) {

}

func GetMessage(id string) (*Message, error) {

}

func GetMessages() {

}

func DeleteMessage(id string) error {

}*/
