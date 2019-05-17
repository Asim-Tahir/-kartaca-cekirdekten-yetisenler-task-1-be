# go-graphql-example

An example go web server implementing a GraphQL api.

## Installation

```bash
make
```

## Usage

```bash
./example
2019/xx/xx xx:xx:xx connect to http://localhost:8080/ for GraphQL playground
```

Then just open this url in your favourite browser: [http://localhost:8080/](http://localhost:8080/).

## Examples

### Create a new User

```graphql
mutation {
  createUser(input: {
    name: "Bryan"
  }) {
    id
    name
  }
}
```

### Create a new Todo

```graphql
mutation {
  createTodo(input: {
    text: "hello world!"
    userId: "1"
  }) {
    id
    text
    done
    user {
      id
      name
    }
  }
}
```

### List Users

```graphql
query {
  users {
    id
    name
  }
}
```

### List Todos

```graphql
query {
  todos {
    id
    text
    done
    user {
      id
      name
    }
  }
}
```
