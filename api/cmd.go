package api

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os/exec"
	"strings"
)

type cmd struct {
	CMD   string `json:"cmd"`
	Param string `json:"param"`
}

func CMD(c *gin.Context) {
	var command cmd
	c.ShouldBindJSON(&command)
	cmd := exec.Command(command.CMD, command.Param)
	stdout, err := cmd.StdoutPipe();
	if err != nil {     //获取输出对象，可以从该对象中读取输出结果
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		return
	}
	defer stdout.Close()   // 保证关闭输出流
	if err := cmd.Start(); err != nil {   // 运行命令
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		return
	}
	opBytes, err := ioutil.ReadAll(stdout);
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		return
	}
	msg := strings.Split(string(opBytes), string('\n'))
	msg = append(msg, "success")
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}