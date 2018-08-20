# image-to-text-api
Recognize text on JPEG images using tesseract ocr (gosseract)

[![Build Status](https://travis-ci.org/gbnk0/image-to-text-api.svg?branch=master)](https://travis-ci.org/gbnk0/image-to-text-api)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/48c74f6d0f274b54af19263e5d9bb26d)](https://www.codacy.com/project/gbnk0/image-to-text-api/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=gbnk0/image-to-text-api&amp;utm_campaign=Badge_Grade_Dashboard)
[![](https://images.microbadger.com/badges/image/gbnk0/image-to-text-api.svg)](https://microbadger.com/images/gbnk0/image-to-text-api "Get your own image badge on microbadger.com")
[![](https://images.microbadger.com/badges/version/gbnk0/image-to-text-api.svg)](https://microbadger.com/images/gbnk0/image-to-text-api "Get your own version badge on microbadger.com")
[![](https://images.microbadger.com/badges/commit/gbnk0/image-to-text-api.svg)](https://microbadger.com/images/gbnk0/image-to-text-api "Get your own commit badge on microbadger.com")

### QUICK START

```
docker run -p8000:8000 gbnk0/image-to-text-api:latest
```

### API ENDPOINTS

#### POST /text

Recognize text on an image from url

- Example:

```
#> curl -X POST   http://localhost:8000/text -F file=@/home/user/image.jpg
#> {"status":"success","text":"HELLO"}
```

#### GET /version

Get the service build number

- Example:

```
#> curl -X GET   http://localhost:8000/version
#> {"status":"success","version":"15"}
```
