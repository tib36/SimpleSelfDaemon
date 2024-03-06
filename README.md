# SimpleSelfDaemon
单文件自我守护进程的简单demo

代码很简单，直接运行该程序会形成守护进程

每隔几秒会检测进程列表，如果子进程不存在则会创建带参数的子进程并记录进程pid

直接运行时为守护进程，带参数运行时为子进程

作为子进程时才执行程序功能，守护进程正常情况下仅监测子进程存活

Example:
```
PS > go build .\SimpleSelfDaemon.go
PS > .\SimpleSelfDaemon.exe
[*]Process does not exist. Restarting...
[+]Process exists. pid:9896
[+]Process exists. pid:9896
[+]Process exists. pid:9896
[+]Process exists. pid:9896
[+]Process exists. pid:9896
[+]Process exists. pid:9896
[*]Process does not exist. Restarting...
[+]Process exists. pid:11384
[+]Process exists. pid:11384
```
