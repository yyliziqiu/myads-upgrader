package svc

import (
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/yyliziqiu/zlib/zfile"
	"github.com/yyliziqiu/zlib/zlog"
)

func BasePath(path string) string {
	return filepath.Join("C:\\Users\\Administrator\\Myads", path)
}

func Boot() {
	zlog.Info("Boot upgrader.")

	ticker := time.NewTicker(time.Second)
	for {
		t := <-ticker.C
		if t.Unix()%30 == 0 {
			upgrade()
		}
	}
}

func upgrade() {
	var (
		exeFile = "sender.exe"
		tmpFile = "sender.tmp"
		exePath = BasePath(exeFile)
		tmpPath = BasePath(tmpFile)
	)

	ok, err := zfile.Exist(tmpPath)
	if err != nil {
		zlog.Warnf("Exist file failed, path: %s, error: %v,", tmpPath, err)
		return
	}
	if !ok {
		// zlog.Debug("No upgrade file")
		return
	}

	err = ExecCmd(exePath, "stop")
	if err != nil {
		zlog.Warnf("Excute stop failed, path: %s, error: %v,", exePath, err)
		return
	}
	time.Sleep(3 * time.Second)

	if ok2, _ := zfile.Exist(exePath); ok2 {
		if err = os.Remove(exePath); err != nil {
			zlog.Warnf("Remove file failed, path: %s, error: %v,", exePath, err)
			return
		}
	}

	err = os.Rename(tmpPath, exeFile)
	if err != nil {
		zlog.Warnf("Rename file failed, path: %s, error: %v,", tmpPath, err)
		return
	}

	err = ExecCmd(exePath, "start")
	if err != nil {
		zlog.Warnf("Excute start failed, path: %s, error: %v,", exePath, err)
		return
	}
}

func ExecCmd(cmd string, action string) error {
	c := exec.Command("cmd", "/c", cmd, action)
	c.Stdout = nil
	return c.Run()
}
