package handler

import (
	"context"
	"log"

	adminpb "github.com/anjush-bhargavan/golib_admin/pkg/pb"
)

func (a *AdminHandler) UserCreateBook(ctx context.Context, p *adminpb.ABookModel) (*adminpb.AdminResponse, error) {
	result, err := a.AdminService.CreateBookServie(p)
	if err != nil {
		log.Println("error while creating book")
		return nil, err
	}
	return result, nil
}


func (a *AdminHandler) UserFetchBookByID(ctx context.Context,p *adminpb.ABookID) (*adminpb.ABookModel, error) {
	result, err := a.AdminService.FetchBookByIDService(p)
	if err != nil {
		log.Println("error while fetching book by id")
		return nil, err
	}
	return result, nil
}


func (a *AdminHandler) UserFetchBookByName(ctx context.Context, p *adminpb.ABookName) (*adminpb.ABookModel, error) {
	result, err := a.AdminService.FetchBookByNameService(p)
	if err != nil {
		log.Println("error while fetching book by name")
		return nil, err
	}
	return result, nil
}

func (a *AdminHandler) USerFetchAllBooks(ctx context.Context,p *adminpb.AdminNoParam) (*adminpb.ABookList, error) {
	result, err := a.AdminService.FetchAllBooksService(p)
	if err != nil {
		log.Println("error while fetching all books")
		return nil, err
	}
	return result, nil
}