package main

import (
	"os"
	"reflect"
	"testing"
)

func TestLog(t *testing.T) {
	file, _ := os.Create("test.txt")
	defer os.Remove(file.Name())
	defer file.Close()

	fileLogger := FileLogger{file: file}
	logsystem := NewLogSystem(WithLogger(fileLogger))
	expectedLogsystem := &LogSystem{fileLogger: FileLogger{file: file}}
	if !reflect.DeepEqual(logsystem, expectedLogsystem) {
		t.Errorf("constructor don't work: get %+v, want %+v", logsystem, expectedLogsystem)
	}

	mess := "Hellow world!"
	fileLogger.Log(mess)
	buff, err := os.ReadFile(file.Name())
	if err != nil {
		t.Errorf("can not reading file")
	}
	if string(buff) != mess {
		t.Errorf("log don't work: get %v, want %v", string(buff), mess)
	}
	logsystem.Log(mess)
	buff, err = os.ReadFile(file.Name())
	if err != nil {
		t.Errorf("can not reading file")
	}
	if string(buff) != mess+mess {
		t.Errorf("log don't work: get %v, want %v", string(buff), mess)
	}
}
