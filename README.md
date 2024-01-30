<h1 align="center">禁用Windows-Defender</h1>
<p align="center">一键永久禁用烦人的WD</p>

<p align="center">
<a href="./README.md">中文介绍</a> |
<a href="./README.en.md">English Description</a> 
</p>

<p align="center"> 
<img src="https://img.shields.io/badge/Author-孙子烧烤-orange.svg" title="Author: 孙子烧烤">
<img src="https://img.shields.io/badge/Go-1.21.6-brightgreen.svg" title="Go" />
<img src="https://img.shields.io/badge/version-v1.0-brightgreen.svg" title="Version: v1.0">
<img src="https://img.shields.io/badge/GPL-3.0-brightgreen.svg" title="GPL-3.0">
<img src="https://gitee.com/szsk/dwd/badge/star.svg?theme=dark" title="Star Count">  
<img src="https://gitee.com/szsk/dwd/badge/fork.svg?theme=dark" title="Fork Count">  

<p align="center">
<a href="https://www.sunzishaokao.com/">官方网址</a> 
</p>

<p align="center">源码地址：<a href="https://gitee.com/szsk/dwd">Gitee</a> | 
<a href="https://github.com/szsk2022/dwd">GitHub</a>
</p>

#### 介绍
一键永久禁用烦人的WD
![](https://www.sunzishaokao.com/wp-content/uploads/2024/01/20240130135058531-CA2B4965-6847-452e-A3CC-43FFF26203AB.png)
>Tips:本项目基于[windows-defender-remover](https://github.com/ionuttbara/windows-defender-remover)二次开发

#### 使用说明
1. 下载我们最新编译的程序[Gitee-Releases](https://gitee.com/szsk/kms/releases "Releases")
2. 直接打开运行即可，此时程序会自动执行禁用安全中心命令
3. 执行后，重启计算机即可生效

#### 支持的系统版本
* Windows 11
* Windows 10
* Windows Server 2016
* Windows Server 2019
* Windows Server 2022

#### 编译说明
1. 将源代码克隆到本地  
```
git clone https://gitee.com/szsk/kms.git
````
2. cd到项目根目录，根据自己的情况，修改Remover目录下的文件，然后执行嵌入命令
```
go-bindata -pkg main -o assets.go Remover/...
```
3. 编译源代码 
```
go build
```

#### 参与贡献

1.  Fork 本仓库
2.  新建 Feat_xxx 分支
3.  提交代码
4.  新建 Pull Request

