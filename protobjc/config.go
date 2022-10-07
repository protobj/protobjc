package protobjc

import "errors"

type LanguageType int32

const (
	JAVA LanguageType = 0
	TS   LanguageType = 1
)

func ToLanguageType(value string) (LanguageType, error) {
	if value == "Java" {
		return JAVA, nil
	} else if value == "Ts" {
		return TS, nil
	}
	return JAVA, errors.New("unknown language")
}

func (receiver LanguageType) FileSuffix() (string, error) {
	switch receiver {
	case JAVA:
		return "java", nil
	case TS:
		return "ts", nil
	}

	return "", errors.New("unknown language")
}

type OutputType int32

const (
	O_ALL     OutputType = 0
	O_SCHEMA  OutputType = 1
	O_MESSAGE OutputType = 2
)

func ToOutputType(value string) OutputType {
	if value == "message" {
		return O_MESSAGE
	} else if value == "schema" {
		return O_SCHEMA
	}
	return O_ALL
}

type ParsedArgs struct {
	SourceDir    string
	OutputDir    string
	LanguageType LanguageType
	OutputType   OutputType
}
