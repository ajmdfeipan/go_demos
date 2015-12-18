package main
import ("fmt"
	"log"
	"encoding/json"
	"sync"
	"github.com/toolkits/file"

)

func main() {
	ParseConfig("/Users/jp/PycharmProjects/go_demos/jsonconfig/cfg.json")//路径

	fmt.Println(Config().Hostname)
	fmt.Println(Config().Plugin.Enabled)
	fmt.Println(Config().IgnoreMetrics["cpu.busy"])
}

type PluginConfig struct {
	Enabled bool   `json:"enabled"`
}


type GlobalConfig struct {
	Debug         bool             `json:"debug"`
	Hostname      string           `json:"hostname"`
	IP            string           `json:"ip"`
	Plugin        *PluginConfig    `json:"plugin"`
	IgnoreMetrics map[string]bool  `json:"ignore"`
}

var (
	ConfigFile string
	config     *GlobalConfig
	lock       = new(sync.RWMutex)
)

func Config() *GlobalConfig {
	lock.RLock()
	defer lock.RUnlock()
	return config
}

func ParseConfig(cfg string) {
	if cfg == "" {
		log.Fatalln("use -c to specify configuration file")
	}

	if !file.IsExist(cfg) {
		log.Fatalln("config file:", cfg, "is not existent. maybe you need `mv cfg.json cfg.json`")
	}

	ConfigFile = cfg

	configContent, err := file.ToTrimString(cfg)
	if err != nil {
		log.Fatalln("read config file:", cfg, "fail:", err)
	}

	var c GlobalConfig
	err = json.Unmarshal([]byte(configContent), &c)
	if err != nil {
		log.Fatalln("parse config file:", cfg, "fail:", err)
	}

	lock.Lock()
	defer lock.Unlock()

	config = &c

	log.Println("read config file:", cfg, "successfully")
}
