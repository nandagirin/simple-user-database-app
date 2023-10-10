
# Simple User Database App

This is a simple application that serves client request to store and get the list of users. This application consists of two backend services, which are auth service and user service. There is no GUI serves by this application, so every interaction with the client is done via REST API call.


## Demo

The demo application could be accessed in [https://35.201.94.31.sslip.io](https://35.201.94.31.sslip.io). The interaction flow could be explained like below.

#### Getting the token

```bash
curl https://34.149.143.26.sslip.io/auth -X POST -d '{ "username":"admin","password":"0v_R$>nYv>wT+x$2"}' -H 'Content-type: application/json'
```

Record the response and get the token from it.

```bash
{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjoiYWRtaW4iLCJpYXQiOjE2OTY4MDcxMDd9.IhBQQ5K-u0Sz-ctk4GippVWjPNt-QJz7GbXfddQNjPw"}
```

#### Inserting a user

```bash
curl https://34.149.143.26.sslip.io/users -X POST -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjoiYWRtaW4iLCJpYXQiOjE2OTY4MDcxMDd9.IhBQQ5K-u0Sz-ctk4GippVWjPNt-QJz7GbXfddQNjPw' --form 'user="Fiber Optic"'
```

Response that is received.

```bash
{"success":true,"user":{"name":"Fiber Optic"}}
```

#### Getting the list of users

```bash
curl https://34.149.143.26.sslip.io/users -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjoiYWRtaW4iLCJpYXQiOjE2OTY4MDcxMDd9.IhBQQ5K-u0Sz-ctk4GippVWjPNt-QJz7GbXfddQNjPw'
```

Response that is received.

```bash
{"success":true,"users":[{"name":"Fiber Optic"},{"name":"Coaxial"},{"name":"Wireless"}]}
```
## Tech Stack

**Server:** Node, Express, Golang

**CI/CD**: GitHub Actions

**Deployment Orchestration**: Kubernetes, Kustomize, Containerization (e.g Docker)

**Infrastructure Provisioning**: Terraform


## Documentation

Refer to below links to check service documentations.

- [Auth service](services/auth/README.md)

- [User service](services/user/README.md)

Refer to below links to check additional documentations.

- [Architecture documentation](docs/architectures.md)

- [Development workflow and CI/CD](docs/development-workflow-cicd.md)


## Roadmap

- Add token expiry

- Migrate from self-managed in-memory database using struct to external in-memory database, such as Redis and Memcached

- Provide linter pipeline for auth service using ESlint

- Provide security pipeline during pull request to check security issue in the code

- Provide container scanning pipeline during pull request


## License

[MIT](https://choosealicense.com/licenses/mit/)

