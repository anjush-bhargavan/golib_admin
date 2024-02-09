package interfaces

import "github.com/anjush-bhargavan/golib_admin/pkg/dom"


type AdminRepoInter interface {
	FindAdmin(username string) (*dom.AdminModel,error)
}