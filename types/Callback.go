package types

import "net/http"

type Callback func(http.ResponseWriter, *http.Request)