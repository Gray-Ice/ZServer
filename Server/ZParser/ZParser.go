package ZParser

import (
	bytes2 "bytes"
	"errors"
	"strings"
)

type ZParser struct {
	Protocol string
	Args     map[string]string
}

func NewZParser() *ZParser {
	return &ZParser{}
}

var endFlag = []byte("\n\n\n\r")
var protocolSeparatorFlag = []byte("?")

/*
 * 解析请求头。请求头之后就全是二进制文件了。
 */
func (p *ZParser) ExtractRequestHeader(bytes []byte) (string, map[string]string, error) {
	endPos := bytes2.Index(bytes, endFlag)
	if endPos <= 0 {
		return "", nil, errors.New("the length of header less than or equal zero")
	}

	requestInfo := bytes[:endPos]
	if len(requestInfo) > 253 {
		return "", nil, errors.New("request header too big. It should be less than 256 bytes")
	}

	ptcolSeparatorPos := bytes2.Index(requestInfo, protocolSeparatorFlag)
	if ptcolSeparatorPos <= 0 {
		return "", nil, errors.New("the length of protocol less than or equal zero")
	}

	protocol := string(requestInfo[:ptcolSeparatorPos])
	args := make(map[string]string)

	// 没有参数，返回
	if ptcolSeparatorPos-endPos == 1 {
		return protocol, args, nil
	}

	// 解析参数，例如: arg1=123&arg2=34122
	body := string(requestInfo[ptcolSeparatorPos+len(protocolSeparatorFlag):])
	pairs := strings.Split(body, "&")
	for _, p := range pairs {
		kv := strings.Split(p, "=")
		if len(kv) != 2 {
			continue
		}

		args[kv[0]] = kv[1]
	}
	return protocol, args, nil
}
