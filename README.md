# Go Web App

A lightweight, multi-page web app built with pure Go — no frameworks, no external dependencies. Built to be containerised and deployed.

![App Preview](screenshot.png)

## Pages

- `/` — Home
- `/about` — About
- `/contact` — Contact
- `/health` — Health check endpoint (returns `{"status":"ok"}`)

## Run locally

```bash
go run .
```

## Run with Docker

```bash
docker build -t go-web-app .
docker run -p 8080:8080 go-web-app
```

## Project Structure

```
go-web-app/
├── main.go
├── go.mod
├── Dockerfile
├── templates/
│   ├── base.html
│   ├── home.html
│   ├── about.html
│   └── contact.html
└── static/
    └── style.css
```
