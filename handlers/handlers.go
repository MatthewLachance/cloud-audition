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

// Message represents the model for a message
type Message struct {
	Msg string `json:"msg"`
}

// CreateMessage godoc
// @Summary Create a message
// @Description Create a new message with the content
// @Tags messages
// @Accept  json
// @Produce  json
// @Param Message body Message true "Create message"
// @Success 201 {object} messagemap.InternalMessage
// @Failure 415 "Content-Type header is not application/json"
// @Failure 400 "Failed to decode request body"
// @Failure 500 "Interal server failure"
// @Router /messages [post]
func CreateMessage(w http.ResponseWriter, r *http.Request) {
	var m Message

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

// GetMessage godoc
// @Summary Get details for a given messageID
// @Description Get details of message corresponding to the input messageID
// @Tags messages
// @Accept  json
// @Produce  json
// @Param messageID path int true "ID of the message"
// @Success 200 {object} messagemap.InternalMessage
// @Failure 400 "Failed to get valid parameter of messageID"
// @Failure 404 "Invalid id that doesn't exit in messages map"
// @Failure 500 "Interal server failure"
// @Router /messages/{messageID} [get]
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

// GetMessages godoc
// @Summary Get details of all messages
// @Description Get details of all messages
// @Tags messages
// @Accept  json
// @Produce  json
// @Success 200 {array} messagemap.InternalMessage
// @Router /messages [get]
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

// UpdateMessage godoc
// @Summary Update message identified by the given messageID
// @Description Update the message corresponding to the input messageID
// @Tags messages
// @Accept  json
// @Produce  json
// @Param messageID path int true "ID of the message to be updated"
// @Param Message body Message true "Update message"
// @Success 200 {object} messagemap.InternalMessage
// @Failure 415 "Content-Type header is not application/json"
// @Failure 400 "Invalid parameter or request body"
// @Failure 404 "Invalid id that doesn't exit in messages map"
// @Failure 500 "Interal server failure"
// @Router /messages/{messageID} [put]
func UpdateMessage(w http.ResponseWriter, r *http.Request) {

	if r.Header.Get("Content-Type") != "" {
		value, _ := header.ParseValueAndParams(r.Header, "Content-Type")
		if value != "application/json" {
			http.Error(w, "Content-Type header is not application/json", http.StatusUnsupportedMediaType)
			return
		}
	}

	var m Message

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

// DeleteMessage godoc
// @Summary Delete message identified by the given messageID
// @Description Delete the message corresponding to the input messageID
// @Tags messages
// @Accept  json
// @Produce  json
// @Param messageID path int true "ID of the message to be deleted"
// @Success 204 "No Content"
// @Failure 400 "Invalid parameter messageID"
// @Failure 404 "Invalid id that doesn't exit in messages map"
// @Router /messages/{messageID} [delete]
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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
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
