# Terraform for configuring istio authn & authn

## Troubleshooting
### Jwks doesn't have key to match kid or alg from Jwt
Istio automatically refreshed the cache on the JWKs specified in the RequestAuthorization object, if for some reason you created this object before the service, you will need to wait like 10~15 minutes for it to refresh
