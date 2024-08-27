package logger

import (
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"log/slog"
	"os"
	"server/pkg/trace"
)

type Config struct {
	FilePath         string `yaml:"file_path" env:"FILE_PATH" default:"logs/logs.json"`
	UseLocalTime     bool   `yaml:"use_local_time" env:"USE_LOCAL_TIME" default:"false"`
	FileMaxSizeInMB  int    `yaml:"file_max_size_in_mb" env:"FILE_MAX_SIZE_IN_MB" default:"10"`
	FileMaxAgeInDays int    `yaml:"file_max_age_in_days" env:"FILE_MAX_AGE_IN_DAYS"  default:"30"`
}

var l *slog.Logger
var dbL *slog.Logger

func NewLogger(config Config) {
	fileWriter := &lumberjack.Logger{
		Filename:  config.FilePath,
		LocalTime: config.UseLocalTime,
		MaxSize:   config.FileMaxSizeInMB,
		MaxAge:    config.FileMaxAgeInDays,
	}
	l = slog.New(
		slog.NewJSONHandler(io.MultiWriter(fileWriter, os.Stdout), &slog.HandlerOptions{}),
	)
}

func L() *slog.Logger {
	t := trace.Parse()

	return l.With(slog.String("path", t.File), slog.Int("line", t.Line))
}

//
//// database log writer
//
//type logRepo interface {
//	Insert(log *entity.Log) error
//}
//
//type DatabaseWriter struct {
//	data []byte
//
//	repo logRepo
//}

//
//func NewDatabaseWriter(repo logRepo) DatabaseWriter {
//	return DatabaseWriter{repo: repo}
//}
//
//func NewDBLogger(writer DatabaseWriter) {
//
//	dbL = slog.New(
//		slog.NewJSONHandler(io.MultiWriter(writer, os.Stdout), nil),
//	)
//}
//
//func DBL() *slog.Logger {
//	t := trace.Parse()
//
//	return dbL.With(slog.String("path", t.File), slog.Int("line", t.Line))
//}
//
//func (dw DatabaseWriter) Write(data []byte) (int, error) {
//
//	lo := new(entity.Log)
//
//	log.Println(string(data))
//
//	if err := json.Unmarshal(data, lo); err != nil {
//		return 0, err
//	}
//
//	if err := dw.repo.Insert(lo); err != nil {
//		return 0, err
//	}
//
//	return len(data), nil
//}
