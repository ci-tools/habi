package file

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

type Spec struct {
	Src     string      `yaml:"src"`
	SrcType string      `yaml:"src-type"`
	Dst     string      `yaml:"dst"`
	Sha256  string      `yaml:"sha256"`
	Owner   string      `yaml:"owner"`
	Mode    os.FileMode `yaml:"mode"`
	Content string      `yaml:"content"`
}

type RawSpec struct {
	Src     string `yaml:"src"`
	SrcType string `yaml:"src-type"`
	Dst     string `yaml:"dst"`
	Sha256  string `yaml:"sha256"`
	Mode    string `yaml:"mode"`
	Owner   string `yaml:"owner"`
	Content string `yaml:"content"`
}

func Hash(filename string) (hash string, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return hash, err
	}
	defer file.Close()

	sha256hash := sha256.New()
	if _, err := io.Copy(sha256hash, file); err != nil {
		return hash, err
	}
	hash = fmt.Sprintf("%x", sha256hash.Sum(nil))

	return hash, nil
}

func ModeConv(modestr string) (mode os.FileMode, err error) {
	modeint, err := strconv.ParseInt(modestr, 8, 64)
	if err != nil {
		return mode, err
	}
	mode = os.FileMode(modeint)
	return mode, nil
}

type FileState struct {
	Exists  bool
	Name    string
	Size    int64
	Mode    os.FileMode
	ModTime time.Time
	IsDir   bool
	Sha256  string
}

func Meta(filename string) (FileState, error) {
	state := FileState{}
	info, err := os.Stat(filename)
	if !os.IsNotExist(err) {
		state.Name = info.Name()
		state.Size = info.Size()
		state.Mode = info.Mode()
		state.IsDir = info.IsDir()
		state.Exists = true
		state.ModTime = info.ModTime()
		sha256str, err := Hash(filename)
		if err != nil {
			return state, err
		}
		state.Sha256 = sha256str
	}
	return state, nil
}
