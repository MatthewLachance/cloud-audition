package messagemap

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateMessage(t *testing.T) {

	expectedID := 1
	expectedMsg := "message"
	expectedIsPalindrome := false

	message := CreateMessage(expectedMsg, expectedIsPalindrome)

	actualID := message.ID
	actualMsg := message.Msg
	actualIsPalindrome := message.IsPalindrome

	str := fmt.Sprintf("Expected the id to be %d, actually got %d!", expectedID, actualID)
	assert.Equal(t, expectedID, actualID, str)

	str = fmt.Sprintf("Expected the message to be %s, actually got %s!", expectedMsg, actualMsg)
	assert.Equal(t, expectedMsg, actualMsg, str)

	str = fmt.Sprintf("Expected the sPalindrome to be %t, actually got %t!", expectedIsPalindrome, actualIsPalindrome)
	assert.Equal(t, expectedIsPalindrome, actualIsPalindrome, str)

	secondMessage := CreateMessage(expectedMsg, expectedIsPalindrome)

	str = fmt.Sprintf("Expected the id to be %d, actually got %d!", 2, secondMessage.ID)
	assert.Equal(t, 2, secondMessage.ID, str)
	CleanMap()
}

func TestGetMessage(t *testing.T) {

	expectedID := 1
	expectedMsg := "message"
	expectedIsPalindrome := false

	expectedMessage := CreateMessage(expectedMsg, expectedIsPalindrome)
	actualMessage, err := GetMessage(expectedID)

	assert.Nil(t, err, "Expected to get the message without error")

	str := fmt.Sprintf("Expected the id to be %d, actually got %d!", expectedMessage.ID, actualMessage.ID)
	assert.Equal(t, expectedMessage.ID, actualMessage.ID, str)

	str = fmt.Sprintf("Expected the message to be %s, actually got %s!", expectedMessage.Msg, actualMessage.Msg)
	assert.Equal(t, expectedMessage.Msg, actualMessage.Msg, str)

	str = fmt.Sprintf("Expected the isPalindrome to be %t, actually got %t!", expectedMessage.IsPalindrome, actualMessage.IsPalindrome)
	assert.Equal(t, expectedMessage.IsPalindrome, actualMessage.IsPalindrome, str)

	nonexitingID := 2
	_, err = GetMessage(nonexitingID)

	assert.EqualError(t, err, ErrorNoSuchKey.Error(), "Expected the error to be ErrorNoSuchKey")

	CleanMap()
}

func TestUpdateMessage(t *testing.T) {

	expectedID := 1
	expectedMsg := "message"
	expectedIsPalindrome := false

	expectedMessage := CreateMessage(expectedMsg, expectedIsPalindrome)

	updatedMsg := "aaa"
	updatedIsPalindrome := true

	_, err := UpdateMessage(updatedMsg, expectedID, updatedIsPalindrome)

	assert.Nil(t, err, "Expected to update the message without error")

	updatedMessage, _ := GetMessage(expectedID)

	str := fmt.Sprintf("Expected the id to be %d, actually got %d!", expectedMessage.ID, updatedMessage.ID)
	assert.Equal(t, expectedMessage.ID, updatedMessage.ID, str)

	str = fmt.Sprintf("Expected the message to be %s, actually got %s!", updatedMsg, updatedMessage.Msg)
	assert.Equal(t, updatedMsg, updatedMessage.Msg, str)

	str = fmt.Sprintf("Expected the isPalindrome to be %t, actually got %t!", updatedIsPalindrome, updatedMessage.IsPalindrome)
	assert.Equal(t, updatedIsPalindrome, updatedMessage.IsPalindrome, str)

	nonexitingID := 2
	_, err = UpdateMessage(updatedMsg, nonexitingID, updatedIsPalindrome)

	assert.EqualError(t, err, ErrorNoSuchKey.Error(), "Expected the error to be ErrorNoSuchKey")

	CleanMap()
}

func TestGetMessages(t *testing.T) {
	expectedMsg := "message"
	expectedIsPalindrome := false

	numberMsgs := 3

	var expectedMessage *Message
	for i := 0; i < numberMsgs; i++ {
		expectedMessage = CreateMessage(expectedMsg, expectedIsPalindrome)
	}

	str := fmt.Sprintf("Expected the id to be %d, actually got %d!", numberMsgs, expectedMessage.ID)
	assert.Equal(t, numberMsgs, expectedMessage.ID, str)

	messages := GetMessages()

	str = fmt.Sprintf("Expected the number of messages to be %d, actually got %d!", numberMsgs, len(messages))
	assert.Equal(t, numberMsgs, len(messages), str)

	for _, m := range messages {
		str = fmt.Sprintf("Expected the message to be %s, actually got %s!", expectedMsg, m.Msg)
		assert.Equal(t, expectedMsg, m.Msg, str)

		str = fmt.Sprintf("Expected the isPalindrome to be %t, actually got %t!", expectedIsPalindrome, m.IsPalindrome)
		assert.Equal(t, expectedIsPalindrome, m.IsPalindrome, str)
	}
	CleanMap()
}

func TestDeleteMessage(t *testing.T) {

	expectedMsg := "message"
	expectedIsPalindrome := false

	numberMsgs := 3

	var expectedMessage *Message
	for i := 0; i < numberMsgs; i++ {
		expectedMessage = CreateMessage(expectedMsg, expectedIsPalindrome)
	}

	err := DeleteMessage(expectedMessage.ID)

	assert.Nil(t, err, "Expected to delete the message without error")

	numberRest := len(GetMessages())
	str := fmt.Sprintf("Expected the number of rest messages to be %d but instead got %d", numberMsgs-1, numberRest)
	assert.Equal(t, numberMsgs-1, numberRest, str)

	err = DeleteMessage(expectedMessage.ID)

	assert.EqualError(t, err, ErrorNoSuchKey.Error(), "Expected the error to be ErrorNoSuchKey")

	CleanMap()
}
