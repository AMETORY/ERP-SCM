package objects

type LocationPointRequest struct {
	Name        string  `json:"name"`
	Lat         float64 `json:"latitude"`
	Lng         float64 `json:"longitude"`
	Address     string  `json:"address"`
	Description string  `json:"description"`
	ProvinceID  *string `json:"province_id"`
	RegencyID   *string `json:"regency_id"`
	DistrictID  *string `json:"district_id"`
	VillageID   *string `json:"village_id"`
	ZipCode     *string `json:"zip_code"`
	Type        string  `json:"type"`
}
type WarehouseRequest struct {
	Name    string   `json:"name"`
	Lat     *float64 `json:"lat"`
	Lng     *float64 `json:"lng"`
	Address string   `json:"address"`
}
