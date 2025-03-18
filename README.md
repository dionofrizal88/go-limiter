## Rate Limiter

[![Go](https://img.shields.io/badge/go-1.23-green.svg)](https://golang.org/)

## 🏃 Getting Started
Hello, in this time I want to share about how to implement rate limiter using go. When building APIs, one of the biggest challenges is protecting the system from abuse or excessive load. Here are some scenarios where a rate limiter is crucial.

Clone this repo:
```shell script
  git clone
```
On limiter folder I create rate limiter using redis with package `redis_rate`. I am using token bucket method for handel rate limit by identifier on this case is by `IP`. Besides that I handle global rate limit to limit request on my APIs and avoid `downtime`.

I create unit testing for how to use `redis_rate` package and custom the limit value. You can try the unit test and implement rate limit for your project.

Thank you for read me