package collegescontroller

import (
	"html/template"
	"net/http"
	"src/entites"
	"src/models"
	"strconv"
)

func Index(response http.ResponseWriter, request *http.Request) {

	var collegesModel models.CollegesModel
	colleges, _ := collegesModel.FindAll()

	data := map[string]interface{}{
		"colleges": colleges,
	}
	tmp, _ := template.ParseFiles("views/colleges/index.html")
	tmp.Execute(response, data)
}
func Add(response http.ResponseWriter, request *http.Request) {
	tmp, _ := template.ParseFiles("views/colleges/add.html")
	tmp.Execute(response, nil)
}

func ProcessAdd(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	var colleges entites.Colleges
	colleges.COL_TYPE = request.Form.Get("col_type")
	colleges.COL_ADD = request.Form.Get("COLLEGE address")
	var collegesModel models.CollegesModel

	collegesModel.Create(&colleges)
	http.Redirect(response, request, "/colleges", http.StatusSeeOther)
}
func Delete(response http.ResponseWriter, request *http.Request) {
	query:= request.URL.Query()

	id,_ := strconv.ParseInt(query.Get("id"),10,64)
	var collegesModel models.StudentModel
	collegesModel.Delete(id)
	http.Redirect(response,request,"/college",http.StatusSeeOther)


}
func Edit(response http.ResponseWriter, request *http.Request) {
	query:= request.URL.Query()

	id,_ := strconv.ParseInt(query.Get("id"),10,64)
	var collegesModel models.StudentModel
	colleges,_ := collegesModel.Find(id),
    data :=map[string]interface{}{

	"colleges":colleges,
}
	tmp, _ := template.ParseFiles("views/colleges/add.html")
	tmp.Execute(response, nil)

}

func update(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	var colleges entites.Colleges
	colleges.COL_TYPE = request.Form.Get("col_type")
	colleges.COL_ADD = request.Form.Get("COLLEGE address")
	var collegesModel models.CollegesModel

	collegesModel.Update(&colleges)
	http.Redirect(response, request, "/colleges", http.StatusSeeOther)
}