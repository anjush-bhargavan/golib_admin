package handler

import (
	"context"
	"log"

	bookpb "github.com/anjush-bhargavan/golib_admin/pkg/books/pb"
	adminpb "github.com/anjush-bhargavan/golib_admin/pkg/pb"
)

func CreateBookHandler(client bookpb.BookServicesClient, p *adminpb.ABookModel) (*bookpb.BookResponse, error) {
	ctx := context.Background()

	response, err := client.CreateBook(ctx, &bookpb.BookModel{
		BookName:     p.BookName,
		Author:       p.Author,
		NoOfCopies:   p.NoOfCopies,
		Year:         p.Year,
		Publications: p.Publications,
		Category:     p.Category,
		Description:  p.Description,
	})
	if err != nil {
		log.Printf("error while creating book")
		return nil, err
	}
	return response, nil
}

func FetchBookByIDHandler(client bookpb.BookServicesClient, p *adminpb.ABookID) (*bookpb.BookModel, error) {
	ctx := context.Background()

	response, err := client.FetchBookByID(ctx, &bookpb.BookID{Id: p.Id})
	if err != nil {
		log.Printf("error while fetching book by name")
		return nil, err
	}
	return response, nil
}

func FetchBookByNameHandler(client bookpb.BookServicesClient, p *adminpb.ABookName) (*bookpb.BookModel, error) {
	ctx := context.Background()

	response, err := client.FetchBookByName(ctx, &bookpb.BookName{Name: p.Name})
	if err != nil {
		log.Printf("error while fetching book by id")
		return nil, err
	}
	return response, nil
}

func FetchAllBooksHandler(client bookpb.BookServicesClient, p *adminpb.AdminNoParam) (*bookpb.BookList, error) {
	ctx := context.Background()

	response, err := client.FetchAllBooks(ctx, &bookpb.NoParam{})
	if err != nil {
		log.Printf("error while fetching books")
		return nil, err
	}
	return response, nil
}
