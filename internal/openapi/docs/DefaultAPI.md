# \DefaultAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetActivities**](DefaultAPI.md#GetActivities) | **Get** /activities | Retorna una lista de actividades
[**GetActivitiesGrouped**](DefaultAPI.md#GetActivitiesGrouped) | **Get** /activities/grouped | Retorna las actividades agrupadas por tipo de actividad
[**PopulateActivities**](DefaultAPI.md#PopulateActivities) | **Post** /activities/populate | Rellena la tabla de actividades con fechas faltantes desde el 1 de enero hasta hoy para cada actividad distinta
[**UpdateActivity**](DefaultAPI.md#UpdateActivity) | **Put** /activities | Inserta o actualiza una actividad



## GetActivities

> GetActivitiesResponse GetActivities(ctx).Execute()

Retorna una lista de actividades

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
	resp, r, err := apiClient.DefaultAPI.GetActivities(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetActivities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetActivities`: GetActivitiesResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetActivities`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetActivitiesRequest struct via the builder pattern


### Return type

[**GetActivitiesResponse**](GetActivitiesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetActivitiesGrouped

> GetActivitiesGroupedResponse GetActivitiesGrouped(ctx).Execute()

Retorna las actividades agrupadas por tipo de actividad

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
	resp, r, err := apiClient.DefaultAPI.GetActivitiesGrouped(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetActivitiesGrouped``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetActivitiesGrouped`: GetActivitiesGroupedResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetActivitiesGrouped`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetActivitiesGroupedRequest struct via the builder pattern


### Return type

[**GetActivitiesGroupedResponse**](GetActivitiesGroupedResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PopulateActivities

> PopulateActivities201Response PopulateActivities(ctx).Execute()

Rellena la tabla de actividades con fechas faltantes desde el 1 de enero hasta hoy para cada actividad distinta

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
	resp, r, err := apiClient.DefaultAPI.PopulateActivities(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PopulateActivities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PopulateActivities`: PopulateActivities201Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.PopulateActivities`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPopulateActivitiesRequest struct via the builder pattern


### Return type

[**PopulateActivities201Response**](PopulateActivities201Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateActivity

> UpdateActivityResponse UpdateActivity(ctx).Activity(activity).Execute()

Inserta o actualiza una actividad

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	activity := *openapiclient.NewActivity(time.Now(), "Leer un libro", "Status_example") // Activity | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdateActivity(context.Background()).Activity(activity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateActivity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateActivity`: UpdateActivityResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateActivity`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateActivityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **activity** | [**Activity**](Activity.md) |  | 

### Return type

[**UpdateActivityResponse**](UpdateActivityResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

