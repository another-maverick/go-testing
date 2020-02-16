package main

import "testing"

//Define an interface  that implements the email send method
type Messager interface {
	Send(email, subject string, body []byte) error
}

func Alert(m Messager, problem []byte) error {
	return m.Send("user@abc.com", "I want to verify this", problem)
}

type MockMessage struct {
	email string
	subject string
	body []byte
}

// This stores the data in the MockMessage struct instead of sending the message
func (m *MockMessage) Send(email, subject string, body []byte)  error {
	m.body = body
	m.email  = email
	m.subject = subject

	return nil
}

func TestAlert(t *testing.T) {
	myMock := new(MockMessage)

	msgBody := []byte("I want to verify this")

	// This will call myMock.Send and it will set the struct vars to ""user@abc.com", "I want to verify this", problem" because we hardcoded  them above
	Alert(myMock, msgBody)

	if myMock.subject != "I want to verify this" {
		t.Errorf("expected ' want to verify this' got %s \n", myMock.subject)
	}


}


