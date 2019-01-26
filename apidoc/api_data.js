define({ "api": [
  {
    "type": "post",
    "url": "/auth/action/add",
    "title": "添加方法",
    "description": "<p>添加权限方法</p>",
    "version": "0.1.0",
    "group": "AUTH",
    "parameter": {
      "fields": {
        "请求参数：": [
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
    "title": "删除方法",
    "description": "<p>删除权限方法</p>",
    "version": "0.1.0",
    "group": "AUTH",
    "parameter": {
      "fields": {
        "请求参数：": [
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
    "title": "修改方法",
    "description": "<p>修改权限方法</p>",
    "version": "0.1.0",
    "group": "AUTH",
    "parameter": {
      "fields": {
        "请求参数：": [
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
    "url": "/auth/model/add",
    "title": "添加模块",
    "description": "<p>新加权限模块</p>",
    "version": "0.1.0",
    "group": "AUTH",
    "parameter": {
      "fields": {
        "请求参数：": [
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
    "title": "删除模块",
    "description": "<p>删除权限模块</p>",
    "version": "0.1.0",
    "group": "AUTH",
    "parameter": {
      "fields": {
        "请求参数：": [
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
    "title": "修改模块",
    "description": "<p>修改权限模块</p>",
    "version": "0.1.0",
    "group": "AUTH",
    "parameter": {
      "fields": {
        "请求参数：": [
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
    "url": "/login",
    "title": "登录",
    "description": "<p>登录接口</p>",
    "version": "0.1.0",
    "group": "AUTH",
    "parameter": {
      "fields": {
        "请求参数：": [
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
        "请求参数：": [
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
        "请求参数：": [
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
  }
] });
