# ori

TODO: clean up readme and finish answering "event store" question

This repo contains a microservice for counting words and lines of text. 
The service provides a gRPC interface with two endpoints. 
The gRPC service is defined in `counter/proto/counter.proto`. 
The implementation for the server is in `counter/server/main.go`.

I provide a simple command line client for the service, also written in Go. 
The source code for the CLI is in `cli/main.go`.

I provide a `Dockerfile` for the server and example Kubernetes deployment manifests in `k8s/ori.yaml`. 
See below for usage instructions.

There are unit tests for the server implementation as well as a suite of e2e tests. The e2e tests could be run premerge or postmerge against a Kubernetes cluster. 

```
ori
├── cli
│   ├── Dockerfile
│   └── main.go
├── counter
│   ├── Dockerfile
│   ├── main.go
│   ├── e2e
│   │   ├── e2e_test.go
│   │   └── runner.sh
│   ├── k8s
│   │   └── ori.yaml
│   ├── proto
│   │   ├── counter.pb.go
│   │   └── counter.proto
│   └── server
│       ├── counters.go
│       ├── counters_test.go
│       └── handlers.go
├── LICENSE
├── proto.Dockerfile
├── README.md
└── spec.md
```

## Design considerations: 

### Testing

There are two types of tests in this repo, unit tests and e2e tests.

If the repo were to get much bigger, I would introduce a testing framework such as testify https://github.com/stretchr/testify

As it stands, I just used the go testing package in the tests to reduce the overhead for the reviewer. 

### CI/CD

This application and repository are well suited to an automated CI/CD pipeline, thanks to Docker and the provided tests.

1. Run premerge tests
2. Build Docker images
3. Store images in registry
4. Deploy to dev environment
5. Run automated postmerge tests
6. Deploy same images to staging etc.

### Kubernetes

I decided to create only minimal Kubernetes manifests for this application, including YAML files for the Deployment and Service. Other Kubernetes resources, such as ConfigMaps, Volumes and StatefulSets, would be overkill for such a simple and stateless application.

### TODO

If I had more time or the project had a broader scope...

- Use a CLI framework such as Cobra -- that said, you can get very far with the Go standard library `flags` package
- Use Testify testing framework
- Integrate with a CI/CD system -- this should be as simple as a writing a `gitlab.ci.yaml` or `circle.ci.yaml` file
- Emit structured logs and use Elasticsearch-Fluentd-Kibana logging stack
- Add metrics and tracing with Prometheus and Zipkin respectively
- Use a build system (such as please.build or Bazel) or Go Modules to manage dependencies
- As the project grows, broaden the suite of tests -- e.g. for the CLI and `handlers.go` as they accumulate more logic

## Usage

Check out the repo into `$GOPATH/github.com/calumball/ori/` or run `go get github.com/calumball/ori`.

Run all commands in this section from the root of the repo.

### Three ways to run the server 

#### Build and run the binary (not recommended)

Ensure that the proto compiler and the gRPC Go package are installed on your system.

1. First, generate the proto files with
`protoc --go_out=plugins=grpc:. counter/proto/counter.proto`

2. Then build the app with
`go build -o server github.com/calumball/ori/counter/`

3. Run the server with
`./server`

Now the server should be running on `localhost:8888`.

#### Build and run the Docker image (recommended)

The above process is too much work and could result in divergence in the way the app is built and run locally and remotely. 
I've include a Dockerfile to handle the dependencies and the building of the app. 
Docker can then run the app in a container. 

1. Build the app with
```
docker build . -t proto -f proto.Dockerfile
docker build -t ori-app ./counter
```

2. Run the app with
`docker run -it -p 8888:8888 ori-app`

Now the server should be running on `localhost:8888`.

#### Run on Kubernetes (bonus)

What if you want to run the app alongside other services? 
What if you want to scale it or run it in the cloud? 
Or get any of the benefits of a container orchestration system?

Then you can run it on Kubernetes using the Kubernetes manifests provided in `counter/k8s/ori.yaml`.

I have included instructions for running the service on Minikube. 
Read about Minikube here: 
https://kubernetes.io/docs/setup/learning-environment/minikube/

Run the helper script `counter/k8s/deploycountercluster.sh` to start Minikube and deploy the service. You must have Minikube installed for this script to work.
Now the counter service should be running in your local Kubernetes cluster!

But how can you interact with the service? The service exposes a `NodePort` https://kubernetes.io/docs/concepts/services-networking/service/#nodeport

Get Minikube IP with `export MINIKUBE_IP=$(minikube ip)`

Get the NodePort of the service with `export NODE_PORT=$(kubectl get svc counter-service -o go-template='{{(index .spec.ports 0).nodePort}}')`

Now you can connect to the service on `${MINIKUBE_IP}:${NODE_PORT}`

Or simply run `minikube service counter-service --url`

### Using the provided CLI

Install the CLI app with `go build -o counter-cli github.com/calumball/ori/cli`

Run the app with, for example,
`./counter-cli --addr=${MINIKUBE_IP}:${NODE_PORT} lines "$(< README.md)"`

Usage: `counter-cli --addr=localhost:8888 [command: words or lines] [text]`

Alternatively, to avoid dealing with protobuf and gRPC dependencies, you can build and run the CLI app with Docker

```
docker build . -t proto -f proto.Dockerfile
docker build -t ori-cli cli
docker run ori-cli
```

### Running the tests

Use the helper script `counter/e2e/runner.sh` to run the end-to-end tests.

The unit tests can simply be run with `go test` in the relevant directory.


## Additional considerations raised in the project spec

### 12factor app best practices

This app was created with the 12factor app best practices in mind. 

https://12factor.net/

Many of the best practices are satisfied simply by using Docker and Kubernetes.
Of the twelve factors,  Docker and Kubernetes contribute to:
- explicit declaration of dependencies in the Dockerfile (although in a bigger project, you should use a proper build system or Go Modules to handle deps)
- explicit separation of building and running the container
- disposability, with containers/pods that are quick to start and die (helped by running a native binary vs e.g. on the JVM)
- dev/prod parity by running the same container in all envs -- developing against a local k8s cluster using Minikube can help further
- port binding and config, with ports passed into the service with environment variables
- processes and concurrency through the Kubernetes model (and through the Go gRPC server's handling of concurrent requests)

Some of the factors don't yet apply to this app due to its very narrow scope. For example, the codebase factor essentially suggests using a monorepo. Since this is a standalone service, the idea of a monorepo is not relevant. However, the app does sit in its own directory so could easily be part of a large monorepo.

Likewise, the service doesn't currently have any backing services, so that factor is less relevant. 

### cloud native

This app is fully cloud native in the sense of being easily containerisable and deployable in the cloud via Kubernetes. 
Being a stateless lightweight microservice, the app is also highly scalable.

### eventstore

### external clients

There are many ways of exposing this app to clients outside of the cluster, depending on the use case.
In Minikube, the service exposes a `NodePort` so that the CLI already acts as an external client. 
However, in a real Kubernetes cluster, you would need a LoadBalancer service or similar to provide an external IP for the service.

However, realistically you wouldn't expose this service directly outside of the cluster.
For a start, you could run a JSON+REST gateway in front of the gRPC server to allow external clients to communicate with a more familiar interface and to avoid sharing `.proto` files outside of the repository. 
This way, outside clients can be written in languages without protobuf support and familiar tools such as `curl` and Postman can be used with the service. 
See here for details:
https://github.com/grpc-ecosystem/grpc-gateway 

To serve clients outside of the cluster, it would be best to expose a separate public API service which can then call the Counter srvice.
