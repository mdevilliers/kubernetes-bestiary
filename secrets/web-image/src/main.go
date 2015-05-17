package main
 
import(
    "encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
)
 
func indexHandler( w http.ResponseWriter, r *http.Request){

	path := "/etc/foo"

	usernameEncoded, err := ioutil.ReadFile(path + "/username")

	if err != nil {
		fmt.Fprintf(w, "error reading username : %s", err.Error())
		return
	}

	username, _ := base64.StdEncoding.DecodeString(string(usernameEncoded))

	passwordEncoded , err := ioutil.ReadFile(path + "/password")

	if err != nil {
		fmt.Fprintf(w, "error reading password : %s", err.Error())
		return
	}

	password, _ := base64.StdEncoding.DecodeString(string(passwordEncoded))

	otherEncoded, err := ioutil.ReadFile(path + "/another-secret")

	if err != nil {
		fmt.Fprintf(w, "error reading another-secret : %s", err.Error())
		return
	}

	other, _ := base64.StdEncoding.DecodeString(string(otherEncoded))

	fmt.Fprintf(w, "ssh!, I'm running on %s. My secrets are username : %s, password : %s, other : %s", 
		runtime.GOOS,string(username), string(password), string(other))
}

func main(){
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080",nil)
}