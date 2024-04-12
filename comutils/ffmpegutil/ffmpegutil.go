package ffmpegutil

import "os/exec"

// GeneratorVideoThumbnail 生成视频缩略图
// sourceFilePath: 视频源文件路径
// thumbnailSavePath: 缩略图保存路径
// time: 生成缩略图的时间点
//
// 例如：GeneratorVideoThumbnail("test.mp4", "test.jpg", "00:00:01")
func GeneratorVideoThumbnail(sourceFilePath, thumbnailSavePath string, time string) error {
	args := []string{
		"-i", sourceFilePath,
		"-ss", time,
		"-vframes", "1",
		"-c:v", "mjpeg",
		"-pix_fmt", "yuvj420p",
		"-color_range", "tv",
		"-y", thumbnailSavePath}
	cmd := exec.Command("ffmpeg", args...)
	cmd.Stderr = cmd.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
