- ServrMux 本质上是一个 HTTP 请求路由器（或者叫多路复用器，Multiplexor）。它把收到的请求与一组预先定义的 URL 路径列表做对比，然后在匹配到路径的时候调用关联的处理器（Handler）。

- 处理器（Handler）负责输出HTTP响应的头和正文。任何满足了http.Handler接口的对象都可作为一个处理器。通俗的说，对象只要有个如下签名的ServeHTTP方法即可：ServeHTTP(http.ResponseWriter, *http.Request)

Go 语言的 HTTP 包自带了几个函数用作常用处理器，比如FileServer，NotFoundHandler 和 RedirectHandler。

### http参数获取
    Form：存储了post、put和get参数，在使用之前需要调用ParseForm方法。

    PostForm：存储了post、put参数，在使用之前需要调用ParseForm方法。

    MultipartForm：存储了包含了文件上传的表单的post参数，在使用前需要调用ParseMultipartForm方法。

 

 