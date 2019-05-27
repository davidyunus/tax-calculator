FROM golang:1.10-alpine as builder

# update lates version
RUN apk add git \
    && go get -u github.com/golang/dep/cmd/dep

# create folder repository
RUN mkdir -p "$GOPATH/src/github.com/tax-calculator"

# set working directory
WORKDIR /go/src/github.com/tax-calculator

# set volume
COPY . /go/src/github.com/tax-calculator

# install and update dep
RUN dep ensure --update

# build docker
RUN go build

# runner
FROM golang:1.10-alpine as runner

# set workdir
WORKDIR /app

# copy build file to runner
COPY --from=builder /go/src/github.com/tax-calculator/tax-calculator .

CMD [ "./tax-calculator" ]