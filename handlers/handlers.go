package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	messagemap "github.com/DragonSSS/cloud-audition-interview/messagemap"
	"github.com/golang/gddo/httputil/header"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

var errorMissingMessageID = errors.New("messageID parameter is missing in request")

// CreateMessage is the handler that creates message
func CreateMessage(w http.ResponseWriter, r *http.Request) {
	var m messagemap.Message

	if r.Header.Get("Content-Type") != "" {
		value, _ := header.ParseValueAndParams(r.Header, "Content-Type")
		if value != "application/json" {
			http.Error(w, "Content-Type header is not application/json", http.StatusUnsupportedMediaType)
			return
		}
	}

	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
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
	logger(w.Write(msgJSON))
}

// GetMessage is the handler that gets message with id
func GetMessage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	param, ok := vars["messageID"]

	if !ok {
		http.Error(w, errorMissingMessageID.Error(), http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(param)

	if err != nil {
		http.Error(w, "Failed to get valid parameter of messageID", http.StatusBadRequest)
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
	w.WriteHeader(http.StatusOK)
	logger(w.Write(msgJSON))
}

// GetMessages is the handler that gets all messages
func GetMessages(w http.ResponseWriter, r *http.Request) {
	resMsgs := messagemap.GetMessages()

	msgJSON, err := json.Marshal(resMsgs)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	logger(w.Write(msgJSON))
}

// UpdateMessage is the handler that updates with id and message
func UpdateMessage(w http.ResponseWriter, r *http.Request) {

	if r.Header.Get("Content-Type") != "" {
		value, _ := header.ParseValueAndParams(r.Header, "Content-Type")
		if value != "application/json" {
			http.Error(w, "Content-Type header is not application/json", http.StatusUnsupportedMediaType)
			return
		}
	}

	var m messagemap.Message

	vars := mux.Vars(r)
	param, ok := vars["messageID"]

	if !ok {
		http.Error(w, errorMissingMessageID.Error(), http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(param)

	if err != nil {
		http.Error(w, "Failed to get valid parameter of messageID", http.StatusBadRequest)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	resMsg, err := messagemap.UpdateMessage(m.Msg, id, isPalindrome(m.Msg))

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
	w.WriteHeader(http.StatusOK)
	logger(w.Write(msgJSON))
}

// DeleteMessage is the handler that deletes message with id
func DeleteMessage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	param, ok := vars["messageID"]

	if !ok {
		http.Error(w, errorMissingMessageID.Error(), http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(param)

	if err != nil {
		http.Error(w, "Failed to get valid parameter of messageID", http.StatusBadRequest)
		return
	}

	if messagemap.DeleteMessage(id) != nil {
		http.Error(w, messagemap.ErrorNoSuchKey.Error(), http.StatusNotFound)
		return
	}

	msg := map[string]string{"result": "success"}

	res, _ := json.Marshal(msg)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	logger(w.Write(res))
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

func logger(n int, err error) {
	if err != nil {
		log.Error(err)
	}
}
