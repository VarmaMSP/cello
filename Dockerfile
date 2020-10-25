FROM golang:1.15.3-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN GO111MODULE=on && go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
	&& go build -o ./server ./cmd/cello/main.go

FROM scratch
COPY --from=build /app/server /app/
EXPOSE 8080
ENTRYPOINT ["/app/server"]
