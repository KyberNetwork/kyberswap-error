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
```

## DomainError
These errors should be used in the domain layer of your service

## RestAPIError
These errors should be used in the application interface of your service
