### 题目
请仔细看题：一个生成目录树哈希的小工具

1. 用golang开发，代码放到github上，用github进行问题跟踪
2. 对整个目录下的所有文件进行遍历，获取所有文件的大小和计算文件的sha1哈希值，记录在一个文件里面
3. 结果文件格式：每一行一个文件，用逗号隔开，前面是文件名称，后面是哈希值，文件大小
4. 需要可以指定忽略哪些目录、文件，需要支持通配符
5. 代码实现简洁，能够最大限度利用上多核性能，运行性能高得分高
6. 要求通过测试代码自我证明代码能够可靠运行并正确实现上述功能
7. 要求写安全稳定可靠的，可读性强的代码
8. 要求支持go get, go run, go build命令
9. 要求命令行友好：支持help命令，root参数，filter参数, 同时支持
10. 不限时间，但是越快完成越好，做好了就提交github，告诉我你的链接
