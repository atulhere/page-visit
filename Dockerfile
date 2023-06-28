# Specify base image 
FROM golang:1.15.0-alpine

# Install some dependencies 
RUN mkdir /maa
ADD  . /maa
WORKDIR  /maa
RUN go build -o main .

# Default commands 
CMD ["/maa/main"]