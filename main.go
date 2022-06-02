package main

import (
	"net/http"
	"src/controllers/collegescontroller"
	//"src/controllers/studentcontroller"
	//"src/controllers/coursecontroller"
)

func main() {
	// http.HandleFunc("/", studentcontroller.Index)
	// http.HandleFunc("/colleges", studentcontroller.Index)
	// http.HandleFunc("/colleges/index", studentcontroller.Index)
	// http.HandleFunc("/colleges/add", studentcontroller.Add)
	// http.HandleFunc("/colleges/processadd", studentcontroller.ProcessAdd)

	http.HandleFunc("/", collegescontroller.Index)
	http.HandleFunc("/colleges", collegescontroller.Index)
	http.HandleFunc("/colleges/index", collegescontroller.Index)
	http.HandleFunc("/colleges/add", collegescontroller.Add)
	http.HandleFunc("/colleges/processadd", collegescontroller.ProcessAdd)
	http.HandleFunc("/colleges/delete", collegescontroller.Delete)
	http.HandleFunc("/colleges/edit", collegescontroller.Edit)
	http.HandleFunc("/colleges/update", collegescontroller.update)
	http.ListenAndServe(":3000", nil)

	// http.HandleFunc("/", coursecontroller.Index)
	// http.HandleFunc("/course", coursecontroller.Index)
	// http.HandleFunc("/course/index", coursecontroller.Index)

}
