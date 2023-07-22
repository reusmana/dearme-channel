package tokenService

import (

	// "strconv"
	// "strings"

	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/rals/dearme-channel/helpers"
	"github.com/rals/dearme-channel/models"

	//"github.com/rals/dearme-channel/models/response"

	"github.com/rals/dearme-channel/repositories/bookingRepository"
	"github.com/rals/dearme-channel/repositories/stockRepository"
	"github.com/rals/dearme-channel/repositories/tokenRepository"

	//"github.com/rals/dearme-channel/enums"
	//"github.com/rals/dearme-channel/utils"
	// "time"
	// "github.com/go-playground/validator"
	//"github.com/gin-gonic/gin"

	"github.com/google/uuid"
)

func SaveToken(id string, datas map[string]interface{}, process string) models.TableApiChannel {
	var objs models.TableApiChannel
	objs, _ = tokenRepository.GetDataToken(id)
	msgErr := ""
	for key, value := range datas {
		if key == "access_token" {
			objs.Key1 = "access_token"
			objs.Value1 = value
		}
		if key == "refresh_token" {
			objs.Key2 = "refresh_token"
			objs.Value2 = value
		}

		if key == "message" {
			msgErr = fmt.Sprintf("%v", value)
		}

	}

	objs.Id = id
	objs.CreatedDate = time.Now()
	if datas["message"] == "" {
		result := tokenRepository.SaveTokenShopee(objs)
		if result != nil {
			fmt.Println("GAGAL")
		} else {
			fmt.Println("Save Token Success")
		}
	} else {
		var objError models.TableLogErrChannel
		objError.Id = id
		objError.Message = msgErr
		objError.CreatedDate = time.Now()
		objError.Process = process

		result := tokenRepository.SaveChannelError(objError)
		if result != nil {
			fmt.Println("GAGAL")
		} else {
			fmt.Println("Save Error Success")
		}
	}

	return objs
}

func SaveTokenBukalapak(objToken models.TokenBukalapak, id string) models.TableApiChannel {
	var objs models.TableApiChannel

	objs, _ = tokenRepository.GetDataToken(id)
	msgErr := ""
	objs.Key1 = "access_token"
	objs.Value1 = objToken.AccessToken
	objs.Key2 = "refresh_token"
	objs.Value2 = objToken.RefreshToken

	objs.Id = id
	objs.CreatedDate = time.Now()
	if objs.Key1 != "" {
		result := tokenRepository.SaveTokenShopee(objs)
		if result != nil {
			fmt.Println("GAGAL")
		} else {
			fmt.Println("Save Token Success")
		}
	} else {
		msgErr = "Gagal Save token " + id
		var objError models.TableLogErrChannel
		objError.Id = id
		objError.Message = msgErr
		objError.CreatedDate = time.Now()
		objError.Process = "SaveToken"

		result := tokenRepository.SaveChannelError(objError)
		if result != nil {
			fmt.Println("GAGAL")
		} else {
			fmt.Println("Save Error Success")
		}
	}

	return objs
}

func SaveTokenJdId(objToken models.TokenJdId, id string) models.TableApiChannel {
	var objs models.TableApiChannel

	objs, _ = tokenRepository.GetDataToken(id)
	msgErr := ""
	objs.Key1 = "access_token"
	objs.Value1 = objToken.AccessToken
	objs.Key2 = "refresh_token"
	objs.Value2 = objToken.RefreshToken

	objs.Id = id
	objs.CreatedDate = time.Now()
	if objs.Key1 != "" {
		result := tokenRepository.SaveTokenShopee(objs)
		if result != nil {
			fmt.Println("GAGAL")
		} else {
			fmt.Println("Save Token Success")
		}
	} else {
		msgErr = "Gagal Save token " + id
		var objError models.TableLogErrChannel
		objError.Id = id
		objError.Message = msgErr
		objError.CreatedDate = time.Now()
		objError.Process = "SaveToken"

		result := tokenRepository.SaveChannelError(objError)
		if result != nil {
			fmt.Println("GAGAL")
		} else {
			fmt.Println("Save Error Success")
		}
	}

	return objs
}

func FindToken(id string) models.TableApiChannel {
	var objs models.TableApiChannel
	objs, _ = tokenRepository.GetDataToken(id)
	return objs
}

func CekError(id string, datas map[string]interface{}, process string) models.TableApiChannel {
	objs, _ := tokenRepository.GetDataToken(id)
	msgErr := ""
	for key, value := range datas {

		if key == "message" {
			msgErr = fmt.Sprintf("%v", value)
		}

	}

	if datas["message"] != "" {
		var objError models.TableLogErrChannel
		objError.Id = id
		objError.Message = msgErr
		objError.CreatedDate = time.Now()
		objError.Process = process
		result := tokenRepository.SaveChannelError(objError)
		if result != nil {
			fmt.Println("GAGAL SAVE ERROR")
		}
	}

	return objs
}
func SaveErrorString(id string, pesan string, process string) {
	msgErr := pesan

	if msgErr != "" {
		var objError models.TableLogErrChannel
		objError.Id = id
		objError.Message = msgErr
		objError.CreatedDate = time.Now()
		objError.Process = process
		result := tokenRepository.SaveChannelError(objError)
		if result != nil {
			fmt.Println("GAGAL SAVE ERROR " + id)
		}
	}
}

func SaveSalesOrder(id string, obj models.Header, objpayment models.DetailEscrowsShopee) {

	var objSalesOrder []models.TableSalesOrder
	var objSalesOrderTotalQty models.TableSalesOrder
	createdDate := time.Now()
	totalQty := 0

	var ArrCek []models.ObjSkuZalora

	//cek jika orderan sudah ada di salesorder
	objCek, _ := tokenRepository.FindSalesOrder(obj.Response.OrderList[0].OrderSn)
	if objCek.NoOrder == "" && obj.Response.OrderList[0].ShippingCarrier != "" {
		//ambil channel code shopeee ke wms_channel
		objChannel, _ := tokenRepository.FindChannel("SHOPEE")
		if objChannel.UuidChannel != "" {
			fmt.Println(obj.Response.OrderList[0].OrderSn)

			for _, value := range obj.Response.OrderList[0].ItemList {
				//if obj.Response.OrderList[0].OrderSn == "211125G0P8NJ9X" { //orderan 1 dulu
				id := uuid.New()
				fmt.Println("masuk sini")
				qty := strconv.Itoa(int(value.ModelQuantityPurchased))

				totalQty += int(value.ModelQuantityPurchased)
				skunya := value.ModelSku
				if skunya == "" {
					skunya = value.ItemSku
				}

				fmt.Println("+++++++++++++++++++++++++")
				fmt.Println(objpayment.Response.OrderIncome.Items)
				fmt.Println(skunya)
				fmt.Println("+++++++++++++++++++++++++")
				//cari amount per sku
				amount := 0.0
				for _, valueAmount := range objpayment.Response.OrderIncome.Items {
					if valueAmount.ItemSku == skunya {
						amount = valueAmount.DiscountedPrice

						fmt.Println("amountnyaa " + strconv.Itoa(int(amount)))
						break
					}

				}

				fmt.Println("amountnyaa akhir " + strconv.Itoa(int(amount)))

				// fmt.Println(obj.Response.OrderList[0].OrderSn)
				// fmt.Println(value.ModelId)

				if obj.Response.OrderList[0].OrderSn == "220808JWUY0420" && value.ItemId == 17943782177 {
					skunya = "08162381"
				}

				if obj.Response.OrderList[0].OrderSn == "230126CFFK244V" && value.ItemId == 3762493907 {
					skunya = "08127278"
				}

				objSkuPengganti := tokenRepository.ChangeSkuChannelWMS(objChannel.ChannelCode, obj.Response.OrderList[0].OrderSn, skunya)
				if len(objSkuPengganti) > 0 {

					for _, valueGanti := range objSkuPengganti {
						if valueGanti.Key1 == strconv.Itoa(int(value.ItemId)) && valueGanti.Key2 == strconv.Itoa(int(value.ModelId)) {
							fmt.Println("ada pengganti")
							skunya = valueGanti.SkuReplace
						}
					}
				}

				//sku := []rune(skunya)
				//skunya = string(sku[1:9])
				if len([]rune(skunya)) == 8 {
					//skunya = skunya
				} else if len([]rune(skunya)) > 8 {
					sku := []rune(skunya)
					skunya = string(sku[1:9])
				} else if len([]rune(skunya)) == 7 {
					skunya = "0" + skunya
				} else {
					skunya = "XXXXXXXX"
				}

				cekSku := ""
				if len(ArrCek) > 0 {

					// for _, valCek := range ArrCek {
					// 	if valCek.Sku == skunya {
					// 		cekSku = "ada"
					// 		break
					// 	}

					// }

				}

				if cekSku == "" {
					ArrCekLoop := models.ObjSkuZalora{
						Sku: skunya,
					}

					ArrCek = append(ArrCek, ArrCekLoop)

					//if obj.Response.OrderList[0].CreateTime != 0 {

					sendBeforeDate := createdDate.AddDate(0, 0, +int(obj.Response.OrderList[0].DaysToShip))

					//}
					// fmt.Println("=========== BEFORE DATE ===========")
					// fmt.Println(sendBeforeDate)
					// fmt.Println(sendBeforeDate)
					// fmt.Println("=========== BEFORE DATE ===========")

					// amount := value.ModelDiscountedPrice
					// if amount < 1 {
					// 	amount = value.ModelOriginalPrice
					// }

					objSalesOrderLoop := models.TableSalesOrder{
						UuidSalesOrder:     id.String(),
						NoOrder:            obj.Response.OrderList[0].OrderSn,
						StatusOrder:        "PDKM",
						SkuNo:              skunya,
						ProductName:        value.ItemName,
						Qty:                qty,
						ExpeditionType:     obj.Response.OrderList[0].ShippingCarrier,
						DeliveryType:       "Pickup",
						TotalQtyOrder:      "30",
						StatusProcessOrder: "0",
						WeightProduk:       "0",
						TotalWeight:        "0",

						SentBeforeDate: sendBeforeDate,
						CreatedDate:    createdDate,

						OrderDate: time.Unix(obj.Response.OrderList[0].CreateTime, 0),

						CreatedBy:     "SHOPEE",
						NameVariant:   value.ModelName,
						UsernameBuyer: obj.Response.OrderList[0].BuyerUsername,
						ChannelCode:   objChannel.ChannelCode,

						RecipientsName:  obj.Response.OrderList[0].RecipientAddress.Name,
						Telephone:       obj.Response.OrderList[0].RecipientAddress.Phone,
						ShippingAddress: obj.Response.OrderList[0].RecipientAddress.FullAddress,
						City:            obj.Response.OrderList[0].RecipientAddress.City,
						Province:        obj.Response.OrderList[0].RecipientAddress.State,
						Amount:          amount,
					}
					objSalesOrder = append(objSalesOrder, objSalesOrderLoop)

					//}
				} //if cekSku == ""

			}

			objSalesOrderTotalQty.NoOrder = obj.Response.OrderList[0].OrderSn
			objSalesOrderTotalQty.TotalQtyOrder = strconv.Itoa(totalQty)

			result := tokenRepository.SaveSalesOrder(objSalesOrder, objSalesOrderTotalQty)
			if result != nil {
				fmt.Println("SaveSalesOrder Gagal")
				SaveErrorInfo("SaveSalesOrder", "SHOPEE", objSalesOrderTotalQty.NoOrder+" "+result.Error())
			} else {
				fmt.Println("SaveSalesOrder Success")
				pesan := BookingSku(objSalesOrderTotalQty.NoOrder)
				if pesan != "Sukses" {
					//delete sales order
					//tokenRepository.DeleteSalesOrder(objSalesOrderTotalQty)

				}
				fmt.Println(objSalesOrderTotalQty.NoOrder + " => " + pesan)

				//cari detail no order
				UpdateStockOrder(objSalesOrderTotalQty.NoOrder)

			}
		} else {
			fmt.Println("Channel shopee Tidak Ditemukan")
		}
	} else {

		//cek jika kurir ganti update kurir
		var objUpdateKurir models.TableSalesOrder
		if objCek.NoOrder == obj.Response.OrderList[0].OrderSn {
			if objCek.ExpeditionType != obj.Response.OrderList[0].ShippingCarrier {
				objUpdateKurir.NoOrder = objCek.NoOrder
				objUpdateKurir.ExpeditionType = obj.Response.OrderList[0].ShippingCarrier
				result := tokenRepository.UpdateKurir(objUpdateKurir)
				if result != nil {
					fmt.Println("UpdateKurir Gagal")
				} else {
					fmt.Println("UpdateKurir Success")
				}

			}
		}

		if objCek.StatusProcessOrder == "0" {
			//jika CANCELED UPDATE status_processs_order jadi 3
			if obj.Response.OrderList[0].OrderStatus == "CANCELLED" {
				objSalesOrderLoop := models.TableSalesOrder{
					NoOrder:            obj.Response.OrderList[0].OrderSn,
					StatusProcessOrder: "3",
					CreatedBy:          "SHOPEE",
				}
				result := tokenRepository.UpdateSalesOrderCanc(objSalesOrderLoop)
				if result != nil {
					fmt.Println("UpdateSalesOrderCanc Gagal")
				} else {
					fmt.Println("UpdateSalesOrderCanc Success")
				}
				UpdateStockOrder(obj.Response.OrderList[0].OrderSn)

			}
		} else if objCek.StatusProcessOrder == "1" {
			//jika complete update ke pyoutdate
			if obj.Response.OrderList[0].OrderStatus == "COMPLETED" {
				payoutdate := time.Unix(obj.Response.OrderList[0].UpdateTime, 0)
				objSalesOrderLoop := models.TableSalesOrder{
					NoOrder:    obj.Response.OrderList[0].OrderSn,
					PayoutDate: payoutdate,
					CreatedBy:  "SHOPEE",
				}
				result := tokenRepository.UpdateSalesOrderCanc(objSalesOrderLoop)
				if result != nil {
					fmt.Println("Update Payoutdate Gagal")
				} else {
					fmt.Println("Update Payoutdate Success")
				}
				fmt.Println("======= PAYOUTDATE =======")
				fmt.Println(payoutdate)
				fmt.Println("======= PAYOUTDATE =======")
			}

		}

		fmt.Println("========================")
		fmt.Println("No Order" + obj.Response.OrderList[0].OrderSn + " Sudah Ada")
		fmt.Println(obj.Response.OrderList[0].ItemList[0].ModelDiscountedPrice)
		fmt.Println(obj.Response.OrderList[0].ItemList[0].ModelOriginalPrice)
		fmt.Println("========================")
	}

}

func UpdateAmountShopee(obj models.DetailEscrowsShopee) {

	NoOrder := obj.Response.OrderSn

	objs := obj.Response.OrderIncome.Items
	for _, value := range objs {

		skunya := value.ModelSku

		if len([]rune(skunya)) == 8 {
			//skunya = skunya
		} else if len([]rune(skunya)) > 8 {
			sku := []rune(skunya)
			skunya = string(sku[1:9])
		} else if len([]rune(skunya)) == 7 {
			skunya = "0" + skunya
		} else {
			skunya = "XXXXXXXX"
		}

		Amounts := value.DiscountedPrice

		objSalesOrderLoop := models.TableSalesOrder{
			NoOrder:   NoOrder,
			Amount:    Amounts,
			CreatedBy: "SHOPEE",
			SkuNo:     skunya,
		}
		//result := tokenRepository.UpdateSalesOrderCanc(objSalesOrderLoop)

		//fmt.Println(objSalesOrderLoop)
		result := tokenRepository.UpdateSalesOrderAmount(objSalesOrderLoop)
		if result != nil {
			fmt.Println("UpdateSalesOrderAmount Shopee Gagal")
		} else {
			fmt.Println("UpdateSalesOrderAmount Shopee Success")
		}

	}

}

func CekArrayOrder(obj models.SuccessResponseItemZalora) []models.ObjSkuZalora {

	var objArray []models.ObjSkuZalora

	for _, value := range obj.SuccessResponse.Body.OrderItems.OrderItem {

		objArrayLoop := models.ObjSkuZalora{
			Sku:              value.Sku,
			Qty:              1,
			Name:             value.Name,
			Variation:        value.Variation,
			ShipmentProvider: value.ShipmentProvider,
			CreatedAt:        value.CreatedAt,
			OrderId:          value.OrderId,
		}

		if len(objArray) < 1 {
			objArray = append(objArray, objArrayLoop)
		} else {

			status := ""

			for indexs, _ := range objArray {
				if objArray[indexs].Sku == value.Sku {
					status = "ada"
					qtynew := objArray[indexs].Qty + 1
					objArray[indexs].Qty = qtynew
					objArray[indexs].Sku = value.Sku
				}
			}

			if status == "" {
				objArray = append(objArray, objArrayLoop)
			}

		}

	}

	return objArray

}

func CekArrayOrderV2(obj models.OrderDetailZaolraV2) []models.ObjSkuZalora {

	var objArray []models.ObjSkuZalora

	for _, value := range obj.Items {

		//if value.Status == "pending" {

		amounts := 0.0

		amounts = float64(value.UnitPrice)

		objArrayLoop := models.ObjSkuZalora{
			Sku:              value.Product.SellerSku,
			Qty:              1,
			Name:             value.Product.Name,
			Variation:        value.Product.Variation,
			ShipmentProvider: value.Shipment.Provider.Name,
			CreatedAt:        value.CreatedAt,
			OrderId:          strconv.Itoa(value.OrderID),
			Status:           value.Status,
			Amount:           amounts,
		}

		if len(objArray) < 1 {
			objArray = append(objArray, objArrayLoop)
		} else {

			status := ""

			for indexs, _ := range objArray {
				if objArray[indexs].Sku == value.Product.SellerSku {
					status = "ada"
					qtynew := objArray[indexs].Qty + 1
					//amountnew := objArray[indexs].Amount + float64(value.UnitPrice)
					objArray[indexs].Qty = qtynew
					objArray[indexs].Sku = value.Product.SellerSku
					//objArray[indexs].Amount = amountnew
				}
			}

			if status == "" {
				objArray = append(objArray, objArrayLoop)
			}

		}
		//}

	}

	return objArray

}

func UpdateAmountZalora(noorder string, obj models.OrderDetailZaolraV2) []models.ObjSkuZalora {
	objNew := CekArrayOrderV2(obj)

	for _, value := range objNew {

		objSalesOrderLoop := models.TableSalesOrder{
			NoOrder:   noorder,
			SkuNo:     value.Sku,
			Amount:    value.Amount * float64(value.Qty),
			CreatedBy: "ZALORA",
		}
		result := tokenRepository.UpdateSalesOrderAmount(objSalesOrderLoop)
		if result != nil {
			fmt.Println("UpdateSalesOrderAmount ZALORA Gagal")
		} else {
			fmt.Println("UpdateSalesOrderAmount ZALORA Success")
		}

	}

	return objNew
}

