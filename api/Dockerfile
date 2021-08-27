FROM golang:1.16

WORKDIR /go/src/api
COPY . .

RUN go install -v .
RUN ["chmod", "+x", "/go/src/api/start.sh"]
CMD ["./start.sh"]