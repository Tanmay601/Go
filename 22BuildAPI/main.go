package main

//model for courses - files

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"courseprice"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

//fake DB - files

var courses []Course

//middleware, helper - files

func (c *Course) IsEmpty() bool {

	return c.CourseId == "" && c.CourseName == "" && c.CourseName == ""
}
func main() {

}
