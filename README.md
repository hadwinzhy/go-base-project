# go_starter

go & gin & gorm starter


### LOCAL环境配置

go get -u github.com/cosmtrek/air
然后添加PATH到go的bin目录


### 升级包
go mod edit -require="github.com/gin-gonic/gin@v1.3.0"
go mod tidy #更新现有依赖

### 删除包
go mod