func UpdateAmountLazada(obj models.OrderDetailItemLazada, objCust models.OrdersDetailGenerated) {

	objNew := CekArrayOrderLazada(obj)

	for _, value := range objNew {
		fmt.Println("masuk dalem sinii UpdateAmountLazada")
		NoOrder := strconv.Itoa(int(objCust.Data.OrderNumber))

		skunya := value.Sku

		// if value.Sku == "a08172335" {
		// 	skunya = "a08172336"
		// }

		if len([]rune(skunya)) == 8 {
			//skunya = skunya
		} else if len([]rune(skunya)) > 8 {
			sku := []rune(skunya)
			skunya = string(sku[1:9])
		} else if len([]rune(skunya)) == 7 {
			skunya = "0" + value.Sku
		} else {
			skunya = "XXXXXXXX"
		}

		amount := value.Amount
		objSalesOrderLoop := models.TableSalesOrder{
			NoOrder:   NoOrder,
			SkuNo:     skunya,
			Amount:    amount * float64(value.Qty),
			CreatedBy: "LAZADA",
		}
		result := tokenRepository.UpdateSalesOrderAmount(objSalesOrderLoop)
		if result != nil {
			fmt.Println("UpdateSalesOrderAmount LAZADA Gagal")
		} else {
			fmt.Println("UpdateSalesOrderAmount LAZADA Success")
		}

	}

}

func SaveSalesOrderZalora(id string, obj models.SuccessResponseItemZalora, dataCust models.OrdersListZalora) {

	var objSalesOrder []models.TableSalesOrder
	var objSalesOrderTotalQty models.TableSalesOrder
	createdDate := time.Now()
	objZalora := obj.SuccessResponse.Body.OrderItems.OrderItem
	objNew := CekArrayOrder(obj)

	for _, value := range objNew {

		NoOrder := dataCust.Order.OrderNumber + "-" + value.OrderId

		objCek, _ := tokenRepository.FindSalesOrder(NoOrder)
		if objCek.NoOrder == "" { //&& obj.Response.OrderList[0].ShippingCarrier != "" {
			//ambil channel code shopeee ke wms_channel
			objChannel, _ := tokenRepository.FindChannel("ZALORA")
			if objChannel.UuidChannel != "" {
				skunya := value.Sku

				ProductName := value.Name
				Variant := value.Variation
				Kurir := value.ShipmentProvider
				OrderDate := value.CreatedAt

				FirstName := dataCust.Order.AddressShipping.FirstName
				LastName := dataCust.Order.AddressShipping.LastName
				Address := dataCust.Order.AddressShipping.Address1
				if Address == "" {
					Address = dataCust.Order.AddressShipping.Address2
				}
				if Address == "" {
					Address = dataCust.Order.AddressShipping.Address3
				}
				if Address == "" {
					Address = dataCust.Order.AddressShipping.Address4
				}

				Phone := dataCust.Order.AddressShipping.Phone
				if Phone == "" {
					Phone = dataCust.Order.AddressShipping.Phone2
				}

				City := dataCust.Order.AddressShipping.City
				Region := dataCust.Order.AddressShipping.Region
				CustFirstName := dataCust.Order.CustomerFirstName

				id := uuid.New()

				if len([]rune(skunya)) == 8 {
					//skunya = skunya
				} else if len([]rune(skunya)) > 8 {
					sku := []rune(skunya)
					skunya = string(sku[1:9])
				} else if len([]rune(skunya)) == 7 {
					skunya = "0" + value.Sku
				} else {
					skunya = "XXXXXXXX"
				}

				sendBeforeDate := createdDate.AddDate(0, 0, +int(2))

				objSalesOrderLoop := models.TableSalesOrder{
					UuidSalesOrder: id.String(),
					NoOrder:        NoOrder,
					StatusOrder:    "PDKM",
					SkuNo:          skunya,
					ProductName:    ProductName,
					// Qty:                qty,
					Qty:                strconv.Itoa(value.Qty),
					ExpeditionType:     Kurir,
					DeliveryType:       "Pickup",
					TotalQtyOrder:      "30",
					StatusProcessOrder: "0",
					WeightProduk:       "0",
					TotalWeight:        "0",

					SentBeforeDate: sendBeforeDate,
					CreatedDate:    createdDate,

					OrderDate: OrderDate,

					CreatedBy:     "ZALORA",
					NameVariant:   Variant,
					UsernameBuyer: CustFirstName,
					ChannelCode:   objChannel.ChannelCode,

					RecipientsName:  FirstName + " " + LastName,
					Telephone:       Phone,
					ShippingAddress: Address,
					City:            City,
					Province:        Region,
				}

				objSalesOrder = append(objSalesOrder, objSalesOrderLoop)

				objSalesOrderTotalQty.NoOrder = NoOrder
				objSalesOrderTotalQty.TotalQtyOrder = "0"

			} else {
				fmt.Println("Channel zalora Tidak Ditemukan")
			}

		} else {
			fmt.Println("No Order " + NoOrder + " Sudah Ada")

			//fmt.Println("status customer object" + dataCust.Order.Statuses[0].Status)
			if objCek.StatusProcessOrder == "0" {
				//jika CANCELED UPDATE status_processs_order jadi 3
				if dataCust.Order.Statuses[0].Status == "canceled" {
					objSalesOrderLoop := models.TableSalesOrder{
						NoOrder:            NoOrder,
						StatusProcessOrder: "3",
						CreatedBy:          "ZALORA",
					}
					result := tokenRepository.UpdateSalesOrderCanc(objSalesOrderLoop)
					if result != nil {
						fmt.Println("UpdateSalesOrderCanc ZALORA Gagal")
					} else {
						fmt.Println("UpdateSalesOrderCanc ZALORA Success")
					}
					UpdateStockOrder(NoOrder)

				}

			} else if objCek.StatusProcessOrder == "1" { //delivered
				//jika complete update ke pyoutdate
				if dataCust.Order.Statuses[0].Status == "delivered" {

					payoutdate := objZalora[0].UpdatedAt
					objSalesOrderLoop := models.TableSalesOrder{
						NoOrder:    NoOrder,
						PayoutDate: payoutdate,
						CreatedBy:  "ZALORA",
					}
					result := tokenRepository.UpdateSalesOrderCanc(objSalesOrderLoop)
					if result != nil {
						fmt.Println("Update Payoutdate ZALORA Gagal")
					} else {
						fmt.Println("Update Payoutdate ZALORA Success")
					}
					fmt.Println("======= PAYOUTDATE =======")
					fmt.Println(payoutdate)
					fmt.Println("======= PAYOUTDATE =======")
				}

			}

			break
		}

	}

	if len(objSalesOrder) > 0 {
		result := tokenRepository.SaveSalesOrder(objSalesOrder, objSalesOrderTotalQty)
		if result != nil {
			fmt.Println("SaveSalesOrderZalora Gagal")
			SaveErrorInfo("SaveSalesOrder", "ZALORA", objSalesOrderTotalQty.NoOrder+" "+result.Error())
		} else {
			fmt.Println("SaveSalesOrderZalora Success")
			pesan := BookingSku(objSalesOrderTotalQty.NoOrder)
			fmt.Println(objSalesOrderTotalQty.NoOrder + " => " + pesan)
			UpdateStockOrder(objSalesOrderTotalQty.NoOrder)
			//pesan := BookingSku(objSalesOrderTotalQty.NoOrder)
			//fmt.Println(objSalesOrderTotalQty.NoOrder + " => " + pesan)
		}
	}

}

func SaveSalesOrderZaloraSingle(id string, obj models.SuccessResponseItemZaloras, dataCust models.OrdersListZalora) {

	var objSalesOrder []models.TableSalesOrder
	var objSalesOrderTotalQty models.TableSalesOrder
	createdDate := time.Now()
	// totalQty := 0

	//cek jika orderan sudah ada di salesorder
	temp_sku := ""
	loop_qty := 0

	value := obj.SuccessResponse.Body.OrderItems.OrderItem

	NoOrder := dataCust.Order.OrderNumber + "-" + value.OrderId

	objCek, _ := tokenRepository.FindSalesOrder(NoOrder)
	if objCek.NoOrder == "" { //&& obj.Response.OrderList[0].ShippingCarrier != "" {
		//ambil channel code shopeee ke wms_channel
		objChannel, _ := tokenRepository.FindChannel("ZALORA")
		if objChannel.UuidChannel != "" {

			skunya := value.Sku

			ProductName := value.Name
			Variant := value.Variation
			Kurir := value.ShipmentProvider

			OrderDate := value.CreatedAt

			FirstName := dataCust.Order.AddressShipping.FirstName
			LastName := dataCust.Order.AddressShipping.LastName
			Address := dataCust.Order.AddressShipping.Address1
			if Address == "" {
				Address = dataCust.Order.AddressShipping.Address2
			}
			if Address == "" {
				Address = dataCust.Order.AddressShipping.Address3
			}
			if Address == "" {
				Address = dataCust.Order.AddressShipping.Address4
			}

			Phone := dataCust.Order.AddressShipping.Phone
			if Phone == "" {
				Phone = dataCust.Order.AddressShipping.Phone2
			}

			City := dataCust.Order.AddressShipping.City
			//Ward := dataCust.Order.AddressShipping.Ward
			Region := dataCust.Order.AddressShipping.Region
			//PostCode := dataCust.Order.AddressShipping.PostCode
			//Country := dataCust.Order.AddressShipping.Country
			CustFirstName := dataCust.Order.CustomerFirstName

			id := uuid.New()

			if len([]rune(skunya)) == 8 {
				//skunya = skunya
			} else if len([]rune(skunya)) > 8 {
				sku := []rune(skunya)
				skunya = string(sku[1:9])
			} else if len([]rune(skunya)) == 7 {
				skunya = "0" + value.Sku
			} else {
				skunya = "XXXXXXXX"
			}

			amounts := value.ItemPrice
			amount := 0.0
			if s, err := strconv.ParseFloat(amounts, 64); err == nil {
				amount = s
			}

			sendBeforeDate := createdDate.AddDate(0, 0, +int(2))

			if temp_sku == "" || temp_sku == skunya {
				loop_qty++
				temp_sku = skunya
			}

			amount = amount * float64(loop_qty)

			objSalesOrderLoop := models.TableSalesOrder{
				UuidSalesOrder: id.String(),
				NoOrder:        NoOrder,
				StatusOrder:    "PDKM",
				SkuNo:          skunya,
				ProductName:    ProductName,
				// Qty:                qty,
				Qty:                strconv.Itoa(loop_qty),
				ExpeditionType:     Kurir,
				DeliveryType:       "Pickup",
				TotalQtyOrder:      "30",
				StatusProcessOrder: "0",
				WeightProduk:       "0",
				TotalWeight:        "0",

				SentBeforeDate: sendBeforeDate,
				CreatedDate:    createdDate,

				OrderDate: OrderDate,

				CreatedBy:     "ZALORA",
				NameVariant:   Variant,
				UsernameBuyer: CustFirstName,
				ChannelCode:   objChannel.ChannelCode,

				RecipientsName:  FirstName + " " + LastName,
				Telephone:       Phone,
				ShippingAddress: Address,
				City:            City,
				Province:        Region,
				Amount:          amount,
			}

			objSalesOrder = append(objSalesOrder, objSalesOrderLoop)

			objSalesOrderTotalQty.NoOrder = NoOrder
			objSalesOrderTotalQty.TotalQtyOrder = "0"

		} else {
			fmt.Println("Channel zalora Tidak Ditemukan")
		}

		result := tokenRepository.SaveSalesOrder(objSalesOrder, objSalesOrderTotalQty)
		if result != nil {
			fmt.Println("SaveSalesOrderZaloraSingle Gagal")
			SaveErrorInfo("SaveSalesOrder", "ZALORA", objSalesOrderTotalQty.NoOrder+" "+result.Error())
		} else {
			fmt.Println("SaveSalesOrderZaloraSingle Success")
			pesan := BookingSku(objSalesOrderTotalQty.NoOrder)
			fmt.Println(objSalesOrderTotalQty.NoOrder + " => " + pesan)
			UpdateStockOrder(objSalesOrderTotalQty.NoOrder)
		}

	} else {
		fmt.Println("No Order " + NoOrder + " Sudah Ada")
	}

	//fmt.Println("status customer " + dataCust.Order.Statuses[0].Status)

	if objCek.StatusProcessOrder == "0" {
		//jika CANCELED UPDATE status_processs_order jadi 3
		if dataCust.Order.Statuses[0].Status == "canceled" {
			objSalesOrderLoop := models.TableSalesOrder{
				NoOrder:            NoOrder,
				StatusProcessOrder: "3",
				CreatedBy:          "ZALORA",
			}
			result := tokenRepository.UpdateSalesOrderCanc(objSalesOrderLoop)
			if result != nil {
				fmt.Println("UpdateSalesOrderCanc ZALORA Gagal")
			} else {
				fmt.Println("UpdateSalesOrderCanc ZALORA Success")
			}
			UpdateStockOrder(NoOrder)

		}
	} else if objCek.StatusProcessOrder == "1" { //delivered
		//jika complete update ke pyoutdate
		if dataCust.Order.Statuses[0].Status == "delivered" {

			payoutdate := value.UpdatedAt
			objSalesOrderLoop := models.TableSalesOrder{
				NoOrder:    NoOrder,
				PayoutDate: payoutdate,
				CreatedBy:  "ZALORA",
			}
			result := tokenRepository.UpdateSalesOrderCanc(objSalesOrderLoop)
			if result != nil {
				fmt.Println("Update Payoutdate ZALORA SINGLE Gagal")
			} else {
				fmt.Println("Update Payoutdate ZALORA SINGLE Success")
			}
			fmt.Println("======= PAYOUTDATE =======")
			fmt.Println(payoutdate)
			fmt.Println("======= PAYOUTDATE =======")
		}

	}

}

