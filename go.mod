module act-back

go 1.21.0

toolchain go1.23.6

replace act-back/internal/openapi => ./internal/openapi

require (
	act-back/internal/openapi v0.0.0-00010101000000-000000000000 // indirect
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/go-sql-driver/mysql v1.9.2 // indirect
)
