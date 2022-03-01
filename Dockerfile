FROM golang:1.17-alpine as build
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN cd cmd; CGO_ENABLED=0 go build -ldflags "-s -w" -installsuffix cgo -o /cmd/app

FROM gcr.io/distroless/static
COPY --from=build /cmd/app /
EXPOSE 3001
CMD ["/app"]