package messagemap

import (
	"testing"
)

func TestCreateMessage(t *testing.T) {

	expectedID := 1
	expectedMsg := "message"
	expectedIsPalindrome := false

	message := CreateMessage(expectedMsg, expectedIsPalindrome)

	actualID := message.ID
	actualMsg := message.Msg
	actualIsPalindrome := message.IsPalindrome

	if actualID != expectedID {
		t.Errorf("Expected the id to be %d but instead got %d!", expectedID, actualID)
	}

	if actualMsg != expectedMsg {
		t.Errorf("Expected the message to be %s but instead got %s!", expectedMsg, actualMsg)
	}

	if actualIsPalindrome != expectedIsPalindrome {
		t.Errorf("Expected the sPalindrome to be %t but instead got %t!", expectedIsPalindrome, actualIsPalindrome)
	}

	secondMessage := CreateMessage(expectedMsg, expectedIsPalindrome)

	if secondMessage.ID != 2 {
		t.Errorf("Expected the id to be %d but instead got %d!", 2, secondMessage.ID)
	}
	CleanMap()
}

func TestGetMessage(t *testing.T) {

	expectedID := 1
	expectedMsg := "message"
	expectedIsPalindrome := false

	expectedMessage := CreateMessage(expectedMsg, expectedIsPalindrome)
	actualMessage, err := GetMessage(expectedID)

	if err != nil {
		t.Errorf("Expected to get the message but instead got error: %s", err.Error())
	}

	if actualMessage.ID != expectedMessage.ID {
		t.Errorf("Expected the id to be %d but instead got %d!", expectedMessage.ID, actualMessage.ID)
	}

	if actualMessage.Msg != expectedMessage.Msg {
		t.Errorf("Expected the message to be %s but instead got %s!", expectedMessage.Msg, actualMessage.Msg)
	}

	if actualMessage.IsPalindrome != expectedMessage.IsPalindrome {
		t.Errorf("Expected the isPalindrome to be %t but instead got %t!", expectedMessage.IsPalindrome, actualMessage.IsPalindrome)
	}

	nonexitingID := 2
	_, err = GetMessage(nonexitingID)

	if err != ErrorNoSuchKey {
		t.Error("Expected the error to be ErrorNoSuchKey but instead got nil!")
	}
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

	if err != nil {
		t.Errorf("Expected to update the message but instead got error: %s", err.Error())
	}

	updatedMessage, _ := GetMessage(expectedID)

	if updatedMessage.ID != expectedMessage.ID {
		t.Errorf("Expected the id to be %d but instead got %d!", expectedMessage.ID, updatedMessage.ID)
	}

	if updatedMessage.Msg != updatedMsg {
		t.Errorf("Expected the message to be %s but instead got %s!", expectedMessage.Msg, updatedMessage.Msg)
	}

	if updatedMessage.IsPalindrome != updatedIsPalindrome {
		t.Errorf("Expected the isPalindrome to be %t but instead got %t!", expectedMessage.IsPalindrome, updatedMessage.IsPalindrome)
	}

	nonexitingID := 2
	_, err = UpdateMessage(updatedMsg, nonexitingID, updatedIsPalindrome)

	if err != ErrorNoSuchKey {
		t.Error("Expected the error to be ErrorNoSuchKey but instead got nil!")
	}

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

	if expectedMessage.ID != 3 {
		t.Errorf("Expected the id to be %d but instead got %d!", numberMsgs, expectedMessage.ID)
	}

	messages := GetMessages()

	if len(messages) != numberMsgs {
		t.Errorf("Expected the number of messages to be %d but instead got %d!", numberMsgs, len(messages))
	}

	for _, m := range messages {
		if m.Msg != expectedMsg {
			t.Errorf("Expected the message to be %s but instead got %s!", expectedMsg, m.Msg)
		}

		if m.IsPalindrome != expectedIsPalindrome {
			t.Errorf("Expected the isPalindrome to be %t but instead got %t!", expectedIsPalindrome, m.IsPalindrome)
		}
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

	if err != nil {
		t.Errorf("Expected to delete the message but instead got error: %s", err.Error())
	}

	numberRest := len(GetMessages())
	if numberRest != numberMsgs-1 {
		t.Errorf("Expected the number of rest messages to be %d but instead got %d", numberMsgs-1, numberRest)
	}

	err = DeleteMessage(expectedMessage.ID)

	if err != ErrorNoSuchKey {
		t.Error("Expected the error to be ErrorNoSuchKey but instead got nil!")
	}

	CleanMap()
}
