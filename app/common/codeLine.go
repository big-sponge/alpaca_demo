package common

import (
	"bufio"
	"os"
)

/**
 * 计算代码行数
 * @author ChengCheng
 * @date 2019-05-14 22:10:18
 */
func CurCodeLine() (count int) {
	count, _ = CodeLine(BasePath+"/app")
	return count
}

/**
 * 计算代码行数
 * @author ChengCheng
 * @date 2019-05-14 22:10:18
 */
func CodeLine(baseDir string) (count int, err error) {
	/* 检查是不是目录 */
	checkDir, err := os.Lstat(baseDir)
	if err != nil || !checkDir.IsDir() {
		return count, err
	}

	/* 打开目录 */
	dir, err := os.Open(baseDir)
	if err != nil {
		return count, err
	}

	/* 遍历目录 */
	subItems, err := dir.Readdir(0)
	if err != nil {
		return count, err
	}
	for _, sub := range subItems {
		/* 如果是目录，递归计算 */
		if sub.IsDir() {
			sunCount, err := CodeLine(baseDir + "/" + sub.Name())
			if err != nil {
				return count, err
			}
			count = count + sunCount
			continue
		}
		/* 如果是文件，计算文件内容行数 */
		subFile, err := os.Open(baseDir + "/" + sub.Name())
		if err != nil {
			return count, err
		}
		fileScanner := bufio.NewScanner(subFile)
		lineCount := 0
		for fileScanner.Scan() {
			lineCount++
		}
		count = count + lineCount
	}

	/* 返回结果 */
	return count, err
}
