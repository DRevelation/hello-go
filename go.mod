module hello-go

go 1.15

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/kataras/iris/v12 v12.1.8
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
	gee v0.0.0
)

replace (
	gee => ./gee
)