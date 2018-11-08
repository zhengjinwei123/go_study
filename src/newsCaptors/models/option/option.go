package option

import (
	"strconv"
	"newsCaptors/models"
	"newsCaptors/util"
)

func GetOptions() map[string]string {
	rs, _ := util.Factory.Get("cache")
	cache := rs.(*util.LRUCache)

	if !cache.IsExist("options") {
		var result []*models.Option
		new(models.Option).FindAll(&result)
		options := make(map[string]string)
		for _,v := range result {
			options[v.Name] = v.Value
		}
		cache.Put("options",options,0)
	}

	v := cache.Get("options")
	return v.(map[string]string)
}

func FlushOptions(){
	rs ,_ := util.Factory.Get("cache")
	cache := rs.(*util.LRUCache)
	cache.Delete("options")
}

func Get(key string) string {
	options := GetOptions()
	if v,ok := options[key];ok {
		return v
	}
	return ""
}

func GetInt(key string) int {
	v,_ := strconv.Atoi(Get(key))
	return v
}
