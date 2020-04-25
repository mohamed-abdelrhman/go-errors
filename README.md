# golang-errors-wrapper
Simple Error wrapper that unify your error reponse for go


# Installation

### Go Modules
`require github.com/mohamed-abdelrhman/go-errors`



## Usage

You just to need to:

`err:=go_errors.NewBadRequestError("Error message Goes Here here")`

//If you are using Gin

`c.JSON(err.Status(),err)`

`/* Output
{
   "message":"Error message Goes Here here",
   "status":400,
   "error":"Bad Request"
}

*/`

