FROM golang:1.21-alpine3.18 AS build
WORKDIR /app
COPY . .
RUN go mod download && go mod verify
RUN go build -o rinha ./cmd/rinha.go

FROM scratch AS final
COPY --from=build /app/rinha . 
ENTRYPOINT [ "./rinha" ]