FROM alpine:3.7
MAINTAINER gbnk0 <github.com/gbnk0>
ARG VCS_REF
LABEL org.label-schema.vcs-ref=$VCS_REF \
        org.label-schema.vcs-url="https://github.com/gbnk0/image-to-text-api"

WORKDIR /app
COPY . /app
RUN apk add git gcc g++ go tesseract-ocr tesseract-ocr-dev && \
    go get github.com/gin-gonic/gin && \
    go get github.com/otiai10/gosseract && \
    go build main.go && \
    apk del git gcc g++ go && \
    rm main.go
ENV GIN_MODE release
CMD ["./main"]
