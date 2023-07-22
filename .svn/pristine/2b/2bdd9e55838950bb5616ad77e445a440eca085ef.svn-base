package stockRepository

import (
	"strconv"

	"github.com/rals/dearme-channel/config"
	"github.com/rals/dearme-channel/enums"
	"github.com/rals/dearme-channel/models"
)

func CariSkuParentbychannel(param string, param2 string) ([]models.TableSkuMapping, error) {
	var obj []models.TableSkuMapping

	db := config.ConnectionPg()
	db.Exec(`set search_path='public'`)
	wms_sku_mapping := enums.SCHEMA_PUBLIC + "wms_sku_mapping"

	err := db.Table(wms_sku_mapping).
		Select(wms_sku_mapping+".sku_no,"+wms_sku_mapping+".id_sku_parent,"+wms_sku_mapping+".id_sku_child1,"+wms_sku_mapping+".channel_code, "+wms_sku_mapping+".uuid_status ").
		Where(wms_sku_mapping+".sku_no = ?  and "+wms_sku_mapping+".channel_code = ?  and "+wms_sku_mapping+".uuid_status = ? ", param, param2, "3936a00a-e321-f5ff-33b5-2dfb8d68ba45").
		// Order(wms_sku_mapping + ".id_sku_child1 ").
		Order(wms_sku_mapping + ".channel_code desc , " + wms_sku_mapping + ".created_date ").
		Find(&obj).Error
	return obj, err
}

func CariSkuParent(param string) ([]models.TableSkuMapping, error) {
	var obj []models.TableSkuMapping

	db := config.ConnectionPg()
	db.Exec(`set search_path='public'`)
	wms_sku_mapping := enums.SCHEMA_PUBLIC + "wms_sku_mapping"

	err := db.Table(wms_sku_mapping).
		Select(wms_sku_mapping+".sku_no,"+wms_sku_mapping+".id_sku_parent,"+wms_sku_mapping+".id_sku_child1,"+wms_sku_mapping+".channel_code , "+wms_sku_mapping+".uuid_status ").
		Where(wms_sku_mapping+".sku_no = ? and "+wms_sku_mapping+".uuid_status =?  ", param, "3936a00a-e321-f5ff-33b5-2dfb8d68ba45").
		// Order(wms_sku_mapping + ".id_sku_child1 ").
		Order(wms_sku_mapping + ".channel_code desc , " + wms_sku_mapping + ".created_date ").
		Find(&obj).Error
	return obj, err
}

func CariSkuBlock(param string, param2 string, param3 string) models.TableSkuMapping {
	var obj models.TableSkuMapping

	db := config.ConnectionPg()
	db.Exec(`set search_path='public'`)
	wms_sku_block := enums.SCHEMA_PUBLIC + "wms_sku_block_channel"

	where_tgl := ` and (date_from between '` + param3 + `' and '` + param3 + `' or
	date_to between '` + param3 + `' and '` + param3 + `' or
	'` + param3 + `' between date_from and date_to or
	'` + param3 + `' between date_from and date_to)`

	db.Table(wms_sku_block).
		Where(wms_sku_block+".sku_no = ? and "+wms_sku_block+".channel_code = ? and "+wms_sku_block+".uuid_status = ?  "+where_tgl, param, param2, "3936a00a-e321-f5ff-33b5-2dfb8d68ba45").
		Order(wms_sku_block + ".created_date desc ").
		First(&obj)
	return obj
}

func CekChannelSku(param string, param2 string) models.TableSkuMapping {
	var obj models.TableSkuMapping

	db := config.ConnectionPg()
	db.Exec(`set search_path='public'`)
	wms_channel := enums.SCHEMA_PUBLIC + "wms_channel"
	wms_sku_mapping := enums.SCHEMA_PUBLIC + "wms_sku_mapping"

	db.Table(wms_sku_mapping).
		Select(wms_channel+".uuid_status ").
		Joins(" JOIN "+wms_channel+" ON "+wms_channel+".channel_code = "+wms_sku_mapping+".channel_code ").
		Where(wms_sku_mapping+".sku_no = ? and "+wms_sku_mapping+".channel_code = ?  ", param, param2).
		First(&obj)
	return obj
}

