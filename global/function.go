package global

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"
)

func ExecPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	return filepath.Abs(file)
}

// 清除目录
func CleanPath(path string) {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	Rmr(filepath.Join(cwd, path))
	CreateDir(path)
}

// 循环删除目录
func Rmr(paths ...string) {
	for _, path := range paths {
		log.Println("rm -r", path)
		os.RemoveAll(path)
	}
}

// 创建目录
func CreateDir(dirName string) {
	err := os.MkdirAll(dirName, 0777)
	if err != nil {
		log.Fatal(err)
	}
}

// 列出目录下所有一级目录名
func ListChildDir(path string) (dir []string, err error) {
	dirs, err := ioutil.ReadDir(path)
	if err != nil {
		return
	}
	for _, d := range dirs {
		if d.IsDir() {
			dir = append(dir, d.Name())
		}
	}
	return dir, nil
}

// 列出目录下所有文件(不包含文件夹)
func ListChildFiles(path string) (files []string, err error) {
	dirs, err := ioutil.ReadDir(path)
	if err != nil {
		return
	}
	for _, d := range dirs {
		if !d.IsDir() {
			files = append(files, d.Name())
		}
	}
	return files, nil
}

// 判断是否是一个文件
func IsFile(filePath string) bool {
	f, e := os.Stat(filePath)
	if e != nil {
		return false
	}
	return !f.IsDir()
}

// 多协程复制文件
func SyncCopyFiles(srcNames, dstNames []string) {
	time1 := time.Now().UnixNano()
	var wg sync.WaitGroup
	for i := 0; i < len(srcNames); i++ {
		wg.Add(1)
		go CopyFile(srcNames[i], dstNames[i], &wg)
	}
	wg.Wait()
	log.Printf("耗时%d毫秒\n", (time.Now().UnixNano()-time1)/1e6)
}

// 复制目录
func CopyFiles(srcNames, dstNames []string) {
	time1 := time.Now().UnixNano()
	for i := 0; i < len(srcNames); i++ {
		CopyFile(srcNames[i], dstNames[i])
	}
	log.Printf("耗时%d毫秒\n", (time.Now().UnixNano()-time1)/1e6)
}

// 同步复制文件
func CopyFile(srcName, dstName string, wg ...*sync.WaitGroup) {
	if len(wg) > 0 {
		defer wg[0].Done()
	}
	src, err := os.Open(srcName)
	if err != nil {
		log.Fatal(err, 153)
		return
	}
	defer src.Close()
	_, err = os.Open(dstName)
	if err == nil {
		//文件存在不修改
		return
	}
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err, 358)
		return
	}
	defer dst.Close()
	_, err = io.Copy(dst, src)
	if err != nil {
		log.Fatal(err, 364)
	}
}

// 零点时间戳
func DateZeroUnix(currentTime time.Time) int64 {
	timeStr := currentTime.Format(TimeSortParse)
	t, _ := time.ParseInLocation(TimeSortParse, timeStr, Location)
	return t.Unix()
}

// byte 转Int64
func ByteToInt(b []byte) int64 {
	v, e := strconv.ParseInt(string(b), 10, 64)
	if e == nil {
		return v
	}
	return 0
	/*	fmt.Println("string:",string(b))

		buf := bytes.NewBuffer(b)
		var x int64
		binary.Read(buf, binary.BigEndian, &x)
		fmt.Println("x:" ,x)
		return x
	*/
}

//数组根据指定的字符串分割
func Implode(glue string, activice []string) string {
	return strings.Replace(strings.Trim(fmt.Sprint(activice), "[]"), " ", glue, -1)
}
