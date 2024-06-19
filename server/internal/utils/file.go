package utils

import (
	"context"
	"github.com/gogf/gf/v2/encoding/gurl"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"net/http"
	"os"
	"path"
	"server/internal/consts"
	"strings"
)

// WriteToFile 写入文件
func WriteToFile(fileName string, content string) error {
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	n, _ := f.Seek(0, os.SEEK_END)
	_, err = f.WriteAt([]byte(content), n)
	defer f.Close()
	return err
}

// FileIsExisted 文件或文件夹是否存在
func FileIsExisted(filename string) bool {
	existed := true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		existed = false
	}
	return existed
}

// ParseFilePath 解析路径获取文件名称及后缀
func ParseFilePath(pathStr string) (fileName string, fileType string) {
	fileNameWithSuffix := path.Base(pathStr)
	fileType = path.Ext(fileNameWithSuffix)
	fileName = strings.TrimSuffix(fileNameWithSuffix, fileType)
	return
}

// IsNotExistMkDir 检查文件夹是否存在
// 如果不存在则新建文件夹
func IsNotExistMkDir(src string) error {
	if exist := !FileIsExisted(src); exist == false {
		if err := MkDir(src); err != nil {
			return err
		}
	}

	return nil
}

// MkDir 新建文件夹
func MkDir(src string) error {
	err := os.MkdirAll(src, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

// GetExt 获取文件后缀
func GetExt(fileName string) string {
	return path.Ext(fileName)
}

// GetType 获取文件类型
func GetType(p string) (result string, err error) {
	file, err := os.Open(p)
	if err != nil {
		g.Log().Error(context.TODO(), err)
		return
	}
	buff := make([]byte, 512)

	_, err = file.Read(buff)

	if err != nil {
		g.Log().Error(context.TODO(), err)
		return
	}
	filetype := http.DetectContentType(buff)
	return filetype, nil
}

// GetFilesPath 获取附件相对路径
func GetFilesPath(ctx context.Context, fileUrl string) (path string, err error) {
	upType := g.Cfg().MustGet(ctx, "upload.default").Int()
	if upType != 0 || (upType == 0 && !gstr.ContainsI(fileUrl, consts.UploadPath)) {
		path = fileUrl
		return
	}
	pathInfo, err := gurl.ParseURL(fileUrl, 32)
	if err != nil {
		g.Log().Error(ctx, err)
		err = gerror.New("解析附件路径失败")
		return
	}
	pos := gstr.PosI(pathInfo["path"], consts.UploadPath)
	if pos >= 0 {
		path = gstr.SubStr(pathInfo["path"], pos)
	}
	return
}
