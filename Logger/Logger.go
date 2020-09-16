package Logger

import (
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

var Logger = logrus.New()

func init() {
	log, _ := os.OpenFile(time.Now().Format("2006-01-02_log.txt"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
	Logger.Out = log
}
