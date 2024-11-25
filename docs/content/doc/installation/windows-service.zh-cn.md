---
date: "2016-12-21T15:00:00-02:00"
title: "注册为Windows服务"
slug: "windows-service"
weight: 10
toc: true
draft: false
menu:
  sidebar:
    parent: "installation"
    name: "Windows服务"
    weight: 30
    identifier: "windows-service"
---

# 注册为Windows服务

要注册为Windows服务，首先以Administrator身份运行 `cmd`，然后执行以下命令：

```
sc create nxgit start= auto binPath= ""C:\nxgit\nxgit.exe" web --config "C:\nxgit\custom\conf\app.ini""
```

别忘了将 `C:\nxgit` 替换成你的 Nxgit 安装目录。

之后在控制面板打开 "Windows Services"，搜索 "nxgit"，右键选择 "Run"。在浏览器打开 `http://localhost:3000` 就可以访问了。（如果你修改了端口，请访问对应的端口，3000是默认端口）。

## 从Windows服务中删除

以Administrator身份运行 `cmd`，然后执行以下命令：

```
sc delete nxgit
```
