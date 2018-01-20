# log

[![Build Status](https://travis-ci.org/NiuStar/log.svg?branch=master)](https://travis-ci.org/NiuStar/log)
[![Build status](https://ci.appveyor.com/api/projects/status/ksii33qx18d94h6v?svg=true)](https://ci.appveyor.com/project/NiuStar/log)
[![Go Report Card](https://goreportcard.com/badge/github.com/NiuStar/log)](https://goreportcard.com/report/github.com/NiuStar/log)



日志系统，可以根据panic精确输出是哪行代码出错，是调试的好帮手

#新增特性：
在高并发下的多线程的日志同步，方便精确定位代码问题

#使用方法

第一步先下载代码到工程环境下


第二步在线程/进程开始的调用初始化方法：fmt.Start()


第三步，在你需要调试的地方fmt.Println("输出内容")

最后，在线程/进程开始的调用初始化方法：fmt.Over()

跨平台应用进行
