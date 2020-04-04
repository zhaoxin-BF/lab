package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type CoursesTable struct {
	Id            int    `orm:"column(order_id);auto"`
	LabId         int    `orm:"column(lab_id);null"`
	LabName       string `orm:"column(lab_name);size(50);null"`
	LabAddress    string `orm:"column(lab_address);size(40);null"`
	LabContext    string `orm:"column(lab_context);size(200);null"`
	Date          int    `orm:"column(date);null"`
	Datestr       string `orm:"column(datestr);size(20);null"`
	Week          string `orm:"column(week);size(10);null"`
	TeachId       int    `orm:"column(teach_id);null"`
	TeachName     string `orm:"column(teach_name);size(20);null"`
	TeachTel      int    `orm:"column(teach_tel);null"`
	ClassName     string `orm:"column(class_name);size(40);null"`
	CourseId      int    `orm:"column(course_id);null"`
	CourseIdstr   string `orm:"column(course_idstr);size(10);null"`
	CourseContext string `orm:"column(course_context);size(200);null"`
}

func (t *CoursesTable) TableName() string {
	return "courses_table"
}

func init() {
	orm.RegisterModel(new(CoursesTable))
}

// AddCoursesTable insert a new CoursesTable into database and returns
// last inserted Id on success.
func AddCoursesTable(m *CoursesTable) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetCoursesTableById retrieves CoursesTable by Id. Returns error if
// Id doesn't exist
func GetCoursesTableById(id int) (v *CoursesTable, err error) {
	o := orm.NewOrm()
	v = &CoursesTable{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllCoursesTable retrieves all CoursesTable matches certain condition. Returns empty list if
// no records exist
func GetAllCoursesTable(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(CoursesTable))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []CoursesTable
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateCoursesTable updates CoursesTable by Id and returns error if
// the record to be updated doesn't exist
func UpdateCoursesTableById(m *CoursesTable) (err error) {
	o := orm.NewOrm()
	v := CoursesTable{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteCoursesTable deletes CoursesTable by Id and returns error if
// the record to be deleted doesn't exist
func DeleteCoursesTable(id int) (err error) {
	o := orm.NewOrm()
	v := CoursesTable{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&CoursesTable{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
