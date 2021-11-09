//********************************************************************************************************************//
//
// Copyright (C) 2018 - 2021 J&J Ideenschmiede GmbH <info@jj-ideenschmiede.de>
//
// This file is part of goerech24.
// All code may be used. Feel free and maybe code something better.
//
// Author: Jonas Kwiedor (aka gowizzard)
//
//********************************************************************************************************************//

package goerecht24

import (
	"encoding/json"
	"fmt"
)

// ClientBody is to create or update a client
type ClientBody struct {
	ClientId   int    `json:"client_id"`
	PushMethod string `json:"push_method"`
	PushUri    string `json:"push_uri"`
	Cms        string `json:"cms"`
	CmsVersion string `json:"cms_version"`
	PluginName string `json:"plugin_name"`
	AuthorMail string `json:"author_mail"`
}

// ListRegisteredClientsReturn is to decode the json data
type ListRegisteredClientsReturn struct {
	ClientId   int    `json:"client_id"`
	ProjectId  int    `json:"project_id"`
	PushMethod string `json:"push_method"`
	PushUri    string `json:"push_uri"`
	Cms        string `json:"cms"`
	CmsVersion string `json:"cms_version"`
	PluginName string `json:"plugin_name"`
	AuthorMail string `json:"author_mail"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

// RegisterClientReturn is to decode the json data
type RegisterClientReturn struct {
	Secret   string `json:"secret"`
	ClientId int    `json:"client_id"`
}

// UpdateRegisteredClientReturn is to decode the json data
type UpdateRegisteredClientReturn struct {
	Secret string `json:"secret"`
}

// DeleteClientReturn is to decode the json data
type DeleteClientReturn struct {
	Message   string      `json:"message"`
	MessageDe string      `json:"message_de"`
	Token     interface{} `json:"token,omitempty"`
}

// ListRegisteredClients is to get a list of all clients from eRecht24
func ListRegisteredClients(r Request) ([]ListRegisteredClientsReturn, error) {

	// Set config for new request
	c := Config{"/clients", "GET", nil}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return nil, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode []ListRegisteredClientsReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return nil, err
	}

	// Return data
	return decode, err

}

// RegisterClient is to register client for push notifications
func RegisterClient(body ClientBody, r Request) (RegisterClientReturn, error) {

	// Convert body
	convert, err := json.Marshal(body)
	if err != nil {
		return RegisterClientReturn{}, err
	}

	// Set config for new request
	c := Config{"/clients", "POST", convert}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return RegisterClientReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode RegisterClientReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return RegisterClientReturn{}, err
	}

	// Return data
	return decode, err

}

// UpdateRegisteredClient is to update registered client
func UpdateRegisteredClient(body ClientBody, r Request) (UpdateRegisteredClientReturn, error) {

	// Convert body
	convert, err := json.Marshal(body)
	if err != nil {
		return UpdateRegisteredClientReturn{}, err
	}

	// Set config for new request
	c := Config{fmt.Sprintf("/clients/%d", body.ClientId), "PUT", convert}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return UpdateRegisteredClientReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode UpdateRegisteredClientReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return UpdateRegisteredClientReturn{}, err
	}

	// Return data
	return decode, err

}

// DeleteClient is to delete registered client
func DeleteClient(id int, r Request) (DeleteClientReturn, error) {

	// Set config for new request
	c := Config{fmt.Sprintf("/clients/%d", id), "DELETE", nil}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return DeleteClientReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode DeleteClientReturn
	json.NewDecoder(response.Body).Decode(&decode)

	// Return data
	return decode, err

}

// TestPush is to send a push notification for development
func TestPush(id int, pushType string, r Request) error {

	// Set config for new request
	c := Config{fmt.Sprintf("/clients/%d/testPush?type=%s", id, pushType), "POST", nil}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return err
	}

	// Close request
	defer response.Body.Close()

	// Check response status
	err = statusCodes(response.Status)
	if err != nil {
		return err
	}

	// Return nothing
	return nil

}
