FROM mcr.microsoft.com/devcontainers/go:0-1-bullseye

WORKDIR /go/src/app

COPY . .

RUN cd app
RUN go get -v ./...
RUN go install -v ./...

CMD ["app"]
