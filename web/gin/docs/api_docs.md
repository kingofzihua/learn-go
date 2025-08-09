# API 中间件文档

| 方法 | 路径 | 中间件执行链 |
|:---|:---|:---|
| GET | `/index` | `gin.Logger -> gin.Recovery` |
| GET | `/hello/gin` | `gin.Logger -> gin.Recovery` |
| GET | `/hello2/gin` | `gin.Logger -> gin.Recovery` |
| GET | `/user/:username` | `gin.Logger -> gin.Recovery -> CorsMiddleware` |
| GET | `/hi:username` | `gin.Logger -> gin.Recovery -> CorsMiddleware` |
| GET | `/helo:username/Name` | `gin.Logger -> gin.Recovery -> CorsMiddleware` |
| GET | `/blog/:category/:post` | `gin.Logger -> gin.Recovery -> CorsMiddleware` |
| GET | `/blog/:category/:post/:name` | `gin.Logger -> gin.Recovery -> CorsMiddleware` |
| GET | `/bg/:category/:post/:name` | `gin.Logger -> gin.Recovery -> CorsMiddleware` |
| GET | `/testing/:name` | `gin.Logger -> gin.Recovery -> CorsMiddleware` |
| GET | `/m3u` | `gin.Logger -> gin.Recovery -> CorsMiddleware` |
| GET | `/api/v1/users` | `gin.Logger -> gin.Recovery -> CorsMiddleware -> AuthMiddleware` |
| POST | `/api/v1/users/create` | `gin.Logger -> gin.Recovery -> CorsMiddleware -> AuthMiddleware -> LoggerMiddleware` |
