# ACS - Auto Clean Service

自动清理指定目录下的过期文件，支持按文件后缀过滤、空目录清理，以 Windows 服务运行。

## 功能

- 按目录清理：可配置多个清理任务，每个任务独立指定目录
- 按后缀过滤：支持 `.log,.tmp` 等多后缀匹配，为空则不限定
- 按天数删除：删除超过指定天数的文件
- 空目录清理：可选在文件删除后自动清理空目录（从最深层逐层往上）
- Windows 服务：自启动、后台运行，无需用户干预

## 配置文件

```yaml
configs:
    - Dir: D:\logs
      Extensions: ".log,.tmp"
      Days: 7
      RemoveEmptyDir: true
      CheckInterval: 1h
      IsEnable: true
```

| 字段 | 说明 |
|---|---|
| `Dir` | 要清理的目录 |
| `Extensions` | 文件后缀，多个用逗号分隔，为空则不限定 |
| `Days` | 文件保留天数，超过该天数将被删除 |
| `RemoveEmptyDir` | 是否删除空目录（`true`/`false`） |
| `CheckInterval` | 清理检测间隔，如 `1h`、`30m`、`10s` |
| `IsEnable` | 是否启用该任务 |

## 使用

### 编译

```bat
build.bat
```

### 安装服务

```bat
acs_install.bat
```

### 启动服务

```bat
acs_start.bat
```

### 停止服务

```bat
acs_stop.bat
```

### 卸载服务

```bat
acs_uninstall.bat
```

## 日志

使用 `go-myzap` + `zap` 输出日志，默认写入 `logs/` 目录（自动滚动），同时输出到控制台。
