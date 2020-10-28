chmod 700 private.key
chmod 700 public.key

export PATH=$(go env GOPATH)/bin:$PATH

export DB_HOST=127.0.0.1
export DB_SSL_MODE=disable
export DB_PORT=5433
export DB_DATABASE=b2b
export DB_USERNAME=postgres
export DB_PASSWORD=12345

export PRIVATE_KEY=$(cat ./private.key)
export PUBLIC_KEY=$(cat ./public.key)

swag init
go run main.go