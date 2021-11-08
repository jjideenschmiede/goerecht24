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

package goerech24

import "encoding/json"

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
