// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag/v2"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "components": {"schemas":{"biz.Menu":{"properties":{"created_at":{"type":"string"},"hidden":{"description":"路由菜单是否隐藏","type":"boolean"},"icon":{"description":"路由菜单图标","type":"string"},"id":{"type":"integer"},"keep_alive":{"description":"路由菜单是否缓存","type":"boolean"},"page_file_path":{"description":"页面文件路径","type":"string"},"parent_id":{"description":"父级菜单、vue文件路径、路由路径、路由name、重定向路径 { meta: { title, icon, hidden, keepAlive,  } }","type":"integer"},"redirect":{"description":"重定向路径","type":"string"},"route_name":{"description":"路由name","type":"string"},"route_path":{"description":"路由路径","type":"string"},"sort":{"description":"路由菜单排序","type":"integer"},"title":{"description":"路由菜单标题","type":"string"},"updated_at":{"type":"string"}},"type":"object"},"biz.PaginationResponse":{"allOf":[{"$ref":"#/components/schemas/list"},{"$ref":"#/components/schemas/total"}],"properties":{"list":{"description":"列表"},"total":{"description":"总数","type":"integer"}},"type":"object"},"biz.Role":{"properties":{"created_at":{"type":"string"},"description":{"type":"string"},"id":{"type":"integer"},"menus":{"items":{"$ref":"#/components/schemas/biz.Menu"},"type":"array","uniqueItems":false},"name":{"type":"string"},"updated_at":{"type":"string"}},"type":"object"},"biz.User":{"description":"系统用户信息","properties":{"created_at":{"description":"创建时间","type":"string"},"email":{"description":"邮箱","type":"string"},"id":{"description":"用户id","type":"integer"},"phone":{"description":"手机号","type":"string"},"roles":{"description":"角色","items":{"$ref":"#/components/schemas/biz.Role"},"type":"array","uniqueItems":false},"status":{"description":"状态","type":"boolean"},"updated_at":{"description":"更新时间","type":"string"},"username":{"description":"用户名 唯一","type":"string"}},"type":"object"},"data":{"properties":{"data":{"allOf":[{"$ref":"#/components/schemas/list"},{"$ref":"#/components/schemas/total"}],"properties":{"list":{"description":"列表"},"total":{"description":"总数","type":"integer"}},"type":"object"}},"type":"object"},"list":{"properties":{"list":{"items":{"$ref":"#/components/schemas/biz.User"},"type":"array"}},"type":"object"},"response.Response":{"allOf":[{"$ref":"#/components/schemas/data"}],"description":"接口统一返回格式","properties":{"code":{"description":"状态码 (200-成功, 500-失败)","example":200,"type":"integer"},"data":{"description":"数据内容"},"message":{"description":"提示信息","example":"success","type":"string"}},"type":"object"},"services.LoginParams":{"properties":{"account":{"default":"zhangsan","description":"账号 邮箱/手机号/用户名","type":"string"},"password":{"default":"123456","description":"密码 6-16位","type":"string"}},"required":["account","password"],"type":"object"},"services.RegisterParams":{"properties":{"account":{"default":"zhangsan","description":"账号 邮箱/手机号/用户名","type":"string"},"password":{"default":"123456","description":"密码 6-16位","type":"string"}},"required":["account","password"],"type":"object"},"total":{"properties":{"total":{"type":"integer"}},"type":"object"}},"securitySchemes":{"":{"description":"通过在请求头中添加 Authorization 字段进行认证","in":"header","name":"Authorization","type":"apiKey"}}},
    "info": {"contact":{"email":"1227379879@qq.com","name":"North"},"description":"{{escape .Description}}","termsOfService":"http://swagger.io/terms/","title":"{{.Title}}","version":"{{.Version}}"},
    "externalDocs": {"description":"","url":""},
    "paths": {"/menu/create":{"post":{"description":"创建菜单","requestBody":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/biz.Menu"}}},"description":"菜单信息","required":true},"responses":{"200":{"content":{"application/json":{"schema":{"allOf":[{"$ref":"#/components/schemas/data"}],"description":"接口统一返回格式","properties":{"code":{"description":"状态码 (200-成功, 500-失败)","example":200,"type":"integer"},"data":{"description":"数据内容"},"message":{"description":"提示信息","example":"success","type":"string"}},"type":"object"}}},"description":"OK"},"500":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/response.Response"}}},"description":"Internal Server Error"}},"summary":"创建菜单","tags":["菜单模块"]}},"/menu/get-role":{"get":{"description":"获取角色菜单","requestBody":{"content":{"application/json":{"schema":{"type":"integer"}}},"description":"角色ID","required":true},"responses":{"200":{"content":{"application/json":{"schema":{"allOf":[{"$ref":"#/components/schemas/data"}],"description":"接口统一返回格式","properties":{"code":{"description":"状态码 (200-成功, 500-失败)","example":200,"type":"integer"},"data":{"description":"数据内容"},"message":{"description":"提示信息","example":"success","type":"string"}},"type":"object"}}},"description":"OK"},"500":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/response.Response"}}},"description":"Internal Server Error"}},"summary":"获取角色菜单","tags":["菜单模块"]}},"/menu/list":{"get":{"description":"获取菜单列表","requestBody":{"content":{"application/json":{"schema":{"type":"object"}}}},"responses":{"200":{"content":{"application/json":{"schema":{"allOf":[{"$ref":"#/components/schemas/data"}],"description":"接口统一返回格式","properties":{"code":{"description":"状态码 (200-成功, 500-失败)","example":200,"type":"integer"},"data":{"description":"数据内容"},"message":{"description":"提示信息","example":"success","type":"string"}},"type":"object"}}},"description":"OK"},"500":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/response.Response"}}},"description":"Internal Server Error"}},"summary":"获取菜单列表","tags":["菜单模块"]}},"/menu/set-role":{"post":{"description":"设置角色菜单","requestBody":{"content":{"application/json":{"schema":{"items":{"type":"integer"},"type":"array"}}},"description":"菜单ID列表","required":true},"responses":{"200":{"content":{"application/json":{"schema":{"allOf":[{"$ref":"#/components/schemas/data"}],"description":"接口统一返回格式","properties":{"code":{"description":"状态码 (200-成功, 500-失败)","example":200,"type":"integer"},"data":{"description":"数据内容"},"message":{"description":"提示信息","example":"success","type":"string"}},"type":"object"}}},"description":"OK"},"500":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/response.Response"}}},"description":"Internal Server Error"}},"summary":"设置角色菜单","tags":["菜单模块"]}},"/menu/update":{"put":{"description":"更新菜单","requestBody":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/biz.Menu"}}},"description":"菜单信息","required":true},"responses":{"200":{"content":{"application/json":{"schema":{"allOf":[{"$ref":"#/components/schemas/data"}],"description":"接口统一返回格式","properties":{"code":{"description":"状态码 (200-成功, 500-失败)","example":200,"type":"integer"},"data":{"description":"数据内容"},"message":{"description":"提示信息","example":"success","type":"string"}},"type":"object"}}},"description":"OK"},"500":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/response.Response"}}},"description":"Internal Server Error"}},"summary":"更新菜单","tags":["菜单模块"]}},"/role/create":{"post":{"description":"创建角色","requestBody":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/biz.Role"}}},"description":"角色信息","required":true},"responses":{"200":{"content":{"application/json":{"schema":{"allOf":[{"$ref":"#/components/schemas/data"}],"description":"接口统一返回格式","properties":{"code":{"description":"状态码 (200-成功, 500-失败)","example":200,"type":"integer"},"data":{"description":"数据内容"},"message":{"description":"提示信息","example":"success","type":"string"}},"type":"object"}}},"description":"OK"},"500":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/response.Response"}}},"description":"Internal Server Error"}},"summary":"创建角色","tags":["角色模块"]}},"/role/delete/{id}":{"delete":{"description":"删除角色","parameters":[{"description":"角色ID","in":"path","name":"id","required":true,"schema":{"type":"integer"}}],"requestBody":{"content":{"application/json":{"schema":{"type":"object"}}}},"responses":{"200":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/response.Response"}}},"description":"OK"},"500":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/response.Response"}}},"description":"Internal Server Error"}},"summary":"删除角色","tags":["角色模块"]}},"/role/list":{"get":{"description":"获取角色列表","parameters":[{"description":"页码","in":"query","name":"page","required":true,"schema":{"type":"integer"}},{"description":"每页数量","in":"query","name":"pageSize","required":true,"schema":{"type":"integer"}}],"requestBody":{"content":{"application/json":{"schema":{"type":"object"}}}},"responses":{"200":{"content":{"application/json":{"schema":{"allOf":[{"$ref":"#/components/schemas/data"}],"description":"接口统一返回格式","properties":{"code":{"description":"状态码 (200-成功, 500-失败)","example":200,"type":"integer"},"data":{"description":"数据内容"},"message":{"description":"提示信息","example":"success","type":"string"}},"type":"object"}}},"description":"OK"},"500":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/response.Response"}}},"description":"Internal Server Error"}},"summary":"获取角色列表","tags":["角色模块"]}},"/user/info":{"get":{"description":"获取用户信息","requestBody":{"content":{"application/json":{"schema":{"type":"object"}}}},"responses":{"200":{"content":{"application/json":{"schema":{"allOf":[{"$ref":"#/components/schemas/data"}],"description":"接口统一返回格式","properties":{"code":{"description":"状态码 (200-成功, 500-失败)","example":200,"type":"integer"},"data":{"description":"数据内容"},"message":{"description":"提示信息","example":"success","type":"string"}},"type":"object"}}},"description":"OK"},"500":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/response.Response"}}},"description":"Internal Server Error"}},"security":[{"ApiKeyAuth":[]}],"summary":"获取用户信息","tags":["用户模块"]}},"/user/list":{"get":{"description":"分页获取用户列表","parameters":[{"description":"每页数量","in":"query","name":"pageSize","required":true,"schema":{"default":10,"description":"每页数量","form":"pageSize","type":"integer"}},{"description":"页码","in":"query","name":"page","required":true,"schema":{"default":1,"description":"页码","form":"page","type":"integer"}}],"requestBody":{"content":{"application/json":{"schema":{"type":"object"}}}},"responses":{"200":{"content":{"application/json":{"schema":{"allOf":[{"$ref":"#/components/schemas/data"}],"description":"接口统一返回格式","properties":{"code":{"description":"状态码 (200-成功, 500-失败)","example":200,"type":"integer"},"data":{"description":"数据内容"},"message":{"description":"提示信息","example":"success","type":"string"}},"type":"object"}}},"description":"OK"},"500":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/response.Response"}}},"description":"Internal Server Error"}},"summary":"分页获取用户列表","tags":["用户模块"]}},"/user/login":{"post":{"description":"用户登录，返回token","requestBody":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/services.LoginParams"}}},"description":"用户信息","required":true},"responses":{"200":{"content":{"application/json":{"schema":{"allOf":[{"$ref":"#/components/schemas/data"}],"description":"接口统一返回格式","properties":{"code":{"description":"状态码 (200-成功, 500-失败)","example":200,"type":"integer"},"data":{"description":"数据内容"},"message":{"description":"提示信息","example":"success","type":"string"}},"type":"object"}}},"description":"OK"},"500":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/response.Response"}}},"description":"Internal Server Error"}},"summary":"用户登录","tags":["用户模块"]}},"/user/register":{"post":{"description":"注册","requestBody":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/services.RegisterParams"}}},"description":"用户信息","required":true},"responses":{"200":{"content":{"application/json":{"schema":{"allOf":[{"$ref":"#/components/schemas/data"}],"description":"接口统一返回格式","properties":{"code":{"description":"状态码 (200-成功, 500-失败)","example":200,"type":"integer"},"data":{"description":"数据内容"},"message":{"description":"提示信息","example":"success","type":"string"}},"type":"object"}}},"description":"OK"},"500":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/response.Response"}}},"description":"Internal Server Error"}},"summary":"注册","tags":["用户模块"]}}},
    "openapi": "3.1.0"
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Title:            "权限管理系统",
	Description:      "权限管理系统api、测试",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
