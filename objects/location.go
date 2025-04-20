package objects

type LocationPointRequest struct {
	Name    string  `json:"name"`
	Lat     float64 `json:"lat"`
	Lng     float64 `json:"lng"`
	Address string  `json:"address"`
	Type    string  `json:"type"`
}
type WarehouseRequest struct {
	Name    string   `json:"name"`
	Lat     *float64 `json:"lat"`
	Lng     *float64 `json:"lng"`
	Address string   `json:"address"`
}
