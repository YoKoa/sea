package log

import (
	"context"
	"io"
	"path/filepath"
	"time"

	"mongodb-service/common/log/filewriter"
)

// level idx
const (
	_infoIdx = iota
	_warnIdx
	_errorIdx
	_totalIdx
)

var _fileNames = map[int]string{
	_infoIdx:  "info.log",
	_warnIdx:  "warning.log",
	_errorIdx: "error.log",
}

// FileHandler .
type FileHandler struct {
	render Render
	fws    [_totalIdx]*filewriter.FileWriter
}

// NewFile crete a file logger.
func NewFile(dir string, bufferSize, rotateSize int64, maxLogFile int) *FileHandler {
	// new info writer
	newWriter := func(name string) *filewriter.FileWriter {
		var options []filewriter.Option
		if rotateSize > 0 {
			options = append(options, filewriter.MaxSize(rotateSize))
		}
		if maxLogFile > 0 {
			options = append(options, filewriter.MaxFile(maxLogFile))
		}
		w, err := filewriter.New(filepath.Join(dir, name), options...)
		if err != nil {
			panic(err)
		}
		return w
	}
	handler := &FileHandler{
		render: newPatternRender("[%D %T] [%L] [%S] %M"),
	}
	for idx, name := range _fileNames {
		handler.fws[idx] = newWriter(name)
	}
	return handler
}

// Log loggint to file .
func (h *FileHandler) Log(ctx context.Context, lv Level, args ...D) {
	d := toMap(args...)
	d[_time] = time.Now().Format(_timeFormat)

	addExtraField(ctx, d)

	var w io.Writer
	switch lv {
	case _warnLevel:
		w = h.fws[_warnIdx]
	case _errorLevel:
		w = h.fws[_errorIdx]
	default:
		w = h.fws[_infoIdx]
	}
	h.render.Render(w, d)
	w.Write([]byte("\n"))
}

// Close log middleware
func (h *FileHandler) Close() error {
	for _, fw := range h.fws {
		// ignored error
		fw.Close()
	}
	return nil
}

// SetFormat set log format
func (h *FileHandler) SetFormat(format string) {
	h.render = newPatternRender(format)
}

func addExtraField(ctx context.Context, fields map[string]interface{}) {

	if traceId := String(ctx, "traceId");  traceId != "" {
		fields[_tid] = traceId
	}
	if userId := String(ctx, "user_id"); userId != "" {
		fields[_userId] = userId
	}
	if customerId := String(ctx, "customer_id"); customerId != "" {
		fields[_customerId] = customerId
	}
}

func String(ctx context.Context, key string) string {
	md, ok := ctx.Value(key).(string)
	if !ok {
		return ""
	}
	//str, _ := md[key].(string)
	return md
}

type MD map[string]interface{}

type mdKey struct{}

