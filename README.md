## This project is developing, it still can't be used.

## 正在开发中，还不能用


## 数据交换设计
插件在服务端注册自己的服务URL，需要声明两个URL:
 - Client URL: 和客户端交互用的URL
 - Phone URL: 与手机交流的URL。这里需要注意以下几点。
    - 当客户端想要与手机端交流时，就会向服务器发送信息，然后服务器会转发信息给手机端，手机端会访问这条URL(访问方式以及该URL应该已经在手机端插件上定义过)以获取客户端想要发送给手机端的数据。
    - 当手机端主动想要与客户端或服务端沟通时，也会访问该URL。
    - ZServer会将是来自手机端的请求还是来自客户端的请求塞进Context.Param[\"MessageFrom\"]里，如果是手机，值就是phone, 是客户端值就是client, 插件可通过值来判断请求是来自哪里。
