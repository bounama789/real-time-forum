package service

var (
	AuthSrvice AuthService
	PostSrvice PostService
	ComSrvice  CommentService
)

func init() {
	AuthSrvice.init()
	PostSrvice.init()
	ComSrvice.init()
}
