package bookingRepository

import (
	"github.com/rals/dearme-channel/config"
	"github.com/rals/dearme-channel/enums"
	"github.com/rals/dearme-channel/models"
)

// config.ConnectionPg()
func GroupSku(param string) ([]models.GroupSkuPicking, error) {
	var objs []models.GroupSkuPicking
	var obj models.GroupSkuPicking

	query := `select wso.sku_no,sum(wso.qty) qty from ` + enums.SCHEMA_PUBLIC + `wms_sales_order wso 
	where wso.no_order ='` + param + `'
	group by wso.sku_no
	order by sku_no `

	db := config.ConnectionPg()
	//set schema postgresql
	db.Exec(`set search_path='public'`)

	objs = nil

	rows, err := db.Raw(query).Rows()
	defer rows.Close()
	if err != nil {
		return objs, nil
	}

	for rows.Next() {

		rows.Scan(&obj.SkuNo, &obj.Qty)

		objs = append(objs, obj)

	}

	return objs, nil
}
func DetailSudahPicking(no_order string, sku string) ([]models.TableProductUidBasket, error) {
	var obj []models.TableProductUidBasket

	db := config.ConnectionPg()
	db.Exec(`set search_path='public'`)
	// err := db.Table(enums.SCHEMA_PUBLIC+"wms_picking").
	err := db.Table(enums.SCHEMA_PUBLIC+"wms_product_uid").
		Select(enums.SCHEMA_PUBLIC+"wms_product_uid.*,"+enums.SCHEMA_PUBLIC+"wms_basket.basket_id,"+enums.SCHEMA_PUBLIC+"wms_product_sku.product_name ").
		Joins("JOIN "+enums.SCHEMA_PUBLIC+"wms_picking ON "+enums.SCHEMA_PUBLIC+"wms_picking.uuid_product_uid = "+enums.SCHEMA_PUBLIC+"wms_product_uid.uuid_product_uid ").
		Joins("JOIN "+enums.SCHEMA_PUBLIC+"wms_basket ON "+enums.SCHEMA_PUBLIC+"wms_basket.uuid_basket_id = "+enums.SCHEMA_PUBLIC+"wms_picking.uuid_basket_id ").
		Joins("JOIN "+enums.SCHEMA_PUBLIC+"wms_product_sku ON "+enums.SCHEMA_PUBLIC+"wms_product_sku.sku_no = "+enums.SCHEMA_PUBLIC+"wms_product_uid.sku_no ").
		//Joins("JOIN "+enums.SCHEMA_PUBLIC+"wms_sales_order ON "+enums.SCHEMA_PUBLIC+"wms_sales_order.no_order = "+enums.SCHEMA_PUBLIC+"wms_picking.no_order ").
		Where(enums.SCHEMA_PUBLIC+"wms_picking.flag_process_picking =? and "+enums.SCHEMA_PUBLIC+"wms_picking.no_order=? and "+enums.SCHEMA_PUBLIC+"wms_product_uid.sku_no =? ", "0", no_order, sku).
		Find(&obj).Error
	return obj, err
}
func SumQtySalesOrder(no_order string, sku string) (models.GroupSkuPicking, error) {
	var obj models.GroupSkuPicking

	db := config.ConnectionPg()
	db.Exec(`set search_path='public'`)
	// err := db.Table(enums.SCHEMA_PUBLIC+"wms_picking").
	err := db.Table(enums.SCHEMA_PUBLIC+"wms_sales_order").
		Select("sum(qty) as qty").
		Where(enums.SCHEMA_PUBLIC+"wms_sales_order.no_order =? and "+enums.SCHEMA_PUBLIC+"wms_sales_order.sku_no = ?", no_order, sku).
		Find(&obj).Error
	return obj, err
}
func ArrayKosongRacking() ([]models.TotalRacking, error) {
	var obj []models.TotalRacking

	db := config.ConnectionPg()
	db.Exec(`set search_path='public'`)
	wms_racking := enums.SCHEMA_PUBLIC + "wms_racking"

	db.Table(wms_racking).
		Where(wms_racking+".racking_used = ? ", "XXXX").
		Find(&obj).RecordNotFound()
	return obj, nil
}
func CekSalesBooking(no_order string) ([]models.TableSalesBooking, error) {
	var obj []models.TableSalesBooking
	db := config.ConnectionPg()
	db.Exec(`set search_path='public'`)
	wms_sales_booking := enums.SCHEMA_PUBLIC + "wms_sales_booking"

	db.Table(wms_sales_booking).
		Where(wms_sales_booking+".no_order =?  ", no_order).
		Find(&obj)

	return obj, nil
}
func CariTotalBooking(no_order string, sku_no string) ([]models.TotalRacking, error) {
	var objs []models.TotalRacking

	wms_sales_booking := enums.SCHEMA_PUBLIC + "wms_sales_booking"

	db := config.ConnectionPg()
	//set schema postgresql
	db.Exec(`set search_path='public'`)

	err := db.Table(wms_sales_booking).
		Select(wms_sales_booking+".racking_no, count(0) as total").
		Where(wms_sales_booking+".sku_no = ? and   "+wms_sales_booking+".no_order = ?", sku_no, no_order).
		Group(wms_sales_booking + ".racking_no").
		Find(&objs).Error

	return objs, err
}
func GetDetailOrder(param string, limit string, offset string, status_pick string) ([]models.DetailOrderPraPicking, error) {
	var objs []models.DetailOrderPraPicking
	var obj models.DetailOrderPraPicking
	status_product_uid := "AVLB"
	if status_pick == "interpid" {
		status_product_uid = "PTFO"

	}
	status_product_uid = "AVLB"

	if limit == "0" { //biar bisa ambil data
		status_product_uid = "PICK"
		limit = "1"
	} //end biar bisa ambil data

	query := `select wpu.uid,wr.racking_no,wps.sku_no,wps.product_name,rcv_no,rcv_date,wr.aisle_shelves from ` + enums.SCHEMA_PUBLIC + `wms_product_uid wpu 
	left join ` + enums.SCHEMA_PUBLIC + `wms_product_sku wps on wps.sku_no =wpu.sku_no  
	left join ` + enums.SCHEMA_PUBLIC + `wms_product_racking wpr on wpr.uuid_product_uid = wpu.uuid_product_uid 
	left join ` + enums.SCHEMA_PUBLIC + `wms_putway wp on wp.uuid_product_racking =wpr.uuid_product_racking 
	left join ` + enums.SCHEMA_PUBLIC + `wms_racking wr on wr.uuid_racking_no = wp.uuid_racking_no 
	where wpu.sku_no ='` + param + `'
	and wp.uuid_status in (select uuid_status from ` + enums.SCHEMA_PUBLIC + `wms_status where status_code='PTWY' ) --status putway yang selesai/sukses
	and wpu.uuid_status in (select uuid_status from ` + enums.SCHEMA_PUBLIC + `wms_status where status_code='` + status_product_uid + `' ) --status available

	order by rcv_date asc,wr.racking_no asc
	limit ` + limit + ` offset ` + offset + ` `
	//fmt.Println(query)
	db := config.ConnectionPg()
	//set schema postgresql
	db.Exec(`set search_path='public'`)

	objs = nil

	rows, err := db.Raw(query).Rows()
	defer rows.Close()
	if err != nil {
		return objs, nil
	}

	for rows.Next() {

		rows.Scan(&obj.Uid, &obj.RackingNo, &obj.SkuNo, &obj.ProductName, &obj.RcvNo, &obj.RcvDate, &obj.AisleShelves)
		obj.BasketId = ""
		objs = append(objs, obj)

	}

	return objs, nil
}
func CariTotal(param string, limit string, status_pick string) ([]models.TotalRacking, error) {
	var objs []models.TotalRacking
	var obj models.TotalRacking

	status_product_uid := "AVLB"
	if status_pick == "interpid" {
		status_product_uid = "PTFO"

	}
	status_product_uid = "AVLB"

	query := `select a.racking_no,count(0) as total from (
			select wr.aisle_shelves ,wpu.uid,wr.racking_no,wps.sku_no,wps.product_name,rcv_no,rcv_date,wr.aisle_shelves from ` + enums.SCHEMA_PUBLIC + `wms_product_uid wpu 
			left join ` + enums.SCHEMA_PUBLIC + `wms_product_sku wps on wps.sku_no =wpu.sku_no  
			left join ` + enums.SCHEMA_PUBLIC + `wms_product_racking wpr on wpr.uuid_product_uid = wpu.uuid_product_uid 
			left join ` + enums.SCHEMA_PUBLIC + `wms_putway wp on wp.uuid_product_racking =wpr.uuid_product_racking 
			left join ` + enums.SCHEMA_PUBLIC + `wms_racking wr on wr.uuid_racking_no = wp.uuid_racking_no 
			where wpu.sku_no ='` + param + `'
			and wp.uuid_status in (select uuid_status from ` + enums.SCHEMA_PUBLIC + `wms_status where status_code='PTWY' ) --status putway yang selesai/sukses
			and wpu.uuid_status in (select uuid_status from ` + enums.SCHEMA_PUBLIC + `wms_status where status_code='` + status_product_uid + `' ) --status available
		
			order by rcv_date asc,wr.racking_no asc
			limit ` + limit + ` ) a 
			group by a.racking_no
		
		 `
	//fmt.Println(query)
	db := config.ConnectionPg()
	//set schema postgresql
	db.Exec(`set search_path='public'`)

	objs = nil

	rows, err := db.Raw(query).Rows()
	defer rows.Close()
	if err != nil {
		return objs, nil
	}

	for rows.Next() {

		rows.Scan(&obj.RackingNo, &obj.Total)
		objs = append(objs, obj)

	}

	return objs, nil
}

