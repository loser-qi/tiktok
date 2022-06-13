# tiktok

## 运行环境

go >= 1.18

MySQL >= 8.0

## 运行

1. 项目目录下新建conf.ini，按如下格式进行配置：

   ```
   [db]
   Host = 127.0.0.1
   Port = 3306
   Username = xxx
   Password = xxx
   Dbname = db_tiktok
   ```

2. MySQL运行create_table.sql建表

3. 安装依赖后，运行`go build && ./tiktok`
