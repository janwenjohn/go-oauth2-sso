# 关于
这是一个用go写的，实现了OAuth2协议及SSO单点登录协议的基础框架，可用于在此基础上搭建基于SSO单点登录的OAuth2受控安全协议

# 接口发布
## oauth2
test.yourhost.com:3000/oauth/authorize?client_id=abc&response_type=code&redirect_uri=http://www.baidu.com&scope=read
test.yourhost.com:3000/oauth/token?client_id=abc&client_secret=abc&grant_type=authorization_code&code=lM4RU738&redirect_uri=http://www.baidu.com
test.yourhost.com:3000/oauth/check?access_token=xxxxx&username=xxxxxx

## sso
test.yourhost.com:5000/login?service=http://www.baidu.com
test.yourhost.com:5000/serviceValidate?service=http://www.baidu.com&ticket=xxxxxxxxx

# 技术
选用以下框架：
go-martini
martini-contrib
go-redis
go-sql-driver

# 配置
## sso
ST有效期时间：constant.go:SERVICE_TICKET_TIME_TO_LIVE
TGT有效期时间：constant.go:TICKET_GRANTING_TICKET_TIME_TO_LIVE
cookie domain：constant.go:COOKIE_DOMAIN
日志目录：config.go:LOG_FILE
redis连接信息：config.go:Redis_Config

## oauth2
code超时时间：constant.go:REDIS_CODE_TIMEOUT
token有效时间：constant.go:TOKEN_VALIDATE_TIME_IN_SECONDS
日志目录：config.go:LOG_FILE
redis连接信息：config.go:Redis_Config
mysql连接信息：config.go:DB_Config
服务器路径配置：config.go:Server_Config

# 数据库
## sso
只使用redis存储数据
存储数据包括：
1. ST:  ST_xxxxxxxxxxxxxxx
2. TGT: TGT_xxxxxxxxxxxxxx
tgt及st超时时间使用redis ttl机制
## oauth2
使用redis和mysql存储数据
其中，redis存储包括：
1. code:    OAUTH_CODE_xxxxxxxx
code超时时间使用redis ttl机制
mysql存储包括：
oauth_client_details：
CREATE TABLE `oauth_client_details` (
  `client_id` varchar(100) NOT NULL,
  `client_secret` varchar(255) NOT NULL,
  `redirect_uri` varchar(255) DEFAULT NULL,
  `scope` varchar(10) DEFAULT NULL,
  `client_name` varchar(255) DEFAULT NULL,
  `client_logo` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`client_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

oauth_token:
CREATE TABLE `oauth_token` (
  `token` varchar(255) NOT NULL,
  `client_id` varchar(100) NOT NULL,
  `username` varchar(100) NOT NULL,
  `expiration` varchar(100) DEFAULT NULL,
  `token_type` varchar(20) DEFAULT NULL,
  `token_scope` varchar(20) DEFAULT NULL,
  `refresh_token` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`token`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