func ViewCariTotal(param string, limit string, status_pick string) ([]models.DetailOrderPraPicking, error) {
	var objs []models.DetailOrderPraPicking

	status_product_uid := "AVLB"
	if status_pick == "interpid" {
		status_product_uid = "PTFO"

	}
	status_product_uid = "AVLB"

	wms_product_uid := enums.SCHEMA_PUBLIC + "wms_product_uid"
	wms_racking := enums.SCHEMA_PUBLIC + "wms_racking"
	wms_putway := enums.SCHEMA_PUBLIC + "wms_putway"
	wms_product_racking := enums.SCHEMA_PUBLIC + "wms_product_racking"
	wms_product_sku := enums.SCHEMA_PUBLIC + "wms_product_sku"
	wms_status := enums.SCHEMA_PUBLIC + "wms_status"

	db := config.ConnectionPg()
	//set schema postgresql
	db.Exec(`set search_path='public'`)

	err := db.Table(wms_product_uid).
		Select(wms_racking+".aisle_shelves, "+wms_product_uid+".uid, "+wms_racking+".racking_no, "+wms_product_sku+".product_name, "+wms_product_uid+".rcv_no, "+wms_product_uid+".rcv_date ").
		Joins("JOIN "+wms_product_sku+" ON "+wms_product_sku+".sku_no = "+wms_product_uid+".sku_no ").
		Joins("JOIN "+wms_product_racking+" ON "+wms_product_racking+".uuid_product_uid = "+wms_product_uid+".uuid_product_uid ").
		Joins("JOIN "+wms_putway+" ON "+wms_putway+".uuid_product_racking = "+wms_product_racking+".uuid_product_racking ").
		Joins("JOIN "+wms_racking+" ON "+wms_racking+".uuid_racking_no = "+wms_putway+".uuid_racking_no ").
		Where(wms_product_uid+".sku_no =  ? and "+wms_putway+".uuid_status in ( select uuid_status from "+wms_status+" where status_code='PTWY' ) and "+wms_product_uid+".uuid_status in ( select uuid_status from "+wms_status+" where status_code='"+status_product_uid+"'  ) ", param).
		Order(wms_product_uid + ".rcv_date asc,  " + wms_racking + ".racking_no asc ").
		Limit(limit).
		Find(&objs).Error

	return objs, err
}
func UidCariTotal(param string, limit string, status_pick string) ([]models.DetailOrderPraPicking, error) {
	var objs []models.DetailOrderPraPicking
	db := config.ConnectionPg()
	//set schema postgresql
	db.Exec(`set search_path='public'`)

	status_product_uid := "AVLB"
	if status_pick == "interpid" {
		status_product_uid = "PTFO"

	}
	status_product_uid = "AVLB"

	err := db.Table(enums.SCHEMA_PUBLIC+"wms_product_uid").
		Select(enums.SCHEMA_PUBLIC+"wms_product_uid.uid,"+enums.SCHEMA_PUBLIC+"wms_racking.racking_no,"+enums.SCHEMA_PUBLIC+"wms_product_sku.sku_no,"+enums.SCHEMA_PUBLIC+"wms_product_sku.product_name,"+enums.SCHEMA_PUBLIC+"wms_product_uid.rcv_no,"+enums.SCHEMA_PUBLIC+"wms_product_uid.rcv_date,"+enums.SCHEMA_PUBLIC+"wms_racking.aisle_shelves").
		Joins("JOIN "+enums.SCHEMA_PUBLIC+"wms_product_sku ON "+enums.SCHEMA_PUBLIC+"wms_product_sku.sku_no = "+enums.SCHEMA_PUBLIC+"wms_product_uid.sku_no ").
		Joins("JOIN "+enums.SCHEMA_PUBLIC+"wms_product_racking ON "+enums.SCHEMA_PUBLIC+"wms_product_racking.uuid_product_uid = "+enums.SCHEMA_PUBLIC+"wms_product_uid.uuid_product_uid ").
		Joins("JOIN "+enums.SCHEMA_PUBLIC+"wms_putway ON "+enums.SCHEMA_PUBLIC+"wms_putway.uuid_product_racking = "+enums.SCHEMA_PUBLIC+"wms_product_racking.uuid_product_racking ").
		Joins("JOIN "+enums.SCHEMA_PUBLIC+"wms_racking ON "+enums.SCHEMA_PUBLIC+"wms_racking.uuid_racking_no = "+enums.SCHEMA_PUBLIC+"wms_putway.uuid_racking_no ").
		Where(enums.SCHEMA_PUBLIC+"wms_product_uid.sku_no =? and "+enums.SCHEMA_PUBLIC+"wms_putway.uuid_status in (select uuid_status from "+enums.SCHEMA_PUBLIC+"wms_status where status_code='PTWY' ) and  "+enums.SCHEMA_PUBLIC+"wms_product_uid.uuid_status in (select uuid_status from "+enums.SCHEMA_PUBLIC+"wms_status where status_code='"+status_product_uid+"' ) ", param).
		//Group(enums.SCHEMA_PUBLIC + "wms_product_uid.uid," + enums.SCHEMA_PUBLIC + "wms_racking.racking_no," + enums.SCHEMA_PUBLIC + "wms_product_sku.sku_no," + enums.SCHEMA_PUBLIC + "wms_product_sku.product_name," + enums.SCHEMA_PUBLIC + "wms_product_uid.rcv_no," + enums.SCHEMA_PUBLIC + "wms_product_uid.rcv_date," + enums.SCHEMA_PUBLIC + "wms_racking.aisle_shelves").
		Order(enums.SCHEMA_PUBLIC + "wms_product_uid.rcv_date asc," + enums.SCHEMA_PUBLIC + "wms_racking.racking_no asc").
		Limit(limit).
		Find(&objs).Error

	return objs, err
}
func ArrayKosong() ([]models.TableProductUidBasket, error) {
	var obj []models.TableProductUidBasket

	db := config.ConnectionPg()
	db.Exec(`set search_path='public'`)
	wms_sales_order := enums.SCHEMA_PUBLIC + "wms_sales_order"

	db.Table(wms_sales_order).
		Where(wms_sales_order+".no_order = ? and "+wms_sales_order+".status_process_order = '0' ", "XXXX").
		Find(&obj).RecordNotFound()
	return obj, nil
}

func StoreBooking(objBooking []models.TableSalesBooking, objStatusProduct []models.UpdateProductByUid) error {
	tx := config.ConnectionPg()

	db := tx.Begin()
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

	//wms_sales_booking
	for _, objStatusProduct := range objStatusProduct {
		if err := db.Table(enums.SCHEMA_PUBLIC+"wms_product_uid").Where("uid = ?", objStatusProduct.Uid).
			Update(&objStatusProduct).Error; err != nil {
			db.Rollback()
			return err
		}
	}

	for _, objBooking := range objBooking {
		if err := db.Table(enums.SCHEMA_PUBLIC + "wms_sales_booking").Create(&objBooking).Error; err != nil {
			db.Rollback()
			return err
		}
	}

	result := db.Commit().Error
	return result

}
