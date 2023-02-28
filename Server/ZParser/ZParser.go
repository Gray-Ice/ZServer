package ZParser

import (
	bytes2 "bytes"
	"errors"
	"strings"
)

var endFlag = []byte("\n\n\n\r")
var protocolSeparatorFlag = []byte("?")

// Parser 网络解析器应该遵循该Parser接口，参数是用户第一次发来的网络请求，返回的string是提取出的协议名，map是参数，error是是否有错误
type Parser interface {
	ExtractRequestHeader([]byte) (string, map[string]string, error)
}

// ZParser 请求解析器
type ZParser struct {
	Protocol       string
	Args           map[string]string
	BinaryStartPos int
}

func NewZParser() *ZParser {
	return &ZParser{}
}

// ExtractRequestHeader 用于从请求头提取信息
// 返回的第一个参数是协议名，第二个是从网络请求中提取的参数，第三个是error
func (p *ZParser) ExtractRequestHeader(bytes []byte) error {
	endPos := bytes2.Index(bytes, endFlag)
	if endPos <= 0 {
		return errors.New("the length of header less than or equal zero")
	}

	requestInfo := bytes[:endPos]
	if len(requestInfo) > 256 {
		return errors.New("request header too big. It should be less than 256 bytes")
	}

	ptcolSeparatorPos := bytes2.Index(requestInfo, protocolSeparatorFlag)
	if ptcolSeparatorPos <= 0 {
		return errors.New("the length of protocol less than or equal zero")
	}

	protocol := string(requestInfo[:ptcolSeparatorPos])
	args := make(map[string]string)

	// 没有参数，返回
	if ptcolSeparatorPos-endPos == 1 {
		p.Protocol = protocol
		p.BinaryStartPos = endPos + len(endFlag)
		return nil
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
	p.Protocol = protocol
	p.Args = args
	return nil
}
