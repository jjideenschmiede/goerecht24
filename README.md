# goerech24

[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/jjideenschmiede/goerecht24.svg)](https://golang.org/) [![Go](https://github.com/jjideenschmiede/goerech24/actions/workflows/go.yml/badge.svg)](https://github.com/jjideenschmiede/goerech24/actions/workflows/go.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/jjideenschmiede/goerecht24)](https://goreportcard.com/report/github.com/jjideenschmiede/goerecht24) [![Go Doc](https://godoc.org/github.com/jjideenschmiede/goerecht24?status.svg)](https://pkg.go.dev/github.com/jjideenschmiede/goerecht24) ![Lines of code](https://img.shields.io/tokei/lines/github/jjideenschmiede/goerecht24) [![Developed with <3](https://img.shields.io/badge/Developed%20with-%3C3-19ABFF)](https://jj-dev.de/)

# Install

```console
go get github.com/jjideenschmiede/goerecht24
```

# How to use?

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