package registration

import (
	"net/http"

	"github.com/egoholic/router"
	"github.com/egoholic/router/params"
)

func Extend(n *router.Node) (err error) {
	n.GET(RenderRegistrationForm, "renders signup form")
	n.POST(PerformRegistration, "registers new company and user")
	return
}

func RenderRegistrationForm(w http.ResponseWriter, r *http.Request, params *params.Params) {

}

func PerformRegistration(w http.ResponseWriter, r *http.Request, params *params.Params) {

}
