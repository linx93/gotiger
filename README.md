# gotiger
go tiger

### 搞一个后端脚手架

### 开发计划
1. 项目结构设计，参考https://github.com/golang-standards/project-layout
2. 项目***IOC***、***DI***、***AOP***、***代理***的设计，参考阿里的IOC-golang等，这个需要花时间多思考
   - 暂定为手动SET方法注入
   - GET方法获取，包含GET和GETProxy两种
      - GET获取到的就是原始对象
      - GETProxy获取到的就是原生对象增强过的代理对象，在这里提供AOP能力
         - d ddd
      - 考虑需不需要支持单列模式和原型模式
3. 配置相关，基于viper、cobra
    - 基础配置读取
    - 环境变量读取
    - cobra实现自己的命令，具体功能待定
    - 其他
4. 日志，基于zap实现
    - zap
    - 其他
5. rabc相关，基于casbin实现
    - rabc基础表
    - 不支持多租户
    - 支持多租户
    - 其他
6. grpc的支持
7. 路由封装，基于gin来做二次封装
8. 常用组件封装
    - websocket
    - redis
    - mysql
    - postgresql
    - uploader
    - 其他
9. 常用工具封装
    - crypto 关于加解密
    - id生成器 
        - 雪花算法
        - uuid
        - 其他
    - http
    - 文件操作
    - 字符串操作
    - 时间日期操作
    - 其他。。。
10. 代码生成器
    - 通过cobra实现命令来执行生成
    - 其他
11. 最后再考虑搞不搞前端