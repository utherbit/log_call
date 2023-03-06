# log_call
***

### example:
```go
package main

import "github.com/utherbit/log_call"

var logDocument = log_call.LogCall.AddGroup("Document")
var logUsers = log_call.LogCall.AddGroup("Users")

func main() {
    log_call.LogCall.SetConfig(map[string]bool{
        "Document":      true,
        "Users":         false,
    })

    logDocument.Log("OnCallLogDocument") // show log "Document: OnCallLogDocument"
    logUsers.Log("OnCallLogUser")        // not show 
}



```