func SaveSalesOrderZaloraV2(id string, obj models.OrderDetailZaolraV2) {
	var objSalesOrder []models.TableSalesOrder
	var objSalesOrderTotalQty models.TableSalesOrder
	createdDate := time.Now()

	objNew := CekArrayOrderV2(obj)
	fmt.Println("********************")
	NoOrder := obj.Number + "-" + strconv.Itoa(obj.ID)
	fmt.Println(NoOrder)
	FirstName := obj.Address.Shipping.FirstName
	LastName := obj.Address.Shipping.LastName
	CustFirstName := obj.Customer.FirstName

	Address := ""
	for indexxx, _ := range obj.Address.Shipping.Address {
		Address += obj.Address.Shipping.Address[indexxx]
	}

	Phone := ""
	for indexxx, _ := range obj.Address.Shipping.Phone {
		if Phone == "" {
			Phone += obj.Address.Shipping.Address[indexxx]
		}
	}

	City := obj.Address.Shipping.City
	Region := obj.Address.Shipping.Region

	objCek, _ := tokenRepository.FindSalesOrder(NoOrder)
	if objCek.NoOrder == "" {
		for _, value := range objNew {

			objChannel, _ := tokenRepository.FindChannel("ZALORA")
			if objChannel.UuidChannel != "" && value.Status == "pending" {

				skunya := value.Sku

				ProductName := value.Name
				Variant := value.Variation
				Kurir := value.ShipmentProvider

				OrderDate := value.CreatedAt

				id := uuid.New()

				if len([]rune(skunya)) == 8 {
					//skunya = skunya
				} else if len([]rune(skunya)) > 8 {
					sku := []rune(skunya)
					skunya = string(sku[1:9])
				} else if len([]rune(skunya)) == 7 {
					skunya = "0" + value.Sku
				} else {
					skunya = "XXXXXXXX"
				}

				sendBeforeDate := createdDate.AddDate(0, 0, +int(2))

				objSalesOrderLoop := models.TableSalesOrder{
					UuidSalesOrder:     id.String(),
					NoOrder:            NoOrder,
					StatusOrder:        "PDKM",
					SkuNo:              skunya,
					ProductName:        ProductName,
					Qty:                strconv.Itoa(value.Qty),
					ExpeditionType:     Kurir,
					DeliveryType:       "Pickup",
					TotalQtyOrder:      "30",
					StatusProcessOrder: "0",
					WeightProduk:       "0",
					TotalWeight:        "0",

					SentBeforeDate: sendBeforeDate,
					CreatedDate:    createdDate,

					OrderDate: OrderDate,

					CreatedBy:     "ZALORA",
					NameVariant:   Variant,
					UsernameBuyer: CustFirstName,
					ChannelCode:   objChannel.ChannelCode,

					RecipientsName:  FirstName + " " + LastName,
					Telephone:       Phone,
					ShippingAddress: Address,
					City:            City,
					Province:        Region,
					Amount:          value.Amount * float64(value.Qty),
				}

				objSalesOrder = append(objSalesOrder, objSalesOrderLoop)

				objSalesOrderTotalQty.NoOrder = NoOrder
				objSalesOrderTotalQty.TotalQtyOrder = "0"
				//fmt.Println(objSalesOrder)

			} else {
				fmt.Println("Channel zalora Tidak Ditemukan")
			}

		}

		if len(objSalesOrder) > 0 {
			result := tokenRepository.SaveSalesOrder(objSalesOrder, objSalesOrderTotalQty)
			if result != nil {
				fmt.Println("SaveSalesOrderZaloraSingle Gagal")
				SaveErrorInfo("SaveSalesOrder", "ZALORA", objSalesOrderTotalQty.NoOrder+" "+result.Error())
			} else {
				fmt.Println("SaveSalesOrderZaloraSingle Success")
				pesan := BookingSku(objSalesOrderTotalQty.NoOrder)
				fmt.Println(objSalesOrderTotalQty.NoOrder + " => " + pesan)
				UpdateStockOrder(objSalesOrderTotalQty.NoOrder)
			}

		}

		//spew.Dump(objSalesOrder)
	} else {
		fmt.Println("No Order " + NoOrder + " Sudah Ada")

		if objCek.StatusProcessOrder == "0" {
			//jika CANCELED UPDATE status_processs_order jadi 3
			for _, value := range objNew {
				if value.Status == "canceled" {

					skunya := value.Sku

					if len([]rune(skunya)) == 8 {
						//skunya = skunya
					} else if len([]rune(skunya)) > 8 {
						sku := []rune(skunya)
						skunya = string(sku[1:9])
					} else if len([]rune(skunya)) == 7 {
						skunya = "0" + value.Sku
					} else {
						skunya = "XXXXXXXX"
					}

					//cari salesorder uuid sales order
					ObjCekSales, _ := tokenRepository.FindSalesOrderBySku(NoOrder, "0", skunya)
					if ObjCekSales.UuidSalesOrder != "" {

						objSalesOrderLoop := models.TableSalesOrder{
							UuidSalesOrder:     ObjCekSales.UuidSalesOrder,
							NoOrder:            "XX" + NoOrder,
							StatusProcessOrder: "3",
							CreatedBy:          "ZALORA",
							SkuNo:              ObjCekSales.SkuNo,
						}
						//result := tokenRepository.UpdateSalesOrderCanc(objSalesOrderLoop)
						result := tokenRepository.UpdateSalesOrderCancNew(NoOrder, objSalesOrderLoop)
						if result != nil {
							fmt.Println("UpdateSalesOrderCanc ZALORA Gagal")
						} else {
							fmt.Println("UpdateSalesOrderCanc ZALORA Success")
						}
						UpdateStockOrder(NoOrder)

					}

				}

			}

		} else if objCek.StatusProcessOrder == "1" && objCek.PayoutDate == "" {

			for _, value := range objNew {
				if value.Status == "delivered" {

					payoutdate := obj.UpdatedAt
					// objSalesOrderLoop := models.TableSalesOrder{
					// 	NoOrder:    NoOrder,
					// 	PayoutDate: payoutdate,
					// 	CreatedBy:  "ZALORA",
					// }
					// result := tokenRepository.UpdateSalesOrderCanc(objSalesOrderLoop)
					// if result != nil {
					// 	fmt.Println("Update Payoutdate ZALORA Gagal")
					// } else {
					// 	fmt.Println("Update Payoutdate ZALORA Success")
					// }
					fmt.Println("======= PAYOUTDATE =======")
					fmt.Println(payoutdate)
					fmt.Println("======= PAYOUTDATE =======")
					break

				}

			}

		}

	}

	//
	fmt.Println("********************")
}
func SaveSalesOrderTiktok(id string, obj []models.OrderListDetailTiktok) {
	var objSalesOrder []models.TableSalesOrder
	var objSalesOrderTotalQty models.TableSalesOrder
	createdDate := time.Now()
	totalQty := 0
	fmt.Println("SaveSalesOrderTiktok")
	index := 0
	for _, valOrderList := range obj {

		no_order := valOrderList.OrderId
		objCek, _ := tokenRepository.FindSalesOrder(no_order)

		if objCek.NoOrder == "" && obj[index].ShippingProvider != "" && strconv.Itoa(int(obj[index].OrderStatus)) == "111" {
			objChannel, _ := tokenRepository.FindChannel("TIKTOK")
			if objChannel.UuidChannel != "" {
				for _, value := range obj[index].ItemList {

					id := uuid.New()
					qty := strconv.Itoa(value.Quantity)
					totalQty += value.Quantity

					skunya := value.SellerSku
					if skunya == "" {
						skunya = value.SellerSku
					}

					if value.SellerSku == "346831" {
						skunya = "08171374"
					}
					if value.SellerSku == "397574" {
						skunya = "08169676"
					}

					if len([]rune(skunya)) == 8 {
						//skunya = skunya
					} else if len([]rune(skunya)) > 8 {
						sku := []rune(skunya)
						skunya = string(sku[1:9])
					} else if len([]rune(skunya)) == 7 {
						skunya = "0" + skunya
					} else {
						skunya = "XXXXXXXX"
					}

					amount := value.SkuSalePrice

					sendBeforeDate := time.Unix(int64(obj[index].RtsSla), 0)
					createtimes, _ := strconv.Atoi(obj[index].CreateTime[0:10])

					objSalesOrderLoop := models.TableSalesOrder{
						UuidSalesOrder:     id.String(),
						NoOrder:            no_order,
						StatusOrder:        "PDKM",
						SkuNo:              skunya,
						ProductName:        value.ProductName,
						Qty:                qty,
						ExpeditionType:     obj[index].ShippingProvider,
						DeliveryType:       "Pickup",
						TotalQtyOrder:      "30",
						StatusProcessOrder: "0",
						WeightProduk:       "0",
						TotalWeight:        "0",

						SentBeforeDate: sendBeforeDate,
						CreatedDate:    createdDate,

						OrderDate: time.Unix(int64(createtimes), 0),

						CreatedBy: "TIKTOK",
						//NameVariant:   value.ModelName,
						UsernameBuyer: obj[index].RecipientAddress.Name,
						ChannelCode:   objChannel.ChannelCode,

						RecipientsName:  obj[index].RecipientAddress.Name,
						Telephone:       obj[index].RecipientAddress.Phone,
						ShippingAddress: obj[index].RecipientAddress.FullAddress,
						City:            obj[index].RecipientAddress.City,
						Province:        obj[index].RecipientAddress.State,
						Amount:          amount * float64(value.Quantity),
					}

					objSalesOrder = append(objSalesOrder, objSalesOrderLoop)

					fmt.Println(" SKU ID : " + value.SkuId)
					fmt.Println(" SKUNYA : " + value.SellerSku)
					fmt.Println(" QTY : " + strconv.Itoa(value.Quantity))
				}

				objSalesOrderTotalQty.NoOrder = no_order
				objSalesOrderTotalQty.TotalQtyOrder = strconv.Itoa(totalQty)
				//fmt.Println(objSalesOrder)

				result := tokenRepository.SaveSalesOrder(objSalesOrder, objSalesOrderTotalQty)
				if result != nil {
					fmt.Println("SaveSalesOrder Gagal")
					SaveErrorInfo("SaveSalesOrder", "TIKTOK", objSalesOrderTotalQty.NoOrder+" "+result.Error())
				} else {
					fmt.Println("SaveSalesOrder Success")
				}

				pesan := BookingSku(objSalesOrderTotalQty.NoOrder)
				fmt.Println(objSalesOrderTotalQty.NoOrder + " => " + pesan)
				UpdateStockOrder(objSalesOrderTotalQty.NoOrder)

				objSalesOrder = nil
			}

		} else {

			//cek jika kurir ganti update kurir
			var objUpdateKurir models.TableSalesOrder
			if objCek.NoOrder == no_order {
				if objCek.ExpeditionType != obj[index].ShippingProvider {
					objUpdateKurir.NoOrder = objCek.NoOrder
					objUpdateKurir.ExpeditionType = obj[index].ShippingProvider
					result := tokenRepository.UpdateKurir(objUpdateKurir)
					if result != nil {
						fmt.Println("UpdateKurir Gagal")
					} else {
						fmt.Println("UpdateKurir Success")
					}

				}
			}

			if objCek.StatusProcessOrder == "0" && objCek.NoOrder == no_order {

				//jika CANCELED UPDATE status_processs_order jadi 3
				if strconv.Itoa(int(obj[index].OrderStatus)) == "140" { //CANCELLED
					objSalesOrderLoop := models.TableSalesOrder{
						NoOrder:            objCek.NoOrder,
						StatusProcessOrder: "3",
						CreatedBy:          "TIKTOK",
					}
					result := tokenRepository.UpdateSalesOrderCanc(objSalesOrderLoop)
					if result != nil {
						fmt.Println("UpdateSalesOrderCanc Gagal")
					} else {
						fmt.Println("UpdateSalesOrderCanc Success")
					}
					UpdateStockOrder(no_order) //sini

				}
			}

			if objCek.StatusProcessOrder == "1" && objCek.NoOrder == no_order {
				if strconv.Itoa(int(obj[index].OrderStatus)) == "130" { //CANCELLED

					payoutdate := time.Unix(int64(obj[index].UpdateTime), 0)
					objSalesOrderLoop := models.TableSalesOrder{
						NoOrder:    objCek.NoOrder,
						PayoutDate: payoutdate,
						CreatedBy:  "TIKTOK",
					}
					result := tokenRepository.UpdateSalesOrderCanc(objSalesOrderLoop)
					if result != nil {
						fmt.Println("Update Payoutdate TIKTOK Gagal")
					} else {
						fmt.Println("Update Payoutdate TIKTOK Success")
					}
					fmt.Println("======= PAYOUTDATE =======")
					fmt.Println(payoutdate)
					fmt.Println("======= PAYOUTDATE =======")

				}
			}

			// objSalesOrder = nil //kosongin
			index++
		}
	}
}

func CariOrderByResi(param string) []models.TableSalesOrder {
	var objSalesOrder []models.TableSalesOrder
	objSalesOrder, _ = tokenRepository.CariOrderByResi(param)

	return objSalesOrder
}

func SearchProductNull(param string) []models.TableSkuMapping {
	var objProduct []models.TableSkuMapping
	objProduct = tokenRepository.SearchProductNull(param)

	return objProduct

}

func SaveSkuMapping(isiModel models.ModelProductHeader, IdItem int64, itemname string, status string) {
	var objMapping []models.TableSkuMapping
	indexnya := 0
	for _, elementModel := range isiModel.ModelProductHeaderResp.ModelProductDetail {
		if len(elementModel.ModelSku) > 5 {

			skunya := elementModel.ModelSku
			if len([]rune(skunya)) == 8 {
				//skunya = skunya
			} else if len([]rune(skunya)) > 8 {
				sku := []rune(skunya)
				skunya = string(sku[1:9])
			} else {
				skunya = "XXXXXXXX"
			}
			//fmt.Println("SaveSkuMapping " + skunya)

			skuid := elementModel.ModelId
			CekSkuWMS, _ := tokenRepository.CariSkuWMS(skunya)

			statusnya := "3936a00a-e321-f5ff-33b5-2dfb8d68ba45"
			if status != "NORMAL" {
				statusnya = "c895af78-1488-cf0b-d17a-33770f2b20bf"
			}

			if CekSkuWMS.SkuNo != "" {

				//update jadi nonaktif jika hanya punya parent
				objUpdateParent, _ := tokenRepository.CekSkuMappingChild(skunya, "R996", strconv.Itoa(int(IdItem)))
				if len(objUpdateParent) > 0 {

					for _, elementVal := range objUpdateParent {
						objMappingUpdParent := models.TableSkuMapping{
							UuidSkuMapping: elementVal.UuidSkuMapping,
							UuidStatus:     "c895af78-1488-cf0b-d17a-33770f2b20bf",
							CreatedBy:      "PARENT",
						}
						//fmt.Println(objMappingUpd)
						result_upd := tokenRepository.UpdateSaveSkuMapping(objMappingUpdParent)
						if result_upd != nil {
							fmt.Println("Update SaveSkuMapping objMappingUpdParent Gagal")
						}
					}

				}

				CekMapping, _ := tokenRepository.CariSkuMapping(skunya, "R996")
				//fmt.Println(CekMapping)

				varian := ""
				color := ""
				indextier := 0
				for _, elementTier := range elementModel.TierIndex {
					if strings.ToUpper(isiModel.ModelProductHeaderResp.TierVariation[indextier].Name) == "COLOR" {
						color = isiModel.ModelProductHeaderResp.TierVariation[indextier].OptionList[elementTier].Option
					} else {
						varian = isiModel.ModelProductHeaderResp.TierVariation[indextier].OptionList[elementTier].Option
					}
					// if strings.ToUpper(isiModel.ModelProductHeaderResp.TierVariation[indextier].Name) == "SIZE"{
					// 	varian = isiModel.ModelProductHeaderResp.TierVariation[indextier].OptionList[elementTier].Option
					// }
					// fmt.Println(isiModel.ModelProductHeaderResp.TierVariation[indextier].OptionList[elementTier])
					indextier++
				}

				if CekMapping.ChannelCode == "" {
					objMappingLoop := models.TableSkuMapping{
						UuidSkuMapping: uuid.New().String(),
						SkuNo:          skunya,
						ChannelCode:    "R996",
						CreatedBy:      "SYSTEM",
						CreatedDate:    time.Now(),
						IdSkuParent:    strconv.Itoa(int(IdItem)),
						IdSkuChild1:    strconv.Itoa(int(skuid)),
						ProductName:    itemname,
						Varian:         varian,
						Color:          color,
						UuidStatus:     statusnya,
					}

					objMapping = append(objMapping, objMappingLoop)
					//fmt.Println("Model SKU SERVICE " + elementModel.ModelSku)
				} else {
					//update varian
					//fmt.Println("update varian " + strconv.Itoa(int(IdItem)))

					if CekMapping.IdSkuParent != strconv.Itoa(int(IdItem)) {

						//cek jika sudah disave
						CekParent, _ := tokenRepository.CariSkuMappingByParent(skunya, "R996", strconv.Itoa(int(IdItem)))
						if CekParent.IdSkuParent == "" {

							CekChild1, _ := tokenRepository.CariSkuMappingByModel(skunya, "R996", strconv.Itoa(int(skuid)))
							if CekChild1.IdSkuParent == "" {
								objMappingLoop := models.TableSkuMapping{
									UuidSkuMapping: uuid.New().String(),
									SkuNo:          skunya,
									ChannelCode:    "R996",
									CreatedBy:      "SYSTEM",
									CreatedDate:    time.Now(),
									IdSkuParent:    strconv.Itoa(int(IdItem)),
									IdSkuChild1:    strconv.Itoa(int(skuid)),
									ProductName:    itemname,
									Varian:         varian,
									Color:          color,
									UuidStatus:     statusnya,
								}
								objMapping = append(objMapping, objMappingLoop)
								SaveErrorString("shopee", "SKU:"+skunya+" OLD:"+CekMapping.IdSkuParent+" | NEW:"+strconv.Itoa(int(IdItem)), "SKU SUDAH ADA")
							}

						}

						//SaveErrorString("shopee", "SKU:"+skunya+" OLD:"+CekMapping.IdSkuParent+" | NEW:"+strconv.Itoa(int(IdItem)), "Id Parent")
						//SaveErrorString("shopee", "SKU:"+skunya+" OLD:"+CekMapping.IdSkuChild1+" | NEW:"+strconv.Itoa(int(skuid)), "Id Child")
					}

					objMappingUpd := models.TableSkuMapping{
						UuidSkuMapping: CekMapping.UuidSkuMapping,
						//IdSkuParent:    strconv.Itoa(int(IdItem)),
						//IdSkuChild1:    strconv.Itoa(int(skuid)),
						SkuNo:       skunya,
						Varian:      varian,
						ProductName: itemname,
						Color:       color,
						UuidStatus:  statusnya,
					}
					//fmt.Println(objMappingUpd)
					result_upd := tokenRepository.UpdateSaveSkuMapping(objMappingUpd)
					if result_upd != nil {
						fmt.Println("Update SaveSkuMapping Gagal")
					}
					indexnya++
				}

			} else {
				//fmt.Println("SaveSkuMapping " + skunya + " gak ada ")
			}
			//insert ke wms_sku_mapping

		}

	}

	//fmt.Println(objMapping)
	result := tokenRepository.SaveSkuMapping(objMapping)
	if result != nil {
		fmt.Println("SaveSkuMapping Gagal")
	} else {
		//fmt.Println("SaveSkuMapping Success")
	}

}

func SaveSkuMappingString(sku string, Itemid int64, itemname string, status string) {
	var objMapping []models.TableSkuMapping

	if len(sku) > 5 {

		skunya := sku
		if len([]rune(skunya)) == 8 {
			//skunya = skunya
		} else if len([]rune(skunya)) > 8 {
			sku := []rune(skunya)
			skunya = string(sku[1:9])
		} else {
			skunya = "XXXXXXXX"
		}

		skuid := Itemid
		CekSkuWMS, _ := tokenRepository.CariSkuWMS(skunya)

		statusnya := "3936a00a-e321-f5ff-33b5-2dfb8d68ba45"
		if status != "NORMAL" {
			statusnya = "c895af78-1488-cf0b-d17a-33770f2b20bf"
		}

		if CekSkuWMS.SkuNo != "" {
			CekMapping, _ := tokenRepository.CariSkuMapping(skunya, "R996")
			//fmt.Println(CekMapping)
			if CekMapping.ChannelCode == "" {
				objMappingLoop := models.TableSkuMapping{
					UuidSkuMapping: uuid.New().String(),
					SkuNo:          skunya,
					ChannelCode:    "R996",
					CreatedBy:      "SYSTEM",
					CreatedDate:    time.Now(),
					IdSkuParent:    strconv.Itoa(int(skuid)),
					ProductName:    itemname,
					UuidStatus:     statusnya,
				}

				objMapping = append(objMapping, objMappingLoop)

			} else {

				if CekMapping.IdSkuParent != strconv.Itoa(int(skuid)) {

					//cek jika sudah disave
					CekParent, _ := tokenRepository.CariSkuMappingByParent(skunya, "R996", strconv.Itoa(int(skuid)))
					if CekParent.IdSkuParent == "" {

						objMappingLoop := models.TableSkuMapping{
							UuidSkuMapping: uuid.New().String(),
							SkuNo:          skunya,
							ChannelCode:    "R996",
							CreatedBy:      "SYSTEM",
							CreatedDate:    time.Now(),
							IdSkuParent:    strconv.Itoa(int(skuid)),
							ProductName:    itemname,
							UuidStatus:     statusnya,
						}
						objMapping = append(objMapping, objMappingLoop)
						SaveErrorString("shopee", "SKU:"+skunya+" OLD:"+CekMapping.IdSkuParent+" | NEW:"+strconv.Itoa(int(skuid)), "SKU SUDAH ADA")

					}

					//SaveErrorString("shopee", "SKU:"+skunya+" OLD:"+CekMapping.IdSkuParent+" | NEW:"+strconv.Itoa(int(skuid)), "Id Parent")
				}

				objMappingUpd := models.TableSkuMapping{
					UuidSkuMapping: CekMapping.UuidSkuMapping,
					SkuNo:          skunya,
					ProductName:    itemname,
					//IdSkuParent:    strconv.Itoa(int(skuid)),
					IdSkuChild1: "",
					UuidStatus:  statusnya,
				}

				result_upd := tokenRepository.UpdateSaveSkuMapping(objMappingUpd)
				if result_upd != nil {
					fmt.Println("Update SaveSkuMappingString Gagal")
				}
			}

		}
		//insert ke wms_sku_mapping

	}

	//fmt.Println(objMapping)
	result := tokenRepository.SaveSkuMapping(objMapping)
	if result != nil {
		fmt.Println("SaveSkuMappingString Gagal")
	} else {
		//fmt.Println("SaveSkuMappingString Success")
	}
}

func UpdateStatusSkuMappingShopeeNull(Itemid int64, status string) {
	skuid := Itemid

	statusnya := "3936a00a-e321-f5ff-33b5-2dfb8d68ba45"
	if status != "NORMAL" {
		statusnya = "c895af78-1488-cf0b-d17a-33770f2b20bf"
	}

	objMappingUpd := models.TableSkuMapping{
		IdSkuParent: strconv.Itoa(int(skuid)),
		ChannelCode: os.Getenv("KODE_SHOPEE"),
		UuidStatus:  statusnya,
	}

	result_upd := tokenRepository.UpdateStatusSkuMappingNull(objMappingUpd)
	if result_upd != nil {
		fmt.Println("UpdateStatusSkuMappingShopeeNull Gagal")
	}

}

