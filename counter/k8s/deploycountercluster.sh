#!/bin/bash
# A script to deploy the Counter service on a local Kubernetes cluster with Minikube 

echo "Starting Minikube with default settings"
minikube start
echo "Interact with Minikube via the `kubectl` command or open the dashboard with `minikube dashboard`"

echo "Building Docker images"
eval $(minikube docker-env)
docker build . -t proto -f proto.Dockerfile
docker build -t ori-app ./counter

echo "Deploying Counter service"
kubectl apply -f counter/k8s/ori.yaml
