FROM alpine:3.7
WORKDIR /app
ADD . /app
RUN apk add git gcc g++ go tesseract-ocr tesseract-ocr-dev && \
    go get github.com/gin-gonic/gin && \
    go get github.com/otiai10/gosseract && \
    go build main.go && \
    apk del git gcc g++ go
ENV GIN_MODE release
CMD ["./main"]