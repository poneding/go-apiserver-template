package db

import (
	"go-apiserver-template/pkg/util"

	"gorm.io/gorm"
)

// instance is a singleton instance of the database repository
var repositoryInstance *repository

// Repository returns a singleton instance of the database repository
func Repository() *repository {
	if repositoryInstance == nil {
		repositoryInstance = &repository{
			DB: instance(),
		}
	}
	return repositoryInstance
}

// Ping is a helper function to ping the database repository
func Ping() error {
	sqlDB, err := Repository().DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Ping()
}

// repository is a wrapper around *gorm.DB
// It provides helper functions to perform CRUD operations
type repository struct {
	*gorm.DB
}

// IsExist checks if the entry with the given id exists in the database
func (r *repository) IsExist(id any, entry any) bool {
	if err := r.DB.First(entry, "id = ?", id).Error; err != nil {
		return false
	}
	return true
}

// Get retrieves the entry with the given id from the database
// If the entry does not exist, it returns an error of type gorm.ErrRecordNotFound
func (r *repository) Get(id, entry any) error {
	if err := r.DB.First(entry, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}

// FirstOption is a optional struct to retrieve the first entry from the database
type FirstOption struct {
	Condition *condition
}

// First retrieves the first entry from the database
// If the entry does not exist, it returns an error of type gorm.ErrRecordNotFound
func (r *repository) First(entry any, opt ...FirstOption) error {
	var firstOption FirstOption
	if len(opt) > 0 {
		firstOption = opt[0]
	}
	query := r.DB.Model(entry)
	if firstOption.Condition != nil {
		query = query.Where(firstOption.Condition.Expression, firstOption.Condition.Arguments...)
	}
	return query.First(entry).Error
}

// Create creates a new entry in the database
func (r *repository) Create(entry any) error {
	return r.DB.Model(entry).Create(entry).Error
}

// UpdateOption is a optional struct to update the entry in the database
type UpdateOption struct {
	Condition *condition
	Columns   []string
}

// Update updates the entry in the database
func (r *repository) Update(entry any, opt ...UpdateOption) error {
	var updateOption UpdateOption
	if len(opt) > 0 {
		updateOption = opt[0]
	}
	query := r.DB.Model(entry)
	if updateOption.Condition != nil {
		query = query.Where(updateOption.Condition.Expression, updateOption.Condition.Arguments...)
	}
	return query.Select(updateOption.Columns).Updates(entry).Error
}

// DeleteByID deletes the entry with the given id from the database
func (r *repository) DeleteByID(id any, entry any) error {
	return r.DB.Where("id = ?", id).Delete(entry).Error
}

// DeleteOption is a optional struct to delete the entry from the database
type DeleteOption struct {
	Condition *condition
}

// Delete deletes the entry from the database
func (r *repository) Delete(entry any, opt DeleteOption) error {
	if opt.Condition == nil {
		return nil
	}
	return r.DB.Where(opt.Condition.Expression, opt.Condition.Arguments...).Delete(entry).Error
}

// condition is a optional struct to specify a sql condition
type condition struct {
	// Expression is a sql condition expression, e.g. "id = ? AND name = ?"
	Expression string
	// Arguments are the arguments for the sql condition
	// expression, e.g. []interface{}{1, "test"}
	Arguments []any
}

// Condition is a helper function to create a new condition
func Condition(exp string, args ...any) *condition {
	return &condition{
		Expression: exp,
		Arguments:  args,
	}
}

// ListOption is a optional struct to list entries from the database
type ListOption struct {
	Condition *condition
	OrderBy   string
	Desc      bool
}

// List lists entries from the database
func (r *repository) List(entries any, opt ...ListOption) error {
	var listOption ListOption
	if len(opt) > 0 {
		listOption = opt[0]
	}
	query := r.DB

	if listOption.Condition != nil {
		query = query.Where(listOption.Condition.Expression, listOption.Condition.Arguments...)
	}
	if len(listOption.OrderBy) > 0 {
		if listOption.Desc {
			listOption.OrderBy = listOption.OrderBy + " DESC"
		}
		query = query.Order(listOption.OrderBy)
	}
	return query.Find(entries).Error
}

// PageOption is a optional struct to page entries from the database
type PageOption struct {
	Condition *condition
	No        int
	Size      int
	OrderBy   string
	Desc      bool
}

// Page retrieves paged entries from the database
func (r *repository) Page(entries any, opt ...PageOption) (int64, error) {
	var pageOption PageOption
	if len(opt) > 0 {
		pageOption = opt[0]
		pageOption.Size = util.ValueIf(pageOption.Size != 0, pageOption.Size, 10)
	} else {
		pageOption = PageOption{
			No:   1,
			Size: 10,
		}
	}
	query := r.DB

	if pageOption.Condition != nil {
		query = query.Where(pageOption.Condition.Expression, pageOption.Condition.Arguments...)
	}
	if len(pageOption.OrderBy) > 0 {
		if pageOption.Desc {
			pageOption.OrderBy = pageOption.OrderBy + " DESC"
		}
		query = query.Order(pageOption.OrderBy)
	}
	var total int64
	query.Model(entries).Count(&total)
	err := query.Offset((pageOption.No - 1) * pageOption.Size).Limit(pageOption.Size).Find(entries).Error
	return total, err
}
