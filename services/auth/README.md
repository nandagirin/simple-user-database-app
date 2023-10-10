
# Auth Service

This simple service handles client request to authenticate and get access token to be used while accessing the other service. This service is built using [ExpressJS boilerplate](https://expressjs.com/en/starter/generator.html).


## Environment Variables

To run this project, you will need to add the following environment variables:

- `ADMIN_PASS`: The password to be used to authenticate (required)

- `JWT_SECRET`: JWT secret key to be used to generate token (required)

- `PORT`: The port which the service runs on (optional, default is 3000)
## API Reference

#### Initiate authentication

```http
  POST /auth
```

|Body       |Type     |Description                                                                       |
|:--------- |:------- |:-------------------------------------------------------------------------------- |
|`username` |`string` |**Required**. Currently only supports value of `admin`                            |
|`password` |`string` |**Required**. Password belonged to the user, currently only supports `admin` user |



## Run Locally

To run the project locally, firstly clone the project.

```bash
  git clone https://github.com/nandagirin/simple-user-database-app
```

Go to the service directory

```bash
  cd simple-user-database-app/services/auth
```

Install dependencies

```bash
  npm install
```

Start the server

```bash
  npm start
```

Please note that the service needs some environment variables to be set. Refer to the previous section to see the list of environment variables.

We could also run the service using container runtime, such as Docker. To do that, we should firstly build the container image. Assuming that we already in the service's directory inside the project, we could run this command.

```bash
  docker build -t auth:latest .
```

After the build process has finished, execute `docker run` to spawn the container. Fill the value of the environment variables `ADMIN_PASS` and `JWT_SECRET` before executing the command.

```bash
  docker run --rm -e ADMIN_PASS="" -e JWT_SECRET="" -p 3000:3000 auth:latest
```

The command above will spawn a container using previously built container image. The container spawned will have port 3000 opened and mapped to also port 3000 in the host, so we could access it using `http://localhost:3000`.





## Running Tests

To run tests, run the following command.

```bash
  npm test
```


## License

[MIT](https://choosealicense.com/licenses/mit/)

