package component

import (
	"bytes"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"math/rand"
	"sync"
	"testing"
)

func TestViperReader(t *testing.T) {
	viper.SetConfigType("yaml")
	var yamlExample = []byte(`
Hacker: true
name: steve
hobbies:
- skateboarding
- snowboarding
- go
clothing:
  jacket: leather
  trousers: denim
age: 35
eyes : brown
beard: true
`)
	_ = viper.ReadConfig(bytes.NewBuffer(yamlExample))
	fmt.Println(viper.GetString("name"))
}

type Config struct {
	Name         string
	ClothingTest struct {
		Jacket string
	} `mapstructure:"clothing"`
}

func TestViperFile(t *testing.T) {
	// 会自动以这个顺序去尝试
	// "json", "toml", "yaml", "yml", "properties", "props", "prop", "hcl", "dotenv", "env", "ini"
	//viper.SetConfigType("yaml")
	viper.SetConfigName("test")
	viper.AddConfigPath("./")
	_ = viper.ReadInConfig()
	fmt.Println(viper.GetString("name"))
	fmt.Println(viper.GetString("clothing.jacket"))
	fmt.Println(viper.GetIntSlice("hobbies")) // 空Slice
	fmt.Println(viper.GetStringSlice("hobbies"))
	sub := viper.Sub("clothing")
	fmt.Println(sub.Get("jacket"))
	var config Config
	if e := viper.Unmarshal(&config); e != nil {
		panic(e)
	}
	fmt.Println(config.Name)
	fmt.Println(config.ClothingTest.Jacket)
}

func TestViperWriteFile(t *testing.T) {
	viper.SetConfigType("yaml")
	viper.SetConfigName("test")
	viper.AddConfigPath("./")
	_ = viper.ReadInConfig()
	viper.Set("name", "test")
	// 修改的值会在尾部写入
	if e := viper.WriteConfig(); e != nil {
		panic(e)
	}
	fmt.Println(viper.GetString("name"))
}

func TestViperWatchFile(t *testing.T) {
	viper.SetConfigType("yaml")
	viper.SetConfigName("test")
	viper.AddConfigPath("./")
	_ = viper.ReadInConfig()
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name, e.Op.String())
	})
	fmt.Println(viper.GetString("name"))
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		viper.Set("name", fmt.Sprintf("test_%d", rand.Intn(255)))
		if e := viper.WriteConfig(); e != nil {
			panic(e)
		}
		wg.Done()
	}()
	wg.Wait()
	fmt.Println(viper.GetString("name"))
}