func SaveSkuMappingZalora(isiModel models.ProductsListZalora) {
	var objMapping []models.TableSkuMapping
	kode_zalora := "R991"
	tempsku := ""

	for _, elementModel := range isiModel.Product {

		if len(elementModel.SellerSku) > 5 {

			skunya := elementModel.SellerSku

			statusnya := "3936a00a-e321-f5ff-33b5-2dfb8d68ba45"
			if elementModel.Status != "active" {
				statusnya = "c895af78-1488-cf0b-d17a-33770f2b20bf"
			}

			if len([]rune(skunya)) == 8 {
				//skunya = skunya
			} else if len([]rune(skunya)) > 8 {
				sku := []rune(skunya)
				skunya = string(sku[1:9])
			} else {
				skunya = elementModel.ParentSku
				if len([]rune(skunya)) < 8 {
					skunya = "0" + skunya
				} else {
					skunya = "XXXXXXXX"
				}

			}

			//fmt.Println("SaveSkuMapping " + skunya)

			// if skunya == "08108714" {
			// 	skunya = "08139207"
			// }
			if tempsku != skunya {
				//fmt.Println(skunya + " | " + elementModel.ParentSku + " | " + tempsku)

				CekSkuWMS, _ := tokenRepository.CariSkuWMS(skunya)
				if CekSkuWMS.SkuNo != "" {

					CekMapping, _ := tokenRepository.CariSkuMapping(skunya, kode_zalora)
					//fmt.Println(CekMapping)

					varian := elementModel.Variation
					color := ""

					if CekMapping.ChannelCode == "" {
						objMappingLoop := models.TableSkuMapping{
							UuidSkuMapping: uuid.New().String(),
							SkuNo:          skunya,
							ChannelCode:    kode_zalora,
							CreatedBy:      "SYSTEM",
							CreatedDate:    time.Now(),
							// IdSkuParent:    elementModel.ParentSku,
							IdSkuParent: elementModel.SellerSku,
							// IdSkuChild1:    strconv.Itoa(int(skuid)),
							ProductName: elementModel.Name,
							Varian:      varian,
							Color:       color,
							UuidStatus:  statusnya,
						}

						objMapping = append(objMapping, objMappingLoop)

					} else {
						objMappingUpd := models.TableSkuMapping{
							UuidSkuMapping: CekMapping.UuidSkuMapping,
							SkuNo:          skunya,
							Varian:         varian,
							ProductName:    elementModel.Name,

							IdSkuParent: elementModel.SellerSku,

							Color:      color,
							UuidStatus: statusnya,
						}
						//fmt.Println(objMappingUpd)
						result_upd := tokenRepository.UpdateSaveSkuMapping(objMappingUpd)
						if result_upd != nil {
							fmt.Println("Update SaveSkuMappingZalora Gagal")
						}
					}

				}

			} //tempsku != skunya

			if tempsku != skunya {
				tempsku = skunya
			}
			//insert ke wms_sku_mapping

		}

	}

	//fmt.Println(objMapping)
	result := tokenRepository.SaveSkuMapping(objMapping)
	if result != nil {
		fmt.Println("SaveSkuMappingZalora Gagals")
	} else {
		//fmt.Println("SaveSkuMapping Success")
	}

}
func SaveSkuMappingZaloraV2OLD(obj models.ProductV2Zalora, status string) {
	var objMapping []models.TableSkuMapping
	kode_zalora := "R991"

	for _, elementModel := range obj.Items {
		// fmt.Println("=======================")
		// fmt.Println(status)
		// fmt.Println(elementModel.ID)
		// fmt.Println(elementModel.Name)
		// fmt.Println(elementModel.ParentSku)
		// fmt.Println(elementModel.SellerSku)
		// fmt.Println("=======================")

		// if len(elementModel.SellerSku) > 5 && elementModel.SellerSku == "08173729" {
		if len(elementModel.SellerSku) > 5 {

			skunya := elementModel.SellerSku

			statusnya := "3936a00a-e321-f5ff-33b5-2dfb8d68ba45"
			if status != "active" {
				statusnya = "c895af78-1488-cf0b-d17a-33770f2b20bf"
			}

			if len([]rune(skunya)) == 8 {
				//skunya = skunya
			} else if len([]rune(skunya)) > 8 {
				sku := []rune(skunya)
				skunya = string(sku[1:9])
			} else {
				skunya = elementModel.ParentSku
				if len([]rune(skunya)) < 8 {
					skunya = "0" + skunya
				} else {
					skunya = "XXXXXXXX"
				}

			}

			CekSkuWMS, _ := tokenRepository.CariSkuWMS(skunya)
			if CekSkuWMS.SkuNo != "" {

				CekMapping, _ := tokenRepository.CariSkuMapping(skunya, kode_zalora)

				if CekMapping.ChannelCode == "" {
					objMappingLoop := models.TableSkuMapping{
						UuidSkuMapping: uuid.New().String(),
						SkuNo:          skunya,
						ChannelCode:    kode_zalora,
						CreatedBy:      "SYSTEM",
						CreatedDate:    time.Now(),
						IdSkuParent:    elementModel.SellerSku,
						IdSkuChild1:    strconv.Itoa(elementModel.ID),
						ProductName:    elementModel.Name,
						// Varian:      varian,
						// Color:       color,
						UuidStatus: statusnya,
					}

					objMapping = append(objMapping, objMappingLoop)

				} else {
					objMappingUpd := models.TableSkuMapping{
						UuidSkuMapping: CekMapping.UuidSkuMapping,
						SkuNo:          skunya,
						ProductName:    elementModel.Name,
						IdSkuParent:    elementModel.SellerSku,
						IdSkuChild1:    strconv.Itoa(elementModel.ID),
						UuidStatus:     statusnya,
					}
					//spew.Dump(objMappingUpd)
					result_upd := tokenRepository.UpdateSaveSkuMapping(objMappingUpd)
					if result_upd != nil {
						fmt.Println("Update SaveSkuMappingZaloraV2 Gagal")
					}
				}

			}

		}
	}

	result := tokenRepository.SaveSkuMapping(objMapping)
	if result != nil {
		fmt.Println("SaveSkuMappingZaloraV2 Gagals")
	}

}

func SaveSkuMappingZaloraV2(obj models.DetailStockZaloraV2, status string) {
	var objMapping []models.TableSkuMapping
	kode_zalora := "R991"

	for _, elementModel := range obj {
		// fmt.Println("=======================")
		// fmt.Println(status)
		// fmt.Println(elementModel.ID)
		// fmt.Println(elementModel.Name)
		// fmt.Println(elementModel.ParentSku)
		// fmt.Println(elementModel.SellerSku)
		// fmt.Println("=======================")

		// if len(elementModel.SellerSku) > 5 && elementModel.SellerSku == "08173729" {
		if len(elementModel.SellerSku) > 5 {

			skunya := elementModel.SellerSku

			statusnya := "3936a00a-e321-f5ff-33b5-2dfb8d68ba45"
			if status != "active" {
				statusnya = "c895af78-1488-cf0b-d17a-33770f2b20bf"
			}

			if status == "sold-out" {
				statusnya = "3936a00a-e321-f5ff-33b5-2dfb8d68ba45"
			}

			if skunya == "08118419" || skunya == "8118419" {
				fmt.Println(skunya + " | " + statusnya)
			}

			if len([]rune(skunya)) == 8 {
				//skunya = skunya
			} else if len([]rune(skunya)) > 8 {
				sku := []rune(skunya)
				skunya = string(sku[1:9])
			} else {
				//skunya = elementModel.ParentSku
				if len([]rune(skunya)) < 8 {
					skunya = "0" + skunya
				} else {
					skunya = "XXXXXXXX"
				}

			}

			CekSkuWMS, _ := tokenRepository.CariSkuWMS(skunya)
			if CekSkuWMS.SkuNo != "" {

				CekMapping, _ := tokenRepository.CariSkuMapping(skunya, kode_zalora)

				if CekMapping.ChannelCode == "" {
					objMappingLoop := models.TableSkuMapping{
						UuidSkuMapping: uuid.New().String(),
						SkuNo:          skunya,
						ChannelCode:    kode_zalora,
						CreatedBy:      "SYSTEM",
						CreatedDate:    time.Now(),
						IdSkuParent:    elementModel.SellerSku,
						IdSkuChild1:    strconv.Itoa(elementModel.ProductID),
						ProductName:    elementModel.Name,
						// Varian:      varian,
						// Color:       color,
						UuidStatus: statusnya,
					}

					objMapping = append(objMapping, objMappingLoop)

				} else {
					objMappingUpd := models.TableSkuMapping{
						UuidSkuMapping: CekMapping.UuidSkuMapping,
						SkuNo:          skunya,
						ProductName:    elementModel.Name,
						IdSkuParent:    elementModel.SellerSku,
						IdSkuChild1:    strconv.Itoa(elementModel.ProductID),
						UuidStatus:     statusnya,
						UpdatedDate:    time.Now(),
					}
					//spew.Dump(objMappingUpd)
					result_upd := tokenRepository.UpdateSaveSkuMapping(objMappingUpd)
					if result_upd != nil {
						fmt.Println("Update SaveSkuMappingZaloraV2 Gagal")
					}
				}

			}

		}
	}

	result := tokenRepository.SaveSkuMapping(objMapping)
	if result != nil {
		fmt.Println("SaveSkuMappingZaloraV2 Gagals")
	}

}

func SaveSkuMappingTiktok(isiModel []models.ProductsListTiktok) {
	kode_tiktok := os.Getenv("KODE_TIKTOK")
	// fmt.Println(os.Getenv("KODE_TIKTOK"))
	var objMapping []models.TableSkuMapping

	for _, Isiproducts := range isiModel {
		// fmt.Println("===================================")

		// fmt.Println("Product ID :" + Isiproducts.Id)
		//cari detail sku
		for _, detailSKU := range Isiproducts.Skus {
			// fmt.Println("ID SKUNYA :" + detailSKU.Id)
			// fmt.Println("SKUNYA :" + detailSKU.SellerSku)

			skunya := detailSKU.SellerSku

			if len(detailSKU.SellerSku) > 7 {

				if len([]rune(skunya)) == 8 {

				} else if len([]rune(skunya)) > 8 {
					sku := []rune(skunya)
					skunya = string(sku[1:9])
				} else {
					//skunya = elementModel.ParentSku
					if len([]rune(skunya)) < 8 {
						skunya = "0" + skunya
					} else {
						skunya = "XXXXXXXX"
					}

				}

				CekSkuWMS, _ := tokenRepository.CariSkuWMS(skunya)
				if CekSkuWMS.SkuNo != "" {
					CekMapping, _ := tokenRepository.CariSkuMapping(skunya, kode_tiktok)

					varian := ""
					color := ""
					//fmt.Println(len(detailSKU.SalesAttributes))
					if len(detailSKU.SalesAttributes) > 0 {

						for _, attrSKU := range detailSKU.SalesAttributes {
							// fmt.Println("*******************")
							// fmt.Println(attrSKU.ValueName)
							// fmt.Println("*******************")
							if strings.ToUpper(attrSKU.Name) == "WARNA" || strings.ToUpper(attrSKU.Name) == "COLOR" ||
								strings.ToUpper(attrSKU.Name) == "COLOUR" {
								color = attrSKU.ValueName
							}

							if strings.ToUpper(attrSKU.Name) == "SIZE" || strings.ToUpper(attrSKU.Name) == "UKURAN" ||
								strings.ToUpper(attrSKU.Name) == "VARIAN" || strings.ToUpper(attrSKU.Name) == "VARIASI" {
								varian = attrSKU.ValueName
							}

						}

					}

					statusnya := "3936a00a-e321-f5ff-33b5-2dfb8d68ba45"
					if Isiproducts.Status != 4 { //4 status live
						statusnya = "c895af78-1488-cf0b-d17a-33770f2b20bf"
					}

					if CekMapping.ChannelCode == "" && CekMapping.IdSkuParent != Isiproducts.Id {
						//insert
						objMappingLoop := models.TableSkuMapping{
							UuidSkuMapping: uuid.New().String(),
							SkuNo:          skunya,
							ChannelCode:    kode_tiktok,
							CreatedBy:      "SYSTEM",
							CreatedDate:    time.Now(),
							IdSkuParent:    Isiproducts.Id,
							IdSkuChild1:    detailSKU.Id,
							ProductName:    Isiproducts.Name,
							Varian:         varian,
							Color:          color,
							UuidStatus:     statusnya,
						}

						objMapping = append(objMapping, objMappingLoop)

					} else {
						//insert
						CekChild1, _ := tokenRepository.CariSkuMappingByModel(skunya, kode_tiktok, detailSKU.Id)
						if CekChild1.IdSkuParent == "" {
							objMappingLoop := models.TableSkuMapping{
								UuidSkuMapping: uuid.New().String(),
								SkuNo:          skunya,
								ChannelCode:    kode_tiktok,
								CreatedBy:      "SYSTEM",
								CreatedDate:    time.Now(),
								IdSkuParent:    Isiproducts.Id,
								IdSkuChild1:    detailSKU.Id,
								ProductName:    Isiproducts.Name,
								Varian:         varian,
								Color:          color,
								UuidStatus:     statusnya,
							}
							objMapping = append(objMapping, objMappingLoop)
							SaveErrorString("tiktok", "SKU:"+skunya+" OLD:"+CekMapping.IdSkuParent+" | NEW:"+Isiproducts.Id, "SKU SUDAH ADA")
						} else {
							//update
							objMappingUpd := models.TableSkuMapping{
								UuidSkuMapping: CekMapping.UuidSkuMapping,
								SkuNo:          skunya,
								Varian:         varian,
								ProductName:    Isiproducts.Name,
								Color:          color,
								UuidStatus:     statusnya,
							}
							//fmt.Println(objMappingUpd)
							result_upd := tokenRepository.UpdateSaveSkuMapping(objMappingUpd)
							if result_upd != nil {
								fmt.Println("Update SaveSkuMapping Tiktok Gagal")
							}

						}

					}

					if len(objMapping) > 0 {
						result := tokenRepository.SaveSkuMapping(objMapping)
						if result != nil {
							fmt.Println("SaveSkuMappingTiktok Gagal")
						}
					}
					objMapping = nil

				}
			}

		}
		//fmt.Println("===================================")

	}

}

func CariSkuMappingObj(id string) []models.TableSkuMapping {
	CekSkuMapping, _ := tokenRepository.CariSkuMappingObj(id)
	return CekSkuMapping
}
func CariSkuMappingObjGroup(id string) []models.TableSkuMapping {
	CekSkuMapping, _ := tokenRepository.CariSkuMappingObjGroup(id)
	return CekSkuMapping
}

func CariUidAvlb(sku string) []models.DetailOrderPraPicking {
	ObjAvlb, _ := tokenRepository.CariUidAvlb(sku)
	return ObjAvlb
}
func CariUidSalesOrder(sku string, param string) int64 {
	TotalSalesOrder, _ := tokenRepository.CariUidSalesOrder(sku, param)
	return TotalSalesOrder
}

func BookingSku(param string) string {
	pesan := ""
	status_pick := ""
	var objs []models.GroupSkuPicking
	//group sku lihat qty
	objs, _ = bookingRepository.GroupSku(param)
	//var skulastpick []models.TableProductUidBasket
	total_qty := 0
	var objSalesBooking []models.TableSalesBooking
	var objStatusProduct []models.UpdateProductByUid

	for _, element1 := range objs {
		//cek ke picking jika sudah pernah di pick harus basket sama

		qty_pick := 0
		obj_detail_pick, _ := bookingRepository.DetailSudahPicking(param, element1.SkuNo)

		qty_pick = len(obj_detail_pick)
		qty_order, _ := strconv.Atoi(element1.Qty)

		CekBooking, _ := bookingRepository.CekSalesBooking(param)
		if len(CekBooking) > 0 {
			pesan = " Sudah Ada Di Booking"
		} else {

			total_qty += qty_order //return total qty
			qty_belum_pick := qty_order - qty_pick

			element1.Qty = strconv.Itoa(qty_belum_pick)
			//obj_detail, _ := pickingRepository.GetDetailOrder(element1.SkuNo, "1", "0")//asli 13-10-2021
			obj_detail, _ := bookingRepository.GetDetailOrder(element1.SkuNo, element1.Qty, "0", status_pick)

			if len(obj_detail) < 1 {
				// res.ResponseCode = http.StatusOK
				// res.ResponseDesc = enums.ERROR
				// res.ResponseTime = utils.DateToStdNow()
				// res.Message = "SKU " + element1.SkuNo + " Stock Tidak Mencukupi "
				// res.Result = nil
				// return res, nil
				pesan = "SKU " + element1.SkuNo + " Stock Tidak Mencukupi "
				return pesan
			}

			hit_qty, _ := strconv.Atoi(element1.Qty)

			//tambahan jika sku sudah picking semua 01-11-2021
			if hit_qty == 0 {
				hit_qty = 1
			}
			//end tambahan sku jika sudah picking semua

			if len(obj_detail) != hit_qty { //tambahan 13-10-2021 cek qty masih ada atau tidak
				// res.ResponseCode = http.StatusOK
				// res.ResponseDesc = enums.ERROR
				// res.ResponseTime = utils.DateToStdNow()
				// res.Message = "SKU " + element1.SkuNo + " Stock Tidak Mencukupi, "
				// res.Result = nil
				// return res, nil

				pesan = "SKU " + element1.SkuNo + " Stock Tidak Mencukupi, "
				return pesan

			}

			//cari total qty di putway berdasarkan racking
			TotalSkuRacking, _ := bookingRepository.CariTotal(element1.SkuNo, element1.Qty, status_pick)
			fmt.Println(len(TotalSkuRacking))

			ObjBooking, _ := bookingRepository.ViewCariTotal(element1.SkuNo, element1.Qty, status_pick)
			//spew.Dump(ObjBooking)
			DateBooking := time.Now()
			for _, elementBooking := range ObjBooking {

				objStatusProductLoop := models.UpdateProductByUid{
					Uid:        elementBooking.Uid,
					UuidStatus: "eff09170-9310-42d2-8c5d-a1f510f4a809", //RSRV
				}
				objStatusProduct = append(objStatusProduct, objStatusProductLoop)

				objSalesBookingLoop := models.TableSalesBooking{
					UuidSalesBooking: uuid.New().String(),
					NoOrder:          param,
					SkuNo:            element1.SkuNo,
					Uid:              elementBooking.Uid,
					RackingNo:        elementBooking.RackingNo,
					AisleShelves:     elementBooking.AisleShelves,
					CreatedBy:        "SYSTEM",
					CreatedDate:      DateBooking,
					ProductName:      obj_detail[0].ProductName,
				}
				objSalesBooking = append(objSalesBooking, objSalesBookingLoop)
			}

		}

		//}
	}

	if len(objSalesBooking) > 0 {
		//cek jika sudah di ambil orderan lain
		errs := bookingRepository.StoreBooking(objSalesBooking, objStatusProduct)
		pesan = "Sukses"
		if errs != nil {
			if errs.Error() == `pq: duplicate key value violates unique constraint "wms_sales_booking_pk"` {
				pesan = "Coba Picking Kembali"
			} else {
				pesan = "Check Table"
			}

			return pesan
		}
	}

	return pesan
}

func UpdateStockOrder(param string) {
	objCariOrder, _ := tokenRepository.FindSalesOrderArray(param)
	for _, valueCariOrder := range objCariOrder {
		//helpers.UpdateStock(valueCariOrder.SkuNo, "API_CHANNEL", "")
		GroupChannelSkuMap(valueCariOrder.SkuNo)
	}
}

func SkuPengganti(channel string, noorder string, skunya string) string {
	skureplace := ""
	objSkuPengganti := tokenRepository.ChangeSkuChannelWMS(channel, noorder, skunya)
	if len(objSkuPengganti) > 0 {

		for _, values := range objSkuPengganti {

			if values.SkuOrigin == skunya {
				fmt.Println("ada pengganti")
				skureplace = values.SkuReplace
			}

		}

	}
	return skureplace
}

