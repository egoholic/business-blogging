package authentication

import (
	"net/http"

	"github.com/egoholic/router"
	"github.com/egoholic/router/params"
)

func Extend(n *router.Node) (err error) {
	n.GET(RenderAuthenticationForm, "renders signin form")
	n.POST(PerformAuthentication, "authenticates user")
	return
}

func RenderAuthenticationForm(w http.ResponseWriter, r *http.Request, params *params.Params) {

}

func PerformAuthentication(w http.ResponseWriter, r *http.Request, params *params.Params) {

}
