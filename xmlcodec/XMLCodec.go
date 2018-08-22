/*
XMLCodec

websocket通信时候使用的Message和JSON对象都是Codec接口的实例。Codec接口定义如下

type Codec struct {
	Marshal func(v interface{})(data []byte, payloadType byte, err error)
	Unmarshal func(data []byte, payloadType byte, v interface{})(err error)
}

我们也可以自己实现一个Codec的接口，来完成特定的Marshal和Unmarshal格式
本代码用于实现XML Codec
*/
package xmlcodec

import (
	"encoding/xml"
	"golang.org/x/net/websocket"
)

func xmlMarshal(v interface{}) (msg []byte, payloadType byte, err error) {
	msg, err = xml.Marshal(v)
	return msg, websocket.TextFrame, nil
}

func xmlUnmarshal(msg []byte, payloadType byte, v interface{}) (err error) {
	err = xml.Unmarshal(msg, v)
	return err
}

var XMLCodec = websocket.Codec{xmlMarshal, xmlUnmarshal}
