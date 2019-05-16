# socketdemo
A socket server which achieve following function by native Go http package:

  1. A custom communication protocol between server and client;
  
  2. A custom connecting mechanism of heartbeating;
  
  3. Read operational parameters from config files;
  
  4. A router-controller structure to decouple codes of server.
  


# socketdemo
一个通过Go语言原生包实现了如下功能的socket server：

 1. 自定义通讯协议
 
 2. 通过心跳机制维护连接
 
 3. 从配置文件中读取系统参数
 
 4. 通过 router-controller 机制解耦服务器
 
 

# Example: 
```
//modify a controller in /utils/router.go
type EchoController struct  {

}
//create a new controller
func (this *EchoController) Excute(message Msg)[]byte {
	mirrormsg,err :=json.Marshal(message)
	Log("echo the message:", string(mirrormsg))
	CheckError(err)
	return mirrormsg
}

//register this controller 
func init() {
	var echo EchoController
	routers = make([][2]interface{} ,0 , 20)
	Route(func(entry Msg)bool{
		if entry.Meta["meta"]=="test"{
			return true}
		return  false
	},&echo)
}

//after setting parameters in config.yaml, run server.go
func main() {
	startServer("./conf/config.yaml")
}

```