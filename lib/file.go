/**
* @Author:Tristan
* @Date: 2021/12/23 9:46 上午
 */

package lib

import "os"

// 判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// 文件夹不存在，创建
func checkAndCreatePath(dirPath string) error {
	exist, err := PathExists(dirPath)
	if err != nil {
		return err
	}

	if !exist {
		// 创建文件夹
		err = os.Mkdir(dirPath, os.ModePerm)
	}
	return nil
}

