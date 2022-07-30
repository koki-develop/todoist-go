[![go.mod](https://img.shields.io/github/go-mod/go-version/koki-develop/go-todoist)](https://github.com/koki-develop/go-todoist/blob/main/go.mod)
[![release](https://img.shields.io/github/v/release/koki-develop/go-todoist)](https://github.com/koki-develop/go-todoist/releases/latest)
[![GitHub Actions](https://github.com/koki-develop/go-todoist/actions/workflows/main.yml/badge.svg)](https://github.com/koki-develop/go-todoist/actions/workflows/main.yml)
[![Maintainability](https://api.codeclimate.com/v1/badges/0ee29ff12f800bcb7001/maintainability)](https://codeclimate.com/github/koki-develop/go-todoist/maintainability)
[![codecov](https://codecov.io/gh/koki-develop/go-todoist/branch/main/graph/badge.svg)](https://codecov.io/gh/koki-develop/go-todoist)
[![LICENSE](https://img.shields.io/github/license/koki-develop/go-todoist)](./LICENSE)

# go-todoist ( :warning: In development. :warning: )

This is a Go client library for accessing the [Todoist APIs](https://developer.todoist.com/guides/#our-apis).

## Contents

- [Installation](#installation)
- [REST API Client](#rest-api-client)
  - [Import](#import)
  - [Example](#example)
    - [Get all projects](#get-all-projects)
    - [Create a new task](#create-a-new-task)
  - [Documentation](#documentation)
- [Sync API Client](#sync-api-client)
- [LICENSE](#license)

## Installation

```sh
go get github.com/koki-develop/go-todoist
```

## REST API Client

`go-todoist/todoist` is a package for accessing the [Todoist REST API](https://developer.todoist.com/rest/v1).

### Import

```go
import "github.com/koki-develop/go-todoist/todoist"
```

### Example

#### Get all projects

```go
package main

import (
	"fmt"

	"github.com/koki-develop/go-todoist/todoist"
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

	"github.com/koki-develop/go-todoist/todoist"
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

	"github.com/koki-develop/go-todoist/todoist"
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

For more information, see [go-todoist/todoist](https://pkg.go.dev/github.com/koki-develop/go-todoist/todoist).

## Sync API Client

<!-- TODO: add -->
wip

## LICENSE

[MIT](./LICENSE)
