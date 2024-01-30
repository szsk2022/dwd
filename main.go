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
	tklFile := filepath.Join(tempDir, "TKL.txt")
	data, err := Asset("Remover/TKL.txt")
	if err != nil {
		fmt.Println("提取TKL文件时出错:", err)
		return
	}
	err = ioutil.WriteFile(tklFile, data, 0644)
	if err != nil {
		fmt.Println("写入TKL文件时出错:", err)
		return
	}

	tasks := readLines(tklFile)
	for _, task := range tasks {
		cmd := exec.Command("taskkill.exe", "/f", "/im", task)
		cmd.Run()
	}

	// 应用注册表文件
	regFiles, _ := AssetDir("Remover/REGS")
	for _, regFile := range regFiles {
		regFilePath := filepath.Join(tempDir, regFile)
		data, err := Asset("Remover/REGS/" + regFile)
		if err != nil {
			fmt.Println("提取注册表文件时出错:", err)
			continue
		}
		err = ioutil.WriteFile(regFilePath, data, 0644)
		if err != nil {
			fmt.Println("写入注册表文件时出错:", err)
			continue
		}
		exec.Command("regedit.exe", "/s", regFilePath).Run()
	}

	// 禁用Defender/安全组件文件
	fdlFile := filepath.Join(tempDir, "FDL.txt")
	data, err = Asset("Remover/FDL.txt")
	if err != nil {
		fmt.Println("提取FDL文件时出错:", err)
		return
	}
	err = ioutil.WriteFile(fdlFile, data, 0644)
	if err != nil {
		fmt.Println("写入FDL文件时出错:", err)
		return
	}

	fdlTasks := readLines(fdlFile)
	for _, file := range fdlTasks {
		cmd := exec.Command("cmd.exe", "/c", "del", "/f", "/q", file)
		cmd.Run()
	}

	ddlFile := filepath.Join(tempDir, "DDL.txt")
	data, err = Asset("Remover/DDL.txt")
	if err != nil {
		fmt.Println("提取DDL文件时出错:", err)
		return
	}
	err = ioutil.WriteFile(ddlFile, data, 0644)
	if err != nil {
		fmt.Println("写入FDL文件时出错:", err)
		return
	}

	ddlTasks := readLines(ddlFile)
	for _, folder := range ddlTasks {
		cmd := exec.Command("cmd.exe", "/c", "rmdir", "/s", "/q", folder)
		cmd.Run()
	}

	fmt.Println("禁用命令执行完成，请重启系统！")
	pause()
}
