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
	"errors"
)

// statusCodes
func statusCodes(status string) error {

	// Check each status codes
	switch status {
	case "400 Bad Request":
		return errors.New("push was not successful")
	case "401 Unauthorized":
		return errors.New("unauthorized request. No API key specified, the specified key is invalid or your eRecht24 membership has expired")
	case "424 Failed Dependency":
		return errors.New("push request not successfully executed due to errors in requests to clients")
	default:
		return nil
	}

}
