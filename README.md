# Istio JWT Authentication & Authorization at the edge
Companion github repo for the article TBD.
This repo is a playground for JWT authentication using istio's request authentication and JWT authorization using GRPC/HTTP external authorizers

## Structure
- [auth-service](./auth-service)
  Go fiber http microservice that manages JWKs/JWTs for user auth/authz
- [authz-service](./authz-service)
  GRPC/HTTP service that handles authorization based a the roles claim of a JWT
- [terraform](./terraform)
  Terraform codebase for installing/managing istio using terraform's helm & terraform providers
- [manage-cluster.sh](./manage-cluster.sh)
  Convenience script to manage a k3d cluster including container registries
