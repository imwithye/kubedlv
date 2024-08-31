package main

import (
	"fmt"
	"os"
	"os/exec"
)

func GetCmdAbsPath(cmd string) (string, error) {
	workingDir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	originalPath := os.Getenv("PATH")
	newPath := fmt.Sprintf("%s%c%s", workingDir, os.PathListSeparator, originalPath)
	err = os.Setenv("PATH", newPath)
	if err != nil {
		return "", err
	}
	cmdAbsPath, err := exec.LookPath(cmd)
	if err != nil {
		return "", err
	}
	return cmdAbsPath, nil
}

func GetDlvArgs(cmd string, cmdArgs []string, config *DlvConfig) ([]string, error) {
	absCmd, err := GetCmdAbsPath(cmd)
	if err != nil {
		return nil, err
	}

	if config == nil {
		config = NewDlvConfig()
	}

	dlvArgs := []string{
		fmt.Sprintf("--api-version=%d", config.APIVersion),
		fmt.Sprintf("--listen=:%d", config.Port),
		"--headless=true",
	}
	if config.Continue {
		//goland:noinspection ALL
		dlvArgs = append(dlvArgs, "--continue", "--accept-multiclient")
	}
	if config.Log {
		dlvArgs = append(dlvArgs, "--log")
	}
	dlvArgs = append(dlvArgs, "exec", absCmd)
	if len(cmdArgs) > 0 {
		dlvArgs = append(dlvArgs, "--")
		dlvArgs = append(dlvArgs, cmdArgs...)
	}
	return dlvArgs, nil
}

func RunDlv(dlvArgs []string) error {
	dlvCmd := exec.Command("dlv", dlvArgs...)
	dlvCmd.Stdout = os.Stdout
	dlvCmd.Stderr = os.Stderr
	dlvCmd.Stdin = os.Stdin
	return dlvCmd.Run()
}
