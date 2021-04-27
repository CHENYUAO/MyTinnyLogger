学习Go语言的过程中，跟随[李文周老师Go语言bilibili教学视频](https://www.bilibili.com/video/BV16E411H7og)所写的一个自定义的日志库。

## 实现功能如下

1. 支持往不同的地方输出日志

2. 日志分级别

   （1） Trace

   （2） Debug

   （3） Info

   （4） Warning

   （5） Error

   （6） Fatal

3. 日志要支持开关

4. 完整的日志记录要包含时间、行号、文件名、日志级别、日志信息

5. 日志文件要切割

6. 日志可以支持格式化输出

   （1）按文件大小切割：每次记录之前判断以下文件大小，如果达到最大值就写新文件

   （2）按日期切割

7. 异步方式写日志，提升性能

