package utils

import (
	"Innovation/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"os/exec"
	"strconv"
)

func ToUint(value any) uint {
	id, _ := value.(uint)
	return id
}

func GetContextData(c *gin.Context, s string) uint {
	if value, exists := c.Get(s); exists {
		id, _ := value.(uint)
		return id
	}
	c.AbortWithStatusJSON(500, model.NewResponse(0, "请传参:"+s, nil))
	return 0
}

func QueryAndParseUint(c *gin.Context, s string) uint {
	value := c.Query(s)
	atoi, err := strconv.Atoi(value)
	if err != nil {
		c.AbortWithStatusJSON(500, model.NewResponse(0, "传参不合法:"+s+"请传数字", nil))
		return 0
	}
	return uint(atoi)

}

// Ffmpeg 视频封面截取
func Ffmpeg(videoURL string, frameNum int) ([]byte, error) {
	// 创建一个临时文件来存储输出图像
	outputFile := "output.jpg"

	// 使用 ffmpeg 从视频中获取指定帧并将其输出到临时文件
	cmd := exec.Command("ffmpeg",
		"-i", videoURL,
		"-vf", fmt.Sprintf("select='gte(n,%d)'", frameNum),
		"-vframes", "1",
		outputFile)

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	// 读取临时文件的内容到缓冲区
	buf, err := os.ReadFile(outputFile)
	if err != nil {
		log.Fatal(err)
	}

	// 删除临时文件
	err = os.Remove(outputFile)
	if err != nil {
		log.Println("Error removing temporary file:", err)
	}

	return buf, nil
}
