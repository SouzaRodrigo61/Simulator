package router

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"strings"
)

// MARK: - Typing
type Router struct {
	ID 		 	string   `json:"routerId"`
	ClientID  	string   `json:"clientId"`
	Positions []Position `json:"positions"`
}

type Position struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}

type PartialRoutePosition struct {
	ID string 			`json:"routeId"`
	ClientID string 	`json:"clientId"`
	Position []float64  `json:"position"`
	Finished bool 		`json:"finished"`
}

// MARK: - Method

// NewRoute creates a *Route struct
func NewRoute() *Router {
	return &Router {}
}

// LoadPositions load all positions from file in path destinations
func (r *Router) LoadPositions() error {

	if r.ID == "" {
		return errors.New("router id not informed")
	}

	f, err := os.Open("destinations/" + r.ID + ".txt")
	if err != nil {
		return err
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		data := strings.Split(scanner.Text(), ",")

		lat, err := strconv.ParseFloat(data[0], 64)
		if err != nil {
			return err
		}

		long, err := strconv.ParseFloat(data[0], 64)
		if err != nil {
			return err
		}

		r.Positions = append(r.Positions, Position{
			Lat: lat,
			Long: long,
		})

	}

	return nil
}

// ExportJsonPositions transform all objects em json type
func (r *Router) ExportJsonPositions() ([]string, error) {
	var route PartialRoutePosition
	var result []string

	total := len(r.Positions)

	for k, v := range r.Positions {
		route.ID = r.ID
		route.ClientID = r.ClientID
		route.Position = []float64{v.Lat, v.Long}
		route.Finished = false

		if total-1 == k {
			route.Finished = true
		}

		jsonRoute, err := json.Marshal(route)
		if err != nil {
			return nil, err
		}

		result = append(result, string(jsonRoute))
	}

	return result, nil
}


