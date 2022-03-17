FROM golang:1.17.7-alpine3.15 as build

WORKDIR /builder

COPY go.* ./
RUN go mod download

COPY . ./

RUN go build -o main cmd/app/main.go


FROM alpine:latest as production

COPY --from=build /builder/main ./
COPY --from=build /builder/docs ./
COPY --from=build /builder/config/config.yml ./config/config.yml

CMD ["./main"]
