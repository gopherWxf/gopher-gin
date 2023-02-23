package models

type ArticleModel struct {
	NewsID      int    `gorm:"column:id"`
	NewsTitle   string `gorm:"column:title"`
	NewsContent string `gorm:"column:content"`
}

func NewArticleModel() *ArticleModel {
	return &ArticleModel{}
}
func (this *ArticleModel) String() string {
	return "ArticleModel"
}
