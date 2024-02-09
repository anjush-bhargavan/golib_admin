package interfaces
import (
	pb "github.com/anjush-bhargavan/golib_admin/pkg/pb"

)

type AdminServiceInter interface {
	AdminLoginService(p *pb.AdminRequest) (*pb.AdminResponse, error)
	EditUserService(p *pb.UserModel) (*pb.AdminResponse, error)
	DeleteUserService(p *pb.AUserID) (*pb.AdminResponse, error)
	FindUserByEmailService(p *pb.UserRequest) (*pb.User, error)
	FindAllUserService(p *pb.AdminNoParam) (*pb.UserList, error)
	FindUserByIDService(p *pb.AUserID) (*pb.User, error)

	CreateBookServie(p *pb.ABookModel) (*pb.AdminResponse,error)
	FetchBookByIDService(p *pb.ABookID) (*pb.ABookModel,error)
	FetchBookByNameService(p *pb.ABookName) (*pb.ABookModel,error)
	FetchAllBooksService(p *pb.AdminNoParam) (*pb.ABookList,error)
}
