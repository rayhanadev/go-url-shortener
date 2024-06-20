# Go URL Shortener

A very basic URL shortener written in Go using Gin, GORM, and SQLite3.

## Installation

To run this repository clone it locally.

```sh
git clone https://github.com/rayhanadev/go-url-shortener
cd go-url-shortener
```

## Usage

To start the server run the following command:

```sh
go run main.go
```

Afterwards, you can use the following endpoints:

```
POST /register
POST /login
POST /logout
POST /shorten
GET /:shortUrl
```