# KyberSwap Error Lib

## Overview
This is the library that defines errors and error handlers for KyberSwap

## Install
```
$ export GOPRIVATE=github.com/KyberNetwork/kyberswap-error
$ go get -u github.com/KyberNetwork/kyberswap-error
```

## How to use

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

### Register new DomainError and new RestAPIError
- There are functions can be used to create customized errors. With `DomainError`, it's `errors.NewDomainError(code int, message string, entities []string, rootCause error)` and it's `errors.NewRestAPIError(httpStatus int, code int, message string, entities []string, rootCause error)` with `RestAPIError`.
- You can use these above functions to create your customized errors directly. But you should define new constructors for your custom errors, so it can be reused. For example:
```
package main

import (
	"github.com/KyberNetwork/kyberswap-error/pkg/errors"
	"net/http"
)

const (
	DomainErrCodeCustomized = 40099
	DomainErrMsgCustomized  = "Customized domain error"

	ApiErrCodeCustomized = 40099
	ApiErrMsgCustomized  = "Customized api error"
) 

func NewDomainErrCustomized(rootCause error, entities ...string) *errors.DomainError {
	return errors.NewDomainError(DomainErrCodeCustomized, DomainErrMsgCustomized, entities, rootCause)
}

func NewRestAPIErrCustomized(rootCause error, entities ...string) *errors.RestAPIError {
	message := errors.AppendEntitiesToErrMsg(ApiErrMsgCustomized, entities)
	return errors.NewRestAPIError(http.StatusBadRequest, ApiErrCodeCustomized, message, entities, rootCause)
}
```
- After defining your custom errors, you have to register the function used to transform your custom `DomainError` to your custom `RestAPIError`. For example:
```
package main

import (
	"github.com/KyberNetwork/kyberswap-error/pkg/errors"
	"github.com/KyberNetwork/kyberswap-error/pkg/transformers"
	"net/http"
)

const (
	DomainErrCodeCustomized = 40099
	DomainErrMsgCustomized  = "Customized domain error"

	ApiErrCodeCustomized = 40099
	ApiErrMsgCustomized  = "Customized api error"
)

func NewDomainErrCustomized(rootCause error, entities ...string) *errors.DomainError {
	return errors.NewDomainError(DomainErrCodeCustomized, DomainErrMsgCustomized, entities, rootCause)
}

func NewRestAPIErrCustomized(rootCause error, entities ...string) *errors.RestAPIError {
	message := errors.AppendEntitiesToErrMsg(ApiErrMsgCustomized, entities)
	return errors.NewRestAPIError(http.StatusBadRequest, ApiErrCodeCustomized, message, entities, rootCause)
}

func main() {
	transformer := transformers.RestTransformerInstance()
	transformer.RegisterTransformFunc(DomainErrCodeCustomized, NewRestAPIErrCustomized)
}
```
- NOTE: 
  - Each error should have a unique error code. Otherwise, it can lead to unexpected results when transforming `DomainError` to `RESTAPIError`. So You should not define your custom error code as one of the predefined error codes in "kyberswap-error". The list of those predefined error codes can be found at https://www.notion.so/kybernetwork/API-Standards-proposal-draft-e8d8bf2dc5f647e89d2bf1b5f0ef8bdf
  - The error code should contain information about HTTP Status. It makes the error code more meaningful and makes it easier to recognize which `RestAPIError` a `DomainError` should be transformed to
  - The domain error code is not required to be the same as the API error code. But they should be the same 

### For gin framework
- This lib provides the function `ValidationErrToRestAPIErr(err error)` which can be used when binding and validating the request. This function just handles these tags: `required`, `oneof`, `min`, `max`. Feel free to contribute more by making PRs.
