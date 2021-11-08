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

// FetchMessageReturn is to decode the json data
type FetchMessageReturn struct {
	Message       string      `json:"message"`
	MessageDe     string      `json:"message_de"`
	Call2Action   string      `json:"call2action"`
	Call2ActionDe string      `json:"call2action_de"`
	Link          string      `json:"link"`
	Token         interface{} `json:"token,omitempty"`
}

// FetchMessage is to fetch a message from eRecht24
func FetchMessage(r Request) (FetchMessageReturn, error) {

	// Set config for new request
	c := Config{"/message", "GET", nil}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return FetchMessageReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode FetchMessageReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return FetchMessageReturn{}, err
	}

	// Return data
	return decode, err

}
