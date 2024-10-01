# 服务治理

--- 

## 超时控制

## 限流

### 为什么使用限流

针对某些业务场景，需要限制对后端服务的访问，防止后端服务过载。例如本项目中的：发布视频接口。
在本项目中发布视频接口是耗时较长，如果后端服务过载，会导致发布视频接口响应时间过长，影响用户体验，
且存在依赖于外部服务的调用，为了避免第三方的并发调用导致后端服务过载，需要使用限流。

### 限流分类

- api 限流
- 服务限流

## 熔断
1. 目的：
服务熔断的主要目的是防止系统雪崩，即当一个服务开始失败时，防止这种失败影响到其他服务，最终导致整个系统瘫痪。
2. 触发机制：
服务熔断通常由断路器模式触发，当检测到下游服务的失败率超过一定阈值时，断路器会“打开”，阻止进一步的请求到达故障服务，直到该服务恢复。
3. 行为：
当断路器打开时，任何对该服务的请求将不会被转发，而是立即返回一个错误或预定义的响应，避免了对故障服务的无效调用。
4. 恢复机制：
断路器在一段时间后会自动进入“半开”状态，允许少量请求通过以检查服务是否已恢复。如果服务确实已恢复，断路器会回到“关闭”状态，恢复正常的请求转发。

> 服务熔断专注于快速失败和隔离故障服务，以防止连锁反应。


## 降级

1. 目的：
服务降级的目标是在资源紧张或服务压力过大时，通过牺牲非核心功能或降低服务质量来保证核心服务的正常运行。
2. 触发机制：
服务降级可能由多种因素触发，包括但不限于资源使用率过高、服务响应时间过长或系统负载过大。
3. 行为：
在服务降级时，系统可能会选择不提供某些功能，返回静态或缓存数据，或者简化服务接口，以减少资源消耗和提升响应速度。
4. 恢复机制：
服务降级通常是一种临时措施，当系统压力减轻或资源得到补充时，服务可以逐步恢复到正常状态，重新启用之前降级的功能。

> 服务降级侧重于在资源受限的情况下，通过调整服务级别来维持系统的关键功能。

## 负载均衡



