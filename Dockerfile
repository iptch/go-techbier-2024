FROM golang:1.21.5-bookworm AS builder

WORKDIR /app/

COPY ./ ./

RUN go build .

# runtime image
FROM scratch

COPY --from=builder /app/stargazer /app/stargazer

ENTRYPOINT [ "/app/stargazer" ]
