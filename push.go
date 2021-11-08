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

import (
	"fmt"
)

// SendPushNotification is to send a push notification to all registered clients
func SendPushNotification(document string, r Request) error {

	// Set config for new request
	c := Config{fmt.Sprintf("/push/%s", document), "POST", nil}

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
