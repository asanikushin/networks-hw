FROM golang:1.16

RUN mkdir /hw3/
WORKDIR /hw3/

RUN mkdir src
COPY src/go.mod go.mod
COPY src/go.sum go.sum

RUN go mod download
COPY src .


ENTRYPOINT ["go", "run", "./..."]
