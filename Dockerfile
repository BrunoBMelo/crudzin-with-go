FROM golang:latest as build
WORKDIR /app
COPY . .
RUN GCO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -a -installsuffix cgo -o main

FROM scratch as final
WORKDIR /app
COPY --from=build /app/main .

EXPOSE 3001
CMD [ "./main" ]