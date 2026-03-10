package main

import "os"

var authServerURL = os.Getenv("AUTH_SERVER_URL")          // e.g. "https://auth.example.com/validate"
var customJWTJWKURL = os.Getenv("CUSTOM_JWT_JWK_URL")     // e.g. "https://xxx.supabase.co/auth/v1/.well-known/jwks.json"
