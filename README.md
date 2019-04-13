# moreal.kr-links

## How to work

1. 유저가 `moreal.kr/~<keyword>`로 접근한다
2. cloudflare의 page-rule을 통해서 `links.moreal.kr/<keyword>`로 redirect합니다
3. `links.yml`을 파싱하여 있으면 그 url로 redirect 합니다

## Stack

- Go
- CloudBuild
- AppEngine (golang 1.11 standard, service-route)