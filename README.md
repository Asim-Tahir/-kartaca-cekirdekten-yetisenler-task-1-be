# go-graphql-sqlite3

An example go web server implementing a GraphQL api.

## Installation

## Default installation

```bash
go get https://github.com/Asim-Tahir/kartaca-cekirdekten-yetisenler-task-1-be/cmd/server
```

## Manual installation

```bash
make
```

## Usage

## From default installation

```bash
server
```

## From manual installation

```bash
./server
2019/xx/xx xx:xx:xx 🚀 GraphQL Playground Server running on http://localhost:4000/
```

Then just open this URL in your favorite browser: [http://localhost:4000/](http://localhost:4000/).

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
