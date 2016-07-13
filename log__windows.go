package log

import (
	"fmt"
	"log"
	//"flag"
	"os"
	"nqc.cn/utils"
	"time"
	"runtime"
	"syscall"
	//"github.com/golang/glog"

)

func SetStdHandle(stdhandle int32, handle syscall.Handle) error {

	kernel32         := syscall.MustLoadDLL("kernel32.dll")
	procSetStdHandle := kernel32.MustFindProc("SetStdHandle")

	r0, _, e1 := syscall.Syscall(procSetStdHandle.Addr(), 2, uintptr(stdhandle), uintptr(handle), 0)
	if r0 == 0 {
		if e1 != 0 {
			return error(e1)
		}
		return syscall.EINVAL
	}
	return nil
}

func redirect_err(f *os.File) {

	err = SetStdHandle(syscall.STD_ERROR_HANDLE, syscall.Handle(f.Fd()))
	if err != nil {
		fmt.Println("SetStdHandle failed: %v", err)
	}
}

func GetLogFile() *os.File {
	t3 := time.Now().Format("2006_01_02_15_04_05")
	fmt.Println(t3)
	path := utils.GetCurrPath() + "log/log_" + t3 + ".txt"
	logfile, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Printf("%s\r\n", err.Error())
		os.Exit(-1)
	}





	defer logfile.Close()
	return logfile
}

var errList []error
var strList []string
var ch chan bool

func Try(fun func(), handler func(interface{})) {
	defer func() {
		if err := recover(); err != nil {
			handler(err)
		}
	}()
	fun()
}
func writeAll() {



		//timeout := time.After(30 * time.Second)
		for {
			select {
			case <-ch:

				{

					//logger.SetPrefix("[ERROR]")
					//fmt.Println("打印错误日志Start")
					for _,value := range errList {
						//glog.Errorf(value)
						//fmt.Println(1)
						//delete(errList,key)
						print(value)


					}
					//fmt.Println(2)
					logger.SetPrefix("[Info]")
					for _,value := range strList {
						//delete(strList,key)
						logger.Println(value)
					}
					//fmt.Println("删除错误日志Start")
					errList=append(errList[:0],errList[len(errList):]...)
					strList=append(strList[:0],strList[len(strList):]...)
					//errList=append(errList[:len(errList)],errList[len(errList):]...)
					//strList=append(strList[:len(strList)],strList[len(strList):]...)
					//fmt.Println("打印错误日志End")
				}
				break

			}
		}

	//<-ch//阻塞协程
	//ch <- 1//释放协程



}

func Init() {

	//glog.Infoln("this is a test")
	ch = make(chan bool)
	//var ch chan int、
	t3 := time.Now().Format("2006_01_02_15_04_05")
	fmt.Println(t3)
	path := utils.GetCurrPath() + "log/log_" + t3 + ".txt"
	logfile, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Printf("%s\r\n", err.Error())
		os.Exit(-1)
	}
	redirect_err(logfile)
	defer logfile.Close()
	logger := log.New(logfile, "\r\n", log.Ldate|log.Ltime|log.Llongfile)

	logger.SetFlags(logger.Flags() | log.LstdFlags)
	logger.Println(int(time.Now().Unix()))

	logger.Println("log初始化完成")

	go writeAll()
	//flag.Parse()

}
func Write(err error) {
	//fmt.Println("错误长度")
	//glog.ErrorDepth(3,err)
	//glog.Flush()
	/*go func() {
		//logger.Fatal(err)
	}()*/
	fmt.Println("error: ",err)
	panic(err)
	//errList = append(errList,err)
	//fmt.Println(errList)
	//ch <- true
	//close(ch)
	//fmt.Println("错误长度")
	//fmt.Println(len(errList))
	//logger.Panicln(err)

}
func WriteString(info string) {
	/*go func() {
		logger.Printf(info)
	}()*/
	fmt.Println(info)
	strList = append(strList,info)
	//fmt.Println(strList)
	ch <- true
	//close(ch)
	//logger.Printf(info)

}
