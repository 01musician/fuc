# syntax=docker/dockerfile:1
FROM alpine:latest

# install app
RUN apk add --no-cache redis

EXPOSE 6379


# final configuration
CMD ["redis-server"]
