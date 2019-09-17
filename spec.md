# Technical task

## Implementation

Create a micro service in Go
- Be TDD aware
- Be CI/CD aware
- Understand the 12factor app impacts

Service should:
- Have a basic set of ‘implementation’ (summing, returning basic calculated value from inputs etc.)
- Provide a gRPC interface for more than one method with more than one payload and response type

Provide example Kubernetes deployment manifests for the service
- Consider resource types that are relevant (Deployment, StatefulSet, Service, Ingress, ConfigMap etc.)

Provide a client for this service
- CLI based
- Should be able to call one or more of the service endpoints
- Print out service response(s)

## Additional thinking

- Prove how it aligns to 12factor app best practices
- Prove how it fits and uses the best cloud native understanding
- How would you expand on this service to allow for the use of an eventstore?
- How would this service be accessed and used from an external client from the cluster?