FROM golang:alpine as builder
RUN apk --no-cache add build-base git bzr mercurial gcc
ADD . /src
RUN cd /src && go build -o tangelo cmd/main.go

FROM alpine
WORKDIR /app
COPY --from=builder /src/tangelo /app/
ENTRYPOINT ./tangelo