func SaveUploadImage(obj models.UploadImageShopeeInfo, path string, namafile string, username string) {

	var objImage models.TableUploadImageChannel
	if obj.ImageId != "" {

		imageUrl := ""
		for _, value := range obj.ImageUrlList {
			if value.ImageUrlRegion == "ID" {
				imageUrl = value.ImageUrl
			}
		}

		// Read the entire file into a byte slice
		bytes, err := ioutil.ReadFile(path)
		if err != nil {
			fmt.Println(err)
		}

		base64Encoding := base64.StdEncoding.EncodeToString(bytes)

		objImage.UuidUploadImage = uuid.New().String()
		objImage.NameFile = namafile
		objImage.FileBase64 = base64Encoding
		objImage.CreatedBy = username
		objImage.CreatedDate = time.Now()
		objImage.PathImage = path
		objImage.ImageId = obj.ImageId
		objImage.ImageUrl = imageUrl
		objImage.ChannelCode = os.Getenv("KODE_SHOPEE")

		result := tokenRepository.SaveUploadImage(objImage)
		if result != nil {
			fmt.Println("SaveUploadImage Gagal")
		}
	}

}

func SaveOrderObjBlibli(obj []models.OrderItemsBlibli, objCust models.ListDetailOrderBlibli) { //sini
	var objSalesOrder []models.TableSalesOrder
	var objSalesOrderTotalQty models.TableSalesOrder
	createdDate := time.Now()
	totalQty := 0

	NoOrder := ""
	TempIdItem := ""
	TempDeliveryType := ""
	StatusOrder := ""

	//cari orderid paling kecil
	TempIdItemOrder := 0
	objOrderID := obj
	for _, valOrderIdTemp := range objOrderID {
		idItemOrder, _ := strconv.Atoi(valOrderIdTemp.Order.ItemId)
		if TempIdItemOrder == 0 {
			TempIdItemOrder = idItemOrder
		} else if TempIdItemOrder >= idItemOrder {
			TempIdItemOrder = idItemOrder
		}

	}

	for _, valOrderItems := range obj {
		// NoOrder = valOrderItems.Order.Id + "-" + valOrderItems.Order.ItemId
		// NoOrder = valOrderItems.Order.Id
		if TempIdItem == "" {
			//TempIdItem = valOrderItems.Order.ItemId
			TempIdItem = strconv.Itoa(TempIdItemOrder)
			TempDeliveryType = valOrderItems.SellerDeliveryType
			StatusOrder = valOrderItems.Order.ItemStatus
		}
		CodeKurir := TempDeliveryType
		NoOrder = valOrderItems.Order.Id + "-" + TempIdItem
		// fmt.Println("=======================")
		// fmt.Println(NoOrder)
		// fmt.Println("=======================")
		objCek, _ := tokenRepository.FindSalesOrder(NoOrder)
		if objCek.NoOrder == "" && StatusOrder == "FP" {

			//objChannel, _ := tokenRepository.FindChannel("BLIBLI")
			//if objChannel.UuidChannel != "" {

			//}
			receiptname := objCust.Content.Recipient.Name
			telp := objCust.Content.Recipient.PhoneNumber
			address := objCust.Content.Recipient.StreetAddress
			city := objCust.Content.Recipient.City
			province := objCust.Content.Recipient.State

			skunya := valOrderItems.Product.SellerSku
			qty := valOrderItems.Order.Quantity
			if len([]rune(skunya)) == 8 {
				//skunya = skunya
			} else if len([]rune(skunya)) > 8 {
				sku := []rune(skunya)
				skunya = string(sku[1:9])
			} else if len([]rune(skunya)) == 7 {
				skunya = "0" + skunya
			} else {
				skunya = "XXXXXXXX"
			}
			sendBeforeDate := createdDate.AddDate(0, 0, +int(3))

			if DigitsCount(int(valOrderItems.Order.AutoCancelTimestamp)) > 10 {
				s := valOrderItems.Order.AutoCancelTimestamp
				IntString := strconv.Itoa(int(s))

				IntStringNew := IntString[0:10]
				// snew := strconv.Itoa(IntString[0:10])
				ss, _ := strconv.Atoi(IntStringNew)
				sendBeforeDate = time.Unix(int64(ss), 0)
			}

			OrderDate := createdDate
			if DigitsCount(int(valOrderItems.CreatedDate)) > 10 {
				s := valOrderItems.CreatedDate
				IntString := strconv.Itoa(int(s))

				IntStringNew := IntString[0:10]
				// snew := strconv.Itoa(IntString[0:10])
				ss, _ := strconv.Atoi(IntStringNew)

				OrderDate = time.Unix(int64(ss), 0)
			}

			//}
			// fmt.Println("=========== BEFORE DATE ===========")
			// fmt.Println(sendBeforeDate)
			// fmt.Println(sendBeforeDate)
			// fmt.Println("=========== BEFORE DATE ===========")

			totalQty = totalQty + int(qty)
			id := uuid.New()

			objSalesOrderLoop := models.TableSalesOrder{
				UuidSalesOrder:     id.String(),
				NoOrder:            NoOrder,
				StatusOrder:        "PDKM",
				SkuNo:              skunya,
				ProductName:        valOrderItems.Product.ItemName,
				Qty:                strconv.Itoa(int(qty)),
				ExpeditionType:     valOrderItems.Logistic.ProductName,
				DeliveryType:       TempDeliveryType,
				TotalQtyOrder:      "30",
				StatusProcessOrder: "0",
				WeightProduk:       "0",
				TotalWeight:        "0",

				SentBeforeDate: sendBeforeDate,
				CreatedDate:    OrderDate,

				CreatedBy:   "BLIBLI",
				NameVariant: "", //value.ModelName,
				//UsernameBuyer: obj.Response.OrderList[0].BuyerUsername,
				ChannelCode: os.Getenv("KODE_BLIBLI"),

				RecipientsName:  receiptname,
				Telephone:       telp,
				ShippingAddress: address,
				City:            city,
				Province:        province,
			}
			// fmt.Println(objSalesOrderLoop)
			// fmt.Println("==========================")
			objSalesOrder = append(objSalesOrder, objSalesOrderLoop)

			objSalesOrderTotalQty.NoOrder = NoOrder
			objSalesOrderTotalQty.TotalQtyOrder = strconv.Itoa(totalQty)

		} else {

			if NoOrder == objCek.NoOrder {

				if objCek.ExpeditionType != CodeKurir { //update kurir

					var objUpdateKurir models.TableSalesOrder
					objUpdateKurir.NoOrder = objCek.NoOrder
					objUpdateKurir.ExpeditionType = CodeKurir
					objUpdateKurir.CreatedBy = "BLIBLI"
					result := tokenRepository.UpdateKurirNew(objUpdateKurir)
					if result != nil {
						fmt.Println("UpdateKurir Gagal")
					} else {
						fmt.Println("UpdateKurir Success")
					}

				}

				if objCek.StatusProcessOrder == "0" {

					if StatusOrder == "X" { //cancel

						objSalesOrderLoop := models.TableSalesOrder{
							NoOrder:            objCek.NoOrder,
							StatusProcessOrder: "3",
							CreatedBy:          "BLIBLI",
						}
						result := tokenRepository.UpdateSalesOrderCanc(objSalesOrderLoop)
						if result != nil {
							fmt.Println("UpdateSalesOrderCanc Gagal")
						} else {
							fmt.Println("UpdateSalesOrderCanc Success")
						}
						UpdateStockOrder(objCek.NoOrder)

					}
				}

				if objCek.StatusProcessOrder == "1" {

					if StatusOrder == "D" {
						payoutdate := time.Now()

						objSalesOrderLoop := models.TableSalesOrder{
							NoOrder:    objCek.NoOrder,
							PayoutDate: payoutdate,
							CreatedBy:  "BLIBLI",
						}
						result := tokenRepository.UpdateSalesOrderCanc(objSalesOrderLoop)
						if result != nil {
							fmt.Println("Update Payoutdate Gagal")
						} else {
							fmt.Println("Update Payoutdate Success")
						}
					}

				}
			}

		}

	}

	if len(objSalesOrder) > 0 {
		//fmt.Println(len(objSalesOrder))
		//fmt.Println("==== SELSAI ====")
		result := tokenRepository.SaveSalesOrder(objSalesOrder, objSalesOrderTotalQty)
		if result != nil {
			fmt.Println("SaveOrderObjBlibli Gagal")
		} else {
			fmt.Println("SaveOrderObjBlibli Success")
			pesan := BookingSku(objSalesOrderTotalQty.NoOrder)
			if pesan != "Sukses" {
				//delete sales order
				//tokenRepository.DeleteSalesOrder(objSalesOrderTotalQty)

			}
			fmt.Println(objSalesOrderTotalQty.NoOrder + " => " + pesan)

			//cari detail no order
			UpdateStockOrder(objSalesOrderTotalQty.NoOrder)

		}
	}

}

func SaveSkuMappingBlibli(objData models.ListProductObjBlibli) {
	var objMapping []models.TableSkuMapping

	skunya := objData.MerchantSku

	statusnya := "3936a00a-e321-f5ff-33b5-2dfb8d68ba45"
	if objData.Buyable != true {
		//if objData.Displayable != true {
		statusnya = "c895af78-1488-cf0b-d17a-33770f2b20bf"
	}

	if len([]rune(skunya)) == 8 {
		//skunya = skunya
	} else if len([]rune(skunya)) > 8 {
		sku := []rune(skunya)
		skunya = string(sku[1:9])
	} else {

		if len([]rune(skunya)) < 8 {
			skunya = "0" + skunya
		} else {
			skunya = "XXXXXXXX"
		}

	}

	CekSkuWMS, _ := tokenRepository.CariSkuWMS(skunya)
	if CekSkuWMS.SkuNo != "" {
		CekMapping, _ := tokenRepository.CariSkuMapping(skunya, os.Getenv("KODE_BLIBLI"))
		if CekMapping.ChannelCode == "" { //insert
			objMappingLoop := models.TableSkuMapping{
				UuidSkuMapping: uuid.New().String(),
				SkuNo:          skunya,
				ChannelCode:    os.Getenv("KODE_BLIBLI"),
				CreatedBy:      "SYSTEM",
				CreatedDate:    time.Now(),
				IdSkuParent:    objData.ProductSku, //objData.GdnSku,
				IdSkuChild1:    objData.GdnSku,     //objData.ProductSku,
				ProductName:    objData.ProductName,
				// Varian:      varian,
				// Color:       color,
				UuidStatus: statusnya,
			}

			//fmt.Println("uuid " + objMappingLoop.UuidSkuMapping + " | " + strconv.Itoa(IDnya) + " | " + skunya)
			objMapping = append(objMapping, objMappingLoop)

		} else {

			//CekParent, _ := tokenRepository.CariSkuMappingByParent(skunya, os.Getenv("KODE_BLIBLI"), objData.GdnSku)
			CekParent, _ := tokenRepository.CariSkuMappingByParent(skunya, os.Getenv("KODE_BLIBLI"), objData.ProductSku)
			if CekParent.IdSkuParent == "" {
				//CekChild1, _ := tokenRepository.CariSkuMappingByModel(skunya, os.Getenv("KODE_BLIBLI"), objData.ProductSku)
				CekChild1, _ := tokenRepository.CariSkuMappingByModel(skunya, os.Getenv("KODE_BLIBLI"), objData.GdnSku)
				if CekChild1.IdSkuParent == "" {
					objMappingLoop := models.TableSkuMapping{
						UuidSkuMapping: uuid.New().String(),
						SkuNo:          skunya,
						ChannelCode:    os.Getenv("KODE_BLIBLI"),
						CreatedBy:      "SYSTEM",
						CreatedDate:    time.Now(),
						IdSkuParent:    objData.ProductSku, //objData.GdnSku,
						IdSkuChild1:    objData.GdnSku,     //objData.ProductSku,
						ProductName:    objData.ProductName,
						// Varian:      varian,
						// Color:       color,
						UuidStatus: statusnya,
					}
					objMapping = append(objMapping, objMappingLoop)
					//SaveErrorString("blibli", "SKU:"+skunya+" OLD:"+CekChild1.IdSkuParent+" | NEW:"+objData.GdnSku, "SKU SUDAH ADA")
					SaveErrorString("blibli", "SKU:"+skunya+" OLD:"+CekChild1.IdSkuParent+" | NEW:"+objData.ProductSku, "SKU SUDAH ADA")
				} else {

					objMappingUpd := models.TableSkuMapping{
						UuidSkuMapping: CekChild1.UuidSkuMapping,
						SkuNo:          skunya,
						ProductName:    objData.ProductName,
						UuidStatus:     statusnya,
					}
					//fmt.Println(objMappingUpd)
					result_upd := tokenRepository.UpdateSaveSkuMapping(objMappingUpd)
					if result_upd != nil {
						fmt.Println("Update SaveSkuMappingBlibli Gagal")
					}

				}
			} else {
				objMappingUpd := models.TableSkuMapping{
					UuidSkuMapping: CekParent.UuidSkuMapping,
					SkuNo:          skunya,
					ProductName:    objData.ProductName,
					UuidStatus:     statusnya,
				}
				//fmt.Println(objMappingUpd)
				result_upd := tokenRepository.UpdateSaveSkuMapping(objMappingUpd)
				if result_upd != nil {
					fmt.Println("Update SaveSkuMappingBlibli Gagal")
				}
			}

		}

		result := tokenRepository.SaveSkuMapping(objMapping)
		if result != nil {
			fmt.Println("SaveSkuMappingBlibli Gagals")
		}
		objMapping = nil
	}

}

func DigitsCount(number int) int {
	count := 0
	for number != 0 {
		number /= 10
		count += 1
	}
	return count

}

func SaveSkuMappingJdId(obj models.DetailProductJdId, IDnya int, Statussku int) {
	var objMapping []models.TableSkuMapping
	objData := obj.JingdongSellerProductGetSkuInfoBySpuIDAndVenderIDResponse.ReturnType.Model

	kode_jdid := os.Getenv("KODE_JDID")
	// tempsku := ""
	for _, elementModel := range objData {
		//fmt.Println("========= DETAIL SKU =========")
		//fmt.Println(elementModel.SellerSkuID)
		skunya := elementModel.SellerSkuID
		if len([]rune(skunya)) == 8 {
			//skunya = skunya
		} else if len([]rune(skunya)) > 8 {
			sku := []rune(skunya)
			skunya = string(sku[1:9])
		} else {

			if len([]rune(skunya)) < 8 {
				skunya = "0" + skunya
			} else {
				skunya = "XXXXXXXX"
			}

		}

		helpers.UpdateStockNol(skunya, "AUTO", kode_jdid)
		//update jadi inaktif
		//fmt.Println("SaveSkuMappingJdId " + skunya)
		statusnya := "c895af78-1488-cf0b-d17a-33770f2b20bf"
		objInaTokped := models.TableSkuMapping{
			SkuNo:       skunya,
			ChannelCode: kode_jdid,
			UuidStatus:  statusnya,
		}

		resultIna := tokenRepository.UpdateSkuMappingByChannelSKu(objInaTokped)
		if resultIna != nil {
			fmt.Println("Update SaveSkuMappingJdId Gagal")
		}

		statusnya = "3936a00a-e321-f5ff-33b5-2dfb8d68ba45"
		//if Statussku != 1 {
		if elementModel.Status != 1 {

			statusnya = "c895af78-1488-cf0b-d17a-33770f2b20bf"
		}

		CekSkuWMS, _ := tokenRepository.CariSkuWMS(skunya)

		//inaktfid
		if CekSkuWMS.SkuNo != "" {
			CekMapping, _ := tokenRepository.CariSkuMapping(skunya, kode_jdid)
			if CekMapping.ChannelCode == "" { //insert
				objMappingLoop := models.TableSkuMapping{
					UuidSkuMapping: uuid.New().String(),
					SkuNo:          skunya,
					ChannelCode:    kode_jdid,
					CreatedBy:      "SYSTEM",
					CreatedDate:    time.Now(),
					IdSkuParent:    strconv.Itoa(IDnya),
					IdSkuChild1:    strconv.Itoa(int(elementModel.SkuID)),
					ProductName:    elementModel.SkuName,
					// Varian:      varian,
					// Color:       color,
					UuidStatus: statusnya,
				}

				//fmt.Println("uuid " + objMappingLoop.UuidSkuMapping + " | " + strconv.Itoa(IDnya) + " | " + skunya)
				objMapping = append(objMapping, objMappingLoop)

			} else {
				CekParent, _ := tokenRepository.CariSkuMappingByParent(skunya, kode_jdid, strconv.Itoa(IDnya))
				if CekParent.IdSkuParent == "" {
					CekChild1, _ := tokenRepository.CariSkuMappingByModel(skunya, kode_jdid, strconv.Itoa(int(elementModel.SkuID)))
					if CekChild1.IdSkuParent == "" {
						objMappingLoop := models.TableSkuMapping{
							UuidSkuMapping: uuid.New().String(),
							SkuNo:          skunya,
							ChannelCode:    kode_jdid,
							CreatedBy:      "SYSTEM",
							CreatedDate:    time.Now(),
							IdSkuParent:    strconv.Itoa(IDnya),
							IdSkuChild1:    strconv.Itoa(int(elementModel.SkuID)),
							ProductName:    elementModel.SkuName,
							// Varian:      varian,
							// Color:       color,
							UuidStatus: statusnya,
						}
						objMapping = append(objMapping, objMappingLoop)
						SaveErrorString("jdid", "SKU:"+skunya+" OLD:"+CekChild1.IdSkuParent+" | NEW:"+strconv.Itoa(IDnya), "SKU SUDAH ADA")
					} else {
						objMappingUpd := models.TableSkuMapping{
							UuidSkuMapping: CekChild1.UuidSkuMapping,
							SkuNo:          skunya,
							ProductName:    elementModel.SkuName,
							UuidStatus:     statusnya,
						}
						//fmt.Println(objMappingUpd)
						result_upd := tokenRepository.UpdateSaveSkuMapping(objMappingUpd)
						if result_upd != nil {
							fmt.Println("Update SaveSkuMappingJdId Gagal")
						}
					}
				} else {

					objMappingUpd := models.TableSkuMapping{
						UuidSkuMapping: CekParent.UuidSkuMapping,
						SkuNo:          skunya,
						ProductName:    elementModel.SkuName,
						UuidStatus:     statusnya,
					}
					//fmt.Println(objMappingUpd)
					result_upd := tokenRepository.UpdateSaveSkuMapping(objMappingUpd)
					if result_upd != nil {
						fmt.Println("Update SaveSkuMappingJdId Gagal")
					}

				}

			}

		}
		//fmt.Println(objMapping)

		//fmt.Println(IDnya)
		//fmt.Println("========= ========== =========")
	}

	if len(objMapping) > 0 {
		result := tokenRepository.SaveSkuMapping(objMapping)
		if result != nil {
			fmt.Println("SaveSkuMappingJdId Gagals")
		}
	}

}

