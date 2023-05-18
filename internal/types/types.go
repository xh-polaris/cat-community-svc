// Code generated by goctl. DO NOT EDIT.
package types

type SignInReq struct {
	AuthType string   `json:"authType,options=phone|email|wechat"`
	AuthId   string   `json:"authId"`
	Password string   `json:"password,optional"`
	Params   []string `json:"params,optional"`
}

type SignInResp struct {
	Status
	UserId       string `json:"userId"`
	AccessToken  string `json:"accessToken"`
	AccessExpire int64  `json:"accessExpire"`
}

type SetPasswordReq struct {
	Password string `json:"password"`
}

type SetPasswordResp struct {
	Status
}

type SendVerifyCodeReq struct {
	AuthType string `json:"authType,options=phone|email"`
	AuthId   string `json:"authId"`
}

type SendVerifyCodeResp struct {
	Status
}

type Status struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type UserPreview struct {
	Id        string `json:"id"`
	Nickname  string `json:"nickname"`
	AvatarUrl string `json:"avatarUrl"`
}

type Role struct {
	RoleType    string `json:"roleType"`
	CommunityId string `json:"communityId,omitempty"`
}

type Post struct {
	Id         string      `json:"id"`
	CreateAt   int64       `json:"createAt"`
	Title      string      `json:"title"`
	Text       string      `json:"text"`
	CoverUrl   string      `json:"coverUrl"`
	Tags       []string    `json:"tags"`
	Likes      int64       `json:"likes"`
	Comments   int64       `json:"comments"`
	User       UserPreview `json:"user"`
	IsOfficial bool        `json:"isOfficial"`
}

type SearchOptions struct {
	Key   *string `json:"key,optional"`
	Title *string `json:"title,optional"`
	Text  *string `json:"text,optional"`
	Tag   *string `json:"tag,optional"`
}

type GetPostPreviewsReq struct {
	Limit         *int64         `json:"limit,optional"`
	Offset        *int64         `json:"offset,optional"`
	LastToken     *string        `json:"lastToken,optional"`
	Backward      *bool          `json:"backward,optional"`
	OnlyOfficial  *bool          `json:"onlyOfficial,optional"`
	OnlyUserId    *string        `json:"onlyUserId,optional"`
	SearchOptions *SearchOptions `json:"searchOptions,optional"`
}

type GetPostPreviewsResp struct {
	Status
	Posts []Post `json:"posts"`
	Total int64  `json:"total"`
	Token string `json:"token"`
}

type GetPostDetailReq struct {
	PostId string `form:"postId"`
}

type GetPostDetailResp struct {
	Post Post `json:"post"`
	Status
}

type NewPostReq struct {
	Id       string   `json:"id,optional"`
	Title    string   `json:"title"`
	Text     string   `json:"text"`
	CoverUrl string   `json:"coverUrl,optional"`
	Tags     []string `json:"tags"`
}

type NewPostResp struct {
	PostId string `json:"postId"`
	Status
}

type DeletePostReq struct {
	Id string `json:"id"`
}

type DeletePostResp struct {
	Status
}

type SetOfficialReq struct {
	PostId   string `json:"postId"`
	IsRemove bool   `json:"isRemove,optional"`
}

type SetOfficialResp struct {
	Status
}

type Cat struct {
	Id           string   `json:"id"`
	CreateAt     int64    `json:"createAt"`
	Age          string   `json:"age"`
	CommunityId  string   `json:"communityId"`
	Color        string   `json:"color"`
	Details      string   `json:"details"`
	Name         string   `json:"name"`
	Popularity   int64    `json:"popularity"`
	Sex          string   `json:"sex"`
	Status       int32    `json:"status"`
	Area         string   `json:"area"`
	IsSnipped    bool     `json:"isSnipped"`
	IsSterilized bool     `json:"isSterilized"`
	Avatars      []string `json:"avatars"`
}

type CatPreview struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Area        string `json:"area"`
	Color       string `json:"color"`
	AvatarUrl   string `json:"avatarUrl"`
	IsCollected bool   `json:"isCollected"`
}

type GetCatPreviewsReq struct {
	CommunityId string `form:"communityId"`
	Page        int64  `form:"page"`
}

