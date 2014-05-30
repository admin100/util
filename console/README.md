console
==========

Golang对console的一些操作。  

Features
==========

 - 通过SetConsoleTitle可以设置windows下console的标题

Document
=========

func SetConsoleTitle(title string) (int, error)  
具体使用方式，请参照 [console_test.go][1]

Install
==========

    go get github.com/admin100/util/console

LICENSE
==========

properties is licensed under the Apache Licence, Version 2.0 (http://www.apache.org/licenses/LICENSE-2.0.html).


  [1]: https://github.com/admin100/util/blob/master/console/console_test.go