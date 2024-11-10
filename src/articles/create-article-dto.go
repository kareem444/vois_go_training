package articles

type createArticleDto struct {
	Urls           []string `json:"urls" binding:"required,dive,url"`
	TitleSelector string `json:"title_selector" binding:"required,max=255"`
}
