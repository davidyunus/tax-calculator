FROM golang:1.10-alpine as builder

# update lates version
RUN apk add git \
    && go get -u github.com/golang/dep/cmd/dep

# create folder repository
RUN mkdir -p "$GOPATH/src/github.com/davidyunus/tax-calculator"

# set working directory
WORKDIR /go/src/github.com/davidyunus/tax-calculator

# set volume
COPY . /go/src/github.com/davidyunus/tax-calculator

# install and update dep
RUN dep ensure --update

# build docker
RUN go build $GOPATH/src/github.com/davidyunus/tax-calculator/cmd/httpserver/main.go

# runner
FROM golang:1.10-alpine as runner

# set workdir
WORKDIR /app

# copy build file to runner
COPY --from=builder /go/src/github.com/davidyunus/tax-calculator .

EXPOSE 9090

CMD [ "go", "run", "$GOPATH/src/github.com/davidyunus/tax-calculator/cmd/httpserver/main.go" ]