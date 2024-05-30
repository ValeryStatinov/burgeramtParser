FROM golang:1.22 as gobuild

ARG TG_TOKEN
ARG CHAT_ID

ENV TG_TOKEN=${TG_TOKEN}
ENV CHAT_ID=${CHAT_ID}

WORKDIR /

COPY ./go.mod ./go.sum /
RUN go mod download

COPY ./ /
RUN go build ./cmd/burgeramtParser/burgeramtParser.go

CMD ["./burgeramtParser"]
