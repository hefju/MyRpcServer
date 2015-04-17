package main
import (
//    "github.com/gorilla/rpc"
//    "github.com/gorilla/rpc/json"
//    "net/http"
//    "fmt"

    "log"
    "net"
    "net/rpc"
    "net/rpc/jsonrpc"

    "github.com/hefju/MyRpcServer/setting"
    "github.com/hefju/MyRpcServer/library"
    "fmt"
)

type Args struct{
    Who string
}
type Reply struct {
    Message string
}
type Hello struct {
}
func (h *Hello)Say(args *Args, reply *Reply)error{
    reply.Message="hello,"+args.Who+"!"
    return  nil
}

func main(){

   w:= new(library.Watcher)
  var  result int
    w.GetInfo(500,&result)
   fmt.Println("result",result)

   st:= new(setting.App)
    st.  Ip="192.168"
    st. Port="8000"
st.SaveToFile()

    hello:=new(Hello)
    s:=rpc.NewServer()
    s.Register(hello)
    s.HandleHTTP(rpc.DefaultRPCPath,rpc.DefaultDebugPath)
    l,e:=net.Listen("tcp",":8081")
    if e!=nil{
        log.Fatal("listen:",e)
    }
    for{
        if conn,err:=l.Accept();err!=nil{
            log.Fatal("accept:",err.Error())
        }else{
            log.Printf("new connect")
            go s.ServeCodec(jsonrpc.NewServerCodec(conn))
        }
    }



//    s:=rpc.NewServer()
//    s.RegisterCodec(json.NewCodec(),"application/json")
//    s.RegisterService(new(Hello),"")
//    http.Handle("/rpc",s)
//    http.ListenAndServe(":8081",nil)
//    fmt.Println("end")
}
