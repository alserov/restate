FROM golang:alpine as build

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN CGO_ENABLED=0 GO_ARCH=amd64 GOOS=linux go build -o ./bin

FROM scratch

WORKDIR /app

COPY --from=build ./build/bin ./bin

CMD ["./bin"]

