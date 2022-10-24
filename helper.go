package logwrapper

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/gommon/log"
)

func utcTimeFunc() time.Time {
	return time.Now().UTC()
}

func openFile(name string, path string) (*os.File, error) {
	file, err := openFileFunc(fileName(name, path), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}

	_, err = file.Stat()
	if err != nil {
		return nil, err
	}
	return file, err
}

func fileName(filename string, filepath string) string {
	if len(filepath) > 0 {
		return fmt.Sprint(filepath, "/", filename)
	}
	return filename
}

func callerMarshalFunc(pc uintptr, file string, line int) string {
	dirs := strings.Split(file, "/")

	if len(dirs) > 4 {
		file = strings.Join(dirs[len(dirs)-4:], "/")
	}
	return file + ":" + strconv.Itoa(line)
}

func jsonMarshal(j log.JSON) string {
	b, err := json.Marshal(j)
	if err != nil {
		return fmt.Sprintf("logwrapper, Marshal error : %s", err.Error())
	}
	return string(b)
}
