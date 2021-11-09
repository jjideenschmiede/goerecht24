# goerecht24

[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/jjideenschmiede/goerecht24.svg)](https://golang.org/) [![Go](https://github.com/jjideenschmiede/goerecht24/actions/workflows/go.yml/badge.svg)](https://github.com/jjideenschmiede/goerecht24/actions/workflows/go.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/jjideenschmiede/goerecht24)](https://goreportcard.com/report/github.com/jjideenschmiede/goerecht24) [![Go Doc](https://godoc.org/github.com/jjideenschmiede/goerecht24?status.svg)](https://pkg.go.dev/github.com/jjideenschmiede/goerecht24) ![Lines of code](https://img.shields.io/tokei/lines/github/jjideenschmiede/goerecht24) [![Developed with <3](https://img.shields.io/badge/Developed%20with-%3C3-19ABFF)](https://jj-dev.de/)

# Install

```console
go get github.com/jjideenschmiede/goerecht24
```

# How to use?

## Client

Here you will find all the functions for the clients.

### List registered clients

If you want to read all clients of a token, you can do this with the following function. You can find the corresponding documentation [here](https://docs.api.e-recht24.de/#/Client/eRecht24%5CApi%5CControllers%5Cv1%5CClientController%3A%3AlistClients).

```go
// Define request
r := goerecht24.Request{
    ApiKey: "",
}

// List registered clients
clients, err := goerecht24.ListRegisteredClients(r)
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(clients)
}
```

### Register client for push notifications

If you want to create a new client, you can do this with the following function. You can find the corresponding documentation [here](https://docs.api.e-recht24.de/#/Client/post_v1_clients).

```go
// Define request
r := goerecht24.Request{
    ApiKey: "",
}

// Define body
body := goerecht24.ClientBody{
    ClientId:   0,
    PushMethod: "GET",
    PushUri:    "https://domain.com/restapi/push-endpoint",
    Cms:        "WordPress",
    CmsVersion: "5.0.3",
    PluginName: "eRecht24.de Rechtstexte für WordPress",
    AuthorMail: "info@jj-ideenschmiede.de",
}

// Register a new client
registerClient, err := goerecht24.RegisterClient(body, r)
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(registerClient)
}
```

### Update client information

If you want to update a registered client, you can do this with the following function. You can find the corresponding documentation [here](https://docs.api.e-recht24.de/#/Client/eRecht24%5CApi%5CControllers%5Cv1%5CClientController%3A%3Aupdate).

```go
// Define request
r := goerecht24.Request{
    ApiKey: "",
}

// Define body
body := goerecht24.ClientBody{
    ClientId:   17828,
    PushMethod: "GET",
    PushUri:    "https://domain2.com/restapi/push-endpoint",
    Cms:        "WordPress",
    CmsVersion: "5.0.3",
    PluginName: "eRecht24.de Rechtstexte für WordPress",
    AuthorMail: "info@jj-ideenschmiede.de",
}

// Update a registered client
updateRegisteredClient, err := goerecht24.UpdateRegisteredClient(body, r)
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(updateRegisteredClient)
}
```

### Delete client

If you want to delete a registered client, you can do this with the following function. You can find the corresponding documentation [here](https://docs.api.e-recht24.de/#/Client/eRecht24%5CApi%5CControllers%5Cv1%5CClientController%3A%3Adelete).

```go
// Define request
r := goerecht24.Request{
    ApiKey: "",
}

// Delete a registered client
deleteClient, err := DeleteClient(17828, r)
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(deleteClient)
}
```

### Send a push notification for development

If you want to send a test notification, you can do so using the following function.

Select which type of push notification you want to receive for testing.

**Available values : ping, message, imprint, privacyPolicy, privacyPolicySocialMedia**

You can find the corresponding documentation [here](https://docs.api.e-recht24.de/#/Client/post_v1_clients__client_id__testPush).

```go
// Define request
r := goerecht24.Request{
    ApiKey: "",
}

// Send a test push
err := goerecht24.TestPush(17828, "ping", r)
if err != nil {
    fmt.Println(err)
}
```

## Legal documents

Here you will find all the information on how to extract the documents from the api.

### Show imprint

If you want to read the imprint, you can do this with the following function. You can find the corresponding documentation [here](https://docs.api.e-recht24.de/#/Legal%20documents/get_v1_imprint).

```go
// Define request
r := goerecht24.Request{
    ApiKey: "",
}

// Get imprint
imprint, err := goerecht24.ShowImprint(r)
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(imprint)
}
```

### Show privacy policy

If you want to read the privacy policy from eRecht24, you can do this with the following function. You can find the corresponding documentation [here](https://docs.api.e-recht24.de/#/Legal%20documents/get_v1_privacyPolicy).

```go
// Define request
r := goerecht24.Request{
    ApiKey: "",
}

// Get privacy policy
privacyPolicy, err := goerecht24.ShowPrivacyPolicy(r)
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(privacyPolicy)
}
```

### Show privacy policy social media

If you want to read the privacy policy social media, you can do this with the following function. You can find the corresponding documentation [here](https://docs.api.e-recht24.de/#/Legal%20documents/get_v1_privacyPolicySocialMedia).

```go
// Define request
r := goerecht24.Request{
    ApiKey: "",
}

// Get privacy policy social media
privacyPolicySocialMedia, err := goerecht24.ShowPrivacyPolicySocialMedia(r)
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(privacyPolicySocialMedia)
}
```

## Messages

Please only call this endpoint after receiving a push notification indicating to fetch a message.

### Fetch message from 

To retrieve a message, you can use the following function. You can find the corresponding documentation [here](https://docs.api.e-recht24.de/#/Messages/get_v1_message).

```go
// Define request
r := goerecht24.Request{
    ApiKey: "",
}

// Fetch message from eRecht24
message, err := goerecht24.FetchMessage(r)
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(message)
}
```

## Push

This endpoint will send a push notification to all clients, which are registered for this api key.

### Send a push notification to registered clients

If you want to push a type, you can do so via this function. Select, which type of document you want to push.

Available values : **imprint, privacyPolicy, privacyPolicySocialMedia, all** 

You can find the corresponding documentation [here](https://docs.api.e-recht24.de/#/Push/post_v1_push__type_).

```go 
// Define request
r := goerecht24.Request{
    ApiKey: "",
}

// Delete a registered client
err := goerecht24.SendPushNotification("all", r)
if err != nil {
    fmt.Println(err)
}
```