func SaveSalesOrderJdId(obj models.OrderDetailJdId) {
	datanya := obj.JingdongSellerOrderGetOrderInfoByOrderIDResponse.Result.Model

	var objSalesOrder []models.TableSalesOrder
	var objSalesOrderTotalQty models.TableSalesOrder
	createdDate := time.Now()
	totalQty := 0

	NoOrder := ""

	for _, valOrderItems := range datanya.OrderSkuinfos {
		if NoOrder == "" {
			fmt.Println("SaveSalesOrderJdId " + strconv.Itoa(datanya.OrderID))
		}
		NoOrder = strconv.Itoa(datanya.OrderID)
		objCek, _ := tokenRepository.FindSalesOrder(NoOrder)
		CodeKurir := strconv.Itoa(datanya.CarrierCode)
		if objCek.NoOrder == "" && CodeKurir != "0" && CodeKurir != "" {
			receiptname := datanya.CustomerName
			telp := datanya.Mobile
			address := datanya.Address
			city := datanya.City
			province := datanya.State

			skunya := valOrderItems.PopSkuID
			qty := valOrderItems.SkuNumber
			if len([]rune(skunya)) == 8 {
				//skunya = skunya
			} else if len([]rune(skunya)) > 8 {
				sku := []rune(skunya)
				skunya = string(sku[1:9])
			} else if len([]rune(skunya)) == 7 {
				skunya = "0" + skunya
			} else {
				skunya = "XXXXXXXX"
			}

			if valOrderItems.PopSkuID == "a111111" {
				skunya = "07909877"
			}

			if valOrderItems.SkuID == 674134376 {
				skunya = "08156215"
			}

			if valOrderItems.SkuID == 674177247 {
				skunya = "08151225"
			}

			if valOrderItems.SkuID == 623301335 {
				skunya = "08128393"
			}

			sendBeforeDate := createdDate.AddDate(0, 0, +int(3))

			OrderDate := time.Unix(datanya.CreateTime, 0)

			if DigitsCount(int(datanya.CreateTime)) > 10 {
				s := datanya.CreateTime
				IntString := strconv.Itoa(int(s))

				IntStringNew := IntString[0:10]
				// snew := strconv.Itoa(IntString[0:10])
				ss, _ := strconv.Atoi(IntStringNew)
				OrderDate = time.Unix(int64(ss), 0)
			}

			id := uuid.New()
			totalQty = totalQty + int(qty)
			objSalesOrderLoop := models.TableSalesOrder{
				UuidSalesOrder:     id.String(),
				NoOrder:            NoOrder,
				StatusOrder:        "PDKM",
				SkuNo:              skunya,
				ProductName:        valOrderItems.SkuName,
				Qty:                strconv.Itoa(int(qty)),
				ExpeditionType:     CodeKurir,
				DeliveryType:       "Pickup",
				TotalQtyOrder:      "30",
				StatusProcessOrder: "0",
				WeightProduk:       "0",
				TotalWeight:        "0",
				OrderDate:          OrderDate,
				SentBeforeDate:     sendBeforeDate,
				CreatedDate:        OrderDate,

				CreatedBy:   "JDID",
				NameVariant: "", //value.ModelName,
				//UsernameBuyer: obj.Response.OrderList[0].BuyerUsername,
				ChannelCode: os.Getenv("KODE_JDID"),

				RecipientsName:  receiptname,
				Telephone:       telp,
				ShippingAddress: address,
				City:            city,
				Province:        province,
			}
			objSalesOrder = append(objSalesOrder, objSalesOrderLoop)

			objSalesOrderTotalQty.NoOrder = NoOrder
			objSalesOrderTotalQty.TotalQtyOrder = strconv.Itoa(totalQty)

		} else if objCek.NoOrder == strconv.Itoa(datanya.OrderID) {

			if objCek.ExpeditionType != datanya.CarrierCompany { //update kurir

				var objUpdateKurir models.TableSalesOrder
				objUpdateKurir.NoOrder = objCek.NoOrder
				objUpdateKurir.ExpeditionType = datanya.CarrierCompany
				objUpdateKurir.CreatedBy = "JDID"
				result := tokenRepository.UpdateKurirNew(objUpdateKurir)
				if result != nil {
					fmt.Println("UpdateKurir Gagal")
				} else {
					fmt.Println("UpdateKurir Success")
				}

			}

			if objCek.StatusProcessOrder == "0" {
				//jika CANCELED UPDATE status_processs_order jadi 3
				if datanya.OrderState == 5 { //5 CANCELLED
					objSalesOrderLoop := models.TableSalesOrder{
						NoOrder:            objCek.NoOrder,
						StatusProcessOrder: "3",
						CreatedBy:          "JDID",
					}
					result := tokenRepository.UpdateSalesOrderCanc(objSalesOrderLoop)
					if result != nil {
						fmt.Println("UpdateSalesOrderCanc Gagal")
					} else {
						fmt.Println("UpdateSalesOrderCanc Success")
					}
					UpdateStockOrder(objCek.NoOrder)

				}
			}

			if objCek.StatusProcessOrder == "1" {
				//jika CANCELED UPDATE status_processs_order jadi 3
				if datanya.OrderState == 6 {
					payoutdate := time.Now()

					if DigitsCount(int(datanya.ModifyTime)) > 10 {
						s := datanya.ModifyTime
						IntString := strconv.Itoa(int(s))
						IntStringNew := IntString[0:10]
						ss, _ := strconv.Atoi(IntStringNew)
						payoutdate = time.Unix(int64(ss), 0)
					}

					objSalesOrderLoop := models.TableSalesOrder{
						NoOrder:    objCek.NoOrder,
						PayoutDate: payoutdate,
						CreatedBy:  "JDID",
					}
					result := tokenRepository.UpdateSalesOrderCanc(objSalesOrderLoop)
					if result != nil {
						fmt.Println("Update Payoutdate Gagal")
					} else {
						fmt.Println("Update Payoutdate Success")
					}
				}

			}

		}

	}
	if len(objSalesOrder) > 0 {
		//fmt.Println(len(objSalesOrder))
		//fmt.Println("==== SELSAI ====")
		result := tokenRepository.SaveSalesOrder(objSalesOrder, objSalesOrderTotalQty)
		if result != nil {
			fmt.Println("SaveSalesOrderJdId Gagal")
		} else {
			fmt.Println("SaveSalesOrderJdId Success")
			pesan := BookingSku(objSalesOrderTotalQty.NoOrder)
			if pesan != "Sukses" {
				//delete sales order
				//tokenRepository.DeleteSalesOrder(objSalesOrderTotalQty)

			}
			fmt.Println(objSalesOrderTotalQty.NoOrder + " => " + pesan)

			//cari detail no order
			UpdateStockOrder(objSalesOrderTotalQty.NoOrder)

		}
	}

	//fmt.Println(NoOrder)

}

func SaveTokenTokped(objToken models.TokenTokped, id string) models.TableApiChannel {
	var objs models.TableApiChannel

	objs, _ = tokenRepository.GetDataToken(id)
	msgErr := ""
	objs.Key1 = "access_token"
	objs.Value1 = objToken.AccessToken
	objs.Key2 = "refresh_token"
	objs.Value2 = ""

	objs.Id = id
	objs.CreatedDate = time.Now()
	if objs.Key1 != "" {
		result := tokenRepository.SaveTokenShopee(objs)
		if result != nil {
			fmt.Println("GAGAL")
		} else {
			fmt.Println("Save Token Success")
		}
	} else {
		msgErr = "Gagal Save token " + id
		var objError models.TableLogErrChannel
		objError.Id = id
		objError.Message = msgErr
		objError.CreatedDate = time.Now()
		objError.Process = "SaveToken"

		result := tokenRepository.SaveChannelError(objError)
		if result != nil {
			fmt.Println("GAGAL")
		} else {
			fmt.Println("Save Error Success")
		}
	}

	return objs
}

func SaveTokenZalora(objToken models.TokenZalora) models.TableApiChannel {
	var objs models.TableApiChannel

	id := "zalora"

	objs, _ = tokenRepository.GetDataToken(id)
	msgErr := ""
	objs.Key1 = "access_token"
	objs.Value1 = objToken.AccessToken
	objs.Key2 = "refresh_token"
	objs.Value2 = ""

	objs.Id = id
	objs.CreatedDate = time.Now()
	if objs.Key1 != "" {
		result := tokenRepository.SaveTokenShopee(objs)
		if result != nil {
			fmt.Println("GAGAL")
		} else {
			fmt.Println("Save Token Success")
		}
	} else {
		msgErr = "Gagal Save token " + id
		var objError models.TableLogErrChannel
		objError.Id = id
		objError.Message = msgErr
		objError.CreatedDate = time.Now()
		objError.Process = "SaveToken"

		result := tokenRepository.SaveChannelError(objError)
		if result != nil {
			fmt.Println("GAGAL")
		} else {
			fmt.Println("Save Error Success")
		}
	}

	return objs
}

func SaveSkuMappingLazada(obj models.ProductDetailLazada, IDnya int) {
	var objMapping []models.TableSkuMapping
	objData := obj.Data.Skus

	kode_lazada := os.Getenv("KODE_LAZADA")
	// tempsku := ""
	for _, elementModel := range objData {
		//fmt.Println("========= DETAIL SKU =========")
		//fmt.Println(elementModel.SellerSkuID)
		skunya := elementModel.SellerSku
		statussku := elementModel.Status
		statusnya := "3936a00a-e321-f5ff-33b5-2dfb8d68ba45"

		if strings.ToUpper(statussku) == "DELETED" || strings.ToUpper(statussku) == "SUSPENDED" {
			statusnya = "c895af78-1488-cf0b-d17a-33770f2b20bf"
		}

		if len([]rune(skunya)) == 8 {
			//skunya = skunya
		} else if len([]rune(skunya)) > 8 {
			sku := []rune(skunya)
			skunya = string(sku[1:9])
		} else {

			if len([]rune(skunya)) < 8 {
				skunya = "0" + skunya
			} else {
				skunya = "XXXXXXXX"
			}

		}

		CekSkuWMS, _ := tokenRepository.CariSkuWMS(skunya)
		if CekSkuWMS.SkuNo != "" {
			CekMapping, _ := tokenRepository.CariSkuMapping(skunya, kode_lazada)
			if CekMapping.ChannelCode == "" { //insert
				objMappingLoop := models.TableSkuMapping{
					UuidSkuMapping: uuid.New().String(),
					SkuNo:          skunya,
					ChannelCode:    kode_lazada,
					CreatedBy:      "SYSTEM",
					CreatedDate:    time.Now(),
					IdSkuParent:    strconv.Itoa(IDnya),
					IdSkuChild1:    strconv.Itoa(int(elementModel.SkuID)),
					ProductName:    obj.Data.Attributes.Name,
					// Varian:      varian,
					// Color:       color,
					UuidStatus: statusnya,
				}

				//fmt.Println("uuid " + objMappingLoop.UuidSkuMapping + " | " + strconv.Itoa(IDnya) + " | " + skunya)
				objMapping = append(objMapping, objMappingLoop)

			} else {

				CekParent, _ := tokenRepository.CariSkuMappingByParent(skunya, kode_lazada, strconv.Itoa(IDnya))
				if CekParent.IdSkuParent == "" {
					objMappingLoop := models.TableSkuMapping{
						UuidSkuMapping: uuid.New().String(),
						SkuNo:          skunya,
						ChannelCode:    kode_lazada,
						CreatedBy:      "SYSTEM",
						CreatedDate:    time.Now(),
						IdSkuParent:    strconv.Itoa(IDnya),
						IdSkuChild1:    strconv.Itoa(int(elementModel.SkuID)),
						ProductName:    obj.Data.Attributes.Name,
						// Varian:      varian,
						// Color:       color,
						UuidStatus: statusnya,
					}
					objMapping = append(objMapping, objMappingLoop)
					SaveErrorString("lazada", "SKU:"+skunya+" OLD:"+CekMapping.IdSkuParent+" | NEW:"+strconv.Itoa(IDnya), "SKU SUDAH ADA")
				} else {
					CekChild1, _ := tokenRepository.CariSkuMappingByModel(skunya, kode_lazada, strconv.Itoa(int(elementModel.SkuID)))
					if CekChild1.IdSkuParent == "" {
						objMappingLoop := models.TableSkuMapping{
							UuidSkuMapping: uuid.New().String(),
							SkuNo:          skunya,
							ChannelCode:    kode_lazada,
							CreatedBy:      "SYSTEM",
							CreatedDate:    time.Now(),
							IdSkuParent:    strconv.Itoa(IDnya),
							IdSkuChild1:    strconv.Itoa(int(elementModel.SkuID)),
							ProductName:    obj.Data.Attributes.Name,
							// Varian:      varian,
							// Color:       color,
							UuidStatus: statusnya,
						}
						objMapping = append(objMapping, objMappingLoop)
						SaveErrorString("lazada", "SKU:"+skunya+" OLD:"+CekChild1.IdSkuParent+" | NEW:"+strconv.Itoa(IDnya), "SKU SUDAH ADA")
					} else {
						objMappingUpd := models.TableSkuMapping{
							UuidSkuMapping: CekParent.UuidSkuMapping,
							SkuNo:          skunya,
							ProductName:    obj.Data.Attributes.Name,
							UuidStatus:     statusnya,
						}
						//fmt.Println(objMappingUpd)
						result_upd := tokenRepository.UpdateSaveSkuMapping(objMappingUpd)
						if result_upd != nil {
							fmt.Println("Update SaveSkuMappingLazada Gagal")
						}
					}
				}
			}

		}
	}

	if len(objMapping) > 0 {
		result := tokenRepository.SaveSkuMapping(objMapping)
		if result != nil {
			fmt.Println("SaveSkuMappingLazada Gagals")
		}
	}

}

func SaveSalesOrderLazada(obj models.OrderDetailItemLazada, objCust models.OrdersDetailGenerated) {

	var objSalesOrder []models.TableSalesOrder
	var objSalesOrderTotalQty models.TableSalesOrder
	createdDate := time.Now()
	objNew := CekArrayOrderLazada(obj)

	expeditionType := obj.Data[0].ShippingType
	fmt.Println("masuk sinii")
	for _, value := range objNew {
		fmt.Println("masuk dalem sinii")
		NoOrder := strconv.Itoa(int(objCust.Data.OrderNumber))
		if obj.Data[0].Status == "pending" {
			//fmt.Println("SaveSalesOrderLazada " + NoOrder)
			// fmt.Println(objNew)
			// fmt.Println(obj.Data)
		}

		objCek, _ := tokenRepository.FindSalesOrder(NoOrder)
		// if objCek.NoOrder == "" && obj.Data[0].Status == "pending" {
		if objCek.NoOrder == "" && value.Status == "pending" {
			fmt.Println("SaveSalesOrderLazada " + NoOrder)
			// objChannel, _ := tokenRepository.FindChannel("LAZADA")
			// if objChannel.UuidChannel != "" {
			skunya := value.Sku

			ProductName := value.Name
			Variant := value.Variation
			Kurir := "KOSONG"
			OrderDate := value.CreatedAt

			FirstName := objCust.Data.AddressShipping.FirstName
			LastName := objCust.Data.AddressShipping.LastName
			Address := objCust.Data.AddressShipping.Address1
			Address += objCust.Data.AddressShipping.Address2
			Address += objCust.Data.AddressShipping.Address3
			Address += objCust.Data.AddressShipping.Address4
			Address += objCust.Data.AddressShipping.Address5

			Phone := objCust.Data.AddressShipping.Phone
			if Phone == "" {
				Phone = objCust.Data.AddressShipping.Phone2
			}

			City := objCust.Data.AddressShipping.City
			Region := "-"
			CustFirstName := FirstName

			if value.Sku == "a08172335" {
				skunya = "a08172336"
			}

			id := uuid.New()

			if len([]rune(skunya)) == 8 {
				//skunya = skunya
			} else if len([]rune(skunya)) > 8 {
				sku := []rune(skunya)
				skunya = string(sku[1:9])
			} else if len([]rune(skunya)) == 7 {
				skunya = "0" + value.Sku
			} else {
				skunya = "XXXXXXXX"
			}

			sendBeforeDate := createdDate.AddDate(0, 0, +int(2))

			amount := value.Amount

			objSalesOrderLoop := models.TableSalesOrder{
				UuidSalesOrder: id.String(),
				NoOrder:        NoOrder,
				StatusOrder:    "PDKM",
				SkuNo:          skunya,
				ProductName:    ProductName,
				// Qty:                qty,
				Qty:                strconv.Itoa(value.Qty),
				ExpeditionType:     Kurir,
				DeliveryType:       expeditionType,
				TotalQtyOrder:      "30",
				StatusProcessOrder: "0",
				WeightProduk:       "0",
				TotalWeight:        "0",

				SentBeforeDate: sendBeforeDate,
				CreatedDate:    createdDate,

				OrderDate: OrderDate,

				CreatedBy:     "LAZADA",
				NameVariant:   Variant,
				UsernameBuyer: CustFirstName,
				ChannelCode:   os.Getenv("KODE_LAZADA"), //objChannel.ChannelCode,

				RecipientsName:  FirstName + " " + LastName,
				Telephone:       Phone,
				ShippingAddress: Address,
				City:            City,
				Province:        Region,
				Amount:          amount * float64(value.Qty),
			}

			objSalesOrder = append(objSalesOrder, objSalesOrderLoop)

			objSalesOrderTotalQty.NoOrder = NoOrder
			objSalesOrderTotalQty.TotalQtyOrder = "0"

			// } else {
			// 	fmt.Println("Channel lazada Tidak Ditemukan")
			// }

		} else {

			// if objCek.ExpeditionType != CodeKurir { //update kurir

			// 	var objUpdateKurir models.TableSalesOrder
			// 	objUpdateKurir.NoOrder = objCek.NoOrder
			// 	objUpdateKurir.ExpeditionType = CodeKurir
			// 	objUpdateKurir.CreatedBy = "LAZADA"
			// 	result := tokenRepository.UpdateKurirNew(objUpdateKurir)
			// 	if result != nil {
			// 		fmt.Println("UpdateKurir Gagal")
			// 	} else {
			// 		fmt.Println("UpdateKurir Success")
			// 	}

			// }

			if objCek.StatusProcessOrder == "0" {
				fmt.Println("masuk dalem siniixxxxxx")
				//fmt.Println("No Order " + NoOrder + " Sudah Ada")
				// if obj.Data[0].Status == "canceled" {
				if value.Status == "canceled" {

					skunya := value.Sku

					if len([]rune(skunya)) == 8 {
						//skunya = skunya
					} else if len([]rune(skunya)) > 8 {
						sku := []rune(skunya)
						skunya = string(sku[1:9])
					} else if len([]rune(skunya)) == 7 {
						skunya = "0" + value.Sku
					} else {
						skunya = "XXXXXXXX"
					}
					fmt.Println("masuk dalem siniixxxxxx " + skunya)
					//cari salesorder uuid sales order
					ObjCekSales, _ := tokenRepository.FindSalesOrderBySku(NoOrder, "0", skunya)
					if ObjCekSales.UuidSalesOrder != "" {

						objSalesOrderLoop := models.TableSalesOrder{
							UuidSalesOrder:     ObjCekSales.UuidSalesOrder,
							NoOrder:            "XX" + NoOrder,
							StatusProcessOrder: "3",
							CreatedBy:          "LAZADA",
							SkuNo:              ObjCekSales.SkuNo,
						}
						//result := tokenRepository.UpdateSalesOrderCanc(objSalesOrderLoop)
						result := tokenRepository.UpdateSalesOrderCancNew(NoOrder, objSalesOrderLoop)
						if result != nil {
							fmt.Println("UpdateSalesOrderCanc LAZADA Gagal")
						} else {
							fmt.Println("UpdateSalesOrderCanc LAZADA Success")
						}
						UpdateStockOrder(NoOrder)

					}

				}

			} else if objCek.StatusProcessOrder == "1" { //delivered
				//jika complete update ke pyoutdate
				//if objCust.Data.Statuses[0] == "delivered" {
				if value.Status == "delivered" {

					payoutdate := objCust.Data.UpdatedAt
					objSalesOrderLoop := models.TableSalesOrder{
						NoOrder:    NoOrder,
						PayoutDate: payoutdate,
						CreatedBy:  "LAZADA",
					}
					result := tokenRepository.UpdateSalesOrderCanc(objSalesOrderLoop)
					if result != nil {
						fmt.Println("Update Payoutdate LAZADA Gagal")
					} else {
						fmt.Println("Update Payoutdate LAZADA Success")
					}
					fmt.Println("======= PAYOUTDATE =======")
					fmt.Println(payoutdate)
					fmt.Println("======= PAYOUTDATE =======")
				}

			}

			//break
		}

	}

	if len(objSalesOrder) > 0 {
		result := tokenRepository.SaveSalesOrder(objSalesOrder, objSalesOrderTotalQty)
		if result != nil {
			fmt.Println("SaveSalesOrderLazada Gagal")
		} else {
			fmt.Println("SaveSalesOrderLazada Success")
			pesan := BookingSku(objSalesOrderTotalQty.NoOrder)
			fmt.Println(objSalesOrderTotalQty.NoOrder + " => " + pesan)
			UpdateStockOrder(objSalesOrderTotalQty.NoOrder)
		}
	}

}
func CekArrayOrderLazada(obj models.OrderDetailItemLazada) []models.ObjSkuZalora { //numpang di zalora

	var objArray []models.ObjSkuZalora

	for _, value := range obj.Data {

		objArrayLoop := models.ObjSkuZalora{
			Sku:              value.Sku,
			Qty:              1,
			Name:             value.Name,
			Variation:        value.Variation,
			ShipmentProvider: value.ShippingType,
			CreatedAt:        value.CreatedAt,
			OrderId:          strconv.Itoa(int(value.OrderID)),
			Status:           value.Status,
			Amount:           value.ItemPrice,
		}

		if len(objArray) < 1 {
			objArray = append(objArray, objArrayLoop)
		} else {

			status := ""

			for indexs, _ := range objArray {
				if objArray[indexs].Sku == value.Sku {
					status = "ada"
					qtynew := objArray[indexs].Qty + 1
					objArray[indexs].Qty = qtynew
					objArray[indexs].Sku = value.Sku
				}
			}

			if status == "" {
				objArray = append(objArray, objArrayLoop)
			}

		}

	}

	return objArray

}

