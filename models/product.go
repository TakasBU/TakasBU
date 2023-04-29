package models

import (
	"gorm.io/gorm"
)

type Product struct {
	ID          uint   `gorm:"primarykey"`
	Productname string `json:"productname" form:"productname" query:"productname"`
	Description string `json:"description" form:"description" query:"description"`
	Owner       uint   `json:"owner" form:"owner" query:"owner"`
}
type Owner struct {
	Owner string
}

// TODO GRUP SİSTEMİ GETİRELECEK CHAT CHANNEL YAPILACAK ÜLKEDE KİMSE AÇ KALMAYACAK FOLLOWER SİSTEMİ
func GetProducts(db *gorm.DB, Product *[]Product) (err error) {
	err = db.Table("Products").Find(Product).Error
	if err != nil {
		return err
	}
	return nil
}
func CreateProduct(db *gorm.DB, Product *Product) (err error) {
	err = db.Table("Products").Create(Product).Error
	if err != nil {
		return err
	}
	return nil
}

func GetProductById(db *gorm.DB, Product *Product, id int) (err error) {
	err = db.Table("Products").Where("id = ?", id).First(Product).Error
	if err != nil {
		return err
	}
	return nil
}

func GetProductByOwner(db *gorm.DB, Product *Product, owner int) (err error) {
	err = db.Table("Products").Where("id = ?", owner).First(Product).Error
	if err != nil {
		return err
	}
	return nil
}
func GetProductsOwner(db *gorm.DB, Owner *Owner, id int) (err error) {
	err = db.Table("Products").Select("Owner").Where("id = ?", id).First(Owner).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateProduct(db *gorm.DB, Product *Product) (err error) {
	db.Table("Products").Save(Product)
	return nil
}

func DeleteProduct(db *gorm.DB, Product *Product, id int) (err error) {
	db.Table("Products").Where("id = ?", id).Delete(Product)
	return nil
}

func TradeProduct(db *gorm.DB, Product *Product, id int, id2 int) (err error) {

	DeleteProduct(db, Product, id)
	DeleteProduct(db, Product, id2)

	return nil
}
