# Specify base image 
FROM golang:1.15.0-alpine

# Install some dependencies 
RUN mkdir /page-visit
ADD  . /page-visit
WORKDIR  /page-visit
RUN go build -o main .

# Default commands 
CMD ["/page-visit/main"]