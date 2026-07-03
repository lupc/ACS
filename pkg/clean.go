package pkg

import (
	"acs/util"
	"os"
	"path/filepath"
	"strings"
	"time"

	"go.uber.org/zap"
)

func CleanProcess(cfg *CleanTask) {
	if cfg == nil {
		return
	}
	defer util.LogRecoverToError("执行清理出错")

	util.GetLogger().Sugar().Infof("开始清理任务[%v]...", cfg.Dir)

	var exts []string
	if cfg.Extensions != "" {
		exts = strings.Split(cfg.Extensions, ",")
		for i, ext := range exts {
			exts[i] = strings.TrimSpace(ext)
			if !strings.HasPrefix(exts[i], ".") {
				exts[i] = "." + exts[i]
			}
		}
	}

	now := time.Now()
	cutoff := now.Add(-time.Duration(cfg.Days) * 24 * time.Hour)

	err := filepath.Walk(cfg.Dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if info.IsDir() {
			return nil
		}
		if info.ModTime().After(cutoff) {
			return nil
		}
		if len(exts) > 0 && !matchExtension(info.Name(), exts) {
			return nil
		}
		err = os.Remove(path)
		if err != nil {
			util.GetLogger().Error("删除文件失败", zap.Any("path", path), zap.Error(err))
		} else {
			util.GetLogger().Info("已删除文件", zap.Any("path", path))
		}
		return nil
	})
	if err != nil {
		util.GetLogger().Error("遍历目录出错", zap.Any("dir", cfg.Dir), zap.Error(err))
	}

	if cfg.RemoveEmptyDir {
		removeEmptyDirs(cfg.Dir)
	}
}

func removeEmptyDirs(root string) {
	for {
		removed := false
		filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if err != nil || !info.IsDir() || path == root {
				return nil
			}
			entries, err := os.ReadDir(path)
			if err != nil || len(entries) > 0 {
				return nil
			}
			if err := os.Remove(path); err == nil {
				util.GetLogger().Info("已删除空目录", zap.Any("path", path))
				removed = true
			}
			return nil
		})
		if !removed {
			break
		}
	}
}

func matchExtension(name string, exts []string) bool {
	name = strings.ToLower(name)
	for _, ext := range exts {
		if strings.HasSuffix(name, strings.ToLower(ext)) {
			return true
		}
	}
	return false
}
