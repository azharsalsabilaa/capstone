package cafe

// type GetHotelID struct {
// 	ID int `uri:"id" binding:"required"`
// }

type InputCafe struct {
	Lokasi string `json:"lokasi"`
	Rating float32 `json:"rating"`
}