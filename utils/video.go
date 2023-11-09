package utils

import (
	"os"
	"os/exec"
)

func ExtractImageFromMP4(input []byte) ([]byte, error) {
	// 将输入字节写入临时文件
	tmpFile, err := os.CreateTemp("", "input.mp4")
	if err != nil {
		return nil, err
	}
	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	_, err = tmpFile.Write(input)
	if err != nil {
		return nil, err
	}

	// 使用 FFmpeg 提取视频帧
	outputFile := "E:\\ffmpeg\\output.jpg" // 更改为你想要保存图像的路径
	cmd := exec.Command("E:\\ffmpeg\\ffmpeg-2023-11-09-git-acf63d5350-essentials_build\\bin\\ffmpeg", "-i", tmpFile.Name(), "-ss", "0", "-vframes", "1", outputFile)
	err = cmd.Run()
	if err != nil {
		return nil, err
	}

	// 读取输出文件的内容
	output, err := os.ReadFile(outputFile)
	if err != nil {
		return nil, err
	}

	return output, nil
}
