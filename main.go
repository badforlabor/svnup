/**
 * Auth :   liubo
 * Date :   2020/9/23 14:07
 * Comment:
 */

package main

import (
	"fmt"
	"os"
	"os/exec"
)


func main() {

	doAction(100)

}

func doAction(maxCount int) {
	if maxCount <= 0 {
		fmt.Println("结束，失败了...")
		return
	}

	var e = runCmd("svn", "up")
	if e != nil {
		e = runCmd("svn", "cleanup", "--include-externals")
		if e == nil {
			doAction(maxCount - 1)
		} else {
			fmt.Println("结束，失败了...")
		}
	} else {
		fmt.Println("结束，成功了...")
	}
}

func runCmd(exeName string, args ...string) error {
	fmt.Println("runCmd:", exeName, args)

	var folder = "./" // *pGitFolder
	var cmd = exec.Command(exeName, args...)
	cmd.Dir = folder
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	var err = cmd.Run()

	if err != nil {
		fmt.Println("错误:" + err.Error())
	}

	return err
}