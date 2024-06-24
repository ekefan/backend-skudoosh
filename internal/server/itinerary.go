package server 

import (
	// "database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ekefan/backend-skudoosh/internal/token"
	db "github.com/ekefan/backend-skudoosh/internal/db/sqlc"
	

	// "fmt"
	// "github.com/lib/pq"
)

// CreateTripRequest model for http post request for creating a trip
type CreateTripRequest struct {
	//Owner int64 `json:"owner"` //when authentication happens only authenticated users can creatTrips
	TakeOffDate time.Time 	`json:"take_off_date"`
	ReturnDate time.Time `json:"return_date"`
	Destination string `json:"destination"`
}

func (server *Server) createTrip(ctx *gin.Context) {
	var req CreateTripRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}


	authPayload  := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	args := db.CreateTripParams{
		Owner:  authPayload.UserID,
		TakeOffDate: req.TakeOffDate,
		ReturnDate: req.ReturnDate,
		Destination: req.Destination,
	}
	itinerary, err := server.store.CreateTrip(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	resp :=struct{
		TakeOffDate time.Time `json:"take_off_date"`
		ReturnDate time.Time `json:"return_date"`
		Destination string `json:"destination"`
		Wakanow string `json:"wakanow"`
	}{
		TakeOffDate: itinerary.TakeOffDate,
		ReturnDate: itinerary.ReturnDate,
		Destination: itinerary.Destination,
		Wakanow: server.config.Wakanow,
	}
	ctx.JSON(http.StatusOK, resp)
}

func (server *Server) deleteTrip(ctx *gin.Context){
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	err := server.store.DeleteItinerary(ctx, authPayload.UserID)
	if err != nil {
		// if err == sql.ErrNoRows {
		// 	ctx.JSON(http.StatusNotFound, errorResponse(err))
		// 	return
		// }
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}
}


func (server *Server) updateTrip(ctx *gin.Context){
	// authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	// trip, err := server.store.GetTripUpdate(ctx, authPayload.UserID)
	// if err != nil {
	// 	ctx.JSON(http.StatusOK, errorResponse(err))
	// 	return
	// }
	//create Transaction to update trip


	// TODO: create transaction to handle trip update
}

type ListTripsRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}
func (server *Server) getItineraries(ctx *gin.Context){
	var req ListTripsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	arg := db.ListTripsParams{
		Owner : authPayload.UserID,
		Limit: req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}
	//
	itineraries, err := server.store.ListTrips(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	ctx.JSON(http.StatusOK, itineraries)
}