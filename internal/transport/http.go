package transport

import (
	"encoding/json"
	"net/http"

	"github.com/marktsarkov/avito/internal/storage"
)

// HttpServer is a HTTP server
type HttpServer struct {
	storage storage.SegmentStorage
}

// NewHttpServer creates a new HTTP server
func NewHttpServer(storage *storage.SegmentStorage) HttpServer {
	return HttpServer{
		storage: storage,
	}
}

// CreateSegment returns new segment
func (h HttpServer) CreateSegment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	var segment Segment
	err := json.NewDecoder(r.Body).Decode(&segment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.storage.CreateSegment(r.Context(), segment.Slug)

	w.WriteHeader(http.StatusCreated)

}

// DeleteSegment delete segment
func (h HttpServer) DeleteSegment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	var segment Segment
	err := json.NewDecoder(r.Body).Decode(&segment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//TODO: передать DELETE в бд

	w.WriteHeader(http.StatusNoContent)
}

func (h HttpServer) AddAndRemoveSegmentsOnUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	var segmentAdd Segment
	// var segmentRem Segment
	// var id UserID

	err := json.NewDecoder(r.Body).Decode(&segmentAdd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//TODO: передать POST и DELETE в бд

	w.WriteHeader(http.StatusNoContent)
}

func (h HttpServer) ShowUserSegments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	var id UserID

	err := json.NewDecoder(r.Body).Decode(&id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//TODO: передать GET команду в бд

	w.WriteHeader(http.StatusNoContent)
}
