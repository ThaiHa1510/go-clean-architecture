package valueobjects

type Address struct{
	Address1 string
	Address2 string
	District uint64
	Province uint64
	Country  uint64
}

func (addr *Address) GetFullAddress() string {
	return addr.Address1 + " "  + addr.Address2
}

func (addr *Address) GetProvinceName() string {
	Redis.Get("province-"+addr.Province)
}

func (addr *Address) GetDistrictName() string {
	Redis.Get("distrince-"+addr.District)
}

func (addr *Address) GetCountruyName() string {
	Redis.Get("country-"+addr.Country)
}