func CariSkuStockWMS(param string) models.TableSkuStock {
	var obj models.TableSkuStock
	db := config.ConnectionPg()

	db.Exec(`set search_path='public'`)
	wms_sku_stock := enums.SCHEMA_PUBLIC + "wms_sku_stock"

	db.Table(wms_sku_stock).
		Where(wms_sku_stock+".sku_no = ? ", param).
		First(&obj)
	return obj
}

func SaveSkuStock(ObjSkuStockIns models.TableSkuStock, ObjSkuStockUpd models.TableSkuStock) error {

	tx := config.ConnectionPg()

	db := tx.Begin()

	// db := db.CreateCon().Begin()

	defer func() {
		if r := recover(); r != nil {
			db.Rollback()
		}
	}()

	if err := db.Error; err != nil {
		return err
	}

	//set schema postgresql
	db.Exec(`set search_path='public'`)

	wms_sku_stock := enums.SCHEMA_PUBLIC + "wms_sku_stock"

	if ObjSkuStockIns.UuidSkuStock != "" {

		if err := db.Table(wms_sku_stock).Create(&ObjSkuStockIns).Error; err != nil {
			db.Rollback()
			return err
		}

	} else {

		if err := db.Table(wms_sku_stock).
			Where(wms_sku_stock+".uuid_sku_stock = ? ", ObjSkuStockUpd.UuidSkuStock).
			Update(&ObjSkuStockUpd).Error; err != nil {
			db.Rollback()
			return err
		}
	}

	result := db.Commit().Error

	return result

}

func CariUidAvlbBuffer(param string) (int, error) {
	type Totals struct {
		Total string `json:"total"`
	}
	var obj Totals

	db := config.ConnectionPg()
	db.Exec(`set search_path='public'`)
	wms_putway := enums.SCHEMA_PUBLIC + "wms_putway"
	wms_product_racking := enums.SCHEMA_PUBLIC + "wms_product_racking"
	wms_product_uid := enums.SCHEMA_PUBLIC + "wms_product_uid"
	//wms_racking := enums.SCHEMA_PUBLIC + "wms_racking"
	wms_status := enums.SCHEMA_PUBLIC + "wms_status"

	err := db.Table(wms_putway).
		Select(wms_product_uid+".sku_no,count(0) total ").
		Joins("JOIN "+wms_product_racking+" ON "+wms_product_racking+".uuid_product_racking = "+wms_putway+".uuid_product_racking ").
		Joins("JOIN "+wms_product_uid+" ON "+wms_product_uid+".uuid_product_uid = "+wms_product_racking+".uuid_product_uid ").
		//Joins("JOIN "+wms_racking+" ON "+wms_racking+".uuid_racking_no = "+wms_putway+".uuid_racking_no ").
		Where(wms_product_uid+".uuid_status in ( select ws.uuid_status from "+wms_status+" ws where status_code='AVLB') and "+wms_putway+".uuid_status in ( select ws.uuid_status from "+wms_status+" ws where status_code='PTWY') and "+wms_product_uid+".sku_no = ? ", param).
		Group(wms_product_uid + ".sku_no ").
		First(&obj).Error

	count := 0
	if obj.Total != "" {
		count, _ = strconv.Atoi(obj.Total)
	}
	return count, err
}

func CariBufferStock(param string) (models.TableBufferStock, error) {
	var obj models.TableBufferStock

	db := config.ConnectionPg()
	db.Exec(`set search_path='public'`)
	wms_buffer_stock := enums.SCHEMA_PUBLIC + "wms_buffer_stock"

	err := db.Table(wms_buffer_stock).
		Where(wms_buffer_stock+".sku_no = ?", param).
		First(&obj).Error
	return obj, err
}

func GroupChannelSkuMap(param string) ([]models.TableSkuMapping, error) {
	var obj []models.TableSkuMapping

	db := config.ConnectionPg()
	db.Exec(`set search_path='public'`)
	wms_sku_mapping := enums.SCHEMA_PUBLIC + "wms_sku_mapping"

	err := db.Table(wms_sku_mapping).
		Select(wms_sku_mapping+".sku_no,"+wms_sku_mapping+".channel_code").
		Where(wms_sku_mapping+".sku_no = ? and "+wms_sku_mapping+".uuid_status =?  ", param, "3936a00a-e321-f5ff-33b5-2dfb8d68ba45").
		Group(wms_sku_mapping + ".sku_no," + wms_sku_mapping + ".channel_code").
		Find(&obj).Error
	return obj, err
}
