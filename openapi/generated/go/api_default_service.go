// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Example API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 */

package openapi

import (
	"context"
	"net/http"
	"errors"
)

// DefaultAPIService is a service that implements the logic for the DefaultAPIServicer
// This service should implement the business logic for every endpoint for the DefaultAPI API.
// Include any external packages or services that will be required by this service.
type DefaultAPIService struct {
}

// NewDefaultAPIService creates a default api service
func NewDefaultAPIService() *DefaultAPIService {
	return &DefaultAPIService{}
}

// GetActivities - Retorna una lista de actividades
func (s *DefaultAPIService) GetActivities(ctx context.Context) (ImplResponse, error) {
	// TODO - update GetActivities with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, GetActivities200Response{}) or use other options such as http.Ok ...
	// return Response(200, GetActivities200Response{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetActivities method not implemented")
}
