package souvenir

// type GetHotelID struct {
// 	ID int `uri:"id" binding:"required"`
// }

type InputSouvenir struct {
	Lokasi string `json:"lokasi"`
	Rating float32 `json:"rating"`
}