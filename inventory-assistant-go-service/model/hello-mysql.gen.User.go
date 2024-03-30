package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type _UserMgr struct {
	*_BaseMgr
}

// UserMgr open func
func UserMgr(db *gorm.DB) *_UserMgr {
	if db == nil {
		panic(fmt.Errorf("UserMgr need init by db"))
	}
	return &_UserMgr{_BaseMgr: &_BaseMgr{DB: db, isRelated: gloabIsRelated}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_UserMgr) GetTableName() string {
	return "User"
}

// Get 获取
func (obj *_UserMgr) Get() (result User, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_UserMgr) Gets() (results []*User, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_UserMgr) WithID(ID int) Option {
	return optionFunc(func(o *options) { o.query["id"] = ID })
}

// WithEmail email获取
func (obj *_UserMgr) WithEmail(Email string) Option {
	return optionFunc(func(o *options) { o.query["email"] = Email })
}

// WithName name获取
func (obj *_UserMgr) WithName(Name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = Name })
}

// GetByOption 功能选项模式获取
func (obj *_UserMgr) GetByOption(opts ...Option) (result User, err error) {
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
func (obj *_UserMgr) GetByOptions(opts ...Option) (results []*User, err error) {
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
func (obj *_UserMgr) GetFromID(ID int) (result User, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", ID).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_UserMgr) GetBatchFromID(IDs []int) (results []*User, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", IDs).Find(&results).Error

	return
}

// GetFromEmail 通过email获取内容
func (obj *_UserMgr) GetFromEmail(Email string) (result User, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("email = ?", Email).Find(&result).Error

	return
}

// GetBatchFromEmail 批量唯一主键查找
func (obj *_UserMgr) GetBatchFromEmail(Emails []string) (results []*User, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("email IN (?)", Emails).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_UserMgr) GetFromName(Name string) (results []*User, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("name = ?", Name).Find(&results).Error

	return
}

// GetBatchFromName 批量唯一主键查找
func (obj *_UserMgr) GetBatchFromName(Names []string) (results []*User, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("name IN (?)", Names).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_UserMgr) FetchByPrimaryKey(ID int) (result User, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", ID).Find(&result).Error

	return
}

// FetchByUnique primay or index 获取唯一内容
func (obj *_UserMgr) FetchByUnique(Email string) (result User, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("email = ?", Email).Find(&result).Error

	return
}