func SaveSkuMappingTokped(obj models.ProductDetailTokped, ProductId string) {
	var objMapping []models.TableSkuMapping
	objData := obj.Data

	kode_tokped := os.Getenv("KODE_TOKPED")
	// tempsku := ""

	for _, elementModel := range objData {

		if ProductId == strconv.Itoa(elementModel.Basic.ProductID) {
			skunya := elementModel.Other.Sku
			//fmt.Println("status detail " + strconv.Itoa(elementModel.Basic.Status) + " sku " + skunya + " productid " + ProductId)

			if skunya != "" {

				statusnya := "c895af78-1488-cf0b-d17a-33770f2b20bf"
				if strconv.Itoa(elementModel.Basic.Status) == "1" || strconv.Itoa(elementModel.Basic.Status) == "3" {
					statusnya = "3936a00a-e321-f5ff-33b5-2dfb8d68ba45"
				}

				if len([]rune(skunya)) == 8 {
					//skunya = skunya
				} else if len([]rune(skunya)) > 8 {
					sku := []rune(skunya)
					skunya = string(sku[1:9])
				} else {

					if len([]rune(skunya)) < 8 {
						skunya = "0" + skunya
					} else {
						skunya = "XXXXXXXX"
					}

				}

				CekSkuWMS, _ := tokenRepository.CariSkuWMS(skunya)
				if CekSkuWMS.SkuNo != "" {
					CekMapping, _ := tokenRepository.CariSkuMapping(skunya, kode_tokped)
					if CekMapping.ChannelCode == "" { //insert
						objMappingLoop := models.TableSkuMapping{
							UuidSkuMapping: uuid.New().String(),
							SkuNo:          skunya,
							ChannelCode:    kode_tokped,
							CreatedBy:      "SYSTEM",
							CreatedDate:    time.Now(),
							IdSkuParent:    ProductId,
							ProductName:    elementModel.Basic.Name,
							// Varian:      varian,
							// Color:       color,
							UuidStatus: statusnya,
						}

						objMapping = append(objMapping, objMappingLoop)

					} else {

						CekParent, _ := tokenRepository.CariSkuMappingByParent(skunya, kode_tokped, ProductId)
						if CekParent.IdSkuParent == "" {
							objMappingLoop := models.TableSkuMapping{
								UuidSkuMapping: uuid.New().String(),
								SkuNo:          skunya,
								ChannelCode:    kode_tokped,
								CreatedBy:      "SYSTEM",
								CreatedDate:    time.Now(),
								IdSkuParent:    ProductId,
								ProductName:    elementModel.Basic.Name,
								// Varian:      varian,
								// Color:       color,
								UuidStatus: statusnya,
							}
							objMapping = append(objMapping, objMappingLoop)
							SaveErrorString("tokopedia", "SKU:"+skunya+" OLD:"+CekParent.IdSkuParent+" | NEW:"+ProductId, "SKU SUDAH ADA")
						} else {
							objMappingUpd := models.TableSkuMapping{
								UuidSkuMapping: CekParent.UuidSkuMapping,
								SkuNo:          skunya,
								ProductName:    elementModel.Basic.Name,
								UuidStatus:     statusnya,
							}
							//fmt.Println(objMappingUpd)
							result_upd := tokenRepository.UpdateSaveSkuMapping(objMappingUpd)
							if result_upd != nil {
								fmt.Println("Update SaveSkuMappingTokopedia Gagal")
							}
						}

						// if ProductId != CekMapping.IdSkuParent {

						// } else {

						// }

					}

				} //CekSkuWMS.SkuNo
			}
		}
	}

	if len(objMapping) > 0 {
		result := tokenRepository.SaveSkuMapping(objMapping)
		if result != nil {
			fmt.Println("SaveSkuMappingTokopedia Gagals")
		}
	}

}

func SaveSkuMappingSingleTokped(obj models.ProductTokped) {
	//fmt.Println("masuk SaveSkuMappingSingleTokped")
	var objMapping []models.TableSkuMapping
	objData := obj

	kode_tokped := os.Getenv("KODE_TOKPED")
	// tempsku := ""

	for _, elementModel := range objData.Data.Products {

		//if ProductId == strconv.Itoa(elementModel.Basic.ProductID) {
		skunya := elementModel.Sku
		//fmt.Println("status detail " + strconv.Itoa(elementModel.Basic.Status) + " sku " + skunya + " productid " + ProductId)

		if skunya != "" {

			statusnya := "c895af78-1488-cf0b-d17a-33770f2b20bf"
			if strconv.Itoa(elementModel.Status) == "1" || strconv.Itoa(elementModel.Status) == "3" {
				// statusnya = "c895af78-1488-cf0b-d17a-33770f2b20bf"
				statusnya = "3936a00a-e321-f5ff-33b5-2dfb8d68ba45"
			}

			if len([]rune(skunya)) == 8 {
				//skunya = skunya
			} else if len([]rune(skunya)) > 8 {
				sku := []rune(skunya)
				skunya = string(sku[1:9])
			} else {

				if len([]rune(skunya)) < 8 {
					skunya = "0" + skunya
				} else {
					skunya = "XXXXXXXX"
				}

			}

			CekSkuWMS, _ := tokenRepository.CariSkuWMS(skunya)
			if CekSkuWMS.SkuNo != "" {
				//fmt.Println("masuk SaveSkuMappingSingleTokped " + skunya)
				CekMapping, _ := tokenRepository.CariSkuMapping(skunya, kode_tokped)
				if CekMapping.ChannelCode == "" { //insert
					objMappingLoop := models.TableSkuMapping{
						UuidSkuMapping: uuid.New().String(),
						SkuNo:          skunya,
						ChannelCode:    kode_tokped,
						CreatedBy:      "SYSTEM",
						CreatedDate:    time.Now(),
						IdSkuParent:    strconv.Itoa(elementModel.ID),
						ProductName:    elementModel.Name,
						// Varian:      varian,
						// Color:       color,
						UuidStatus: statusnya,
					}

					objMapping = append(objMapping, objMappingLoop)

				} else {

					CekParent, _ := tokenRepository.CariSkuMappingByParent(skunya, kode_tokped, strconv.Itoa(elementModel.ID))
					if CekParent.IdSkuParent == "" {
						objMappingLoop := models.TableSkuMapping{
							UuidSkuMapping: uuid.New().String(),
							SkuNo:          skunya,
							ChannelCode:    kode_tokped,
							CreatedBy:      "SYSTEM",
							CreatedDate:    time.Now(),
							IdSkuParent:    strconv.Itoa(elementModel.ID),
							ProductName:    elementModel.Name,
							// Varian:      varian,
							// Color:       color,
							UuidStatus: statusnya,
						}
						objMapping = append(objMapping, objMappingLoop)
						SaveErrorString("tokopedia", "SKU:"+skunya+" OLD:"+CekParent.IdSkuParent+" | NEW:"+strconv.Itoa(elementModel.ID), "SKU SUDAH ADA")
					} else {
						objMappingUpd := models.TableSkuMapping{
							UuidSkuMapping: CekParent.UuidSkuMapping,
							SkuNo:          skunya,
							ProductName:    elementModel.Name,
							UuidStatus:     statusnya,
						}
						//fmt.Println(objMappingUpd)
						result_upd := tokenRepository.UpdateSaveSkuMapping(objMappingUpd)
						if result_upd != nil {
							fmt.Println("Update SaveSkuMappingTokopedia Gagal")
						}
					}

				}

			} else {
				//fmt.Println("masuk SaveSkuMappingSingleTokped kosong")
			} //CekSkuWMS.SkuNo
		}
		//}
	}

	if len(objMapping) > 0 {
		result := tokenRepository.SaveSkuMapping(objMapping)
		if result != nil {
			fmt.Println("SaveSkuMappingTokopedia Gagals")
		}
	}

}

func SaveSkuMappingSingleTokpedNew(obj models.ProductDetailTokped, paramsku string) {
	var objMapping []models.TableSkuMapping

	kode_tokped := os.Getenv("KODE_TOKPED")

	//update jadi inaktif
	//fmt.Println("SaveSkuMappingSingleTokpedNew " + paramsku)
	statusnya := "c895af78-1488-cf0b-d17a-33770f2b20bf"
	objInaTokped := models.TableSkuMapping{
		SkuNo:       paramsku,
		ChannelCode: kode_tokped,
		UuidStatus:  statusnya,
	}
	helpers.UpdateStockNol(paramsku, "AUTO", kode_tokped)

	if len(obj.Data) > 0 {
		resultIna := tokenRepository.UpdateSkuMappingByChannelSKu(objInaTokped)
		if resultIna != nil {
			fmt.Println("Update SaveSkuMappingSingleTokpedNewIna Gagal")
		}
	}

	indexin := 0
	for _, elementModel := range obj.Data {
		//fmt.Println("masuk")
		//if ProductId == strconv.Itoa(elementModel.Basic.ProductID) {
		skunya := elementModel.Other.Sku
		//fmt.Println("status detail " + strconv.Itoa(elementModel.Basic.Status) + " sku " + skunya + " productid " + ProductId)

		if skunya != "" {
			//statusnya := "3936a00a-e321-f5ff-33b5-2dfb8d68ba45"
			statusnya = "c895af78-1488-cf0b-d17a-33770f2b20bf"
			if len(obj.Data) > 1 {
				if strconv.Itoa(elementModel.Basic.Status) == "1" {

					statusnya = "3936a00a-e321-f5ff-33b5-2dfb8d68ba45"
					indexin++
				} else if indexin == 0 && strconv.Itoa(elementModel.Basic.Status) == "3" {
					statusnya = "3936a00a-e321-f5ff-33b5-2dfb8d68ba45"
					indexin++
				}

			} else {
				if strconv.Itoa(elementModel.Basic.Status) == "1" || strconv.Itoa(elementModel.Basic.Status) == "3" {
					//statusnya = "c895af78-1488-cf0b-d17a-33770f2b20bf"
					statusnya = "3936a00a-e321-f5ff-33b5-2dfb8d68ba45"
				}
			}

			// if strconv.Itoa(elementModel.Basic.Status) == "1" || strconv.Itoa(elementModel.Basic.Status) == "3" {
			// 	//statusnya = "c895af78-1488-cf0b-d17a-33770f2b20bf"
			// 	statusnya = "3936a00a-e321-f5ff-33b5-2dfb8d68ba45"
			// }
			//fmt.Println("statusnya " + statusnya)

			if len([]rune(skunya)) == 8 {
				//skunya = skunya
			} else if len([]rune(skunya)) > 8 {
				sku := []rune(skunya)
				skunya = string(sku[1:9])
			} else {

				if len([]rune(skunya)) < 8 {
					skunya = "0" + skunya
				} else {
					skunya = "XXXXXXXX"
				}

			}

			CekSkuWMS, _ := tokenRepository.CariSkuWMS(skunya)
			//fmt.Println(skunya)
			if CekSkuWMS.SkuNo != "" {
				//fmt.Println("masuk SaveSkuMappingSingleTokped " + skunya)
				CekMapping, _ := tokenRepository.CariSkuMapping(skunya, kode_tokped)
				if CekMapping.ChannelCode == "" { //insert
					objMappingLoop := models.TableSkuMapping{
						UuidSkuMapping: uuid.New().String(),
						SkuNo:          skunya,
						ChannelCode:    kode_tokped,
						CreatedBy:      "SYSTEM",
						CreatedDate:    time.Now(),
						IdSkuParent:    strconv.Itoa(elementModel.Basic.ProductID),
						ProductName:    elementModel.Basic.Name,
						// Varian:      varian,
						// Color:       color,
						UuidStatus: statusnya,
					}

					objMapping = append(objMapping, objMappingLoop)

				} else {

					CekParent, _ := tokenRepository.CariSkuMappingByParent(skunya, kode_tokped, strconv.Itoa(elementModel.Basic.ProductID))
					if CekParent.IdSkuParent == "" {
						objMappingLoop := models.TableSkuMapping{
							UuidSkuMapping: uuid.New().String(),
							SkuNo:          skunya,
							ChannelCode:    kode_tokped,
							CreatedBy:      "SYSTEM",
							CreatedDate:    time.Now(),
							IdSkuParent:    strconv.Itoa(elementModel.Basic.ProductID),
							ProductName:    elementModel.Basic.Name,
							UuidStatus:     statusnya,
						}
						objMapping = append(objMapping, objMappingLoop)
						SaveErrorString("tokopedia", "SKU:"+skunya+" OLD:"+CekParent.IdSkuParent+" | NEW:"+strconv.Itoa(elementModel.Basic.ProductID), "SKU SUDAH ADA")
					} else {
						objMappingUpd := models.TableSkuMapping{
							UuidSkuMapping: CekParent.UuidSkuMapping,
							SkuNo:          skunya,
							ProductName:    elementModel.Basic.Name,
							UuidStatus:     statusnya,
						}
						//fmt.Println(objMappingUpd)
						result_upd := tokenRepository.UpdateSaveSkuMapping(objMappingUpd)
						if result_upd != nil {
							fmt.Println("Update SaveSkuMappingSingleTokpedNew Gagal")
						}
					}

				}

			}
		}
		//}
	}

	if len(objMapping) > 0 {
		result := tokenRepository.SaveSkuMapping(objMapping)
		if result != nil {
			fmt.Println("SaveSkuMappingSingleTokpedNew Gagals")
		}
	}

}

