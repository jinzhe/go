package zip

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

// 压缩文件
func In(dir, filename string) error {
	buf := bytes.NewBuffer(make([]byte, 0, 10*1024*1024)) // 创建一个读写缓冲
	myzip := zip.NewWriter(buf)                           // 用压缩器包装该缓冲
	// 用Walk方法来将所有目录下的文件写入zip
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		var file []byte
		if err != nil {
			return filepath.SkipDir
		}
		header, err := zip.FileInfoHeader(info) // 转换为zip格式的文件信息
		if err != nil {
			return filepath.SkipDir
		}
		header.Name, _ = filepath.Rel(filepath.Dir(dir), path)
		if !info.IsDir() {
			// 确定采用的压缩算法（这个是内建注册的deflate）
			header.Method = 8
			file, err = ioutil.ReadFile(path) // 获取文件内容
			if err != nil {
				return filepath.SkipDir
			}
		} else {
			file = nil
		}
		// 上面的部分如果出错都返回filepath.SkipDir
		// 下面的部分如果出错都直接返回该错误
		// 目的是尽可能的压缩目录下的文件，同时保证zip文件格式正确
		w, err := myzip.CreateHeader(header) // 创建一条记录并写入文件信息
		if err != nil {
			return err
		}
		_, err = w.Write(file) // 非目录文件会写入数据，目录不会写入数据
		if err != nil {        // 因为目录的内容可能会修改
			return err // 最关键的是我不知道咋获得目录文件的内容
		}
		return nil
	})
	if err != nil {
		return err
	}
	myzip.Close()                    // 关闭压缩器，让压缩器缓冲中的数据写入buf
	file, err := os.Create(filename) // 建立zip文件
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = buf.WriteTo(file) // 将buf中的数据写入文件
	if err != nil {
		return err
	}
	return nil
}

// 解压文件
func Out(filename string) error {
	r, e := zip.OpenReader(filename)
	if e != nil {
		return e
	}
	defer r.Close()
	for _, f := range r.File {
		fmt.Println(f.Name)
		rc, e := f.Open()
		if e != nil {
			return e
		}
		defer rc.Close()

		dw, e := os.Create(f.Name)
		if e != nil {
			return e
		}
		defer dw.Close()

		if _, e = io.Copy(dw, rc); e != nil {
			return e
		}
	}
	return nil
}
