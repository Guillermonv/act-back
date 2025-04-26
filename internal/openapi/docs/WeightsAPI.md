# \WeightsAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetWeights**](WeightsAPI.md#GetWeights) | **Get** /weights | Obtener pesos semanales agrupados por tipo



## GetWeights

> GetWeightsResponse GetWeights(ctx).Execute()

Obtener pesos semanales agrupados por tipo

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WeightsAPI.GetWeights(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WeightsAPI.GetWeights``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWeights`: GetWeightsResponse
	fmt.Fprintf(os.Stdout, "Response from `WeightsAPI.GetWeights`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWeightsRequest struct via the builder pattern


### Return type

[**GetWeightsResponse**](GetWeightsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