func SaveSalesOrderTokped(objx models.OrderDetailTokpedSingle, Expedisi string) {
	var obj = objx.Data
	var objSalesOrder []models.TableSalesOrder
	var objSalesOrderTotalQty models.TableSalesOrder
	createdDate := time.Now()
	totalQty := 0

	NoOrder := ""

	// fmt.Println("======= ID ==========")
	// fmt.Println(obj.OrderID)

	// fmt.Println("======= STATUS ==========")
	// fmt.Println(obj.OrderStatus)

	NoOrder = strconv.Itoa(obj.OrderID)

	objCek, _ := tokenRepository.FindSalesOrder(NoOrder)
	CodeKurir := obj.OrderInfo.ShippingInfo.LogisticName

	if CodeKurir == "" {
		CodeKurir = Expedisi
	}

	// fmt.Println("=== KURIR TOKPED ===")
	// fmt.Println(CodeKurir)
	// fmt.Println("=== KURIR TOKPED ===")

	if objCek.NoOrder == "" && CodeKurir != "" && (obj.OrderStatus == 400 || obj.OrderStatus == 220) {
		fmt.Println("SaveSalesOrderTokped " + NoOrder)
		for _, elementDetail := range obj.OrderInfo.OrderDetail {
			//	fmt.Println("sku " + elementDetail.Sku)
			skunya := elementDetail.Sku
			qty := elementDetail.Quantity
			if len([]rune(skunya)) == 8 {
				//skunya = skunya
			} else if len([]rune(skunya)) > 8 {
				sku := []rune(skunya)
				skunya = string(sku[1:9])
			} else if len([]rune(skunya)) == 7 {
				skunya = "0" + skunya
			} else {
				skunya = "XXXXXXXX"
			}

			// fmt.Println("sku " + skunya)
			// fmt.Println("qty " + strconv.Itoa(qty))

			if skunya == "08172335" {
				//skunya = "08172336"
			}

			receiptname := obj.OrderInfo.Destination.ReceiverName
			telp := obj.OrderInfo.Destination.ReceiverPhone
			address := obj.OrderInfo.Destination.AddressStreet
			city := obj.OrderInfo.Destination.AddressCity
			province := obj.OrderInfo.Destination.AddressProvince

			id := uuid.New()
			totalQty = totalQty + int(qty)

			amount := elementDetail.ProductPrice

			objSalesOrderLoop := models.TableSalesOrder{
				UuidSalesOrder:     id.String(),
				NoOrder:            NoOrder,
				StatusOrder:        "PDKM",
				SkuNo:              skunya,
				ProductName:        elementDetail.ProductName,
				Qty:                strconv.Itoa(int(qty)),
				ExpeditionType:     CodeKurir,
				DeliveryType:       "Pickup",
				TotalQtyOrder:      "30",
				StatusProcessOrder: "0",
				WeightProduk:       "0",
				TotalWeight:        "0",
				OrderDate:          obj.UpdateTime,
				SentBeforeDate:     obj.ShipmentFulfillment.ConfirmShippingDeadline,
				CreatedDate:        createdDate,

				CreatedBy:   "TOKOPEDIA",
				NameVariant: "", //value.ModelName,
				//UsernameBuyer: obj.Response.OrderList[0].BuyerUsername,
				ChannelCode: os.Getenv("KODE_TOKPED"),

				RecipientsName:  receiptname,
				Telephone:       telp,
				ShippingAddress: address,
				City:            city,
				Province:        province,
				NoInvoice:       obj.InvoiceNumber,
				Amount:          amount,
			}

			objSalesOrder = append(objSalesOrder, objSalesOrderLoop)
			objSalesOrderTotalQty.NoOrder = NoOrder
			objSalesOrderTotalQty.TotalQtyOrder = strconv.Itoa(totalQty)

		}

	} else {

		if NoOrder == objCek.NoOrder {

			//update invoice jika masih kosong
			if objCek.NoInvoice == "" {
				var objUpdateInv models.TableSalesOrder
				objUpdateInv.NoOrder = objCek.NoOrder
				objUpdateInv.NoInvoice = obj.InvoiceNumber
				objUpdateInv.CreatedBy = "TOKOPEDIA"
				//result := tokenRepository.UpdateKurir(objUpdateKurir)
				result := tokenRepository.UpdateKurirNew(objUpdateInv)
				if result != nil {
					fmt.Println("UpdateInvoice Gagal")
				} else {
					fmt.Println("UpdateInvoice Success")
				}
			}

			if objCek.ExpeditionType != CodeKurir { //update kurir

				var objUpdateKurir models.TableSalesOrder
				objUpdateKurir.NoOrder = objCek.NoOrder
				objUpdateKurir.ExpeditionType = CodeKurir
				objUpdateKurir.CreatedBy = "TOKOPEDIA"
				//result := tokenRepository.UpdateKurir(objUpdateKurir)
				result := tokenRepository.UpdateKurirNew(objUpdateKurir)
				if result != nil {
					fmt.Println("UpdateKurir Gagal")
				} else {
					fmt.Println("UpdateKurir Success")
				}

			}

			if objCek.StatusProcessOrder == "0" {

				if obj.OrderStatus == 0 || obj.OrderStatus == 3 || obj.OrderStatus == 5 || obj.OrderStatus == 6 ||
					obj.OrderStatus == 10 || obj.OrderStatus == 15 {

					objSalesOrderLoop := models.TableSalesOrder{
						NoOrder:            objCek.NoOrder,
						StatusProcessOrder: "3",
						CreatedBy:          "TOKOPEDIA",
					}
					result := tokenRepository.UpdateSalesOrderCanc(objSalesOrderLoop)
					if result != nil {
						fmt.Println("UpdateSalesOrderCanc Gagal")
					} else {
						fmt.Println("UpdateSalesOrderCanc Success")
					}
					UpdateStockOrder(objCek.NoOrder)

				}
			}

			if objCek.StatusProcessOrder == "1" {

				if obj.OrderStatus == 700 {
					payoutdate := time.Now()

					objSalesOrderLoop := models.TableSalesOrder{
						NoOrder:    objCek.NoOrder,
						PayoutDate: payoutdate,
						CreatedBy:  "TOKOPEDIA",
					}
					result := tokenRepository.UpdateSalesOrderCanc(objSalesOrderLoop)
					if result != nil {
						fmt.Println("Update Payoutdate Gagal")
					} else {
						fmt.Println("Update Payoutdate Success")
					}
				}

			}
		}
	}

	if len(objSalesOrder) > 0 {
		result := tokenRepository.SaveSalesOrder(objSalesOrder, objSalesOrderTotalQty)
		if result != nil {
			fmt.Println("SaveSalesOrderTokped Gagal")
		} else {
			fmt.Println("SaveSalesOrderTokped Success")
			pesan := BookingSku(objSalesOrderTotalQty.NoOrder)
			if pesan != "Sukses" {
				//delete sales order
				//tokenRepository.DeleteSalesOrder(objSalesOrderTotalQty)

			}
			fmt.Println(objSalesOrderTotalQty.NoOrder + " => " + pesan)

			//cari detail no order
			UpdateStockOrder(objSalesOrderTotalQty.NoOrder)

		}
	}

}

func SaveSkuMappingBukalapak(obj models.ProductDetailBukalapak, SkuBl string) {
	var objMapping []models.TableSkuMapping
	objData := obj.Data

	kode_bukalapak := os.Getenv("KODE_BUKALAPK")
	// tempsku := ""

	skunya := SkuBl
	fmt.Println(" sku " + skunya + " productid " + strconv.Itoa(int(objData.SkuID)))

	if skunya != "" {
		statusnya := "3936a00a-e321-f5ff-33b5-2dfb8d68ba45"
		if objData.Active != true {
			statusnya = "c895af78-1488-cf0b-d17a-33770f2b20bf"
		}

		if len([]rune(skunya)) == 8 {
			//skunya = skunya
		} else if len([]rune(skunya)) > 8 {
			sku := []rune(skunya)
			skunya = string(sku[1:9])
		} else {

			if len([]rune(skunya)) < 8 {
				skunya = "0" + skunya
			} else {
				skunya = "XXXXXXXX"
			}

		}

		CekSkuWMS, _ := tokenRepository.CariSkuWMS(skunya)
		if CekSkuWMS.SkuNo != "" {
			CekMapping, _ := tokenRepository.CariSkuMapping(skunya, kode_bukalapak)
			if CekMapping.ChannelCode == "" { //insert
				objMappingLoop := models.TableSkuMapping{
					UuidSkuMapping: uuid.New().String(),
					SkuNo:          skunya,
					ChannelCode:    kode_bukalapak,
					CreatedBy:      "SYSTEM",
					CreatedDate:    time.Now(),
					IdSkuParent:    objData.ID,
					IdSkuChild1:    strconv.Itoa(int(objData.SkuID)),
					ProductName:    objData.Name,
					// Varian:      varian,
					// Color:       color,
					UuidStatus: statusnya,
				}

				objMapping = append(objMapping, objMappingLoop)

			} else {

				CekParent, _ := tokenRepository.CariSkuMappingByParent(skunya, kode_bukalapak, objData.ID)

				if CekParent.IdSkuParent == "" {

					CekChild1, _ := tokenRepository.CariSkuMappingByModel(skunya, kode_bukalapak, strconv.Itoa(int(objData.SkuID)))
					if CekChild1.IdSkuParent == "" {
						objMappingLoop := models.TableSkuMapping{
							UuidSkuMapping: uuid.New().String(),
							SkuNo:          skunya,
							ChannelCode:    kode_bukalapak,
							CreatedBy:      "SYSTEM",
							CreatedDate:    time.Now(),
							IdSkuParent:    objData.ID,
							IdSkuChild1:    strconv.Itoa(int(objData.SkuID)),
							ProductName:    objData.Name,
							// Varian:      varian,
							// Color:       color,
							UuidStatus: statusnya,
						}
						objMapping = append(objMapping, objMappingLoop)
						SaveErrorString("bukalapak", "SKU:"+skunya+" OLD:"+CekChild1.IdSkuParent+" | NEW:"+objData.ID, "SKU SUDAH ADA")
					} else {
						objMappingUpd := models.TableSkuMapping{
							UuidSkuMapping: CekChild1.UuidSkuMapping,
							SkuNo:          skunya,
							ProductName:    objData.Name,
							UuidStatus:     statusnya,
						}
						//fmt.Println(objMappingUpd)
						result_upd := tokenRepository.UpdateSaveSkuMapping(objMappingUpd)
						if result_upd != nil {
							fmt.Println("Update SaveSkuMappingBukalapak Gagal")
						}
					}

				} else {
					objMappingUpd := models.TableSkuMapping{
						UuidSkuMapping: CekParent.UuidSkuMapping,
						SkuNo:          skunya,
						ProductName:    objData.Name,
						UuidStatus:     statusnya,
					}
					//fmt.Println(objMappingUpd)
					result_upd := tokenRepository.UpdateSaveSkuMapping(objMappingUpd)
					if result_upd != nil {
						fmt.Println("Update SaveSkuMappingBukalapak Gagal")
					}
				}

				// if objData.ID != CekMapping.IdSkuParent {

				// } else {

				// }

			}

		} //CekSkuWMS.SkuNo
	}

	if len(objMapping) > 0 {
		result := tokenRepository.SaveSkuMapping(objMapping)
		if result != nil {
			fmt.Println("SaveSkuMappingBukalapak Gagals")
		}
	}

}

func SaveSkuMappingVarianBukalapak(obj models.ProductDetailSkuBukalapak, ProductName string) {
	var objMapping []models.TableSkuMapping
	objData := obj.Data

	kode_bukalapak := os.Getenv("KODE_BUKALAPK")
	// tempsku := ""

	skunya := objData.SkuName

	if skunya != "" {
		statusnya := "3936a00a-e321-f5ff-33b5-2dfb8d68ba45"
		if objData.State != "active" {
			statusnya = "c895af78-1488-cf0b-d17a-33770f2b20bf"
		}

		if len([]rune(skunya)) == 8 {
			//skunya = skunya
		} else if len([]rune(skunya)) > 8 {
			sku := []rune(skunya)
			skunya = string(sku[1:9])
		} else {

			if len([]rune(skunya)) < 8 {
				skunya = "0" + skunya
			} else {
				skunya = "XXXXXXXX"
			}

		}

		CekSkuWMS, _ := tokenRepository.CariSkuWMS(skunya)
		if CekSkuWMS.SkuNo != "" {
			CekMapping, _ := tokenRepository.CariSkuMapping(skunya, kode_bukalapak)
			if CekMapping.ChannelCode == "" { //insert
				objMappingLoop := models.TableSkuMapping{
					UuidSkuMapping: uuid.New().String(),
					SkuNo:          skunya,
					ChannelCode:    kode_bukalapak,
					CreatedBy:      "SYSTEM",
					CreatedDate:    time.Now(),
					IdSkuParent:    objData.ProductID,
					IdSkuChild1:    strconv.Itoa(int(objData.ID)),
					ProductName:    ProductName,
					Varian:         objData.VariantName,
					// Color:       color,
					UuidStatus: statusnya,
				}

				objMapping = append(objMapping, objMappingLoop)

			} else {

				CekParent, _ := tokenRepository.CariSkuMappingByParent(skunya, kode_bukalapak, objData.ProductID)
				if CekParent.IdSkuParent == "" {

					CekChild1, _ := tokenRepository.CariSkuMappingByModel(skunya, kode_bukalapak, strconv.Itoa(int(objData.ID)))
					if CekChild1.IdSkuParent == "" {
						objMappingLoop := models.TableSkuMapping{
							UuidSkuMapping: uuid.New().String(),
							SkuNo:          skunya,
							ChannelCode:    kode_bukalapak,
							CreatedBy:      "SYSTEM",
							CreatedDate:    time.Now(),
							IdSkuParent:    objData.ProductID,
							IdSkuChild1:    strconv.Itoa(int(objData.ID)),
							ProductName:    ProductName,
							Varian:         objData.VariantName,
							// Color:       color,
							UuidStatus: statusnya,
						}
						objMapping = append(objMapping, objMappingLoop)
						SaveErrorString("bukalapak", "SKU:"+skunya+" OLD:"+CekChild1.IdSkuParent+" | NEW:"+objData.ProductID, "SKU SUDAH ADA")
					} else {
						objMappingUpd := models.TableSkuMapping{
							UuidSkuMapping: CekChild1.UuidSkuMapping,
							SkuNo:          skunya,
							ProductName:    ProductName,
							Varian:         objData.VariantName,
							UuidStatus:     statusnya,
						}
						//fmt.Println(objMappingUpd)
						result_upd := tokenRepository.UpdateSaveSkuMapping(objMappingUpd)
						if result_upd != nil {
							fmt.Println("Update SaveSkuMappingVarianBukalapak Gagal")
						}
					}

				} else {

					objMappingUpd := models.TableSkuMapping{
						UuidSkuMapping: CekParent.UuidSkuMapping,
						SkuNo:          skunya,
						ProductName:    ProductName,
						Varian:         objData.VariantName,
						UuidStatus:     statusnya,
					}
					//fmt.Println(objMappingUpd)
					result_upd := tokenRepository.UpdateSaveSkuMapping(objMappingUpd)
					if result_upd != nil {
						fmt.Println("Update SaveSkuMappingVarianBukalapak Gagal")
					}
				}

				// if objData.ProductID != CekMapping.IdSkuParent {

				// } else {

				// }

			}

		} //CekSkuWMS.SkuNo
	}

	if len(objMapping) > 0 {
		result := tokenRepository.SaveSkuMapping(objMapping)
		if result != nil {
			fmt.Println("SaveSkuMappingVarianBukalapak Gagals")
		}
	}
}

func SaveSalesOrderBukalapak(objx models.OrderDetailBukalapak) {
	var obj = objx.Data
	var objSalesOrder []models.TableSalesOrder
	var objSalesOrderTotalQty models.TableSalesOrder
	createdDate := time.Now()
	totalQty := 0

	NoOrder := ""

	NoOrder = obj.TransactionID

	objCek, _ := tokenRepository.FindSalesOrder(NoOrder)
	CodeKurir := obj.Delivery.RequestedCarrier

	if objCek.NoOrder == "" && CodeKurir != "" && obj.State == "paid" {
		// if objCek.NoOrder == "" && CodeKurir != "" {
		for _, elementDetail := range obj.Items {

			skunya := elementDetail.Stuff.SkuName
			qty := elementDetail.Quantity
			if len([]rune(skunya)) == 8 {
				//skunya = skunya
			} else if len([]rune(skunya)) > 8 {
				sku := []rune(skunya)
				skunya = string(sku[1:9])
			} else if len([]rune(skunya)) == 7 {
				skunya = "0" + skunya
			} else {
				skunya = "XXXXXXXX"
			}

			// fmt.Println("sku " + skunya)
			// fmt.Println("qty " + strconv.Itoa(qty))

			receiptname := obj.Delivery.Consignee.Name
			telp := "" //get saat ambil resi
			address := obj.Delivery.Consignee.Address
			city := obj.Delivery.Consignee.City
			province := obj.Delivery.Consignee.Province

			id := uuid.New()
			totalQty = totalQty + int(qty)
			objSalesOrderLoop := models.TableSalesOrder{
				UuidSalesOrder:     id.String(),
				NoOrder:            NoOrder,
				StatusOrder:        "PDKM",
				SkuNo:              skunya,
				ProductName:        elementDetail.Name,
				Qty:                strconv.Itoa(int(qty)),
				ExpeditionType:     CodeKurir,
				DeliveryType:       "Pickup",
				TotalQtyOrder:      "30",
				StatusProcessOrder: "0",
				WeightProduk:       "0",
				TotalWeight:        "0",
				OrderDate:          obj.UpdatedAt,
				SentBeforeDate:     obj.StateChangedAt.ExpiredAt,
				CreatedDate:        createdDate,

				CreatedBy:   "BUKALAPAK",
				NameVariant: "", //value.ModelName,
				//UsernameBuyer: obj.Response.OrderList[0].BuyerUsername,
				ChannelCode: os.Getenv("KODE_BUKALAPK"),

				RecipientsName:  receiptname,
				Telephone:       telp,
				ShippingAddress: address,
				City:            city,
				Province:        province,
			}

			objSalesOrder = append(objSalesOrder, objSalesOrderLoop)
			objSalesOrderTotalQty.NoOrder = NoOrder
			objSalesOrderTotalQty.TotalQtyOrder = strconv.Itoa(totalQty)

		}

	} else {

		if NoOrder == objCek.NoOrder {

			if objCek.ExpeditionType != CodeKurir { //update kurir

				var objUpdateKurir models.TableSalesOrder
				objUpdateKurir.NoOrder = objCek.NoOrder
				objUpdateKurir.ExpeditionType = CodeKurir
				objUpdateKurir.CreatedBy = "BUKALAPAK"
				result := tokenRepository.UpdateKurirNew(objUpdateKurir)
				if result != nil {
					fmt.Println("UpdateKurir Gagal")
				} else {
					fmt.Println("UpdateKurir Success")
				}

			}

			if objCek.StatusProcessOrder == "0" {

				if obj.State == "cancelled" {

					objSalesOrderLoop := models.TableSalesOrder{
						NoOrder:            objCek.NoOrder,
						StatusProcessOrder: "3",
						CreatedBy:          "BUKALAPAK",
					}
					result := tokenRepository.UpdateSalesOrderCanc(objSalesOrderLoop)
					if result != nil {
						fmt.Println("UpdateSalesOrderCanc Gagal")
					} else {
						fmt.Println("UpdateSalesOrderCanc Success")
					}
					UpdateStockOrder(objCek.NoOrder)

				}
			}

			if objCek.StatusProcessOrder == "1" {

				if obj.State == "remitted" {
					payoutdate := time.Now()

					objSalesOrderLoop := models.TableSalesOrder{
						NoOrder:    objCek.NoOrder,
						PayoutDate: payoutdate,
						CreatedBy:  "BUKALAPAK",
					}
					result := tokenRepository.UpdateSalesOrderCanc(objSalesOrderLoop)
					if result != nil {
						fmt.Println("Update Payoutdate Gagal")
					} else {
						fmt.Println("Update Payoutdate Success")
					}
				}

			}
		}
	}

	if len(objSalesOrder) > 0 {
		result := tokenRepository.SaveSalesOrder(objSalesOrder, objSalesOrderTotalQty)
		if result != nil {
			fmt.Println("SaveSalesOrderBukalapak Gagal")
		} else {
			fmt.Println("SaveSalesOrderBukalapak Success")
			pesan := BookingSku(objSalesOrderTotalQty.NoOrder)
			if pesan != "Sukses" {
				//delete sales order
				//tokenRepository.DeleteSalesOrder(objSalesOrderTotalQty)

			}
			fmt.Println(objSalesOrderTotalQty.NoOrder + " => " + pesan)

			//cari detail no order
			UpdateStockOrder(objSalesOrderTotalQty.NoOrder)

		}
	}

}

func UpdateProductNuse(IdParent string, IdChild1 string, ChannelCode string) {

	CekMapping, _ := tokenRepository.CariSkuMappingParentChild1(IdParent, IdChild1, ChannelCode)
	if CekMapping.ChannelCode != "" {
		objMappingUpd := models.TableSkuMapping{
			UuidSkuMapping: CekMapping.UuidSkuMapping,
			UuidStatus:     "a14d0d35-9fd7-0038-1950-60e8bb905c34", //NUSE
		}
		result_upd := tokenRepository.UpdateSaveSkuMapping(objMappingUpd)
		if result_upd != nil {
			fmt.Println("UpdateProductNuse Gagal")
		}
	}

}

func GroupChannelSkuMap(sku string) {
	objChannel, _ := stockRepository.GroupChannelSkuMap(sku)
	objAvlb, _ := stockRepository.CariUidAvlbBuffer(sku)
	objBuffer, _ := stockRepository.CariBufferStock(sku)

	helpers.InsertStockWms(sku, "AUTO", strconv.Itoa(objAvlb))

	if len(objChannel) > 0 {
		var wg sync.WaitGroup
		wg.Add(len(objChannel))
		for _, val := range objChannel {
			//fmt.Println(val.SkuNo + " | " + val.ChannelCode)
			go func(Skunya string, Channelnya string) {
				//fmt.Println(Skunya + " | " + Channelnya)
				helpers.UpdateStockV2(Skunya, "API_CHANNEL", Channelnya, objAvlb, objBuffer, &wg)

			}(val.SkuNo, val.ChannelCode)

		}
		wg.Wait()

	}

}

func SaveErrorInfo(param, channel, noorder string) {
	//helpers.KirimEmail(param, " "+param+" "+channel+" "+noorder, "")
}
