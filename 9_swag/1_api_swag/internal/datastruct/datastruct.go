package datastruct

import (
	"github.com/Igor-Koniukhov/api_with_swag/dbase"

	)

type Data struct{
	StructHomePage []dbase.HomePageStruct
	Comment        dbase.CommentInfo
	Comments       []dbase.CommentInfo
	Post           dbase.PostInfo
	Posts          []dbase.PostInfo

}
