docker image rm -f b2b:latest
docker build --no-cache -t b2b:latest .
docker image prune -f