Go 语言北京聚会抽奖程序
===================

### 数据格式

在当前目录下放置文件名位 `data.txt` 的数据文件，格式为：

```
<姓名> <手机号或邮箱>
```

在没有数据文件的情况下会模拟数据。

### 下载使用

	go get -d github.com/macaron-contrib/examples/lottery_go_beijing_201501
	go build -o lottery && ./lottery

然后访问 http://localhost:4000