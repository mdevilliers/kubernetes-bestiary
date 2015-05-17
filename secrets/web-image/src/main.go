package main
 
import(
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
)
 
func indexHandler( w http.ResponseWriter, r *http.Request){

	path := "/etc/foo"

	username, err := ioutil.ReadFile(path + "/username")

	if err != nil {
		fmt.Fprintf(w, "error reading username : %s", err.Error())
	}

	password, err := ioutil.ReadFile(path + "/password")

	if err != nil {
		fmt.Fprintf(w, "error reading password : %s", err.Error())
	}

	other, err := ioutil.ReadFile(path + "/another-secret")

	if err != nil {
		fmt.Fprintf(w, "error reading another-secret : %s", err.Error())
	}

	fmt.Fprintf(w, "ssh!, I'm running on %s. My secrets are username : %s, password : %s, other : %s", 
		runtime.GOOS,string(username), string(password), string(other))
}
 
func main(){
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080",nil)
}