type GetCatPreviewsResp struct {
	Status
	Cats  []CatPreview `json:"cats"`
	Total int64        `json:"total"`
}

type GetCatDetailReq struct {
	CatId string `form:"catId"`
}

type GetCatDetailResp struct {
	Status
	Cat Cat `json:"cat"`
}

type DeleteCatReq struct {
	CatId string `json:"catId"`
}

type DeleteCatResp struct {
	Status
}

type NewCatReq struct {
	Id           string   `json:"id,optional"`
	Age          string   `json:"age"`
	CommunityId  string   `json:"communityId"`
	Color        string   `json:"color"`
	Details      string   `json:"details"`
	Name         string   `json:"name"`
	Sex          string   `json:"sex"`
	Area         string   `json:"area"`
	IsSnipped    bool     `json:"isSnipped"`
	IsSterilized bool     `json:"isSterilized"`
	Avatars      []string `json:"avatars"`
}

type NewCatResp struct {
	Status
	CatId string `json:"catId"`
}

type SearchCatReq struct {
	CommunityId string `form:"communityId"`
	Keyword     string `form:"keyword"`
	Page        int64  `form:"page"`
}

type SearchCatResp struct {
	Status
	Cats  []CatPreview `json:"cats"`
	Total int64        `json:"total"`
}

type CreateImageElement struct {
	CatId string `json:"catId"`
	Url   string `json:"url"`
}

type CreateImageReq struct {
	Images []CreateImageElement `json:"images"`
}

type CreateImageResp struct {
	Status
	Id []string `json:"id"`
}

type DeleteImageReq struct {
	Id string `json:"id"`
}

type DeleteImageResp struct {
	Status
}

type GetImageByCatReq struct {
	CatId    string `form:"catId"`
	PrevId   string `form:"prevId,optional"`
	Limit    int64  `form:"limit,default=10"`
	Offset   int64  `form:"offset,optional"`
	Backward bool   `form:"backward,optional"`
}

type Image struct {
	Id    string `json:"id"`
	Url   string `json:"url"`
	CatId string `json:"catId"`
}

type GetImageByCatResp struct {
	Status
	Images []Image `json:"images"`
	Total  int64   `json:"total"`
}

type Moment struct {
	Id          string      `json:"id"`
	CreateAt    int64       `json:"createAt"`
	CatId       string      `json:"catId,optional"`
	Photos      []string    `json:"photos"`
	Title       string      `json:"title"`
	Text        string      `json:"text"`
	User        UserPreview `json:"user"`
	CommunityId string      `json:"communityId"`
}

type GetMomentPreviewsReq struct {
	CommunityId *string `form:"communityId,optional"`
	IsParent    bool    `form:"isParent,default=false"`
	OnlyUserId  *string `form:"onlyUserId,optional"`
	Page        int64   `form:"page"`
	Limit       *int64  `form:"limit,optional"`
	LastToken   *string `form:"lastToken,optional"`
	Backward    *bool   `form:"backward,optional"`
}

type GetMomentPreviewsResp struct {
	Status
	Moments []Moment `json:"moments"`
	Total   int64    `json:"total"`
}

type GetMomentDetailReq struct {
	MomentId string `form:"momentId"`
}

type GetMomentDetailResp struct {
	Status
	Moment Moment `json:"moment"`
}

type DeleteMomentReq struct {
	MomentId string `json:"momentId"`
}

type DeleteMomentResp struct {
	Status
}

type NewMomentReq struct {
	Id          string   `json:"id,optional"`
	Title       string   `json:"title,optional"`
	CatId       string   `json:"catId,optional"`
	Text        string   `json:"text,optional"`
	Photos      []string `json:"photos"`
	CommunityId string   `json:"communityId"`
}

type NewMomentResp struct {
	MomentId string `json:"momentId"`
	Status
}

type SearchMomentReq struct {
	CommunityId *string `form:"communityId,optional"`
	IsParent    bool    `form:"isParent,default=false"`
	OnlyUserId  *string `form:"onlyUserId,optional"`
	Keyword     string  `form:"keyword,optional"`
	Page        int64   `form:"page"`
	Limit       *int64  `form:"limit,optional"`
	LastToken   *string `form:"lastToken,optional"`
	Backward    *bool   `form:"backward,optional"`
}

