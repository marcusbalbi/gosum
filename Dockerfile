FROM golang:alpine3.12 as compiler

WORKDIR /builder
COPY ./src/ ./

RUN go build -o  app .

FROM scratch
COPY --from=compiler /builder/app app
ENTRYPOINT [ "./app" ]