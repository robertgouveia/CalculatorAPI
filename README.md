
# Calculator API

This is a capable API with middleware, validation and execution.


## Installation

Ensure golang is installed on your system

```bash
  git clone https://github.com/robertgouveia/CalculatorAPI
  cd CalculatorAPI
  docker build -t calc
  docker run -p 8080:8080 calc
```
    
## Endpoints

#### Add
```
localhost:8080/add
```
Request
```json
{
    "Number1": 1,
    "Number2": 2
}
```
Response
```
3
```

#### Subtract
```
localhost:8080/subtract
```
Request
```json
{
    "Number1": 2,
    "Number2": 1
}
```
Response
```
1
```

#### Divide
```
localhost:8080/divide
```
Request
```json
{
    "Number1": 2,
    "Number2": 1
}
```
Response
```
2
```

#### Multiply
```
localhost:8080/multiply
```
Request
```json
{
    "Number1": 2,
    "Number2": 1
}
```
Response
```
2
```

#### Sum
```
localhost:8080/sum
```
Request
```json
{
    "Numbers": [1, 2, 3]
}
```
Response
```
6
```

## Edge Cases

#### Home Directory
```bash
  localhost:8080/
```
Response
```bash
Endpoint not found
```

#### Any GET request

```
localhost:8080/
```
Response
```bash
Method Not Allowed
```

#### Not supplying required fields

```
localhost:8080/add
```
Request
```json
{

}
```

Response
```bash
Empty Payload
```

Request
```json
{
    "Number1": 1
}
```

Response
```bash
Missing field: Number2
```

#### Division By 0


```
localhost:8080/divide
```
Request
```json
{
    "Number1": 0,
    "Number2": 1
}
```

Response
```bash
Cannot divide by zero
```

## Implementations

#### Server Implementation

```go
http.ListenAndServe(":8080", middleware.RegisterMiddleware(handlers.SetupHandlers()))
```

This allows for the handlers to be setup first with the server but wrapped in the global middleware.

#### Endpoints

```go
var Requirements = map[string]structs.EndpointRequirements{
	"/add": {
		Handler:          Add,
		Required:         []string{"Number1", "Number2"},
		CustomMiddleware: []func(http.Handler) http.Handler{},
	},
	"/subtract": {
		Handler:          Subtract,
		Required:         []string{"Number1", "Number2"},
		CustomMiddleware: []func(http.Handler) http.Handler{},
	},
	"/multiply": {
		Handler:          Multiply,
		Required:         []string{"Number1", "Number2"},
		CustomMiddleware: []func(http.Handler) http.Handler{},
	},
	"/divide": {
		Handler:  Divide,
		Required: []string{"Number1", "Number2"},
		CustomMiddleware: []func(http.Handler) http.Handler{
			customMiddleware.DivideByZero,
		},
	},
	"/sum": {
		Handler:          Sum,
		Required:         []string{"Numbers"},
		CustomMiddleware: []func(http.Handler) http.Handler{},
	},
}
```
Endpoints are defined like so, they have a handler, required fields and can hold custom middleware.

#### Custom middleware Implementation

```go
func RegisterMiddleware(mux *http.ServeMux) http.Handler
```

The register middleware returns the server wrapped in middleware then wrapped in the custom middleware setup.

```go
return CustomMiddleware(wrappedHandler)
```

The custom middleware will ask the handler setup for the required custom middleware and assign it like so.

```go
for _, m := range middleware {
	next = m(next)
}
```
