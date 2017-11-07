FROM golang:1.9.2-alpine3.6

LABEL Name=simple-contact-list Version=0.0.1 

# Add some packages
RUN apk add --no-cache git bash

# Making directories
RUN mkdir -p /go/src/github.com/MuchChaca/Training/simple-contact-list/bin
RUN mkdir -p /go/src/github.com/MuchChaca/Training/simple-contact-list/src
# Add the source code
ADD src/. /go/src/github.com/MuchChaca/Training/simple-contact-list/src

# Go dependencies
# echo
RUN go get -u github.com/labstack/echo
RUN go get -u github.com/go-sql-driver/mysql
RUN go get -u github.com/dgrijalva/jwt-go

# GOBIN
ENV oldGoBin ${GOBIN}
ENV GOBIN /go/src/github.com/MuchChaca/Training/simple-contact-list/bin

WORKDIR /go/src/github.com/MuchChaca/Training/simple-contact-list/src
RUN go install alpha.go


# Set the GOBIN to the old one
ENV GOBIN ${oldGoBin}}

# Prepare to start the binary file
WORKDIR /go/src/github.com/MuchChaca/Training/simple-contact-list/bin
CMD [ "./alpha" ]

# Exposing port 8080
EXPOSE 8080

RUN echo $DB_HOST

# For more control, you can copy and build manually
# FROM golang:latest 
# LABEL Name=simple-contact-list Version=0.0.1 
# RUN mkdir /app 
# ADD . /app/ 
# WORKDIR /app 
# RUN go build -o main .
# EXPOSE 8080 
# CMD ["/app/main"]
