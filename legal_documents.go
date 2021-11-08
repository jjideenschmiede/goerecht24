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

// ShowImprintReturn is to decode the json data
type ShowImprintReturn struct {
	HtmlEn    string      `json:"html_en"`
	Created   string      `json:"created"`
	Modified  string      `json:"modified"`
	Warnings  string      `json:"warnings"`
	Pushed    string      `json:"pushed"`
	HtmlDe    string      `json:"html_de"`
	Message   string      `json:"message,omitempty"`
	MessageDe string      `json:"message_de,omitempty"`
	Token     interface{} `json:"token,omitempty"`
}

// ShowPrivacyPolicyReturn is to decode the json data
type ShowPrivacyPolicyReturn struct {
	HtmlEn    string      `json:"html_en"`
	Dsgalt    int         `json:"dsgalt"`
	Created   string      `json:"created"`
	Modified  string      `json:"modified"`
	Warnings  string      `json:"warnings"`
	Pushed    string      `json:"pushed"`
	HtmlDe    string      `json:"html_de"`
	Message   string      `json:"message,omitempty"`
	MessageDe string      `json:"message_de,omitempty"`
	Token     interface{} `json:"token,omitempty"`
}

// PrivacyPolicySocialMediaReturn is to decode the json data
type PrivacyPolicySocialMediaReturn struct {
	HtmlEn    string      `json:"html_en"`
	Warnings  string      `json:"warnings"`
	Created   string      `json:"created"`
	Modified  string      `json:"modified"`
	Pushed    string      `json:"pushed"`
	HtmlDe    string      `json:"html_de"`
	Message   string      `json:"message,omitempty"`
	MessageDe string      `json:"message_de,omitempty"`
	Token     interface{} `json:"token,omitempty"`
}

// ShowImprint is to get the imprint from eRecht24
func ShowImprint(r Request) (ShowImprintReturn, error) {

	// Set config for new request
	c := Config{"/imprint", "GET", nil}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return ShowImprintReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode ShowImprintReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return ShowImprintReturn{}, err
	}

	// Return data
	return decode, err

}

// ShowPrivacyPolicy is to get the privacy policy from eRecht24
func ShowPrivacyPolicy(r Request) (ShowPrivacyPolicyReturn, error) {

	// Set config for new request
	c := Config{"/privacyPolicy", "GET", nil}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return ShowPrivacyPolicyReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode ShowPrivacyPolicyReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return ShowPrivacyPolicyReturn{}, err
	}

	// Return data
	return decode, err

}

// ShowPrivacyPolicySocialMedia is to get the privacy policy social media from eRecht24
func ShowPrivacyPolicySocialMedia(r Request) (PrivacyPolicySocialMediaReturn, error) {

	// Set config for new request
	c := Config{"/privacyPolicySocialMedia", "GET", nil}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return PrivacyPolicySocialMediaReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode PrivacyPolicySocialMediaReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return PrivacyPolicySocialMediaReturn{}, err
	}

	// Return data
	return decode, err

}
