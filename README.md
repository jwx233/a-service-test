# Vercel Java 8 API Demo

基于 Vercel 的 Java 8 后端 API 服务。

## 项目结构

```
├── src/main/java/cn/kamakura/aservice/
│   ├── Application.java           # 应用入口
│   └── handler/
│       ├── HelloHandler.java      # Hello World API
│       └── GreetHandler.java      # 带参数的问候 API
├── vercel.json                    # Vercel 配置文件
└── pom.xml
```

## API 接口

| 接口 | 方法 | 说明 |
|------|------|------|
| `/api/hello` | GET | 返回 Hello World |
| `/api/greet?name=xxx` | GET | 返回个性化问候 |

## 部署步骤

1. 安装 Vercel CLI：
   ```bash
   npm install -g vercel
   ```

2. 登录 Vercel：
   ```bash
   vercel login
   ```

3. 部署项目：
   ```bash
   vercel
   ```

4. 生产环境部署：
   ```bash
   vercel --prod
   ```

## 本地测试

```bash
vercel dev
```

## 示例响应

`GET /api/hello`:
```json
{
  "message": "Hello World!",
  "status": "success",
  "platform": "Vercel",
  "java": "8"
}
```

`GET /api/greet?name=Kiro`:
```json
{
  "message": "Hello, Kiro!",
  "timestamp": 1702195200000
}
```
