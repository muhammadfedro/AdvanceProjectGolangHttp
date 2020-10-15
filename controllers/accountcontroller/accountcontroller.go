package accountcontroller

import (
	"net/http"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("mysession"))

func Login(respone http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	user := request.Form.Get("user")
	pass := request.Form.Get("pass")
	if user == "edo" && pass == "123" {
		sessions, _ := store.Get(request, "mysession")
		sessions.Values["user"] = user
		sessions.Save(request, respone)
		http.Redirect(respone, request, "/admin", http.StatusSeeOther)

	} else {

		http.Redirect(respone, request, "/salah", http.StatusSeeOther)

	}

}

func Logout(respone http.ResponseWriter, request *http.Request) {
	sessions, _ := store.Get(request, "mysession")
	sessions.Options.MaxAge = -1
	sessions.Save(request, respone)
	http.Redirect(respone, request, "/index", http.StatusSeeOther)
}
