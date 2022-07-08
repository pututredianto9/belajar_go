package middlewareAuth

type UserData struct {
	UserID          int64  `json:"uid"`
	UserName        string `json:"uname"`
	RepresentedID   int64  `json:"rid"`
	RepresentedName string `json:"rname"`
	RepresentedSlug string `json:"rslug"`
	Role            string `json:"rol"`
}
