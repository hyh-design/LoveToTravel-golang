package utils

import (
	"crypto/md5"
	"encoding/hex"
)

/**
 * 参考文章: https://www.cnblogs.com/dadishi/p/17064495.html
 */

func GetMD5(s string) string {
	// 创建MD5算法
	has := md5.New()
	has.Write([]byte(s)) // 写入需要加密的数据
	// 获取hash值字符切片；
	// Sum函数接受一个字符切片，这个切片的内容会原样的追加到abc123加密后的hash值的前面，这里我们不需要这么做，所以传入nil
	b := has.Sum(nil)
	// 通过hex包的EncodeToString函数将数据转为16进制字符串
	return hex.EncodeToString(b)
}
