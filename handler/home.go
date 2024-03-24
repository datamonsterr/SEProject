package handler

import (
	"net/http"

	"example.com/test/view"
	"example.com/test/view/components"
)

func GetHomePage(w http.ResponseWriter, r *http.Request) {
	view.Index().Render(r.Context(), w)
}

func GetBookingView(w http.ResponseWriter, r *http.Request) {
	components.SeatEditor(12, 24).Render(r.Context(), w)
}
