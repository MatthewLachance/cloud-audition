package messagemap

import (
	"errors"
	"sync"
)

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

// ErrorNoSuchKey is the error of non-existing key
var ErrorNoSuchKey = errors.New("invalid id that doesn't exit in messages map")

func generateID() int {
	mu.Lock()
	id++
	res := id
	mu.Unlock()
	return res
}

// CreateMessage is the func that adds message into map
func CreateMessage(msg string, isPalindrome bool) *Message {
	id := generateID()

	message := &Message{
		ID:           id,
		Msg:          msg,
		IsPalindrome: isPalindrome,
	}

	messagemap.Lock()
	messagemap.m[id] = message
	messagemap.Unlock()

	return message
}

// UpdateMessage is the func that updates message with id and content into map
func UpdateMessage(msg string, id int, isPalindrome bool) (*Message, error) {

	// check if id exist in map
	messagemap.RLock()
	_, found := messagemap.m[id]
	messagemap.RUnlock()

	if !found {
		return &Message{}, ErrorNoSuchKey
	}

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

// GetMessage is the func that gets message with id from map
func GetMessage(id int) (*Message, error) {

	// check if id exist in map
	messagemap.RLock()
	_, found := messagemap.m[id]
	messagemap.RUnlock()

	if !found {
		return &Message{}, ErrorNoSuchKey
	}

	messagemap.RLock()
	message := messagemap.m[id]
	messagemap.RUnlock()

	return message, nil
}

// GetMessages is the func that gets all messages from map
func GetMessages() []Message {
	messagemap.RLock()
	size := len(messagemap.m)

	if size == 0 {
		return []Message{}
	}

	res := make([]Message, len(messagemap.m))
	index := 0

	for _, value := range messagemap.m {
		res[index] = *value
		index++
	}
	messagemap.RUnlock()

	return res
}

// DeleteMessage is the func that deletes message from map
func DeleteMessage(id int) error {
	// check if id exist in map
	messagemap.RLock()
	_, found := messagemap.m[id]
	messagemap.RUnlock()

	if !found {
		return ErrorNoSuchKey
	}

	messagemap.Lock()
	delete(messagemap.m, id)
	messagemap.Unlock()
	return nil
}

// CleanMap is the func that makes map empty
func CleanMap() {
	messagemap.Lock()
	messagemap.m = make(map[int]*Message)
	messagemap.Unlock()

	mu.Lock()
	id = 0
	mu.Unlock()
}
