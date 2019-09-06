## gqlgen with echo sample

sample project of [gqlgen](https://github.com/99designs/gqlgen) with [echo](https://github.com/labstack/echo)

### Installation
```sh
$ git clone https://github.com/ebkn/gqlgen-with-echo-sample.git
```

### Usage
```sh
$ docker-compose up
$ docker exec -it gqlgen-with-echo-sammple_app sh
(inside docker container) # go run .
```

##### get token
```sh
$ curl http://localhost:3000/login -X POST -d username=username -d password=password
> xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
```

##### try GraphQL
open `http://localhost:3000/playground`

set header (at bottom menu`HTTP HEADERS`)

```sh
{
  "Authorization": "Bearer <input your token>"
}
```

write query

```graphql
query {
  user {
    username
  }
}
```
