package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DragonSSS/cloud-audition-interview/messagemap"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	t.Run("abc", testIsPalindromeFunc("abc", false))
	t.Run("1aaa1", testIsPalindromeFunc("aaa", true))
	t.Run("1aa a 1", testIsPalindromeFunc("aa a", true))
	t.Run("1a,a a 1", testIsPalindromeFunc("a,a a ", true))
	t.Run(" ", testIsPalindromeFunc(" ", true))
	t.Run("", testIsPalindromeFunc("", true))
}

func testIsPalindromeFunc(message string, expected bool) func(*testing.T) {
	return func(t *testing.T) {
		actual := isPalindrome(message)
		str := fmt.Sprintf("expect message %s isPalindrome to be %t, got %t", message, expected, actual)
		assert.Equal(t, expected, actual, str)
	}
}

func TestCreateMessageHandler(t *testing.T) {
	var jsonStr = []byte(`{"msg":"aaa"}`)
	r, _ := http.NewRequest("POST", "/messages", bytes.NewBuffer(jsonStr))
	w := httptest.NewRecorder()

	CreateMessage(w, r)

	resp := w.Result()
	assert.Equal(t, 201, resp.StatusCode, "expected 201 status code")
	messagemap.CleanMap()
}

func TestGetMessageHandler(t *testing.T) {
	expectedMsg := "message"
	expectedIsPalindrome := false

	expectedMessage := messagemap.CreateMessage(expectedMsg, expectedIsPalindrome)
	var actualMessage messagemap.InternalMessage

	r, _ := http.NewRequest("GET", "/messages/1", nil)
	w := httptest.NewRecorder()

	vars := map[string]string{
		"messageID": "1",
	}
	r = mux.SetURLVars(r, vars)

	GetMessage(w, r)

	resp := w.Result()
	err := json.NewDecoder(resp.Body).Decode(&actualMessage)
	assert.Nil(t, err)
	assert.Equal(t, expectedMessage.ID, actualMessage.ID, "expected 1 message id")
	assert.Equal(t, 200, resp.StatusCode, "expected 200 status code")
	messagemap.CleanMap()
}

func TestUpdateMessageHandler(t *testing.T) {
	expectedMsg := "message"
	expectedIsPalindrome := false

	expectedMessage := messagemap.CreateMessage(expectedMsg, expectedIsPalindrome)
	var actualMessage messagemap.InternalMessage

	var jsonStr = []byte(`{"msg":"aaa"}`)
	r, _ := http.NewRequest("PUT", "/messages/1", bytes.NewBuffer(jsonStr))
	w := httptest.NewRecorder()

	vars := map[string]string{
		"messageID": "1",
	}
	r = mux.SetURLVars(r, vars)

	UpdateMessage(w, r)

	resp := w.Result()
	err := json.NewDecoder(resp.Body).Decode(&actualMessage)
	assert.Nil(t, err)
	assert.Equal(t, expectedMessage.ID, actualMessage.ID, "expected 1 message id")
	assert.Equal(t, 200, resp.StatusCode, "expected 200 status code")
	assert.Equal(t, true, actualMessage.IsPalindrome, "expected IsPalindrome true")
	assert.Equal(t, "aaa", actualMessage.Msg, "expected msg aaa")
	messagemap.CleanMap()
}

func TestDeleteMessageHandler(t *testing.T) {
	expectedMsg := "message"
	expectedIsPalindrome := false

	messagemap.CreateMessage(expectedMsg, expectedIsPalindrome)

	r, _ := http.NewRequest("DELETE", "/messages/1", nil)
	w := httptest.NewRecorder()

	vars := map[string]string{
		"messageID": "1",
	}
	r = mux.SetURLVars(r, vars)
	DeleteMessage(w, r)

	_, err := messagemap.GetMessage(1)

	resp := w.Result()

	assert.Equal(t, 204, resp.StatusCode, "expected 204 status code")
	assert.Equal(t, messagemap.ErrorNoSuchKey.Error(), err.Error(), "expected ErrorNoSuchKey error")
	messagemap.CleanMap()
}

func TestGetMessagesHandler(t *testing.T) {
	expectedMsg := "message"
	expectedIsPalindrome := false

	for i := 0; i < 3; i++ {
		messagemap.CreateMessage(expectedMsg, expectedIsPalindrome)
	}
	r, _ := http.NewRequest("GET", "/messages", nil)
	w := httptest.NewRecorder()

	GetMessages(w, r)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	var messages []messagemap.InternalMessage
	err := json.Unmarshal(body, &messages)

	assert.Nil(t, err, "messages read successfully")
	assert.Equal(t, 200, resp.StatusCode, "expected 200 status code")
	assert.Equal(t, 3, len(messages), "get all 3 messages")
	messagemap.CleanMap()
}
