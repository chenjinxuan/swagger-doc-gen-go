# API文档自动生成-命令行工具

> 基于beego框架自动化文档，实现从代码注释生成swagger文档。

## 安装

1. `cd $GOPATH/src/`
2. 拉取代码 `git clone git@gitlab.airdroid.com:treasure/swagger-doc-gen.git`
3. `cd swagger-doc-gen`
4. `go install`

## 使用swagger-doc-gen命令

* 第一次生成文档 `swagger-doc-gen --downdoc true --router_path ./router.go --output_path ./`
* 更新文档 `swagger-doc-gen --router_path ./router.go --output_path ./`

###### 参数介绍

* **--downdoc** - 是否下载swagger-ui文件，第一次使用需要下载，更新文档会自动更新swagger.json和swagger.yml文件。默认 `false`
* **--router_path** - 路由文件路径，默认 `./main.go`
* **--output_path** - swagger文档生成目录，默认 `./`

## 代码注释

#### API 全局设置

必须设置在 `--router_path` 路由文件的注释，最顶部
```javascript
// @APIVersion 2.0
// @Title My Test API
// @Description API项目描述
// @TermsOfServiceUrl http://beego.me/
// @Contact astaxie@gmail.com
// @Name http://beego.me/
// @URL http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
// @Schemes http,https
// @Host my.apidomain.com
```

#### router 路由分组设置

必须设置在 `--router_path` 路由文件的注释中
```javascript
// @SubApi 分组别名 [/server/]
```

#### controller 控制器方法设置

设置在controller package中，工具会自动扫描在router文件中引入的所有非系统包中的文件
```javascript
// @Title getStaticBlock
// @Description get all the staticblock by key
// @Summary 接口摘要
// @Deprecated false
// @Accept json
// @Param   key     path    string  true        "The email for login"
// @Param   category_id     query   int false       "category id"
// @Param   brand_id    query   int false       "brand id"
// @Param   query   query   string  false       "query of search"
// @Param   segment query   string  false       "segment"
// @Param   sort    query   string  false       "sort option"
// @Param   dir     query   string  false       "direction asc or desc"
// @Param   offset  query   int     false       "offset"
// @Param   limit   query   int     false       "count limit"
// @Param   price           query   float       false       "price"
// @Param   special_price   query   bool        false       "whether this is special price"
// @Param   size            query   string      false       "size filter"
// @Param   color           query   string      false       "color filter"
// @Param   format          query   bool        false       "choose return format"
// @Success 200 {object} models.ZDTCustomer.Customer
// @Failure 400 Invalid email supplied
// @Failure 404 User not found
// @router /staticblock/:key [get]
```

* @Title
这个 API 所表达的含义，是一个文本，空格之后的内容全部解析为 title
* @Description
这个 API 详细的描述，是一个文本，空格之后的内容全部解析为 Description
* @Param
参数，表示需要传递到服务器端的参数，有五列参数，使用空格或者 tab 分割，五个分别表示的含义如下

		1.参数名
		2.参数类型，可以有的值是 formData、query、path、body、header，
		  formData 表示是 post 请求的数据，
		  query 表示带在 url 之后的参数，
		  path 表示请求路径上得参数，例如上面例子里面的 key，
		  body 表示是一个 raw 数据请求，
		  header 表示带在 header 信息中得参数。
		3.参数类型
		4.是否必须
		5.注释
* @Success
成功返回给客户端的信息，三个参数，第一个是 status code。第二个参数是返回的类型，必须使用 {} 包含，第三个是返回的对象或者字符串信息，如果是 {object} 类型，那么 bee 工具在生成 docs 的时候会扫描对应的对象，这里填写的是想对你项目的目录名和对象，例如 models.ZDTProduct.ProductList 就表示 /models/ZDTProduct 目录下的 ProductList 对象。三个参数必须通过空格分隔
* @Failure
失败返回的信息，包含两个参数，使用空格分隔，第一个表示 status code，第二个表示错误信息
* @router
路由信息，包含两个参数，使用空格分隔，第一个是请求的路由地址，支持正则和自定义路由，和之前的路由规则一样，第二个参数是支持的请求方法,放在 [] 之中，如果有多个方法，那么使用 , 分隔。
