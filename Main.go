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
