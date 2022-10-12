# KyberSwap Error Lib

## Overview
This is the library that defines errors and error handlers for KyberSwap

## Add this lib to your project
- Step 1: 
```
$ export GOPRIVATE=github.com/KyberNetwork/kyberswap-error
```
- Step 2: Add file `tools/tools.go` with content:
```
package tools

import (
	_ "github.com/KyberNetwork/kyberswap-error/tools"
)
```
- Step 3: 
```
$ go mod tidy
$ go mod vendor
```

## Update to latest version
```
$ go get -u github.com/KyberNetwork/kyberswap-error
$ go mod vendor
```

## How to use

### `InfraError`
These errors should be used in the infra layer (repository) of your service

### `DomainError`
These errors should be used in the domain layer of your service

### `RestAPIError`
These errors should be used in the application interface of your service

### Transforms DomainError to RestAPIError
```
package main

import (
	"fmt"
	
	"github.com/KyberNetwork/kyberswap-error/pkg/errors"
	t "github.com/KyberNetwork/kyberswap-error/pkg/transformer"
)

func main() {
	domainErr := errors.NewDomainErrorNotFound(nil)
	transformer := t.RestTransformerInstance()
	apiErr := transformer.DomainErrToRestAPIErr(domainErr)
	fmt.Println(apiErr.Error())
}
```
- NOTE
  - Transforming `InfraError` to `DomainError` is done in a similar way
  - Transform errors flow: `InfraError`(if exists) => `DomainError` => `ClientError (RestAPIError, RPCError)`

### Register new DomainError and new RestAPIError
- There are functions can be used to create customized errors. With `DomainError`, it's `errors.NewDomainError(code string, message string, entities []string, rootCause error)` and it's `errors.NewRestAPIError(httpStatus int, code int, message string, entities []string, rootCause error)` with `RestAPIError`.
- You can use these above functions to create your customized errors directly. But you should define new constructors for your custom errors, so it can be reused. For example:
```
package main

import (
	"net/http"
	
	"github.com/KyberNetwork/kyberswap-error/pkg/errors"
)

const (
	DomainErrCodeCustomized = "DOMAIN:CUSTOMIZED"
	DomainErrMsgCustomized  = "Customized domain error"

	ClientErrCodeCustomized = 40099
	ClientErrMsgCustomized  = "Customized client error"
) 

func NewDomainErrCustomized(rootCause error, entities ...string) *errors.DomainError {
	return errors.NewDomainError(DomainErrCodeCustomized, DomainErrMsgCustomized, entities, rootCause)
}

func NewRestAPIErrCustomized(rootCause error, entities ...string) *errors.RestAPIError {
	message := errors.AppendEntitiesToErrMsg(ClientErrMsgCustomized, entities)
	return errors.NewRestAPIError(http.StatusBadRequest, ClientErrCodeCustomized, message, entities, rootCause)
}
```
- After defining your custom errors, you have to register the function used to transform your custom `DomainError` to your custom `RestAPIError`. For example:
```
package main

import (
	"net/http"
	
	"github.com/KyberNetwork/kyberswap-error/pkg/errors"
)

const (
	DomainErrCodeCustomized = "DOMAIN:CUSTOMIZED"
	DomainErrMsgCustomized  = "Customized domain error"

	ClientErrCodeCustomized = 40099
	ClientErrMsgCustomized  = "Customized client error"
) 

func NewDomainErrCustomized(rootCause error, entities ...string) *errors.DomainError {
	return errors.NewDomainError(DomainErrCodeCustomized, DomainErrMsgCustomized, entities, rootCause)
}

func NewRestAPIErrCustomized(rootCause error, entities ...string) *errors.RestAPIError {
	message := errors.AppendEntitiesToErrMsg(ClientErrMsgCustomized, entities)
	return errors.NewRestAPIError(http.StatusBadRequest, ClientErrCodeCustomized, message, entities, rootCause)
}

func main() {
	transformer := transformers.RestTransformerInstance()
	transformer.RegisterTransformFunc(DomainErrCodeCustomized, NewRestAPIErrCustomized)
}
```
- NOTE: 
  - Each error should have a unique error code. Otherwise, it can lead to unexpected results when transforming. So You should not define your custom error code as one of the predefined error codes in "kyberswap-error". The list of those predefined error codes can be found at https://www.notion.so/kybernetwork/API-Standards-proposal-draft-e8d8bf2dc5f647e89d2bf1b5f0ef8bdf
  - The `ClientErrorCode` should contain information about HTTP Status. It makes the error code more meaningful
  - Registering new `InfraError` is done in a similar way

### For gin framework
- This lib provides the function `ValidationErrToRestAPIErr(err error)` which can be used when binding and validating the request.
