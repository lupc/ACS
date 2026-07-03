package pkg

import (
	"os"
	"acs/util"
	"time"

	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

type CleanTask struct {
	Dir            string        `yaml:"Dir"`            //清理目录
	Extensions     string        `yaml:"Extensions"`     //目标文件后缀(多个用,分割),为空则不限定后缀
	Days           int           `yaml:"Days"`           //文件保留天数,超过该天数将被删除
	RemoveEmptyDir bool          `yaml:"RemoveEmptyDir"` //是否删除空目录
	CheckInterval  time.Duration `yaml:"CheckInterval"`  //清理检测间隔
	IsEnable       bool          `yaml:"IsEnable"`       //是否启用
}
type Config struct {
	Configs []*CleanTask //支持多个清理任务
}

func GetConfig(path string) (cfg *Config) {
	defer util.LogRecoverToError("加载配置出错")

	data, err := os.ReadFile(path)
	if err != nil {
		util.GetLogger().Error("打开配置文件出错", zap.Any("path", path), zap.Error(err))
		if os.IsNotExist(err) {
			cfg = new(Config)
			cfg.Configs = append(cfg.Configs, &CleanTask{
				Dir:            "C:\\temp",
				Extensions:     ".log,.tmp",
				Days:           7,
				RemoveEmptyDir: false,
				CheckInterval:  time.Hour,
				IsEnable:       false,
			})
			data, err := yaml.Marshal(cfg)
			if err != nil {
				util.GetLogger().Error("序列化配置出错", zap.Any("path", path), zap.Error(err))
			} else {
				newFile, err := os.Create(path)
				if err != nil {
					util.GetLogger().Error("创建配置文件出错", zap.Any("path", path), zap.Error(err))
				} else {
					util.GetLogger().Info("成功创建配置文件", zap.Any("path", path))
					newFile.Write(data)
				}
				newFile.Close()
			}
		}
		return
	}
	if len(data) == 0 {
		util.GetLogger().Error("配置文件为空", zap.Any("path", path))
		return
	}
	util.GetLogger().Sugar().Infof("成功加载配置:\n%v", string(data))
	var ccc = &Config{}
	err = yaml.Unmarshal(data, ccc)
	if err != nil {
		util.GetLogger().Error("解析配置文件出错", zap.Any("path", path), zap.Error(err))
		return
	}
	cfg = ccc
	return
}
