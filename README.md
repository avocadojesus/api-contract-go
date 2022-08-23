# Api Contract

API Contract was a concept developed as an attempt to solve a critical problem facing many in the modern web development world who decide to partition their teams into segregated front and back end teams. As these teams build their apps out, it becomes difficult to consistently agree with their backend counterparts about which api endpoint we are expressing, leading to many headaches and production bugs caused by a lack of ability to do proper integration testing.

The solution proposed through `Api Contract` is to develop a single source of truth json file that both teams can share, as well as test helpers that can provide a bridge where proper integration testing is not possible.

To implement such a solution, there are several repos providing sensible implementations of this for teams, depending on what they are doing. For instance, if you are developing a backend JSON api to be consumed by a separate frontend team, Api Contract provides repos for your stack that allow you to granularly test your payload shapes to make sure they follow the correct formatting.

However, if you are on a frontend team writing in react or vuejs, Api Conract provides a repo which sidechains to your feature spec runs, providing a dummy REST API server at a port of your choosing which will listen at the same endpoints exposed in the `api-contract.json` file at the root of your frontend repo, and will serve dummy payloads in the exact shape expressed by the json file.

This allows the front-end team to spike on an agreed feature as though the backend team had already built the endpoint to spec, and ensures that once the backend team and frontend team do a coordinated deployment (after building their features using our provided test helpers to validate the shape of their own endpoints), the apps should gel together perfectly without any headaches or deployment nightmares.

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
