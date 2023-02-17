package utils

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"

	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/yaml.v3"
)

// 优雅关闭
func WaitForGracefullyStop(sig chan os.Signal, funcs ...func()) {
	quit := sig
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	for _, fn := range funcs {
		fn()
	}
}

func LoadJsonConfigFileWithDefaultPath(v any) error {
	bts, err := os.ReadFile("./config.json")
	if err != nil {
		return err
	}
	return json.Unmarshal(bts, v)
}

func LoadYamlConfigFileWithDefaultPath(v any) error {
	bts, err := os.ReadFile("./config.yaml")
	if err != nil {
		return err
	}
	return yaml.Unmarshal(bts, v)
}

// MD5加密
func Hash(str string) string {
	m := md5.New()
	io.WriteString(m, str)
	return fmt.Sprintf("%x", m.Sum(nil))
}

func GetUUID() string {
	return uuid.NewV4().String()
}

func NewStatusError(code int, err any) error {
	return status.Error(codes.Code(code), fmt.Sprintf("%v", err))
}
