package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	messagemap "github.com/DragonSSS/cloud-audition-interview/messagemap"
	"github.com/gorilla/mux"
)

var errorMissingMessageID = errors.New("messageID parameter is missing in request")

// CreateMessage is the handler
func CreateMessage(w http.ResponseWriter, r *http.Request) {
	var m messagemap.Message

	// if there is no body
	// TODO imple

	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resMsg := messagemap.CreateMessage(m.Msg, isPalindrome(m.Msg))

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
	vars := mux.Vars(r)
	param, ok := vars["messageID"]

	if !ok {
		http.Error(w, errorMissingMessageID.Error(), http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(param)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resMsg, err := messagemap.GetMessage(id)

	if err != nil {
		http.Error(w, messagemap.ErrorNoSuchKey.Error(), http.StatusNotFound)
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
	// two pointers, O(len(s))
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
