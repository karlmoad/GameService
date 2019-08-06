FROM alpine:3.7
MAINTAINER Karl Moad <https:/github.com/karlmoad>
RUN apk update
RUN apk add --no-cache bash && apk add --no-cache ca-certificates
RUN update-ca-certificates 2>/dev/null || true
ADD build /service
EXPOSE 30200
WORKDIR /service
ENTRYPOINT /service/GameService