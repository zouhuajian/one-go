package serialize

import (
	"bytes"
	"encoding/gob"
	"github.com/sirupsen/logrus"
)

type Serialize struct {
	Network bytes.Buffer
	Val     interface{}
}

// Encode
// 序列化 注意修改原始数据类型
func (s *Serialize) Encode(val interface{}) error {
	// 1.创建编码器
	enc := gob.NewEncoder(&s.Network)
	// 2.向编码器中写入数据
	err := enc.Encode(val)
	if err != nil {
		logrus.Panicf("failed to encode, err: %#v", err)
	}
	return nil
}

// Decode
//反序列化
func (s *Serialize) Decode(val []byte) error {
	decoder := gob.NewDecoder(bytes.NewReader(val)) //创建解密器
	err := decoder.Decode(&s.Val)                   //解密
	if err != nil {
		logrus.Panicf("failed to decode, err: %#v", err)
	}
	return nil
}
