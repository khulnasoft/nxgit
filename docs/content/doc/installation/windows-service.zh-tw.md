---
date: "2016-12-21T15:00:00-02:00"
title: "註冊為 Windows 服務"
slug: "windows-service"
weight: 10
toc: true
draft: false
menu:
  sidebar:
    parent: "installation"
    name: "Windows 服務"
    weight: 30
    identifier: "windows-service"
---

# 註冊為 Windows 服務

要註冊為 Windows 服務，首先要以管理者身份執行 `cmd`，跳出命令列視窗後執行底下指令：

```
sc create nxgit start= auto binPath= ""C:\nxgit\nxgit.exe" web --config "C:\nxgit\custom\conf\app.ini""
```

別忘記將 `C:\nxgit` 取代為您的 Nxgit 安裝路徑。

之後打開 "Windows Services"，並且搜尋服務名稱 "nxgit"，按右鍵選擇 "Run"。在瀏覽器打開 `http://localhost:3000` 就可以成功看到畫面 (如果修改過連接埠，請自行修正，3000 是預設值)。

## 刪除服務

要刪除 Nxgit 服務，請用管理者身份執行 `cmd` 並且執行底下指令：

```
sc delete nxgit
```
