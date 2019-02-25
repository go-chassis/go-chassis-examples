# A circuit breaker example to show circuit breaker feature
here is the config
```yaml
---
---
cse:
  isolation:
    Consumer:
      timeoutInMilliseconds: 1000
      maxConcurrentRequests: 100
  circuitBreaker:
    Consumer:
      enabled: false
      forceOpen: false
      forceClosed: false
      sleepWindowInMilliseconds: 10000
      requestVolumeThreshold: 20
      errorThresholdPercentage: 1
  #容错处理函数，目前暂时按照开源的方式来不进行区分处理，统一调用fallback函数
  fallback:
    Consumer:
      enabled: true
  fallbackpolicy:
    Consumer:
      policy: throwexception
```
time out is 1000ms

there is 3 APIs, 2 of them has different problem

1. client calls a API which has a dead lock, it will cause timout, so that circuit breaker will isolate this API

2. servcer has a lot of error response, circuit breaker will isolate this API

in additinal, if client has a lot of concurrency, circuit breaker will isolate the API



the runtime metrics is exported by prometheus exporter
check 127.0.0.1:5000/metrics and 127.0.0.1:5001/metrics to observe services
