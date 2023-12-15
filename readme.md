# JWT Test

The objective of this test is to write a small library that can generate Json Web Tokens (JWTs) in Go. A short introduction to JWTs can be found at https://jwt.io/introduction and the full specification can be found at https://www.rfc-editor.org/rfc/rfc7519.

The library should be able to generate a JWT with the public `iss`, `sub`, `aud`, `iat` and `name` claims as well as any private claims the user may wish to define.

It only needs to support the HMACSHA256 algorithm for signing although a design which supports the introduction of other algorithms is preferred.

**Please note that only the generation and not the verification of the JWT is required.**

The test should take no more than 2 hours. Please work in a different branch and once complete a Pull Request should be opened against the main branch containing all the work.

Thank you and good luck!