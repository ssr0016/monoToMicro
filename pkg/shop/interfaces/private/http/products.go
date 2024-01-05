package http

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/ssr0016/monolithToMicro/pkg/common/price"
)

func AddRoutes(router *chi.Mux, repo products_domain.Repository) {
	resource := productsResource{repo}
	router.Get("/products/{id}", resource.Get)
}

type productsResource struct {
	repo products_domain.Repository
}

type Priceview struct {
	Cents    uint   `json:"cents"`
	Currency string `json:"currency"`
}

func priceViewFromPrice(p price.Price) Priceview {
	return Priceview{
		p.Cents(),
		p.Currency(),
	}
}

func (p productsResource) Get(w http.ResponseWriter, r *http.Request) {
	product, err := p.repo.ByID(products_domain.ID(chi.URLParam(r, "id")))
	if err != nil {
		_ = render.Render(w, r, common_http.ErrInternal(err))
		return
	}

	render.Respond(w, r, ProductView{
		string(product.ID()),
		product.Name(),
		product.Description(),
		priceViewFromPrice(product.Price()),
	})
}
