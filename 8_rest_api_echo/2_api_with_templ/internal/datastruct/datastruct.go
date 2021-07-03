package datastruct

import "Nix_practice/8_rest_api_echo/2_api_with_templ/dbase"

type Data struct{
	Int            int
	Int2           int
	Ints           []int
	Str            string
	Strs           []string
	StructHomePage []dbase.HomePageStruct
	Comment        dbase.CommentInfo
	Comments       []dbase.CommentInfo
	Post           dbase.PostInfo
	Posts          []dbase.PostInfo

}
