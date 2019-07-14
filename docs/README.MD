- 基础
Go语言中，main函数结束，整个进程也随之结束。

Go语言中，所有的变量的传递都是值传递

- 导包
import{

}

从标准库中导入代码时，编译器在查找包的时候，总是回到GOROOT和GOPATH环境变量引用的位置去查找
比如：
   举个例子导入 net/http 包。如果 Go 安装在/usr/local/go，并且环境变量 GOPATH 设置为
   /home/myproject:/home/mylibraries，编译器就会按照下面的顺序查找 net/http 包：
   /usr/local/go/src/pkg/net/http 标准库所在位置
   /home/myproject/src/net/http
   /home/mylibraries/src/net/http
导入远程代码：
    github.com/spf13/viper
    用导入路径编译程序时， go build 命令会使用 GOPATH 的设置，在磁盘上搜索这个包。事实上，
    这个导入路径代表一个 URL，指向 GitHub 上的代码库。如果路径包含 URL，可以使用 Go 工具链从
    DVCS 获取包，并把包的源代码保存在 GOPATH 指向的路径里与 URL 匹配的目录里。这个获取过程
    使用 go get 命令完成。 go get 将获取任意指定的 URL 的包，或者一个已经导入的包所依赖的其
    他包。由于 go get 的这种递归特性，这个命令会扫描某个包的源码树，获取能找到的所有依赖包


import _ "package name" 会导致该包的init方法执行,init方法会在main方法之前执行

在Go语言中，所有的变量都被初始化为其对应的零值

":=" 表示创建变量并且赋值[省去了var 类型由编译器推断]

sync包中的WaitGroup可以跟踪启动的gorotine，WaitGroup是一个计数信号量

res,_ := test(),忽略test函数返回的第二个值

闭包访问外部函数里面的变量的时，是直接访问变量本身，并不是变量的副本

- 数组、切片、映射
函数间传递切片和映射类似于传递指针

- 引用类型
Go语言里的引用类型有如下几种：切片、映射、通道、接口、函数
当声明引用变量时，创建的变量被称作标头值，每个引用类型创建的标头值是包含一个指向底层数据结构的指针
每个引用类型还包含一组独特的字段，用于管理底层数据结构。因为标头值是为复制而设计的，所以永远不需要共享一个引用
类型的值。标头值里包含一个指针，因此通过复制来传递一个引用类型的值的副本，本质上就是在共享底层数据结构。


- 通道
// 通道可以共享内置类型、命名类型、结构类型和引用类型的值或者指针
// 无缓冲通道
var unbuffered = make(chan int)
// 有缓冲通道
var buffered = make(chan string, 10)

向通道发送数据 buffered <- "hello"
从通道读取数据 value := <- buffered

无缓冲的通道（unbuffered channel）是指在接收前没有能力保存任何值的通道。这种类型的通
道要求发送 goroutine 和接收 goroutine 同时准备好，才能完成发送和接收操作。如果两个 goroutine
没有同时准备好，通道会导致先执行发送或接收操作的 goroutine 阻塞等待。这种对通道进行发送
和接收的交互行为本身就是同步的。其中任意一个操作都无法离开另一个操作单独存在。

有缓冲的通道（buffered channel）是一种在被接收前能存储一个或者多个值的通道。这种类
型的通道并不强制要求 goroutine 之间必须同时完成发送和接收。通道会阻塞发送和接收动作的
条件也会不同。只有在通道中没有要接收的值时，接收动作才会阻塞。只有在通道没有可用缓冲
区容纳被发送的值时，发送动作才会阻塞。


- ServrMux 本质上是一个 HTTP 请求路由器（或者叫多路复用器，Multiplexor）。它把收到的请求与一组预先定义的 URL 路径列表做对比，然后在匹配到路径的时候调用关联的处理器（Handler）。

- 处理器（Handler）负责输出HTTP响应的头和正文。任何满足了http.Handler接口的对象都可作为一个处理器。通俗的说，对象只要有个如下签名的ServeHTTP方法即可：ServeHTTP(http.ResponseWriter, *http.Request)

Go 语言的 HTTP 包自带了几个函数用作常用处理器，比如FileServer，NotFoundHandler 和 RedirectHandler。

### http参数获取
    Form：存储了post、put和get参数，在使用之前需要调用ParseForm方法。

    PostForm：存储了post、put参数，在使用之前需要调用ParseForm方法。

    MultipartForm：存储了包含了文件上传的表单的post参数，在使用前需要调用ParseMultipartForm方法。

http服务器：
    服务器启动之后，监听指定端口，接受到客户端的请求之后，一般会启动一条线程来处理客户端的连接，
    所以一般处理用户请求是在一个线程中执行的，在自己的处理方法里面再启动线程，需要保证在服务器启动的线程执行结束之前就结束，
    否则，服务器启动的线程结束了[http 请求就处理完毕了]就没意义了。


 

 