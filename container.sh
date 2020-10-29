docker container rm -f b2b-demo
docker run \
-e DB_HOST="host.docker.internal" \
-e DB_SSL_MODE=disable \
-e DB_PORT=5433 \
-e DB_DATABASE=b2b \
-e DB_USERNAME=postgres \
-e DB_PASSWORD=12345 \
-e PRIVATE_KEY="$(cat ./private.key)" \
-e PUBLIC_KEY="$(cat ./public.key)" \
-m 100M --cpus="0.5" \
-p 8080:8080 \
--name b2b-demo b2b:latest