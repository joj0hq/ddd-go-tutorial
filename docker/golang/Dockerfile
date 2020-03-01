# change 1.13 if you understand dep
FROM golang:1.10

#RUN apt-get update -qq
#RUN apt-get install -y

# Set Environments
ENV APP=src/training

# Copy the local package files to the container
ADD $APP /go/$APP

## change for Go Module
RUN go get -u github.com/golang/dep/cmd/dep
WORKDIR /go/$APP
RUN dep ensure

# Install hot deploy tool
RUN go get github.com/pilu/fresh

#RUN apt-get update -qq \
#    && apt-get install -y \
##    && go get -u "github.com/go-sql-driver/mysql" \
##    && go get -u "github.com/gin-gonic/gin" \
##    && go get -u "github.com/jinzhu/gorm" \
##    && go get -u "github.com/mattn/go-sqlite3" \
#    && go get -u "github.com/golang/dep/cmd/dep" \
#    && dep ensure

#RUN go build
