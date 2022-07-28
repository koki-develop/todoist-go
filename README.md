[![go.mod](https://img.shields.io/github/go-mod/go-version/koki-develop/todoist-go)](https://github.com/koki-develop/todoist-go/blob/main/go.mod)
[![release](https://img.shields.io/github/v/release/koki-develop/todoist-go)](https://github.com/koki-develop/todoist-go/releases/latest)
[![GitHub Actions](https://github.com/koki-develop/todoist-go/actions/workflows/main.yml/badge.svg)](https://github.com/koki-develop/todoist-go/actions/workflows/main.yml)
[![Maintainability](https://api.codeclimate.com/v1/badges/76287788f44794b7f701/maintainability)](https://codeclimate.com/github/koki-develop/todoist-go/maintainability)
[![codecov](https://codecov.io/gh/koki-develop/todoist-go/branch/main/graph/badge.svg)](https://codecov.io/gh/koki-develop/todoist-go)
[![LICENSE](https://img.shields.io/github/license/koki-develop/todoist-go)](./LICENSE)

# todoist-go ( :warning: In development. :warning: )

This is a Go client library for accessing the [Todoist APIs](https://developer.todoist.com/guides/#our-apis).

## Installation

```sh
go get github.com/koki-develop/todoist-go
```

## REST API Client

`todoist-go/todoist` is a package for accessing the [Todoist REST API](https://developer.todoist.com/rest/v1).

### import

```go
import "github.com/koki-develop/todoist-go/todoist"
```

### Example

#### Get all projects

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

#### Create a new task

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

### Documentation

- [todoist-go/todoist](https://pkg.go.dev/github.com/koki-develop/todoist-go/todoist)

## Sync API Client

<!-- TODO: add -->
wip

## LICENSE

[MIT](./LICENSE)
