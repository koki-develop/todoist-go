[![go.mod](https://img.shields.io/github/go-mod/go-version/koki-develop/todoist-go)](https://github.com/koki-develop/todoist-go/blob/main/go.mod)
[![release](https://img.shields.io/github/v/release/koki-develop/todoist-go)](https://github.com/koki-develop/todoist-go/releases/latest)
[![GitHub Actions](https://github.com/koki-develop/todoist-go/actions/workflows/main.yml/badge.svg)](https://github.com/koki-develop/todoist-go/actions/workflows/main.yml)
[![codecov](https://codecov.io/gh/koki-develop/todoist-go/branch/main/graph/badge.svg)](https://codecov.io/gh/koki-develop/todoist-go)
[![LICENSE](https://img.shields.io/github/license/koki-develop/todoist-go)](./LICENSE)

# todoist-go

This is an unofficial Go client library for accessing the [Todoist REST API](https://developer.todoist.com/rest/v1).

## Contents

- [Installation](#installation)
- [Import](#import)
- [Example](#example)
  - [Get all projects](#get-all-projects)
  - [Create a new task](#create-a-new-task)
  - [Handling Errors](#handling-errors)
- [Documentation](#documentation)
- [LICENSE](#license)

## Installation

```sh
go get github.com/koki-develop/todoist-go
```

## Import

```go
import "github.com/koki-develop/todoist-go/todoist"
```

## Example

### Get all projects

```go
package main

import (
	"fmt"

	"github.com/koki-develop/todoist-go/todoist"
)

func main() {
	client := todoist.New("TODOIST_API_TOKEN")

	projects, err := client.GetProjects()
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	for _, project := range projects {
		fmt.Printf("ID: %d, Name: %s\n", project.ID, project.Name)
		// ID: 1234567890, Name: Inbox
		// ID: 2345678901, Name: Shopping List
		// ...
	}
}
```

### Create a new task

```go
package main

import (
	"fmt"

	"github.com/koki-develop/todoist-go/todoist"
)

func main() {
	cl := todoist.New("TODOIST_API_TOKEN")

	task, err := cl.CreateTask("task content")
	if err != nil {
		fmt.Printf("%s\n", err)
	}

	fmt.Printf("ID: %d, Content: %s\n", task.ID, task.Content)
	// ID: 3456789012, Content: task content
}
```

With optional parameters:

```go
package main

import (
	"fmt"

	"github.com/koki-develop/todoist-go/todoist"
)

func main() {
	cl := todoist.New("TODOIST_API_TOKEN")

	task, err := cl.CreateTaskWithOptions("task content", &todoist.CreateTaskOptions{
		// Helper functions can be used to specify optional parameters.
		ProjectID: todoist.Int(4567890123),
		SectionID: todoist.Int(5678901234),
		DueString: todoist.String("every 3 months"),
	})
	if err != nil {
		fmt.Printf("%s\n", err)
	}

	fmt.Printf("ID: %d, Content: %s\n", task.ID, task.Content)
	// ID: 6789012345, Content: task content
}
```

### Handling Errors

todoist-go returns a `RequestError` with status code and body when an error response is returned from the Todoist REST API.

```go
package main

import (
	"fmt"
	"io"

	"github.com/koki-develop/todoist-go/todoist"
)

func main() {
	cl := todoist.New("TODOIST_API_TOKEN")

	_, err := cl.GetTask(0)
	if reqerr, ok := err.(todoist.RequestError); ok {
		// The status code of error response can be retrieved from the StatusCode property.
		fmt.Printf("%#v\n", reqerr.StatusCode)
		// => 400

		// The body of the error response can be retrieved from the Body property as io.Reader.
		b, _ := io.ReadAll(reqerr.Body)
		fmt.Printf("%#v\n", string(b))
		// => "task_id is invalid"
	}
}
```

## Documentation

For more information, see [todoist-go/todoist](https://pkg.go.dev/github.com/koki-develop/todoist-go/todoist).

## LICENSE

[MIT](./LICENSE)
