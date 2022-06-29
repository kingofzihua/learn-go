* CHANGELOG : 变更日志
* api :
  * version :  版本
    * openapi : API 文档
    * protobuf : protobuf 文件
* bin :  build 后的二进制文件
* cmd :  整个项目启动的入口文件
  * modules : 模块
    * version : 版本
      * 
* configs :  配置文件目录
* docs : 文档
* examples : 示例代码
* internal :  业务逻辑
  * modules : 模块
    * adapters : 
    * biz :  业务逻辑
    * ports : 对外提供服务
      * grpc.go :  对外提供 grpc
      * http.go :  对外提供 http
* third_party :  依赖的第三方proto
* pkg :  项目内共享的代码
* web : 如果有 web 页面
* test :  测试相关(非单元测试)
* scripts : 脚本