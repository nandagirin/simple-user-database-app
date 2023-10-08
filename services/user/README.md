
# User Service

This simple service handles client request to add users to in-memory database and get the list of those users. This service is built using [GoFiber boilerplate](https://github.com/gofiber/boilerplate).


## Environment Variables

To run this project, you will need to add the following environment variables:

- `JWT_SECRET`: JWT secret key to be used to validate token (required)
## API Reference

Both APIs below needs authentication token generated from auth service. Make sure to include the token in the `Authorization` header.

#### Add user

```http
  POST /users
```

|Body   |Type     |Description                                     |
|:----- |:------- |:---------------------------------------------- |
|`user` |`string` |**Required**. The name of the user to be added. |

#### List users
```http
  GET /users
```




## Run Locally

To run the project locally, firstly clone the project.

```bash
  git clone https://github.com/nandagirin/simple-user-database-app
```

Go to the service directory

```bash
  cd simple-user-database-app/services/user
```

Install dependencies

```bash
  go mod download
```

Start the server

```bash
  go run app.go -port=:8080
```

Please note that the service needs some environment variables to be set. Refer to the previous section to see the list of environment variables.

We could also run the service using container runtime, such as Docker. To do that, we should firstly build the container image. Assuming that we already in the service's directory inside the project, we could run this command.

```bash
  docker build -t user:latest .
```

After the build process has finished, execute `docker run` to spawn the container. Fill the value of the environment variables `ADMIN_PASS` and `JWT_SECRET` before executing the command.

```bash
  docker run --rm -e JWT_SECRET="" -p 8080:8080 user:latest
```

The command above will spawn a container using previously built container image. The container spawned will have port 8080 opened and mapped to also port 8080 in the host, so we could access it using `http://localhost:8080`.





## Running Tests

To run tests, run the following command.

```bash
  go test ./...
```


## License

[MIT](https://choosealicense.com/licenses/mit/)

