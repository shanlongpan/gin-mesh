/**
* @Author:Tristan
* @Date: 2021/12/16 6:02 下午
 */

package xlog

import (
	"github.com/shanlongpan/gin-mesh/config"
	"github.com/shanlongpan/gin-mesh/consts"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"syscall"
)

var stdErrFile *os.File

func init() {
	err := RewriteStderrFile()
	if err != nil {
		log.Fatalln(err)
	}
}

// 标准输出写入到文件
func RewriteStderrFile() error {
	if runtime.GOOS == "windows" {
		return nil
	}

	stdErrFile, err := os.OpenFile(filepath.Join(consts.LogFileDir, config.Conf.LogName.StderrPanicLogFile+consts.Suffix), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	if err = syscall.Dup2(int(stdErrFile.Fd()), int(os.Stderr.Fd())); err != nil {
		return err
	}
	// 内存回收前关闭文件描述符
	runtime.SetFinalizer(stdErrFile, func(fd *os.File) {
		fd.Close()
	})

	return nil
}
