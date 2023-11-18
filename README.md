# happy-translate
Automator+翻译API实现的便捷翻译工具，通过快捷键快速翻译单词并记录到单词本，并将翻译结果写入剪贴板

灵感来自谷歌浏览器插件[划词翻译](https://chrome.google.com/webstore/detail/%E5%88%92%E8%AF%8D%E7%BF%BB%E8%AF%91/ikhdkkncnoglghljlkmcimlnlhkeamad?hl=zh-CN)，只不过它只能在谷歌浏览器使用并且单词本需要付费，本着能白嫖就绝不花钱的原则，利用现有资源做了这个小工具，虽然很粗糙，但是能用就行，有什么好的建议，欢迎沟通！！！

### 使用场景

> 外文文档阅读时的单词查询

选中文本，Command+q，弹出翻译结果通知

> 文本翻译替换

选中文本，Command+q，Command+v，完成替换

> 单词本整理

单词本文件：~/WordBook.txt


### 体验一下
下载translate.workflow，双击安装，设置快捷键Command+q


选中一段文本，按Command+q试试吧！

### 私有化部署
1. 百度翻译平台注册开发者账号，并申请免费体验[通用文本翻译API](https://fanyi-api.baidu.com/product/11)和[语种识别API](https://fanyi-api.baidu.com/product/14)
2. 将[APPID和密钥](https://fanyi-api.baidu.com/api/trans/product/desktop)写入公网服务器环境变量
    ```shell
    export BAIDU_TRANSLATE_APPID=202310xxxxxxxxxx
    export BAIDU_TRANSLATE_APPKEY=XIXXxxxxxxxxxxx
    ```
3. 在公网服务器运行翻译服务`go run main.go`
4. 在Mac上打开./translate.workflow/Contents/document.wflow 替换 http://118.24.149.250:8080 为你自己的IP和端口，双击安装translate.workflow




### TODO
- [x] 百度翻译
- [ ] 谷歌翻译

