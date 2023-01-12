# go-free-email

![Last version](https://img.shields.io/github/v/release/ilkerkorkut/go-free-email)
[![License](https://img.shields.io/github/license/ilkerkorkut/go-free-email.svg)](LICENSE)

> A comprehensive list of all free email domain providers. Based on [HubSpot blocked domains](https://knowledge.hubspot.com/articles/kcs_article/forms/what-domains-are-blocked-when-using-the-forms-email-domains-to-block-feature).

## Install

```bash
$ go get github.com/ilkerkorkut/go-free-email
```

## Usage

```go
gfe := gofreeemail.NewGoFreeEmails()
gfe.IsFreeEmail("gmail.com") // returns true or false
gfe.AddFreeEmail("example.com") // adds a new free email domain on runtime
gfe.RemoveDomainFromList("example.com") // removes a domain from the list
```