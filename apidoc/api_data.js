define({ "api": [
  {
    "type": "delete",
    "url": "/auth/route/:id",
    "title": "路由-5、删除",
    "description": "<p>删除路由</p>",
    "group": "AUTH",
    "version": "0.1.0",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "id",
            "description": "<p>路由id</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "int",
            "optional": false,
            "field": "code",
            "description": "<p>状态值</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "msg",
            "description": "<p>状态描述</p>"
          },
          {
            "group": "Success 200",
            "type": "array",
            "optional": false,
            "field": "data",
            "description": "<p>返回数据</p>"
          }
        ]
      }
    },
    "filename": "api/auth/route.go",
    "groupTitle": "权限相关接口",
    "name": "DeleteAuthRouteId",
    "sampleRequest": [
      {
        "url": "http://127.0.0.1:8000/auth/route/:id"
      }
    ]
  },
  {
    "type": "get",
    "url": "/auth/action/info/:id",
    "title": "模块方法-2.详情",
    "description": "<p>根据方法id获取方法详情</p>",
    "version": "0.1.0",
    "group": "AUTH",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "id",
            "description": "<p>方法id</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "int",
            "optional": false,
            "field": "code",
            "description": "<p>状态值</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "msg",
            "description": "<p>状态描述</p>"
          }
        ]
      }
    },
    "filename": "api/auth/action.go",
    "groupTitle": "权限相关接口",
    "name": "GetAuthActionInfoId",
    "sampleRequest": [
      {
        "url": "http://127.0.0.1:8000/auth/action/info/:id"
      }
    ]
  },
  {
    "type": "get",
    "url": "/auth/action/model/:model_id",
    "title": "模块方法-1.列表",
    "description": "<p>根据模块id获取该模块下的所有方法列表</p>",
    "version": "0.1.0",
    "group": "AUTH",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "model_id",
            "description": "<p>模块id</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "int",
            "optional": false,
            "field": "code",
            "description": "<p>状态值</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "msg",
            "description": "<p>状态描述</p>"
          }
        ]
      }
    },
    "filename": "api/auth/action.go",
    "groupTitle": "权限相关接口",
    "name": "GetAuthActionModelModel_id",
    "sampleRequest": [
      {
        "url": "http://127.0.0.1:8000/auth/action/model/:model_id"
      }
    ]
  },
  {
    "type": "get",
    "url": "/auth/model",
    "title": "模块-1列表",
    "description": "<p>获取模块列表</p>",
    "version": "0.1.0",
    "group": "AUTH",
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "int",
            "optional": false,
            "field": "code",
            "description": "<p>状态值</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "msg",
            "description": "<p>状态描述</p>"
          }
        ]
      }
    },
    "filename": "api/auth/model.go",
    "groupTitle": "权限相关接口",
    "name": "GetAuthModel",
    "sampleRequest": [
      {
        "url": "http://127.0.0.1:8000/auth/model"
      }
    ]
  },
  {
    "type": "get",
    "url": "/auth/model/:id",
    "title": "模块-2详情",
    "description": "<p>获取模块详情</p>",
    "version": "0.1.0",
    "group": "AUTH",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "id",
            "description": "<p>模块id</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "int",
            "optional": false,
            "field": "code",
            "description": "<p>状态值</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "msg",
            "description": "<p>状态描述</p>"
          }
        ]
      }
    },
    "filename": "api/auth/model.go",
    "groupTitle": "权限相关接口",
    "name": "GetAuthModelId",
    "sampleRequest": [
      {
        "url": "http://127.0.0.1:8000/auth/model/:id"
      }
    ]
  },
  {
    "type": "get",
    "url": "/auth/route",
    "title": "路由-1、列表",
    "description": "<p>获取路由列表</p>",
    "version": "0.1.0",
    "group": "AUTH",
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "int",
            "optional": false,
            "field": "code",
            "description": "<p>状态值</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "msg",
            "description": "<p>状态描述</p>"
          },
          {
            "group": "Success 200",
            "type": "array",
            "optional": false,
            "field": "data",
            "description": "<p>返回数据</p>"
          }
        ]
      }
    },
    "filename": "api/auth/route.go",
    "groupTitle": "权限相关接口",
    "name": "GetAuthRoute",
    "sampleRequest": [
      {
        "url": "http://127.0.0.1:8000/auth/route"
      }
    ]
  },
  {
    "type": "get",
    "url": "/auth/route/:id",
    "title": "路由-2、详情",
    "description": "<p>查询路由详情</p>",
    "group": "AUTH",
    "version": "0.1.0",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "id",
            "description": "<p>路由id</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "int",
            "optional": false,
            "field": "code",
            "description": "<p>状态值</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "msg",
            "description": "<p>状态描述</p>"
          },
          {
            "group": "Success 200",
            "type": "array",
            "optional": false,
            "field": "data",
            "description": "<p>返回数据</p>"
          }
        ]
      }
    },
    "filename": "api/auth/route.go",
    "groupTitle": "权限相关接口",
    "name": "GetAuthRouteId",
    "sampleRequest": [
      {
        "url": "http://127.0.0.1:8000/auth/route/:id"
      }
    ]
  },
  {
    "type": "post",
    "url": "/auth/action/add",
    "title": "模块方法-3.添加",
    "description": "<p>添加权限方法</p>",
    "version": "0.1.0",
    "group": "AUTH",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "model_id",
            "description": "<p>模块id</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "name",
            "description": "<p>方法名称</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": true,
            "field": "description",
            "description": "<p>方法描述</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": true,
            "field": "status",
            "description": "<p>方法状态</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "int",
            "optional": false,
            "field": "code",
            "description": "<p>状态值</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "msg",
            "description": "<p>状态描述</p>"
          }
        ]
      }
    },
    "filename": "api/auth/action.go",
    "groupTitle": "权限相关接口",
    "name": "PostAuthActionAdd",
    "sampleRequest": [
      {
        "url": "http://127.0.0.1:8000/auth/action/add"
      }
    ]
  },
  {
    "type": "post",
    "url": "/auth/action/del",
    "title": "模块方法-5.删除",
    "description": "<p>删除权限方法</p>",
    "version": "0.1.0",
    "group": "AUTH",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "id",
            "description": "<p>方法id</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "int",
            "optional": false,
            "field": "code",
            "description": "<p>状态值</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "msg",
            "description": "<p>状态描述</p>"
          }
        ]
      }
    },
    "filename": "api/auth/action.go",
    "groupTitle": "权限相关接口",
    "name": "PostAuthActionDel",
    "sampleRequest": [
      {
        "url": "http://127.0.0.1:8000/auth/action/del"
      }
    ]
  },
  {
    "type": "post",
    "url": "/auth/action/edit",
    "title": "模块方法-4.修改",
    "description": "<p>修改权限方法</p>",
    "version": "0.1.0",
    "group": "AUTH",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "id",
            "description": "<p>方法id</p>"
          },
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "model_id",
            "description": "<p>模块id</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "name",
            "description": "<p>方法名称</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": true,
            "field": "description",
            "description": "<p>方法描述</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": true,
            "field": "status",
            "description": "<p>方法状态</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "int",
            "optional": false,
            "field": "code",
            "description": "<p>状态值</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "msg",
            "description": "<p>状态描述</p>"
          }
        ]
      }
    },
    "filename": "api/auth/action.go",
    "groupTitle": "权限相关接口",
    "name": "PostAuthActionEdit",
    "sampleRequest": [
      {
        "url": "http://127.0.0.1:8000/auth/action/edit"
      }
    ]
  },
  {
    "type": "post",
    "url": "/auth/assignment",
    "title": "路由分配",
    "description": "<p>分配权限</p>",
    "version": "0.0.1",
    "group": "AUTH",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "user_id",
            "description": "<p>用户id</p>"
          },
          {
            "group": "Parameter",
            "type": "Array",
            "optional": false,
            "field": "route_ids",
            "description": "<p>路由id集合</p>"
          }
        ]
      }
    },
    "filename": "api/auth/assignment.go",
    "groupTitle": "权限相关接口",
    "name": "PostAuthAssignment",
    "sampleRequest": [
      {
        "url": "http://127.0.0.1:8000/auth/assignment"
      }
    ]
  },
  {
    "type": "post",
    "url": "/auth/model/add",
    "title": "模块-3添加",
    "description": "<p>新加权限模块</p>",
    "version": "0.1.0",
    "group": "AUTH",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "name",
            "description": "<p>模块名称</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": true,
            "field": "description",
            "description": "<p>描述</p>"
          },
          {
            "group": "Parameter",
            "type": "int",
            "optional": true,
            "field": "parent_id",
            "description": "<p>父节点id</p>"
          },
          {
            "group": "Parameter",
            "type": "int",
            "optional": true,
            "field": "status",
            "description": "<p>模块状态</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "int",
            "optional": false,
            "field": "code",
            "description": "<p>状态值</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "msg",
            "description": "<p>状态描述</p>"
          }
        ]
      }
    },
    "filename": "api/auth/model.go",
    "groupTitle": "权限相关接口",
    "name": "PostAuthModelAdd",
    "sampleRequest": [
      {
        "url": "http://127.0.0.1:8000/auth/model/add"
      }
    ]
  },
  {
    "type": "post",
    "url": "/auth/model/del",
    "title": "模块-5删除",
    "description": "<p>删除权限模块</p>",
    "version": "0.1.0",
    "group": "AUTH",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "id",
            "description": "<p>模块id</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "int",
            "optional": false,
            "field": "code",
            "description": "<p>状态值</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "msg",
            "description": "<p>状态描述</p>"
          }
        ]
      }
    },
    "filename": "api/auth/model.go",
    "groupTitle": "权限相关接口",
    "name": "PostAuthModelDel",
    "sampleRequest": [
      {
        "url": "http://127.0.0.1:8000/auth/model/del"
      }
    ]
  },
  {
    "type": "post",
    "url": "/auth/model/edit",
    "title": "模块-4修改",
    "description": "<p>修改权限模块</p>",
    "version": "0.1.0",
    "group": "AUTH",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "id",
            "description": "<p>模块id</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "name",
            "description": "<p>模块名称</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": true,
            "field": "description",
            "description": "<p>描述</p>"
          },
          {
            "group": "Parameter",
            "type": "int",
            "optional": true,
            "field": "parent_id",
            "description": "<p>父节点id</p>"
          },
          {
            "group": "Parameter",
            "type": "int",
            "optional": true,
            "field": "status",
            "description": "<p>模块状态</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "int",
            "optional": false,
            "field": "code",
            "description": "<p>状态值</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "msg",
            "description": "<p>状态描述</p>"
          }
        ]
      }
    },
    "filename": "api/auth/model.go",
    "groupTitle": "权限相关接口",
    "name": "PostAuthModelEdit",
    "sampleRequest": [
      {
        "url": "http://127.0.0.1:8000/auth/model/edit"
      }
    ]
  },
  {
    "type": "post",
    "url": "/auth/route",
    "title": "路由-3、新建",
    "description": "<p>新建路由</p>",
    "version": "0.1.0",
    "group": "AUTH",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "name",
            "description": "<p>路由名称</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": true,
            "field": "desc",
            "description": "<p>路由描述</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": true,
            "field": "route",
            "description": "<p>路由地址</p>"
          },
          {
            "group": "Parameter",
            "type": "int",
            "optional": true,
            "field": "parent_id",
            "description": "<p>父节点id</p>"
          },
          {
            "group": "Parameter",
            "type": "int",
            "optional": true,
            "field": "status",
            "description": "<p>路由状态</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "int",
            "optional": false,
            "field": "code",
            "description": "<p>状态值</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "msg",
            "description": "<p>状态描述</p>"
          },
          {
            "group": "Success 200",
            "type": "array",
            "optional": false,
            "field": "data",
            "description": "<p>返回数据</p>"
          }
        ]
      }
    },
    "filename": "api/auth/route.go",
    "groupTitle": "权限相关接口",
    "name": "PostAuthRoute",
    "sampleRequest": [
      {
        "url": "http://127.0.0.1:8000/auth/route"
      }
    ]
  },
  {
    "type": "post",
    "url": "/login",
    "title": "登录",
    "description": "<p>登录接口</p>",
    "version": "0.1.0",
    "group": "AUTH",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "username",
            "description": "<p>账号</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "password",
            "description": "<p>密码</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "uuid",
            "description": "<p>客户端唯一标识</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "int",
            "optional": false,
            "field": "code",
            "description": "<p>状态值</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "msg",
            "description": "<p>状态描述</p>"
          }
        ]
      }
    },
    "filename": "api/auth/user.go",
    "groupTitle": "权限相关接口",
    "name": "PostLogin",
    "sampleRequest": [
      {
        "url": "http://127.0.0.1:8000/login"
      }
    ]
  },
  {
    "type": "post",
    "url": "/login.uuid",
    "title": "uuid登录",
    "description": "<p>uuid登录接口</p>",
    "version": "0.1.0",
    "group": "AUTH",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "uuid",
            "description": "<p>客户端唯一标识</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "int",
            "optional": false,
            "field": "code",
            "description": "<p>状态值</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "msg",
            "description": "<p>状态描述</p>"
          }
        ]
      }
    },
    "filename": "api/auth/user.go",
    "groupTitle": "权限相关接口",
    "name": "PostLoginUuid",
    "sampleRequest": [
      {
        "url": "http://127.0.0.1:8000/login.uuid"
      }
    ]
  },
  {
    "type": "post",
    "url": "/register",
    "title": "注册",
    "description": "<p>注册接口</p>",
    "version": "0.1.0",
    "group": "AUTH",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "username",
            "description": "<p>账号</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "password",
            "description": "<p>密码</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "email",
            "description": "<p>邮箱</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "uuid",
            "description": "<p>客户端唯一标识</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "int",
            "optional": false,
            "field": "code",
            "description": "<p>状态值</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "msg",
            "description": "<p>状态描述</p>"
          }
        ]
      }
    },
    "filename": "api/auth/user.go",
    "groupTitle": "权限相关接口",
    "name": "PostRegister",
    "sampleRequest": [
      {
        "url": "http://127.0.0.1:8000/register"
      }
    ]
  },
  {
    "type": "put",
    "url": "/auth/route/:id",
    "title": "路由-4、修改",
    "description": "<p>修改路由信息</p>",
    "version": "0.1.0",
    "group": "AUTH",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "id",
            "description": "<p>路由id</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "name",
            "description": "<p>路由名称</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": true,
            "field": "desc",
            "description": "<p>路由描述</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": true,
            "field": "route",
            "description": "<p>路由地址</p>"
          },
          {
            "group": "Parameter",
            "type": "int",
            "optional": true,
            "field": "parent_id",
            "description": "<p>父节点id</p>"
          },
          {
            "group": "Parameter",
            "type": "int",
            "optional": true,
            "field": "status",
            "description": "<p>路由状态</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "int",
            "optional": false,
            "field": "code",
            "description": "<p>状态值</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "msg",
            "description": "<p>状态描述</p>"
          },
          {
            "group": "Success 200",
            "type": "array",
            "optional": false,
            "field": "data",
            "description": "<p>返回数据</p>"
          }
        ]
      }
    },
    "filename": "api/auth/route.go",
    "groupTitle": "权限相关接口",
    "name": "PutAuthRouteId",
    "sampleRequest": [
      {
        "url": "http://127.0.0.1:8000/auth/route/:id"
      }
    ]
  }
] });
