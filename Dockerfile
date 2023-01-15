FROM golang:alpine
COPY blog-server /blog-server
ENTRYPOINT ["/blog-server"]


