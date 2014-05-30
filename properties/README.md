properties
==========

Golang读取properties文件，根据key返回对应的value。  

Features
==========

 - 非常方便的文件加载方式，只要传入文件名即可
 - 提供Get方法，快速的value获取

Document
=========

func Load(file string) (*properties, error)   
type properties   
　　func (p *properties) Get(key string) string  
具体使用方式，请参照 [properties_test.go][1]

Install
==========

    go get github.com/admin100/util/properties

LICENSE
==========

properties is licensed under the Apache Licence, Version 2.0 (http://www.apache.org/licenses/LICENSE-2.0.html).


  [1]: https://github.com/admin100/util/blob/master/properties/properties_test.go