FROM alpine:3.7

ARG VCS_REF
LABEL maintainer="github.com/gbnk0"
LABEL org.label-schema.vcs-ref=$VCS_REF \
        org.label-schema.vcs-url="https://github.com/gbnk0/image-to-text-api"

WORKDIR /app
COPY main.go /app/
RUN apk add git gcc g++ go tesseract-ocr tesseract-ocr-dev tesseract-ocr-data-fra && \
    go get github.com/gin-gonic/gin && \
    go get github.com/otiai10/gosseract && \
    GOOS=linux go build -ldflags="-w" main.go && \
    apk del git gcc g++ go tesseract-ocr-dev && \
    rm main.go && rm -rf /root/go && \
    rm -rf /var/cache/apk/*

ENV GIN_MODE release
CMD ["./main"]
