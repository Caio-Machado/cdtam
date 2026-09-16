package output

import (
	"errors"
	"fmt"
	"image"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type FileFormat string

const (
	PNG FileFormat = "png"
)

func parseFileFormat(ext string) (FileFormat, error) {
	ext = strings.ToLower(strings.TrimPrefix(ext, "."))

	switch FileFormat(ext) {
	case PNG:
		return PNG, nil
	default:
		return "", fmt.Errorf("unsupported file extension: %s", ext)
	}
}

type FileOutput struct {
	FileName   string
	Dir        string
	FileFormat FileFormat
}

func (fo FileOutput) Write(img image.Image) error

func (fo FileOutput) FullPath() string {
	return filepath.Join(fo.Dir, fo.FileName)
}

func NewFileOutput(filePath string) (FileOutput, error) {
	cleanPath := filepath.Clean(filePath)
	info, err := os.Stat(cleanPath)
	exists := err == nil
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return FileOutput{}, err
	}

	fileName := filepath.Base(cleanPath)
	dir := filepath.Dir(cleanPath)
	if exists && info.IsDir() {
		dir = cleanPath
		fileName = fmt.Sprintf("screenshot_%d.png", time.Now().UnixNano())
	}

	_, err = os.Stat(dir)
	if errors.Is(err, os.ErrNotExist) {
		return FileOutput{}, fmt.Errorf("the directory %s does not exist %w", dir, err)
	}
	if err != nil {
		return FileOutput{}, err
	}

	ext := filepath.Ext(fileName)
	fileFormat, err := parseFileFormat(ext)
	if err != nil {
		return FileOutput{}, err
	}

	return FileOutput{
		FileName:   fileName,
		Dir:        dir,
		FileFormat: fileFormat,
	}, nil
}
