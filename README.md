### md -> html -> pdf的转化功能


golang的确实由于wkhtmktopdf的C库比较难打包，所以暂时就这么着吧，反正手已经练了

随手写一个Python拿来顶着用了



Python版本在服务器上的部署

Python的pdfkit是wkhtmlpdf的wrapper。

因此：
- 需要安装wkhtmltopdf，
- wkhtmltopdf依赖X server，在服务器上可参考：https://github.com/JazzCore/python-pdfkit/wiki/Using-wkhtmltopdf-without-X-server。建立虚拟x server

```bash
apt-get install wkhtmltopdf
apt-get install xvfb
printf '#!/bin/bash\nxvfb-run -a --server-args="-screen 0, 1024x768x24" /usr/bin/wkhtmltopdf -q $*' > /usr/bin/wkhtmltopdf.sh
chmod a+x /usr/bin/wkhtmltopdf.sh
ln -s /usr/bin/wkhtmltopdf.sh /usr/local/bin/wkhtmltopdf
wkhtmltopdf http://www.google.com output.pdf
```


---
依赖于该项目[github.com/adrg/go-wkhtmltopdf](https://github.com/adrg/go-wkhtmltopdf)，而非另外衣蛾go-wkhtmltopdf。

优点：该项目直接依赖于libwkhtmltox，不依赖于command line的wkhtmltopdf。
缺点：c库难配置，跨平台编译复杂。需要人为的将各种依赖库设定好

目前跨平台编译没有好办法，就是虚拟机


---

Kotlin版本

依赖于Java的诸多项目，独立于wkhtmlpdf，理论上不会像上边这些一样有依赖问题

但是目前尚未完工

-[x] markdown -> html
-[ ] html -> pdf
    - 诸多依赖库通过SAXParser处理html问题，SAXParser过于严格
    - 中文字符问题 
