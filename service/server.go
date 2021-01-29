package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday/v2"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

type HeavenDoor interface {
}

type HeavenDoorService struct {
}

func NewHeavenDoorService() *HeavenDoorService {
	return &HeavenDoorService{}
}

func (srv *HeavenDoorService) InitRoot() {
	// Do some cache
}

func (srv *HeavenDoorService) ListDir(ctx *gin.Context) {
	p := ctx.Query("path")
	paths := path.Dir(p)
	filepathNames, err := filepath.Glob(filepath.Join(p, "*"))
	if err != nil {
		panic(err)
	}
	for i := range filepathNames {
		fmt.Println(filepathNames[i])
	}
	ctx.JSON(http.StatusOK, gin.H{"paths": paths})
}

func (srv *HeavenDoorService) Content(ctx *gin.Context) {
	pwd := ctx.Query("path")
	info, err := os.Stat(pwd)
	if err != nil {
		panic(err)
	}
	info.Name()

	b, err := ioutil.ReadFile(pwd)
	if err != nil {
		panic(err)
	}

	output := blackfriday.Run(b)
	fmt.Println(string(output))
	ctx.Data(http.StatusOK, "text/html;charset=utf-8", output)
	//ctx.JSON(http.StatusOK, gin.H{})
}
