# golang-errors-wrapper
Simple Error wrapper that unify your error reponse for go


# Installation

### Go Modules
`require github.com/mohamed-abdelrhman/go-errors`



## Usage

You just to need to:

`err:=go_errors.NewBadRequestError("Error message Goes Here here")`


Format Response:

For Gin Framwork
```
c.JSON(err.Status(),err)
```
Another Framworks

```
w.Header().Set("Content-Type","application/json")
w.WriteHeader(err.Status())
json.NewEncoder(w).Encode(err)
```

OutPut
```

{
   "message":"Error message Goes Here here",
   "status":400,
   "error":"Bad Request"
}




