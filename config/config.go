package config

import (
	"github.com/fsnotify/fsnotify"
	config "github.com/spf13/viper"
	log "github.com/sirupsen/logrus"
	"path"
	"runtime"
)

func SetConfig() {
	config.SetConfigName("App")
	config.SetConfigType("yaml")
	//config.AddConfigPath(p)
	config.AddConfigPath("./env")
	config.AddConfigPath(GetDefaultConfigPath())

	err := config.ReadInConfig()
	if err != nil {
		log.Fatal("config error: ", err)
	}

	config.WatchConfig()
	config.OnConfigChange(func(e fsnotify.Event) {
		log.Warn("Config file changed:", e.Name)
	})

}



// GetDefaultConfigPath location
func GetDefaultConfigPath() string {
	_, filename, _, ok := runtime.Caller(0)
	if ok == false {
		log.Fatal("err")
	}
	filePath := path.Join(path.Dir(filename), "../env/")

	return filePath
}
