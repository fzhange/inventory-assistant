package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type _PostMgr struct {
	*_BaseMgr
}

// PostMgr open func
func PostMgr(db *gorm.DB) *_PostMgr {
	if db == nil {
		panic(fmt.Errorf("PostMgr need init by db"))
	}
	return &_PostMgr{_BaseMgr: &_BaseMgr{DB: db, isRelated: gloabIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_PostMgr) GetTableName() string {
	return "Post"
}

// Get 获取
func (obj *_PostMgr) Get() (result Post, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_PostMgr) Gets() (results []*Post, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_PostMgr) WithID(ID int) Option {
	return optionFunc(func(o *options) { o.query["id"] = ID })
}

// WithTitle title获取
func (obj *_PostMgr) WithTitle(Title string) Option {
	return optionFunc(func(o *options) { o.query["title"] = Title })
}

// WithContent content获取
func (obj *_PostMgr) WithContent(Content string) Option {
	return optionFunc(func(o *options) { o.query["content"] = Content })
}

// WithPublished published获取
func (obj *_PostMgr) WithPublished(Published int8) Option {
	return optionFunc(func(o *options) { o.query["published"] = Published })
}

// WithAuthorID authorId获取
func (obj *_PostMgr) WithAuthorID(AuthorID int) Option {
	return optionFunc(func(o *options) { o.query["authorId"] = AuthorID })
}

// GetByOption 功能选项模式获取
func (obj *_PostMgr) GetByOption(opts ...Option) (result Post, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_PostMgr) GetByOptions(opts ...Option) (results []*Post, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_PostMgr) GetFromID(ID int) (result Post, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", ID).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_PostMgr) GetBatchFromID(IDs []int) (results []*Post, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", IDs).Find(&results).Error

	return
}

// GetFromTitle 通过title获取内容
func (obj *_PostMgr) GetFromTitle(Title string) (results []*Post, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("title = ?", Title).Find(&results).Error

	return
}

// GetBatchFromTitle 批量唯一主键查找
func (obj *_PostMgr) GetBatchFromTitle(Titles []string) (results []*Post, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("title IN (?)", Titles).Find(&results).Error

	return
}

// GetFromContent 通过content获取内容
func (obj *_PostMgr) GetFromContent(Content string) (results []*Post, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("content = ?", Content).Find(&results).Error

	return
}

// GetBatchFromContent 批量唯一主键查找
func (obj *_PostMgr) GetBatchFromContent(Contents []string) (results []*Post, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("content IN (?)", Contents).Find(&results).Error

	return
}

// GetFromPublished 通过published获取内容
func (obj *_PostMgr) GetFromPublished(Published int8) (results []*Post, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("published = ?", Published).Find(&results).Error

	return
}

// GetBatchFromPublished 批量唯一主键查找
func (obj *_PostMgr) GetBatchFromPublished(Publisheds []int8) (results []*Post, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("published IN (?)", Publisheds).Find(&results).Error

	return
}

// GetFromAuthorID 通过authorId获取内容
func (obj *_PostMgr) GetFromAuthorID(AuthorID int) (results []*Post, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("authorId = ?", AuthorID).Find(&results).Error

	return
}

// GetBatchFromAuthorID 批量唯一主键查找
func (obj *_PostMgr) GetBatchFromAuthorID(AuthorIDs []int) (results []*Post, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("authorId IN (?)", AuthorIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_PostMgr) FetchByPrimaryKey(ID int) (result Post, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", ID).Find(&result).Error

	return
}

// FetchByIndex  获取多个内容
func (obj *_PostMgr) FetchByIndex(AuthorID int) (results []*Post, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("authorId = ?", AuthorID).Find(&results).Error

	return
}
