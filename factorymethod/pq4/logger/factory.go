package logger

import "fmt"

func NewLogger(vertical string) (Logger, error) {
	switch vertical {
	case "db":
		return &db{}, nil
	case "file":
		return &file{}, nil
	case "console":
		return &console{}, nil
	default:
		return nil, fmt.Errorf("unsupported logger type %s", vertical)
	}
}
