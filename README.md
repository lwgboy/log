# log
日志系统，可以根据panic精确输出是哪行代码出错，是调试的好帮手


第一步先下载代码到工程环境下


第二步调用初始化方法：log.Init()


第三步，在你需要调试的地方panic或者log.Write(err)



请注意：调用后后面的代码将不能执行，当前线程退出，所以请在子线程中进行调用
