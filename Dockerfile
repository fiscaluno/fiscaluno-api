
FROM golang

ADD . /go/src/github.com/fiscaluno/fiscaluno-api/

COPY . /go/src/github.com/fiscaluno/fiscaluno-api/

RUN go get -u github.com/kataras/iris

RUN go install github.com/fiscaluno/fiscaluno-api/

ENTRYPOINT /go/bin/fiscaluno-api

EXPOSE 8080
