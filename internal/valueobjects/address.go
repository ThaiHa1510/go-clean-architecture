package valueobjects

import (
  "fmt"
)

type Address struct{
  ProvinceId int64
  ProvinceName string
  DistrictId int64
  DistrictName string
  Stress string
  CountryId int64
  ConuntryName string
}

func NewAddress(int64 provinceId ,int64 districtId  , string stress){
  return Address{
    ProinceId proniceId
    DistrictId districtId
    Stress stress
  }
} 
