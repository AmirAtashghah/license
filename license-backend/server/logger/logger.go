package logger

import (
	"encoding/json"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"log/slog"
	"os"
	"server/entity"
	"server/pkg/trace"
)

// TODO: should we transfer these constants to the default config struct also? or not?
const (
	defaultFilePath        = "logs/logs.json"
	defaultUseLocalTime    = false
	defaultFileMaxSizeInMB = 10
	defaultFileAgeInDays   = 30
)

type Config struct {
	FilePath         string
	UseLocalTime     bool
	FileMaxSizeInMB  int
	FileMaxAgeInDays int
}

var l *slog.Logger

func init() {
	fileWriter := &lumberjack.Logger{
		Filename:  defaultFilePath,
		LocalTime: defaultUseLocalTime,
		MaxSize:   defaultFileMaxSizeInMB,
		MaxAge:    defaultFileAgeInDays,
	}
	l = slog.New(
		slog.NewJSONHandler(io.MultiWriter(fileWriter, os.Stdout), &slog.HandlerOptions{}),
	)
}

func L() *slog.Logger {
	return l
}

func WithGroup(groupName string) *slog.Logger {
	t := trace.Parse()

	return l.With(slog.String("group", groupName),
		slog.String("path", t.File),
		slog.Int("line", t.Line),
		slog.String("function", t.Function),
	)
}

func SaveClientLogsOnDatabase() *slog.Logger {

	databaseWriter := DatabaseWriter{}

	logger := slog.New(
		slog.NewJSONHandler(io.MultiWriter(databaseWriter, os.Stdout), nil),
	)

	return logger
}

type logRepo interface {
	Insert(logEntry entity.Log) error
}

type DatabaseWriter struct {
	data []byte

	repo logRepo
}

func New(repo logRepo) *DatabaseWriter {
	return &DatabaseWriter{repo: repo}
}

func (dw DatabaseWriter) Write(data []byte) (int, error) {

	lo := new(entity.Log)

	if err := json.Unmarshal(data, lo); err != nil {
		return 0, err
	}

	if err := dw.repo.Insert(*lo); err != nil {
		return 0, err
	}

	return len(data), nil
}
