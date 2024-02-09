package service

import(
	adminpb "github.com/anjush-bhargavan/golib_admin/pkg/pb"
	book "github.com/anjush-bhargavan/golib_admin/pkg/books/handler"
)

func (a *AdminService) CreateBookServie(p *adminpb.ABookModel) (*adminpb.AdminResponse,error) {
	result, err := book.CreateBookHandler(a.bookClient,p)
	if err != nil {
		return nil, err
	}

	return &adminpb.AdminResponse{
		Status: result.Status,
		Error: result.Error,
		Message: result.Message,
	},nil
}

func (a *AdminService) FetchBookByIDService(p *adminpb.ABookID) (*adminpb.ABookModel,error) {
	result, err := book.FetchBookByIDHandler(a.bookClient,p)
	if err != nil {
		return nil, err
	}

	return &adminpb.ABookModel{
		BookName: result.BookName,
		Author: result.Author,
		NoOfCopies: result.NoOfCopies,
		Year: result.Year,
		Publications: result.Publications,
		Category: result.Category,
		Description: result.Description,

	},nil
}

func (a *AdminService) FetchBookByNameService(p *adminpb.ABookName) (*adminpb.ABookModel,error) {
	result, err := book.FetchBookByNameHandler(a.bookClient,p)
	if err != nil {
		return nil, err
	}

	return &adminpb.ABookModel{
		BookName: result.BookName,
		Author: result.Author,
		NoOfCopies: result.NoOfCopies,
		Year: result.Year,
		Publications: result.Publications,
		Category: result.Category,
		Description: result.Description,

	},nil
}

func (a *AdminService) FetchAllBooksService(p *adminpb.AdminNoParam) (*adminpb.ABookList,error) {
	result, err := book.FetchAllBooksHandler(a.bookClient,p)
	if err != nil {
		return nil, err
	}
	var books []*adminpb.ABookModel
	for _,i := range result.Books {
		books = append(books, &adminpb.ABookModel{
			BookName: i.BookName,
			Author: i.Author,
			NoOfCopies: i.NoOfCopies,
			Year: i.Year,
			Publications: i.Publications,
			Category: i.Category,
			Description: i.Description,
		})
	}

	return &adminpb.ABookList{
		Books: books,
	},nil
}