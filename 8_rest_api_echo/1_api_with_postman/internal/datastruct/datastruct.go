package datastruct

import (

	"Nix_practice/8_rest_api_echo/1_api_with_postman/dbase"
	)

type Data struct{
	StructHomePage []dbase.HomePageStruct
	Comment        dbase.CommentInfo
	Comments       []dbase.CommentInfo
	Post           dbase.PostInfo
	Posts          []dbase.PostInfo

}
