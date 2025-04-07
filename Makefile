openapi:
	openapi-generator-cli generate -i openapi.yaml -g go -o internal/openapi --package-name=openapi --skip-validate-spec --global-property=models
