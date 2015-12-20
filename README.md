# go-input
simple http input validation in go

## sample
``` go
func index(res http.ResponseWriter, req *http.Request) {
    query := req.URL.Query()
    v := input.Validator{}
    foo := v.Required(&query, "foo", "foo:missing")
    
    intField := v.Int(&query, "i", "i:not_int")

    if v.Err != nil {
        http.Error(res, req, v.Error(), http.StatusBadRequest)
    }
}
```
