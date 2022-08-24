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

## Contract API

The contract API refers to the way in which an `api-contract.json` file can be composed. To integrate, one needs to expose all of their endpoints to a json file, carefully exposing the intended types for each param expected in the response payload, like so:

#### Basic example

```json
{
  "GET:/api/v1/users/:id": {
    "payload_shape": {
      "user": {
        "id": "number",
        "email": "string",
        "created_at": "datetime"
      }
    }
  }
}
```

#### Nesting Example

The api supports infinite nested structures, allowing you to express payloads like:

```json
{
  "GET:/api/v1/users/:id": {
    "payload_shape": {
      "user": {
        "preferences": {
          "inapp": {
            "marketing": {
              "send_emails": "bool"
            }
          }
        }
      }
    }
  }
}
```

#### Datatypes

All datatypes supported natively by JSON are supported here, with the addition of the `date` and `datetime` datatypes, which have an expressive api for determining which format to use to validate/generate them.

##### Primary datatypes

* `string`
* `number`
* `bool`
* `date`
* `datetime`

These datatypes can be seen below in their most basic format:

```json
{
  "GET:/api/v1/users/:id": {
    "payload_shape": {
      "user": {
        "id": "number",
        "email": "string",
        "likes_cats": "bool",
        "created_at": "datetime",
        "updated_on": "date"
      }
    }
  }
}
```

All Primary types can additionally represent arrays, like so:

#### Using array values
```json
{
  "GET:/api/v1/users/:id": {
    "payload_shape": {
      "user": {
        "id": "number[]",
        "email": "string[]",
        "likes_cats": "bool[]",
        "created_at": "datetime[]",
        "updated_on": "date[]"
      }
    }
  }
}
```

They can additionally take on extended parameters (called `decorators` in the codebase), which allow them to narrow down in format, like so:

```json
{
  "GET:/api/v1/users/:id": {
    "payload_shape": {
      "amount": "number:float"
    }
  }
}
```

The full list of decorators supported (varying by type) can be found below:

#### Decorator API

* `string:uuid`: formats for uuid
* `string:email`: loosely formats for email. The test server will generate random email addresses using `faker` if this is specified.
* `string:name`: loosely formats for name. The test server will generate random names using `faker` if this is specified.
* `string:name`: Similar to name, but will generate a full name instead of a first name.
* `number:int`: creates an integer
* `number:float`: creates a floating point number (with decimal precision of 2)
* `number:bigint`: creates a large int, i.e. `580405389235143`
* `date:yyyymmdd`: generates dates with the format `YYYY-MM-DD`
* `date:yymmdd`: generates dates with the format `YY-MM-DD`
* `date:mmddyyyy`: generates dates with the format `MM-DD-YYYY`
* `date:mmddyy`: generates dates with the format `MM-DD-YY`
* `datetime:ansic`: generates **ansic**-formatted datetimes (i.e. `Mon Jan 22 15:04:05 2006`)
* `datetime:iso861`: generates **ISO861**-formatted datetimes (i.e. `2022-04-07T00:00:00.000-07:00`)
* `datetime:kitchen`: generates a time similar to a kitchen clock (i.e. `7:04PM`)
* `datetime:rfc1123`: generates **RFC1123**-formatted datetimes (i.e. `Sat, 20 Aug 2022 07:22:19 PDT`)
* `datetime:rfc1123z`: generates **RFC1123Z**-formatted datetimes (i.e. `Sat, 20 Aug 2022 07:24:17 -0700`)
* `datetime:rfcrfc3339`: generates **RFC3339**-formatted datetimes (i.e. `2022-08-20T07:27:56-07:00`)
* `datetime:rfcrfc3339_nano`: generates **RFC3339Nano**-formatted datetimes (i.e. `2022-08-20T07:33:33.671227-07:00`)
* `datetime:rfc822`: generates **RFC822**-formatted datetimes (i.e. `20 Aug 22 07:16 PDT`)
* `datetime:rfc822z`: generates **RFC822Z**-formatted datetimes (i.e. `20 Aug 22 07:17 -0700`)
* `datetime:rfc850`: generates **RFC850**-formatted datetimes (i.e. `Saturday, 20-Aug-22 07:20:10 PDT`)
* `datetime:ruby_date`: generates **RubyDate**-formatted datetimes (i.e. `Sat Aug 20 07:12:29 -0700 2022`)
* `datetime:stamp`: generates **Stamp**-formatted datetimes (i.e. `Aug 20 07:40:43`)
* `datetime:stamp_micro`: generates **StampMicro**-formatted datetimes (i.e. `Aug 20 07:45:26.087422`)
* `datetime:stamp_milli`: generates **StampMilli**-formatted datetimes (i.e. `Aug 20 07:43:36.680`)
* `datetime:stamp_nano`: generates **StampNano**-formatted datetimes (i.e. `Aug 20 07:47:27.520037000`)
* `datetime:unix`: generates **Unix**-formatted datetimes (i.e. `Sat Aug 20 07:06:22 PDT 2022`)
