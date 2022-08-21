package u3

import (
	"encoding/json"
	"goapi/utils"
	"goapi/utils/u2"
	"io/ioutil"
	"log"
	"path"
	"regexp"
	"sync"
)

var goModName string
var apiConfig *utils.ApiConfig
var apiName string
var apiStructName string

var once = sync.Once{}

func prepareData() {
	once.Do(func() {
		// parse config_model.json
		b, err := ioutil.ReadFile("config_model.json")
		if err != nil {
			log.Println(err)
		}

		api := utils.ApiConfig{}
		err = json.Unmarshal(b, &api)
		if err != nil {
			log.Println(err)
		}

		apiConfig = &api
		apiName = api.ModelName
		apiStructName = utils.ToGoUpper(apiName)
		goModName = parseGoMod()

		// mkdir
		err = u2.Madir("", path.Join(".", "api", api.ModelName))
		if err != nil {
			log.Println(err)
		}
	})
}

func parseGoMod() string {

	b, _ := ioutil.ReadFile("go.mod")
	reg, _ := regexp.Compile(`module\s+(.*)`)

	res := reg.FindAllStringSubmatch(string(b), -1)

	tmp := res[0]

	return tmp[len(tmp)-1]
}
