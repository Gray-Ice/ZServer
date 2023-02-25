package ZFlag

import (
	"fmt"
	"strings"
)

const (
	ClipboardFlag = "clipboard"
	FileTransFlag = "filetrans"
	ErrorFlag     = "error"
)

func GetFlag(bytes []byte) (err string) {
	err = ErrorFlag
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	message := string(bytes)
	if strings.HasSuffix(message, ClipboardFlag) {
		return ClipboardFlag
	}
	if strings.HasSuffix(message, FileTransFlag) {
		return FileTransFlag
	}
	return err
}
