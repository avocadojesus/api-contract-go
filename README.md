# Api Contract

API Contract was a concept developed to solve integration testing challenges with split frontend/backend teams on separate repos/tech stacks. For more information on how/why it was developed, see the [Api Contract wiki](https://github.com/avocadojesus/api-contract-test-server/wiki). Additionally, if you would like more information on the json api, that can be found [here (subnested in the wiki)](https://github.com/avocadojesus/api-contract-test-server/wiki/JSON-API).

# api-contract-go

implements test helpers to validate structure of JSON apis for backend teams implementing the `Api Contract` approach.

# Installation

```
go get github.com/avocadojesus/api-contract-go
go mod vendor
```

# Usage

Given a very simple endpoint handler like this:

```go
package handlers

import (
	"net/http"
  "github.com/go-chi/render"
)

type MeResponse struct {
  Name string `json:"name"`
  Email string `json:"email"`
}

func Me(w http.ResponseWriter, r *http.Request) {
  resp := &HamsandwichResponse{Name: "Calvin", Email: "calvin@cooli.dge"}
  render.JSON(w, r, resp)
}
```

One might have an `api-contract.json` file reflecting this structure at the root of your repo, like so:

```json
{
  "GET:/api/v1/me": {
    "payload_shape": {
      "name": "string",
      "email": "string"
    }
  }
}
```

Within your spec file, you can import the `api-contract-go` package, and assert that the payload shape from your endpoint is in compliance with the expected shape described by your `api-contract.json` file.

```go
package handlers_test

import (
  "testing"
  "net/http"
  "net/http/httptest"
	"node-api-contract-go-app/cmd/gateway/handlers"
	"github.com/avocadojesus/api-contract-go"
)

func TestMe(t *testing.T) {
  r:= httptest.NewRequest(http.MethodGet, "/api/v1/me", nil)
  w:= httptest.NewRecorder()

  handlers.Me(w, r)
  api_contract.Comply(w.Body.Bytes(), "GET", "/api/v1/me", t.Error)
}
```

# API

## api_contract.Comply

For the given `httpMethod` and `path`, raises an exception if the payload does not match the expected payload shape expressed in `api-contract.json` file.

### Params:

```go
Comply(
  bytes []byte,
  httpMethod string,
  path string,
  err *testing.T.Error
)
```

### Usage:

```go
func TestMe(t *testing.T) {
  r:= httptest.NewRequest(http.MethodGet, "/api/v1/me", nil)
  w:= httptest.NewRecorder()

  handlers.Me(w, r)
  api_contract.Comply(w.Body.Bytes(), "GET", "/api/v1/me", t.Error)
}
```

## api_contract.Validate

Returns `true` if the payload shape matches the passed bytes, and `false` if it does not match.

### Params:

```go
Validate(
  bytes []byte,
  httpMethod string,
  path string,
) bool
```

### Usage:

```go
func TestMe(t *testing.T) {
  r:= httptest.NewRequest(http.MethodGet, "/api/v1/me", nil)
  w:= httptest.NewRecorder()

  handlers.Me(w, r)
  passedValidation, reason := api_contract.Validate(w.Body.Bytes(), "GET", "/api/v1/me")

  if !passedValidation {
    t.Error(fmt.Sprintf("JSON shape did not match expected. Reason given: %s", reason))
  }
}
```
