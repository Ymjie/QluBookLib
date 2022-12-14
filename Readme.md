### 介绍

Golang 开发的 齐鲁工业大学 图书馆六点自动预约程序，欢迎Push。

**截图**

![](img/01.png)

### 使用方法

1. 在IDlist.txt 选择中意的位置
2. 编辑config.yml
3. 运行程序


```powershell
Book.exe config.yml
```

### 注意事项

登陆成功会挤掉其他设备的登陆状态，5.30程序自动登录成功后，如果在其他设备登陆，此程序会失效哦。

**位于第一个的BookID将会优先并发请求预约，若预约失败将会尝试预约之后的BookID**，Booklist推荐添加3-5个位置

### 运行过程

1. 检查是否能成功登陆。
2. 5:30 尝试登陆
3. 6：00 开始预约
4. 等待第二天的5:30

### 更新历史
1. 2022年10月29日，发布v0.1
2. 2022年11月2日，发布v1.0，添加多账号，自定义线程数，通知功能。
3. 2022年11月6日，发布v1.1，fix bug,完善输出显示。
### TODO

- [x] 添加多账号、通知功能
- [ ] 当天座位监控并捡漏功能
- [ ] 整合IDlist.txt 更好的进行输入输出。
- [ ] 监控在意的BookID的信息（谁什么时候预约，什么时候开始使用、签离···）