<h1 align="center">Disabling Windows Defender</h1>
<p align="center">One-click permanent disabling of pesky WD</p>

<p align="center">
<a href="./README.md">Chinese Introduction</a> | 
<a href="./README.en.md">English Description</a> 
</p>

<p align="center"> 
<img src="https://img.shields.io/badge/Author-SZSK-orange.svg" title="Author: SZSK">
<img src="https://img.shields.io/badge/Go-1.21.6-brightgreen.svg" title="Go Version: 1.21.6">
<img src="https://img.shields.io/badge/version-v2.0-brightgreen.svg" title="Version: v2.0">
<img src="https://img.shields.io/badge/GPL-3.0-brightgreen.svg" title="License: GPL-3.0">
<img src="https://gitee.com/szsk/dwd/badge/star.svg?theme=dark" title="Star Count">  
<img src="https://gitee.com/szsk/dwd/badge/fork.svg?theme=dark" title="Fork Count">  

<p align="center">
<a href="https://www.sunzishaokao.com/">Official Website</a> 
</p>

<p align="center">Source Code Repositories:<br>
<a href="https://gitee.com/szsk/dwd">Gitee</a> | 
<a href="https://github.com/szsk2022/dwd">GitHub</a>
</p>

#### Introduction
This tool provides one-click, permanent disabling of the bothersome Windows Defender.
[![](https://www.sunzishaokao.com/wp-content/uploads/2024/01/20240130135058531-CA2B4965-6847-452e-A3CC-43FFF26203AB.png)](https://www.sunzishaokao.com/wp-content/uploads/2024/01/20240130135058531-CA2B4965-6847-452e-A3CC-43FFF26203AB.png)
>Tips: This project is built upon and further developed from [windows-defender-remover](https://github.com/ionuttbara/windows-defender-remover).

#### Usage Instructions
1. Download our latest compiled program from [Gitee Releases](https://gitee.com/szsk/kms/releases "Releases").
2. Simply run the program; it will automatically execute commands to disable the Security Center.
3. Restart your computer for changes to take effect after execution.

#### Supported Operating Systems
* Windows 11
* Windows 10
* Windows Server 2016
* Windows Server 2019
* Windows Server 2022

#### Compilation Instructions
1. Clone the source code to your local machine:
```
git clone https://gitee.com/szsk/kms.git
``` 
2. Navigate to the root directory of the project, modify files in the Remover directory as needed, then execute the embedding command:
```
go-bindata -pkg main -o assets.go Remover/...
```
3. Compile the source code:
```
go build
```

#### Contributing
1. Fork this repository.
2. Create a new branch named `Feat_xxx`.
3. Commit your changes.
4. Open a new Pull Request.