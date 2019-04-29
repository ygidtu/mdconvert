### md -> html -> pdf的转化功能


golang的确实由于wkhtmktopdf的C库比较难打包，所以暂时就这么着吧，反正手已经练了

随手写一个Python拿来顶着用了

---
依赖于该项目[github.com/adrg/go-wkhtmltopdf](https://github.com/adrg/go-wkhtmltopdf)，而非另外衣蛾go-wkhtmltopdf。

优点：该项目直接依赖于libwkhtmltox，不依赖于command line的wkhtmltopdf。
缺点：c库难配置，跨平台编译复杂。需要人为的将各种依赖库设定好

目前跨平台编译没有好办法，就是虚拟机