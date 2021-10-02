FROM golang:1.16

RUN mkdir /hw3/
WORKDIR /hw3/

COPY src .
RUN go mod download


ENTRYPOINT ["go", "run", "./..."]
