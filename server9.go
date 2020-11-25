package main

import(
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Server struct{
	Materias map[string]map[string]float64
	Alumnos map[string]map[string]float64
	Lista_alumnos []string
	Lista_materias []string
}

var S Server

func (this *Server) verAlumnos() string {
	var html string
	for k:=0;k < len(S.Lista_alumnos);k++{
		html += "<tr>" +
			"<td>" + S.Lista_alumnos[k] + "</td>" +
			"</tr>"
	}

	return html
}

func (this *Server) verMaterias() string {
	var html string
	for k:=0;k < len(S.Lista_materias);k++{
		html += "<tr>" +
			"<td>" + S.Lista_materias[k] + "</td>" +
			"</tr>"
	}

	return html
}

func cargarHtml(a string) string{
	html,_ := ioutil.ReadFile(a)

	return string(html)
}


func form(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	fmt.Fprintf(
		res,
		cargarHtml("form.html"),
	)
}

func agregarAlumno(res http.ResponseWriter, req *http.Request){
	fmt.Println(req.Method)
	switch req.Method{
		case "POST":
			if err := req.ParseForm(); err != nil{
				fmt.Fprintf(res, "ParseForm() error %v", err)
				return
			}
			fmt.Println(req.PostForm)

			fmt.Println("-------------------AgregarAlumno")
			materia := make(map[string]float64)
			S.Alumnos[req.FormValue("alumno")] = materia
			S.Lista_alumnos = append(S.Lista_alumnos, req.FormValue("alumno"))
			fmt.Println("Nueva Lista de alumnos: ",S.Lista_alumnos)

			res.Header().Set(
				"Content-Type",
				"text/html",
			)
			fmt.Fprintf(
				res,
				cargarHtml("respuestaAlumno.html"),
				req.FormValue("alumno"),
			)
		case "GET":
			res.Header().Set(
				"Content-Type",
				"text/html",
			)
			fmt.Fprintf(
				res,
				cargarHtml("tablaAlumnos.html"),
				S.verAlumnos(),
			)
	}
}

func agregarMateria(res http.ResponseWriter, req *http.Request){
	fmt.Println(req.Method)
	switch req.Method{
		case "POST":
			if err := req.ParseForm(); err != nil{
				fmt.Fprintf(res, "ParseForm() error %v", err)
				return
			}
			fmt.Println(req.PostForm)

			fmt.Println("-------------------AgregarMateria")
			alumno := make(map[string]float64)
			S.Materias[req.FormValue("materia")] = alumno
			S.Lista_materias = append(S.Lista_materias, req.FormValue("materia"))
			fmt.Println("Nueva Lista de materias: ",S.Lista_materias)

			res.Header().Set(
				"Content-Type",
				"text/html",
			)
			fmt.Fprintf(
				res,
				cargarHtml("respuestaMateria.html"),
				req.FormValue("materia"),
			)
		case "GET":
			res.Header().Set(
				"Content-Type",
				"text/html",
			)
			fmt.Fprintf(
				res,
				cargarHtml("tablaMaterias.html"),
				S.verMaterias(),
			)
	}
}

func agregarCalifMateria(res http.ResponseWriter, req *http.Request){
	var calif float64
	fmt.Println(req.Method)
	switch req.Method{
		case "POST":
			if err := req.ParseForm(); err != nil{
				fmt.Fprintf(res, "ParseForm() error %v", err)
				return
			}
			if s, err := strconv.ParseFloat(req.FormValue("califCalif"), 64); err == nil {
				//fmt.Println(s)
				calif = s
			}

			fmt.Println("-------------------AgregarCalifMateria")
			fmt.Println("MATERIA: ",req.FormValue("materiaCalif"))
			fmt.Println("ALUMNO: ",req.FormValue("alumnoCalif"))
			fmt.Println("CALIF: ",req.FormValue("califCalif"))
			fmt.Println(calif)

			S.Materias[req.FormValue("materiaCalif")][req.FormValue("alumnoCalif")] = calif
			S.Alumnos[req.FormValue("alumnoCalif")][req.FormValue("materiaCalif")] = calif

			fmt.Println(S.Materias)
			fmt.Println(S.Alumnos)

			fmt.Println("Se agrego la calificacion de "+req.FormValue("califCalif")+" a el/la alumn@ "+req.FormValue("alumnoCalif")+" en la materia de "+req.FormValue("materiaCalif"))


			res.Header().Set(
				"Content-Type",
				"text/html",
			)
			fmt.Fprintf(
				res,
				cargarHtml("respuestaCalifMateria.html"),
				req.FormValue("materiaCalif"),
				req.FormValue("alumnoCalif"),
				req.FormValue("califCalif"),
			)
	}
}

func promedioAlumno(res http.ResponseWriter, req *http.Request){
	fmt.Println(req.Method)
	switch req.Method{
		case "POST":
			if err := req.ParseForm(); err != nil{
				fmt.Fprintf(res, "ParseForm() error %v", err)
				return
			}
			fmt.Println(req.PostForm)

			fmt.Println("-------------------promedioAlumno")
			var i float64
			i = 0
			var promedio float64
			promedio = 0
			fmt.Println(req.FormValue("alumnoPromedio"))
			for materia, calificion := range S.Alumnos[req.FormValue("alumnoPromedio")] {
				fmt.Println(materia + ":", calificion)
				i++
				promedio = promedio + calificion
				} 
			promedio = promedio/i
			str_prom := fmt.Sprintf("%f", promedio)

			res.Header().Set(
				"Content-Type",
				"text/html",
			)
			fmt.Fprintf(
				res,
				cargarHtml("respuestaPromedioAlumno.html"),
				req.FormValue("alumnoPromedio"),
				str_prom,
			)
	}
}

func promedioGeneral(res http.ResponseWriter, req *http.Request){
	fmt.Println(req.Method)
	switch req.Method{
		case "POST":
			if err := req.ParseForm(); err != nil{
				fmt.Fprintf(res, "ParseForm() error %v", err)
				return
			}
			fmt.Println(req.PostForm)

			fmt.Println("-------------------PromedioGeneral")
			var i float64
			i = 0
			var promedio_Alumno float64
			var num_clases float64
			var promedio_General float64
			for k:=0;k < len(S.Lista_alumnos);k++{
				promedio_Alumno = 0
				num_clases = 0
				for _, calificion := range S.Alumnos[S.Lista_alumnos[int(i)]]{
					promedio_Alumno = promedio_Alumno + calificion
					num_clases++
				}
				i++
				promedio_Alumno = promedio_Alumno/num_clases
				promedio_General = promedio_General + promedio_Alumno
			}
			promedio_General = promedio_General/float64(len(S.Lista_alumnos))
			fmt.Println("El promedio general es de: ", promedio_General)

			res.Header().Set(
				"Content-Type",
				"text/html",
			)
			fmt.Fprintf(
				res,
				cargarHtml("respuestaPromedio.html"),
				promedio_General,
			)
	}
}

func promedioMateria(res http.ResponseWriter, req *http.Request){
	fmt.Println(req.Method)
	switch req.Method{
		case "POST":
			if err := req.ParseForm(); err != nil{
				fmt.Fprintf(res, "ParseForm() error %v", err)
				return
			}
			fmt.Println(req.PostForm)

			fmt.Println("-------------------promedioMateria")
			var i float64
			i = 0
			var promedio float64
			promedio = 0
			fmt.Println(req.FormValue("materiaPromedio"))
			for materia, calificion := range S.Materias[req.FormValue("materiaPromedio")] {
				fmt.Println(materia + ":", calificion)
				i++
				promedio = promedio + calificion
				} 
			promedio = promedio/i
			str_prom := fmt.Sprintf("%f", promedio)

			fmt.Println("El promedio de "+req.FormValue("materiaPromedio")+" es de "+str_prom)

			res.Header().Set(
				"Content-Type",
				"text/html",
			)
			fmt.Fprintf(
				res,
				cargarHtml("respuestaPromedioMateria.html"),
				req.FormValue("materiaPromedio"),
				str_prom,
			)
	}
}

func main(){
	S.Materias = make(map[string]map[string]float64)
	S.Alumnos = make(map[string]map[string]float64)
	http.HandleFunc("/form", form)
	http.HandleFunc("/agregarAlumno", agregarAlumno)
	http.HandleFunc("/agregarMateria", agregarMateria)
	http.HandleFunc("/agregarCalifMateria", agregarCalifMateria)
	http.HandleFunc("/promedioAlumno", promedioAlumno)
	http.HandleFunc("/promedioGeneral", promedioGeneral)
	http.HandleFunc("/promedioMateria", promedioMateria)
	fmt.Println("INICIANDO SERVIDOR 09")
	http.ListenAndServe(":9000",nil)
}