package setting
import (
    "encoding/json"
    "log"
)
type App struct {
    Ip string
    Port string
}

func (app App)SaveToFile(){
   data,err:=  json.Marshal(app)
    if err!=nil{
        log.Fatal("setting.App.SaveToFile.json.Marshal:",err)
    }
    log.Println("App to json: ",string(data))
}

