# log_call
***

### example:
```go
package main

import "github.com/utherbit/log_call"

var logDocument = log_call.LogCall.AddGroup("Document")
var logUsers = log_call.LogCall.AddGroup("Users")

func main() {
    log_call.LogCall.SetConfig(
		log_call.Config{
			Groups: {
                "Document":      true,
                "Users":         false,
            },
            TimeFormat: "_2 Jan 15:04:05"
        })

    logDocument.Log("OnCallLogDocument") // show log "2 Jan 15:04:05 Document: OnCallLogDocument"
    logUsers.Log("OnCallLogUser")        // not show 
}



```