type SearchMomentResp struct {
	Status
	Moments []Moment `json:"moments"`
	Total   int64    `json:"total"`
}

type Comment struct {
	Id        string      `json:"id"`
	CreateAt  int64       `json:"createAt"`
	Text      string      `json:"text"`
	User      UserPreview `json:"user"`
	Comments  int64       `json:"comments"`
	ReplyName string      `json:"replyName,optional"`
}

type NewCommentReq struct {
	Text  string `json:"text"`
	Id    string `json:"id,optional"`
	Scope string `json:"scope"`
}

type NewCommentResp struct {
	Status
}

type GetCommentsReq struct {
	Scope string `form:"scope"`
	Page  int64  `form:"page"`
	Id    string `form:"id"`
}

type GetCommentsResp struct {
	Status
	Comments []Comment `json:"comments"`
	Total    int64     `json:"total"`
}

type DeleteCommentReq struct {
	CommentId string `json:"commentId"`
}

type DeleteCommentResp struct {
	Status
}

type News struct {
	Id          string `json:"id"`
	CreateAt    int64  `json:"createAt"`
	CommunityId string `json:"communityId"`
	ImageUrl    string `json:"imageUrl"`
	LinkUrl     string `json:"linkUrl"`
	Type        string `json:"type"`
}

type Admin struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	CommunityId string `json:"communityId"`
	Name        string `json:"name"`
	Phone       string `json:"phone"`
	AvatarUrl   string `json:"avatarUrl"`
	Wechat      string `json:"wechat"`
}

type Notice struct {
	Id       string `json:"id"`
	Text     string `json:"text"`
	CreateAt int64  `json:"createAt"`
}

type GetNewsReq struct {
	CommunityId string `form:"communityId"`
}

type GetNewsResp struct {
	Status
	News []News `json:"news"`
}

type GetAdminsReq struct {
	CommunityId string `form:"communityId"`
}

type GetAdminsResp struct {
	Status
	Admins []Admin `json:"admins"`
}

type DeleteAdminReq struct {
	Id string `json:"id"`
}

type DeleteAdminResp struct {
	Status
}

type NewAdminReq struct {
	Id          string `json:"id,optional"`
	Title       string `json:"title"`
	CommunityId string `json:"communityId"`
	Name        string `json:"name"`
	Phone       string `json:"phone"`
	AvatarUrl   string `json:"avatarUrl"`
	Wechat      string `json:"wechat"`
}

type NewAdminResp struct {
	Status
	Id string `json:"id"`
}

type GetNoticesReq struct {
	CommunityId string `form:"communityId"`
}

type GetNoticesResp struct {
	Status
	Notices []Notice `json:"notices"`
}

type NewNoticeReq struct {
	Id          string `json:"id,optional"`
	CommunityId string `json:"communityId,optional"`
	Text        string `json:"text"`
}

type NewNoticeResp struct {
	Status
	Id string `json:"id"`
}

type NewNewsReq struct {
	Id          string `json:"id,optional"`
	CommunityId string `json:"communityId,optional"`
	ImageUrl    string `json:"imageUrl"`
	LinkUrl     string `json:"linkUrl"`
	Type        string `json:"type"`
}

type NewNewsResp struct {
	Status
	Id string `json:"id"`
}

type DeleteNoticeReq struct {
	Id string `json:"id"`
}

type DeleteNoticeResp struct {
	Status
}

type DeleteNewsReq struct {
	Id string `json:"id"`
}

type DeleteNewsResp struct {
	Status
}

type Community struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ParentId string `json:"parentId,omitempty"`
}

type ListCommunityReq struct {
	ParentId string `form:"parentId,optional"`
}

type ListCommunityResp struct {
	Communities []Community `json:"communities"`
	Status
}

type NewCommunityReq struct {
	Id       string `json:"id,optional"`
	Name     string `json:"name""`
	ParentId string `json:"parentId,optional"`
}

type NewCommunityResp struct {
	Id string `json:"id"`
	Status
}

type DeleteCommunityReq struct {
	Id string `json:"id"`
}

type DeleteCommunityResp struct {
	Status
}

type GetUserRolesReq struct {
}

