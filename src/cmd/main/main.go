package main

import (
	"crypto/sha1"
	"flag"
	"fmt"
	"github.com/ryanuber/go-glob"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"
)

const VERSION = "1.02"

var (
	folderpath  string
	outfile     string
	filter      string
	groutineNum int64
)

func init() {
	flag.StringVar(&folderpath, "p", "./", "遍历的文件夹路径，默认为当前路径")
	flag.StringVar(&outfile, "o", "tree.txt", "输出目录树的文件,默认tree.txt")
	flag.StringVar(&filter, "f", "", "忽略哪些目录、文件，支持通配符,多个用|分隔")
	flag.Int64Var(&groutineNum, "gr", 10, "并发数量.默认10")
	// 控制使用多核数量,获取当前CPU核数-2,减少
	runtime.GOMAXPROCS(runtime.NumCPU() - 1)
}

func main() {
	flag.Parse()

	fmt.Printf("当前版本 : %s", VERSION)
	fmt.Printf("参数 : \n folderpath : %s \n outfile : %s \n filter : %s \n groutineNum : %d \n", folderpath, outfile, filter, groutineNum)

	if len(folderpath) == 0 {
		folderpath = flag.Arg(0)
	}
	filterArr := make([]string, 0)
	if len(filter) != 0 {
		filterArr = strings.Split(filter, "|")
	}
	// fmt.Printf("filterArr : %v \n", filterArr)

	// 输出目录树的文件
	writef, err := os.OpenFile(outfile, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		fmt.Printf("打开key文件失败. Error : %v, Path : %s", err, outfile)
		return
	}
	defer writef.Close()

	// 遍历文件夹
	var wg sync.WaitGroup
	waitChan := make(chan bool, groutineNum)

	t1 := time.Now()
	filepath.Walk(folderpath, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		// 过滤
		for _, f := range filterArr {
			if glob.Glob(f, path) {
				return nil
			}
		}
		// 忽略文件夹
		if f.IsDir() {
			return nil
		}
		wg.Add(1)
		go func() {
			for {
				waitChan <- true
				writef.WriteString(fmt.Sprintf("%s,%s,%d \n", f.Name(), fileSHA1(path), f.Size()))
				writef.Sync()
				<-waitChan
				wg.Done()
				break
				runtime.Gosched()
			}
		}()
		return nil
	})
	wg.Wait()
	t2 := time.Now()

	fmt.Printf("开始时间 : %v \n 结束时间 : %v \n 用时 : %v", t1, t2, t2.Sub(t1))
}

// 生成文件 SHA1
func fileSHA1(path string) string {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND, os.ModePerm)
	defer file.Close()
	if err != nil {
		log.Panicf("生成文件SHA1失败 : %v, 文件路径 : %s", err, path)
		return ""
	}

	h := sha1.New()
	_, err = io.Copy(h, file)
	if err != nil {
		log.Panicf("生成文件SHA1失败 : %v, 文件路径 : %s", err, path)
		return ""
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}

// 获取文件名字,大小
func fileSize(path string) (string, int64, error) {
	file, err := os.Stat(path)
	if err != nil {
		return "", 0, err
	}
	return file.Name(), file.Size(), nil
}

// 检查目录是否存在
func isDirExists(path string) bool {
	dir, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return dir.IsDir()
}
