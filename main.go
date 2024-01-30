package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

/*
	@Author:@szsk2022
	@Email：admin@sunzishaokao.com
	@URL:https://gitee.com/szsk2/defender
*/

// 使用go-bindata嵌入文本文件，您将在终端中运行此命令：
// go-bindata -pkg main -o assets.go Remover/...

// 然后将生成的assets.go文件包括在项目中.

func main() {
	fmt.Println("Author:孙子烧烤")
	fmt.Println("Ver:1.0.0")
	fmt.Println("Email:admin@sunzishaokao.com")
	fmt.Println("URL:https://www.sunzishaokao.com")
	fmt.Println("")
	time.Sleep(2 * time.Second)
	fmt.Println("正在禁用Windows-Defender......")
	// 创建临时目录以提取嵌入的文件
	tempDir, err := ioutil.TempDir("", "defender-removal")
	if err != nil {
		fmt.Println("创建临时目录时出错:", err)
		return
	}
	defer os.RemoveAll(tempDir) // 程序退出后清理临时目录

	// 提取TKL.txt并处理任务
	// tklFile 是存储TKL文件的路径
	tklFile := filepath.Join(tempDir, "TKL.txt")

	// 通过Asset函数获取Remover/TKL.txt文件的数据
	data, err := Asset("Remover/TKL.txt")
	if err != nil {
		fmt.Println("提取TKL文件时出错:", err)
		return
	}

	// 将获取到的TKL文件数据写入到tklFile指定的文件中
	err = ioutil.WriteFile(tklFile, data, 0644)
	if err != nil {
		fmt.Println("写入TKL文件时出错:", err)
		return
	}

	// 从tklFile文件中读取每行数据作为任务，并执行任务kill命令
	tasks := readLines(tklFile)
	for _, task := range tasks {
		// 构建一个执行taskkill命令的Command对象
		cmd := exec.Command("taskkill.exe", "/f", "/im", task)

		// 执行Command对象的任务kill命令
		cmd.Run()
	}

	// 应用注册表文件
	// regFiles 是资源文件中REGS目录下的文件列表
	regFiles, _ := AssetDir("Remover/REGS")
	// 对于每个注册表文件
	for _, regFile := range regFiles {
		// 拼接注册表文件的完整路径
		regFilePath := filepath.Join(tempDir, regFile)
		// 从资源文件中提取注册表文件的内容
		data, err := Asset("Remover/REGS/" + regFile)
		if err != nil {
			fmt.Println("提取注册表文件时出错:", err)
			continue
		}
		// 将提取到的内容写入注册表文件中
		err = ioutil.WriteFile(regFilePath, data, 0644)
		if err != nil {
			fmt.Println("写入注册表文件时出错:", err)
			continue
		}
		// 执行regedit命令，打开注册表文件
		exec.Command("regedit.exe", "/s", regFilePath).Run()
	}

	// 禁用Defender/安全组件文件
	// fdlFile 指向临时目录中FDL.txt文件的路径
	fdlFile := filepath.Join(tempDir, "FDL.txt")
	// 从资源文件中提取FDL.txt文件的内容
	data, err = Asset("Remover/FDL.txt")
	if err != nil {
		fmt.Println("提取FDL文件时出错:", err)
		return
	}
	// 将提取到的内容写入FDL.txt文件中
	err = ioutil.WriteFile(fdlFile, data, 0644)
	if err != nil {
		fmt.Println("写入FDL文件时出错:", err)
		return
	}

	// 从FDL.txt文件中读取任务列表
	fdlTasks := readLines(fdlFile)
	// 对于每个任务，执行删除命令
	for _, file := range fdlTasks {
		cmd := exec.Command("cmd.exe", "/c", "del", "/f", "/q", file)
		cmd.Run()
	}

	// ddlFile 指向临时目录中DDL.txt文件的路径
	ddlFile := filepath.Join(tempDir, "DDL.txt")
	// 从资源文件中提取DDL.txt文件的内容
	data, err = Asset("Remover/DDL.txt")
	if err != nil {
		fmt.Println("提取DDL文件时出错:", err)
		return
	}
	// 将提取到的内容写入DDL.txt文件中
	err = ioutil.WriteFile(ddlFile, data, 0644)
	if err != nil {
		fmt.Println("写入DDL文件时出错:", err)
		return
	}

	// 从DDL.txt文件中读取文件夹列表
	ddlTasks := readLines(ddlFile)
	// 对于每个文件夹，执行删除命令
	for _, folder := range ddlTasks {
		cmd := exec.Command("cmd.exe", "/c", "rmdir", "/s", "/q", folder)
		cmd.Run()
	}

	fmt.Println("禁用命令执行完成，请重启系统！")
	pause()
}
