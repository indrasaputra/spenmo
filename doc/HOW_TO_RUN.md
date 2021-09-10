## How to Run

There are two ways to run the application. The first is to use docker image and docker compose. The second is to setup all dependencies manually. We prefer and encourage using the former approach as it is simpler and the base is already provided.  

### Docker

- Install [Docker Compose](https://docs.docker.com/compose/).

- We assume that the docker can be accessed via localhost. By default, it works this way. If your setting is different, please adjust accordingly.

- Go to this repository directory

    e.g:
    ```
    $ cd go/src/github.com/indrasaputra/spenmo
    ```

- Take a look at [docker-compose.yaml](../docker-compose.yaml). Change the env value as you want, but we encourage you to let it be as it is.

    If you change the values, please, take a second look on PostgreSQL configuration. Make sure that the values are the same as in Spenmo image.

    Adjust the rate limit configuration to check if it works. `RATE_LIMIT_PER_SECOND=1` and `RATE_BURST_PER_SECOND=1` means only 1 request can be served per user per second. It uses [token bucket algorithm](https://en.wikipedia.org/wiki/Token_bucket).

- Run docker compose

    ```
    $ docker-compose up
    ```

- Migrate the database and seed the data

    Make sure you already installed [golang-migrate](https://github.com/golang-migrate/migrate/blob/master/cmd/migrate/README.md).

    ```
    $ make migrate url=<postgres url>
    ```

    e.g:

    ```
    $ make migrate url="postgres://user:password@host:port/dbname"
    ```

    **DON'T FORGET** to supply/change the `user`, `password`, `host`, `port`, and `dbname` according to your database settings.

    e.g:

    ```
    $ make migrate url="postgres://postgresuser:postgrespassword@localhost:5432/spenmo"
    ```

- Make some requests

    Please, adjust the value accordingly.

    - Create a user's card
        ```
        $ curl --request POST 'http://localhost:8081/v1/users/cards' \
            --header 'Authorization: 1' \
            --header 'Content-Type: application/json' \
            --data-raw '{
                "walletId": "oWx0b8DZ1a",
                "limitDaily": 2000000,
                "limitMonthly": 5000000
            }'
        ```

    - Get all user's card
        ```
        $ curl --request GET 'http://localhost:8081/v1/users/cards' --header 'Authorization: 1'
        ```

    - Get a single card
        ```
        $ curl --request GET 'http://localhost:8081/v1/users/cards/:id' --header 'Authorization: 1'
        ```

        e.g:

        ```
        $ curl --request GET 'http://localhost:8081/v1/users/cards/oWx0b8DZ1a' --header 'Authorization: 1'
        ```

    - Update a single card
        ```
        $ curl --request PUT 'http://localhost:8081/v1/users/cards/:id' \
            --header 'Authorization: 1' \
            --header 'Content-Type: application/json' \
            --data-raw '{
                "limitDaily": 2000000,
                "limitMonthly": 3000000
            }'
        ```

    - Delete a single card
        ```
        $ curl --request DELETE 'http://localhost:8081/v1/users/cards/:id' --header 'Authorization: 1'
        ```

### Manual

- Go to this repository directory

    e.g:
    ```
    $ cd go/src/github.com/indrasaputra/spenmo

- Create `.env` file

    You can copy the `env.example` and change its values accordingly

    ```
    $ cp env.example .env
    ```

- Adjust the rate limit configuration to check if it works. `RATE_LIMIT_PER_SECOND=1` and `RATE_BURST_PER_SECOND=1` means only 1 request can be served per user per second. It uses [token bucket algorithm](https://en.wikipedia.org/wiki/Token_bucket).

- Fill all env variables with prefix `POSTGRES_` according to your PostgreSQL settings

- Run or start PostgreSQL

- Fill `PORT_GRPC` and `PORT_REST` value as you wish. We use `8080` as default value for `PORT_GRPC` and `8081` for `PORT_REST`.
    `PORT_GRPC` is a port for HTTP/2 gRPC. `PORT_REST` is port for HTTP/1.1 REST.
    We encourage to let both values as default

- Download the dependencies

    ```
    $ make tidy
    ```

- It is always good to have your database migration up-to-date.
    Run the following command to make your database stays up-to-date with the current migrations.

    ```
    $ make migrate url=<postgres url>
    ```

    e.g:

    ```
    $ make migrate url="postgres://user:password@host:port/dbname"
    ```

    **DON'T FORGET** to supply/change the `user`, `password`, `host`, `port`, and `dbname` according to your database settings.

- Run the application

    ```
    $ go run cmd/api/main.go
    ```
