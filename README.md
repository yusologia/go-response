## LOGIARES

To install this library, you just need to run the following command:
```shell
go get github.com/yusologia/go-response
```

This is the standard response that you can use to return data when you are using Yusologia Backend Service.

Below are some examples of using this library:
```go
package example

import (
	"github.com/yusologia/go-response"
	error2 "github.com/yusologia/go-response/error"
	"net/http"
)

type Handler struct{}

func (handler Handler) Testing(w http.ResponseWriter, r *http.Request) {
	// Set data object
	dataObject := map[string]interface{}{"id": 1, "name": "Logiares"}

	// Set data object
	dataArray := make([]interface{}, 0)
	for i := 0; i < 2; i++ {
		dataArray = append(dataArray, dataObject)
	}

	// Set pagination
	// For examples of its use, see 
	// https://github.com/yusologia/go-backend-service/blob/master/app/Repository/Pagination.go
	paginate := logiares.Pagination{
		Count:       6,
		CurrentPage: 1,
		PerPage:     2,
		TotalPage:   3,
	}

	// Success response
	res1 := logiares.Response{}
	res1.Success(w)

	// Success response with data (object)
	res2 := logiares.Response{Object: dataObject}
	res2.Success(w)

	// Success response with data (array)
	res3 := logiares.Response{Array: dataArray}
	res3.Success(w)

	// Success response with data (array & pagination)
	pagination := paginate.ParsePagination()
	res4 := logiares.Response{Array: dataArray, Pagination: &pagination}
	res4.Success(w)

	// Error response
	ErrLogiaCustomerGet()

	// Error response with attributes or error details (validation)
	errValidation()
}

func ErrLogiaCustomerGet() {
	logiares.Error(http.StatusNotFound, "Customer not found", "Please search by another name!!", nil)
}

func errValidation() {
	// This is just an example of a validation response
	// For validation, you have prepared a code that can be seen and used in the github.com/yusologia/go-backend-service
	attributes := make([]interface{}, 0)
	for i := 0; i < 2; i++ {
		attributes = append(attributes, map[string]interface{}{
			"param":   "Name",
			"message": "Field validation for 'Name' failed on the 'required' tag",
		})
	}

	// This is a function we have prepared that can be used
	error2.ErrLogiaValidation(attributes)
}
```

###

### Example responses that will be generated:
#### Success response:
```json
{
  "status": {
    "code": 200,
    "message": "Success",
    "internalMsg": "",
    "attributes": null
  },
  "result": null
}
```

#### Success response with data (object):
```json
{
  "status": {
    "code": 200,
    "message": "Success",
    "internalMsg": "",
    "attributes": null
  },
  "result": {
    "id": 1,
    "name": "Logiares"
  }
}
```

#### Success response with data (array):
```json
{
  "status": {
    "code": 200,
    "message": "Success",
    "internalMsg": "",
    "attributes": null
  },
  "result": [
    {
      "id": 1,
      "name": "Logiares First"
    },
    {
      "id": 2,
      "name": "Logiares Second"
    }
  ]
}
```

#### Success response with data (array & pagination):
```json
{
  "status": {
    "code": 200,
    "message": "Success",
    "internalMsg": "",
    "attributes": null
  },
  "result": [
    {
      "id": 1,
      "name": "Logiares First"
    },
    {
      "id": 2,
      "name": "Logiares Second"
    }
  ],
  "pagination": {
    "count": 6,
    "currentPage": 1,
    "perPage": 2,
    "totalPage": 3
  }
}
```

#### Error response:
```json
{
  "status": {
    "code": 500,
    "message": "An error occurred!",
    "internalMsg": "Example error response..",
    "attributes": null
  }
}
```

#### Error response with attributes or error details (validation):
```json
{
  "status": {
    "code": 400,
    "message": "Missing Required Parameter!",
    "internalMsg": "",
    "attributes": [
      {
        "param": "Name",
        "message": "Field validation for 'Name' failed on the 'required' tag"
      }
    ]
  }
}
```
