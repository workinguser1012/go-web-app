# Go Web App

A lightweight, multi-page web app built with pure Go — Containerised and deployed to Amazon EKS.

<p align="center">
  <img src="assets/Screenshot%202026-05-12%20200723.png" width="700">
</p>
## Pages

- `/` — Home
- `/about` — About
- `/contact` — Contact
- `/health` — Health check endpoint (returns `{"status":"ok"}`)

## Prerequisites

Before deploying, make sure you have the following installed and configured:

- [Go 1.21+](https://golang.org/dl/)
- [Docker](https://docs.docker.com/get-docker/)
- [kubectl](https://kubernetes.io/docs/tasks/tools/)
- [AWS CLI](https://aws.amazon.com/cli/) — configured with `aws configure`
- [eksctl](https://eksctl.io/) — used to create and manage the EKS cluster

## Run locally

```bash
go run .
```

## Run with Docker

```bash
docker build -t go-web-app .
docker run -p 8080:8080 go-web-app
```
## Devopsifying the Project

After Running with Docker , The Image has to be pushed to a registory

```
docker push workinguser1210/go-web-app:v1

```
## Deployment — Amazon EKS

This app is deployed to a Kubernetes cluster running on Amazon EKS.

1. Create the EKS cluster
2. Push the Docker image to Amazon ECR
3. Update the `image:` field in `deployment.yaml` to point to your ECR image
4. Apply the Kubernetes manifests

```bash
kubectl apply -f k8s/deployment.yaml
kubectl apply -f k8s/service.yaml
kubectl apply -f k8s/ingress.yaml
```

## Project Structure

```
go-web-app/
├── main.go
├── go.mod
├── Dockerfile
├── k8s/
│   ├── deployment.yaml
│   ├── service.yaml
│   └── ingress.yaml
├── templates/
│   ├── base.html
│   ├── home.html
│   ├── about.html
│   └── contact.html
└── static/
    └── style.css
```
