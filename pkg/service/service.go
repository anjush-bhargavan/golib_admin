package service

import (
	"errors"
	"fmt"
	user "github.com/anjush-bhargavan/golib_admin/pkg/user/handler"
	pb "github.com/anjush-bhargavan/golib_admin/pkg/pb"
	inter "github.com/anjush-bhargavan/golib_admin/pkg/repo/interfaces"
	"github.com/anjush-bhargavan/golib_admin/pkg/service/interfaces"
	userpb "github.com/anjush-bhargavan/golib_admin/pkg/user/pb"
	bookpb "github.com/anjush-bhargavan/golib_admin/pkg/books/pb"
)

type AdminService struct {
	AdminRepo inter.AdminRepoInter
	client    userpb.UserServiceClient
	bookClient bookpb.BookServicesClient
}

func NewAdminService(repo inter.AdminRepoInter, client userpb.UserServiceClient,bookClient bookpb.BookServicesClient) interfaces.AdminServiceInter {
	return &AdminService{
		AdminRepo: repo,
		client:    client,
		bookClient: bookClient,
	}
}

func (a *AdminService) AdminLoginService(p *pb.AdminRequest) (*pb.AdminResponse, error) {
	admin, err := a.AdminRepo.FindAdmin(p.Username)
	if err != nil {
		return nil, err
	}
	if admin.Password != p.Password {
		return nil, errors.New("incorrect password")
	}

	token,err := a.GenerateToken(p.Username)
	if err != nil {
		return nil,err
	}

	response := pb.AdminResponse{
		Status:  "Success",
		Error:   "",
		Message: "Logged in successfully, Token: "+ token,
	}
	return &response, nil
}


func (a *AdminService) EditUserService(p *pb.UserModel) (*pb.AdminResponse, error) {
	result, err := user.EditUserHandler(a.client, p)
	if err != nil {
		return nil, err
	}
	response := &pb.AdminResponse{
		Status: result.Status,
		Error: result.Error,
		Message: result.Message,
	}
	return response, nil
}

func (a *AdminService) DeleteUserService(p *pb.AUserID) (*pb.AdminResponse, error) {
	result, err := user.DeleteUserHandler(a.client, p)
	if err != nil {
		return nil, err
	}
	respose := &pb.AdminResponse{
		Status:  result.Status,
		Error:   result.Error,
		Message: result.Message,
	}

	return respose, nil
}

func (a *AdminService) FindUserByEmailService(p *pb.UserRequest) (*pb.User, error) {
	result, err := user.FindUserByEmailHandler(a.client, p)
	if err != nil {
		return nil, err
	}

	return &pb.User{
		Firstname: result.Firstname,
		Lastname:  result.Lastname,
		Username:  result.Username,
		Email:     result.Email,
		Dob:       result.Dob,
		Gender:    result.Gender,
		Phone:     result.Phone,
		Address:   result.Address,
	}, nil
}

func (a *AdminService) FindAllUserService(p *pb.AdminNoParam) (*pb.UserList, error) {
	result, err := user.FindAllUserHandler(a.client, p)
	if err != nil {
		return nil, err
	}

	var users []*pb.User
	for _, i := range result.Userlist {
		users = append(users, &pb.User{
			Firstname: i.Firstname,
			Lastname:  i.Lastname,
			Username:  i.Username,
			Email:     i.Email,
			Dob:       i.Dob,
			Gender:    i.Gender,
			Phone:     i.Phone,
			Address:   i.Address,
		})
	}
	response := &pb.UserList{
		Users: users,
	}


	return response, nil
}

func (a *AdminService) FindUserByIDService(p *pb.AUserID) (*pb.User, error) {
	result, err := user.FindUserByIDHandler(a.client, p)
	if err != nil {
		return nil, err
	}
	fmt.Println("2231")

	return &pb.User{
		Firstname: result.Firstname,
		Lastname:  result.Lastname,
		Username:  result.Username,
		Email:     result.Email,
		Dob:       result.Dob,
		Gender:    result.Gender,
		Phone:     result.Phone,
		Address:   result.Address,
	}, nil
}
