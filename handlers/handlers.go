package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	messagemap "github.com/DragonSSS/cloud-audition-interview/messagemap"
)

// CreateMessage is the handler
func CreateMessage(w http.ResponseWriter, r *http.Request) {
	var m messagemap.Message

	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resMsg, err := messagemap.AddMessage(m.Msg, isPalindrome(m.Msg))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	msgJSON, err := json.Marshal(resMsg)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(msgJSON)
}

// GetMessage is the handler
func GetMessage(w http.ResponseWriter, r *http.Request) {

}

// GetMessages is the handler
func GetMessages(w http.ResponseWriter, r *http.Request) {

}

// UpdateMessage is the handler
func UpdateMessage(w http.ResponseWriter, r *http.Request) {

}

// DeleteMessage is the handler
func DeleteMessage(w http.ResponseWriter, r *http.Request) {

}

func isPalindrome(s string) bool {
	// make sure lower cases
	s = strings.ToLower(s)
	// two pointers
	left, right := 0, len(s)-1
	for left < right {
		for left < right && !isChar(s[left]) {
			left++
		}

		for left < right && !isChar(s[right]) {
			right--
		}

		if s[left] != s[right] {
			return false
		}

		left++
		right--
	}
	return true
}

func isChar(c byte) bool {
	if ('a' <= c && c <= 'z') || ('0' <= c && c <= '9') {
		return true
	}
	return false
}
