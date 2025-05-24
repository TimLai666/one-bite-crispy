FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy the go mod and sum files
COPY go.mod go.sum ./
# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download
# Copy the source code
COPY . .
# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest
WORKDIR /
# Copy the Precompiled binary file from the previous stage
COPY --from=builder /app/main .
# Copy the config file
COPY --from=builder /app/.env .
COPY --from=builder /app/prompt/base_prompt.md ./prompt/base_prompt.md

CMD [ "./main" ]