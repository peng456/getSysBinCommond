
---
# getSysBinCommond

 获取操作系统提供的所有命令（PATH 目录下的命令）

## Go 版本

### 前提条件
- 本地已正确安装并配置了 Go 环境。

### 安装与使用步骤

   ```bash
   go install github.com/peng456/getSysBinCommond@latest && getSysBinCommond
   ```
### 解释   
1. **安装 `getSysBinCommond` 工具**

   使用以下命令将 `getSysBinCommond` 工具安装到本地 Go 的 bin 目录中：
   
   ```bash
   go install github.com/peng456/getSysBinCommond@latest
   ```

2. **运行命令**

   安装完成后，通过以下命令执行 `getSysBinCommond`：
   
   ```bash
   getSysBinCommond
   ```


## PHP 版本

### 前提条件
- 本地已正确安装并配置了 PHP 环境。
### 安装与使用步骤

   ```bash
   wget https://raw.githubusercontent.com/peng456/getSysBinCommond/master/getSysBinCommond.php && php getSysBinCommond.php
   ```
   or 高权限运行  
   ```bash
   wget https://raw.githubusercontent.com/peng456/getSysBinCommond/master/getSysBinCommond.php && sudo php getSysBinCommond.php
   ```
### 解释  
### 下载与运行脚本


1. **下载 `getSysBinCommond.php` 脚本**

   使用 `wget` 命令下载原始 PHP 脚本文件：
   
   ```bash
   wget https://raw.githubusercontent.com/peng456/getSysBinCommond/master/getSysBinCommond.php
   ```

2. **运行脚本**

   下载完成后，可以通过以下命令之一运行脚本：
   
   - 普通权限：
     ```bash
     php getSysBinCommond.php
     ```
   
   - 高权限（如需）：
     ```bash
     sudo php getSysBinCommond.php
     ```

---
