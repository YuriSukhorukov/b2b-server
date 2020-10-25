chmod 700 private.key
chmod 700 public.key

export PATH=$(go env GOPATH)/bin:$PATH
export PRIVATE_KEY=$(cat ./private.key)
export PUBLIC_KEY=$(cat ./public.key)

swag init
go run main.go