type GetUserRolesResp struct {
	Roles []Role `json:"roles"`
	Status
}

type UpdateCommunityAdminReq struct {
	UserId      string `json:"userId"`
	CommunityId string `json:"communityId"`
	IsRemove    bool   `json:"isRemove,default=false"`
}

type UpdateCommunityAdminResp struct {
	Status
}

type UpdateSuperAdminReq struct {
	UserId   string `json:"userId"`
	IsRemove bool   `json:"isRemove,default=false"`
}

type UpdateSuperAdminResp struct {
	Status
}

type RetrieveUserPreviewReq struct {
	RoleType    string `form:"roleType"`
	CommunityId string `form:"communityId,optional"`
}

type RetrieveUserPreviewResp struct {
	Users []UserPreview `json:"users"`
	Status
}

type DoLikeReq struct {
	TargetId   string `json:"targetId"`
	TargetType int64  `json:"targetType"`
}

type DoLikeResp struct {
	Status
}

type GetUserLikedReq struct {
	TargetId   string `form:"targetId"`
	TargetType int64  `form:"targetType"`
}

type GetUserLikedResp struct {
	Status
	Liked bool `json:"liked"`
}

type GetLikedCountReq struct {
	TargetId   string `form:"targetId"`
	TargetType int64  `form:"targetType"`
}

type GetLikedCountResp struct {
	Status
	Count int64 `json:"count"`
}

type GetUserLikesReq struct {
	UserId     string `form:"userId"`
	TargetType int64  `form:"targetType"`
}

type Like struct {
	TargetId     string `json:"targetId"`
	AssociatedId string `json:"associatedId"`
}

type GetUserLikesResp struct {
	Likes []Like `json:"likes"`
	Status
}

type GetLikedUsersReq struct {
	TargetId   string `form:"targetId"`
	TargetType int64  `form:"targetType"`
}

type GetLikedUsersResp struct {
	Status
	Users []UserPreview `json:"users"`
}

type User struct {
	Id        string `json:"id"`
	Nickname  string `json:"nickname"`
	AvatarUrl string `json:"avatarUrl"`
	Motto     string `json:"motto"`
	Follower  int64  `json:"follower"`
	Following int64  `json:"following"`
	Article   int64  `json:"article"`
	Like      int64  `json:"like"`
}

type UserPreviewWithRole struct {
	UserPreview
	Roles []Role `json:"roles"`
}

type GetUserInfoReq struct {
	UserId *string `form:"userId,optional"`
}

type GetUserInfoResp struct {
	Status
	User User `json:"user"`
}

type UpdateUserInfoReq struct {
	AvatarUrl string `json:"avatarUrl,optional"`
	Nickname  string `json:"nickname,optional"`
	Motto     string `json:"motto,optional"`
}

type UpdateUserInfoResp struct {
	Status
}

type SearchUserReq struct {
	Keyword   string `form:"keyword"`
	Page      int64  `form:"page"`
	Limit     int64  `form:"limit,optional"`
	LastToken string `form:"lastToken,optional"`
}

type SearchUserResp struct {
	Status
	Users []UserPreview `json:"users"`
	Total int64         `json:"total"`
	Token string        `json:"token"`
}

type SearchUserForAdminReq struct {
	Keyword   string `form:"keyword"`
	Page      int64  `form:"page"`
	Limit     int64  `form:"limit,optional"`
	LastToken string `form:"lastToken,optional"`
}

type SearchUserForAdminResp struct {
	Status
	Users []UserPreviewWithRole `json:"users"`
	Total int64                 `json:"total"`
	Token string                `json:"token"`
}

type ApplySignedUrlReq struct {
	Prefix string `json:"prefix,optional"`
	Suffix string `json:"suffix,optional"`
}

type ApplySignedUrlResp struct {
	Status
	Url          string `json:"url"`
	SessionToken string `json:"sessionToken"`
}

type ApplySignedUrlAsCommunityReq struct {
	CommunityId string `json:"communityId"`
	Prefix      string `json:"prefix,optional"`
	Suffix      string `json:"suffix,optional"`
}

type ApplySignedUrlAsCommunityResp struct {
	Status
	Url          string `json:"url"`
	SessionToken string `json:"sessionToken"`
}
