package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/rals/dearme-channel/models"
	"github.com/rals/dearme-channel/repositories/stockRepository"
)

func UpdateStockV2(sku string, user string, channel string, objAvlb int, objBuffer models.TableBufferStock, wg *sync.WaitGroup) {
	waktu := time.Now().String()
	fmt.Println("UpdateStockV2 " + sku + " " + waktu + " " + channel)
	Skunya := sku

	var ObjSkuMapping []models.TableSkuMapping

	if channel == "" {
		ObjSkuMapping, _ = stockRepository.CariSkuParent(Skunya)
	} else {
		ObjSkuMapping, _ = stockRepository.CariSkuParentbychannel(Skunya, channel)
	}

	if len(ObjSkuMapping) > 0 {
		StockBuffers := 0
		// objAvlb, _ := stockRepository.CariUidAvlbBuffer(Skunya)
		// objBuffer, _ := stockRepository.CariBufferStock(Skunya)
		if objBuffer.SkuNo != "" {
			StockBuffers = int(objBuffer.BufferStock)
		}

		InsertStockWms(Skunya, user, strconv.Itoa(objAvlb))

		//update stock API
		istok := 0

		for _, ObjSkuMapping := range ObjSkuMapping {
			IdParent := ObjSkuMapping.IdSkuParent
			IdChild1 := ObjSkuMapping.IdSkuChild1
			Channelnya := ObjSkuMapping.ChannelCode

			StockBuffer := objAvlb - int(StockBuffers)

			if StockBuffer < 1 {
				StockBuffer = 0
			}

			StockTemp := StockBuffer //R996

			if istok > 0 { //yg ke2 atau lebih jadi 0
				StockTemp = 0
			}

			//cek ke wms_sku_block_channel
			tgl := time.Now().Format("2006-01-02")
			objBlock := stockRepository.CariSkuBlock(Skunya, Channelnya, tgl)
			if objBlock.ChannelCode != "" {
				StockBuffer = 0
				StockTemp = 0
			}

			//cek jika channel tidak aktif stok 0
			CekChannel := stockRepository.CekChannelSku(Skunya, Channelnya)
			// fmt.Println("status channel " + CekChannel.UuidStatus)
			if CekChannel.UuidStatus != "3936a00a-e321-f5ff-33b5-2dfb8d68ba45" {
				StockBuffer = 0
				StockTemp = 0
			}

			if Channelnya == "R996" { //SHOPEE
				StockBuffer = StockTemp
				cek, _ := StockShopee(IdParent, IdChild1, int64(StockBuffer), sku)
				//spew.Dump(hasil)
				if cek == false {
					fmt.Println("Update stock helper " + Skunya + " gagal " + strconv.Itoa(StockBuffer) + " " + waktu)

				} else {
					fmt.Println("Update stock helper " + Skunya + " sukses " + strconv.Itoa(StockBuffer) + " " + waktu)
				}
			} else if Channelnya == "R991" { //ZALORA
				//StockBuffer = objAvlb - int(StockBuffers)
				//cek, _ := StockZalora(Skunya, int64(StockBuffer))
				//spew.Dump(hasil)

				cek := false
				if IdChild1 != "" {
					cek, _ = StockZaloraV2(IdChild1, int64(StockBuffer), sku)
				} else {
					cek, _ = StockZalora(Skunya, int64(StockBuffer))
				}

				if cek == false {
					fmt.Println("Update stock helper " + Skunya + " gagal " + strconv.Itoa(StockBuffer) + " " + waktu)

				} else {
					fmt.Println("Update stock helper " + Skunya + " sukses " + strconv.Itoa(StockBuffer) + " " + waktu)
				}
			} else if Channelnya == "R990" { //TIKTOK
				StockBuffer = StockTemp
				cek, _ := StockTiktok(IdParent, IdChild1, int64(StockBuffer))
				//spew.Dump(hasil)
				if cek == false {
					fmt.Println("Update stock helper " + Skunya + " gagal " + strconv.Itoa(StockBuffer) + " " + waktu)

				} else {
					fmt.Println("Update stock helper " + Skunya + " sukses " + strconv.Itoa(StockBuffer) + " " + waktu)
				}
			} else if Channelnya == "R995" { //JDID
				StockBuffer = StockTemp
				skuId := IdChild1
				if skuId == "" {
					skuId = IdParent
				}
				cek, _ := StockJdId(skuId, int64(StockBuffer))
				//spew.Dump(hasil)
				if cek == false {
					fmt.Println("Update stock helper " + Skunya + " gagal " + strconv.Itoa(StockBuffer) + " " + waktu)

				} else {
					fmt.Println("Update stock helper " + Skunya + " sukses " + strconv.Itoa(StockBuffer) + " " + waktu)
				}

			} else if Channelnya == "R994" { //BLIBLI

				StockBuffer = StockTemp
				cek, _ := StockBlibli(IdChild1, int64(StockBuffer))
				//spew.Dump(hasil)
				if cek == false {
					fmt.Println("Update stock helper " + Skunya + " gagal " + strconv.Itoa(StockBuffer) + " " + waktu)

				} else {
					fmt.Println("Update stock helper " + Skunya + " sukses " + strconv.Itoa(StockBuffer) + " " + waktu)
				}
			} else if Channelnya == "R992" { //BUKALAPAK
				StockBuffer = StockTemp
				cek, _ := StockBukalapak(IdParent, IdChild1, int64(StockBuffer))
				if cek == false {
					fmt.Println("Update stock helper " + Skunya + " gagal " + strconv.Itoa(StockBuffer) + " " + waktu)

				} else {
					fmt.Println("Update stock helper " + Skunya + " sukses " + strconv.Itoa(StockBuffer) + " " + waktu)
				}
			} else if Channelnya == "R998" { //TOKOPEDIA
				StockBuffer = StockTemp
				cek, _ := StockTokopedia(IdParent, int64(StockBuffer), sku)
				if cek == false {
					fmt.Println("Update stock helper " + Skunya + " gagal " + strconv.Itoa(StockBuffer) + " " + waktu)

				} else {
					fmt.Println("Update stock helper " + Skunya + " sukses " + strconv.Itoa(StockBuffer) + " " + waktu)
				}
			} else if Channelnya == "R997" { //LAZADA
				StockBuffer = StockTemp
				cek, _ := StockLazada(IdParent, IdChild1, int64(StockBuffer), sku)
				if cek == false {
					fmt.Println("Update stock helper " + Skunya + " gagal " + strconv.Itoa(StockBuffer) + " " + waktu)

				} else {
					fmt.Println("Update stock helper " + Skunya + " sukses " + strconv.Itoa(StockBuffer) + " " + waktu)
				}
			} else {
				fmt.Println("Channel SKU " + Skunya + " Tidak Ditemukan")
			}

			istok++
		}

		//end update stock API
	}

	wg.Done()
}

func UpdateStock(sku string, user string, channel string) {
	waktu := time.Now().String()
	fmt.Println("UpdateStock " + sku + " " + waktu)
	Skunya := sku

	var ObjSkuMapping []models.TableSkuMapping

	if channel == "" {
		ObjSkuMapping, _ = stockRepository.CariSkuParent(Skunya)
	} else {
		ObjSkuMapping, _ = stockRepository.CariSkuParentbychannel(Skunya, channel)
	}

	if len(ObjSkuMapping) > 0 {
		StockBuffers := 0
		objAvlb, _ := stockRepository.CariUidAvlbBuffer(Skunya)
		objBuffer, _ := stockRepository.CariBufferStock(Skunya)
		if objBuffer.SkuNo != "" {
			StockBuffers = int(objBuffer.BufferStock)
		}

		InsertStockWms(Skunya, user, strconv.Itoa(objAvlb))

		//update stock API
		istok := 0

		//temp_sku := ""
		temp_parent := ""
		temp_child := ""
		temp_channel := ""

		for _, ObjSkuMapping := range ObjSkuMapping {
			IdParent := ObjSkuMapping.IdSkuParent
			IdChild1 := ObjSkuMapping.IdSkuChild1
			Channelnya := ObjSkuMapping.ChannelCode

			StockBuffer := objAvlb - int(StockBuffers)

			if istok > 0 { //yg ke2 atau lebih jadi 0
				//StockBuffer = 0
			}

			if StockBuffer < 1 {
				StockBuffer = 0
			}

			StockTemp := StockBuffer //R996

			//jika sku ada 2 atau lebih maka yg ke 2 jadi 0
			//if temp_sku != ObjSkuMapping.SkuNo && temp_parent != ObjSkuMapping.IdSkuParent && temp_channel != ObjSkuMapping.ChannelCode {
			if temp_channel != Channelnya {
				//temp_sku = ObjSkuMapping.SkuNo
				temp_parent = ObjSkuMapping.IdSkuParent
				temp_child = ObjSkuMapping.IdSkuChild1
				temp_channel = ObjSkuMapping.ChannelCode
			} else {

				if temp_parent != ObjSkuMapping.IdSkuParent && temp_child != ObjSkuMapping.IdSkuChild1 {
					StockTemp = 0
				}
			}

			//cek ke wms_sku_block_channel
			tgl := time.Now().Format("2006-01-02")
			objBlock := stockRepository.CariSkuBlock(Skunya, Channelnya, tgl)
			if objBlock.ChannelCode != "" {
				StockBuffer = 0
				StockTemp = 0
			}

			//cek jika channel tidak aktif stok 0
			CekChannel := stockRepository.CekChannelSku(Skunya, Channelnya)
			fmt.Println("status channel " + CekChannel.UuidStatus)
			if CekChannel.UuidStatus != "3936a00a-e321-f5ff-33b5-2dfb8d68ba45" {
				StockBuffer = 0
				StockTemp = 0
			}

			if Channelnya == "R996" { //SHOPEE
				StockBuffer = StockTemp
				cek, _ := StockShopee(IdParent, IdChild1, int64(StockBuffer), sku)
				//spew.Dump(hasil)
				if cek == false {
					fmt.Println("Update stock helper " + Skunya + " gagal " + strconv.Itoa(StockBuffer) + " " + waktu)

				} else {
					fmt.Println("Update stock helper " + Skunya + " sukses " + strconv.Itoa(StockBuffer) + " " + waktu)
				}
			} else if Channelnya == "R991" { //ZALORA
				//StockBuffer = objAvlb - int(StockBuffers)
				//cek, _ := StockZalora(Skunya, int64(StockBuffer))
				//spew.Dump(hasil)

				cek := false
				if IdChild1 != "" {
					cek, _ = StockZaloraV2(IdChild1, int64(StockBuffer), sku)
				} else {
					cek, _ = StockZalora(Skunya, int64(StockBuffer))
				}

				if cek == false {
					fmt.Println("Update stock helper " + Skunya + " gagal " + strconv.Itoa(StockBuffer) + " " + waktu)

				} else {
					fmt.Println("Update stock helper " + Skunya + " sukses " + strconv.Itoa(StockBuffer) + " " + waktu)
				}
			} else if Channelnya == "R990" { //TIKTOK
				StockBuffer = StockTemp
				cek, _ := StockTiktok(IdParent, IdChild1, int64(StockBuffer))
				//spew.Dump(hasil)
				if cek == false {
					fmt.Println("Update stock helper " + Skunya + " gagal " + strconv.Itoa(StockBuffer) + " " + waktu)

				} else {
					fmt.Println("Update stock helper " + Skunya + " sukses " + strconv.Itoa(StockBuffer) + " " + waktu)
				}
			} else if Channelnya == "R995" { //JDID
				StockBuffer = StockTemp
				skuId := IdChild1
				if skuId == "" {
					skuId = IdParent
				}
				cek, _ := StockJdId(skuId, int64(StockBuffer))
				//spew.Dump(hasil)
				if cek == false {
					fmt.Println("Update stock helper " + Skunya + " gagal " + strconv.Itoa(StockBuffer) + " " + waktu)

				} else {
					fmt.Println("Update stock helper " + Skunya + " sukses " + strconv.Itoa(StockBuffer) + " " + waktu)
				}

			} else if Channelnya == "R994" { //BLIBLI
				StockBuffer = StockTemp
				// cek, _ := StockBlibli(IdParent, int64(StockBuffer))
				cek, _ := StockBlibli(IdChild1, int64(StockBuffer))
				//spew.Dump(hasil)
				if cek == false {
					fmt.Println("Update stock helper " + Skunya + " gagal " + strconv.Itoa(StockBuffer) + " " + waktu)

				} else {
					fmt.Println("Update stock helper " + Skunya + " sukses " + strconv.Itoa(StockBuffer) + " " + waktu)
				}
			} else if Channelnya == "R992" { //BUKALAPAK
				StockBuffer = StockTemp
				cek, _ := StockBukalapak(IdParent, IdChild1, int64(StockBuffer))
				if cek == false {
					fmt.Println("Update stock helper " + Skunya + " gagal " + strconv.Itoa(StockBuffer) + " " + waktu)

				} else {
					fmt.Println("Update stock helper " + Skunya + " sukses " + strconv.Itoa(StockBuffer) + " " + waktu)
				}
			} else if Channelnya == "R998" { //TOKOPEDIA
				StockBuffer = StockTemp
				cek, _ := StockTokopedia(IdParent, int64(StockBuffer), sku)
				if cek == false {
					fmt.Println("Update stock helper " + Skunya + " gagal " + strconv.Itoa(StockBuffer) + " " + waktu)

				} else {
					fmt.Println("Update stock helper " + Skunya + " sukses " + strconv.Itoa(StockBuffer) + " " + waktu)
				}
			} else if Channelnya == "R997" { //LAZADA
				StockBuffer = StockTemp
				cek, _ := StockLazada(IdParent, IdChild1, int64(StockBuffer), sku)
				if cek == false {
					fmt.Println("Update stock helper " + Skunya + " gagal " + strconv.Itoa(StockBuffer) + " " + waktu)

				} else {
					fmt.Println("Update stock helper " + Skunya + " sukses " + strconv.Itoa(StockBuffer) + " " + waktu)
				}
			} else {
				fmt.Println("Channel SKU " + Skunya + " Tidak Ditemukan")
			}

			istok++
		}

		//end update stock API
	}
}

func UpdateStockNol(sku string, user string, channel string) {
	waktu := time.Now().String()
	fmt.Println("UpdateStockNol " + sku + " " + waktu)
	Skunya := sku

	var ObjSkuMapping []models.TableSkuMapping

	if channel == "" {
		ObjSkuMapping, _ = stockRepository.CariSkuParent(Skunya)
	} else {
		ObjSkuMapping, _ = stockRepository.CariSkuParentbychannel(Skunya, channel)
	}

	if len(ObjSkuMapping) > 0 {

		//update stock API
		istok := 0

		//temp_sku := ""
		temp_parent := ""
		temp_child := ""
		temp_channel := ""

		for _, ObjSkuMapping := range ObjSkuMapping {
			IdParent := ObjSkuMapping.IdSkuParent
			IdChild1 := ObjSkuMapping.IdSkuChild1
			Channelnya := ObjSkuMapping.ChannelCode

			StockBuffer := 0

			if istok > 0 { //yg ke2 atau lebih jadi 0
				//StockBuffer = 0
			}

			if StockBuffer < 1 {
				StockBuffer = 0
			}

			StockTemp := StockBuffer //R996

			//jika sku ada 2 atau lebih maka yg ke 2 jadi 0
			//if temp_sku != ObjSkuMapping.SkuNo && temp_parent != ObjSkuMapping.IdSkuParent && temp_channel != ObjSkuMapping.ChannelCode {
			if temp_channel != Channelnya {
				//temp_sku = ObjSkuMapping.SkuNo
				temp_parent = ObjSkuMapping.IdSkuParent
				temp_child = ObjSkuMapping.IdSkuChild1
				temp_channel = ObjSkuMapping.ChannelCode
			} else {

				if temp_parent != ObjSkuMapping.IdSkuParent && temp_child != ObjSkuMapping.IdSkuChild1 {
					StockTemp = 0
				}
			}

			//cek ke wms_sku_block_channel
			tgl := time.Now().Format("2006-01-02")
			objBlock := stockRepository.CariSkuBlock(Skunya, Channelnya, tgl)
			if objBlock.ChannelCode != "" {
				StockBuffer = 0
				StockTemp = 0
			}

			//cek jika channel tidak aktif stok 0
			CekChannel := stockRepository.CekChannelSku(Skunya, Channelnya)
			fmt.Println("status channel " + CekChannel.UuidStatus)
			if CekChannel.UuidStatus != "3936a00a-e321-f5ff-33b5-2dfb8d68ba45" {
				StockBuffer = 0
				StockTemp = 0
			}

			if Channelnya == "R996" { //SHOPEE
				StockBuffer = StockTemp
				cek, _ := StockShopee(IdParent, IdChild1, int64(StockBuffer), sku)
				//spew.Dump(hasil)
				if cek == false {
					fmt.Println("Update stock helper " + Skunya + " gagal " + strconv.Itoa(StockBuffer) + " " + waktu)

				} else {
					fmt.Println("Update stock helper " + Skunya + " sukses " + strconv.Itoa(StockBuffer) + " " + waktu)
				}
			} else if Channelnya == "R991" { //ZALORA
				//StockBuffer = objAvlb - int(StockBuffers)
				cek := false
				if IdChild1 != "" {
					cek, _ = StockZaloraV2(IdChild1, int64(StockBuffer), sku)
				} else {
					cek, _ = StockZalora(Skunya, int64(StockBuffer))
				}

				//spew.Dump(hasil)
				if cek == false {
					fmt.Println("Update stock helper " + Skunya + " gagal " + strconv.Itoa(StockBuffer) + " " + waktu)

				} else {
					fmt.Println("Update stock helper " + Skunya + " sukses " + strconv.Itoa(StockBuffer) + " " + waktu)
				}
			} else if Channelnya == "R990" { //TIKTOK
				StockBuffer = StockTemp
				cek, _ := StockTiktok(IdParent, IdChild1, int64(StockBuffer))
				//spew.Dump(hasil)
				if cek == false {
					fmt.Println("Update stock helper " + Skunya + " gagal " + strconv.Itoa(StockBuffer) + " " + waktu)

				} else {
					fmt.Println("Update stock helper " + Skunya + " sukses " + strconv.Itoa(StockBuffer) + " " + waktu)
				}
			} else if Channelnya == "R995" { //JDID
				StockBuffer = StockTemp
				skuId := IdChild1
				if skuId == "" {
					skuId = IdParent
				}
				cek, _ := StockJdId(skuId, int64(StockBuffer))
				//spew.Dump(hasil)
				if cek == false {
					fmt.Println("Update stock helper " + Skunya + " gagal " + strconv.Itoa(StockBuffer) + " " + waktu)

				} else {
					fmt.Println("Update stock helper " + Skunya + " sukses " + strconv.Itoa(StockBuffer) + " " + waktu)
				}

			} else if Channelnya == "R994" { //BLIBLI
				StockBuffer = StockTemp
				// cek, _ := StockBlibli(IdParent, int64(StockBuffer))
				cek, _ := StockBlibli(IdChild1, int64(StockBuffer))
				//spew.Dump(hasil)
				if cek == false {
					fmt.Println("Update stock helper " + Skunya + " gagal " + strconv.Itoa(StockBuffer) + " " + waktu)

				} else {
					fmt.Println("Update stock helper " + Skunya + " sukses " + strconv.Itoa(StockBuffer) + " " + waktu)
				}
			} else if Channelnya == "R992" { //BUKALAPAK
				StockBuffer = StockTemp
				cek, _ := StockBukalapak(IdParent, IdChild1, int64(StockBuffer))
				if cek == false {
					fmt.Println("Update stock helper " + Skunya + " gagal " + strconv.Itoa(StockBuffer) + " " + waktu)

				} else {
					fmt.Println("Update stock helper " + Skunya + " sukses " + strconv.Itoa(StockBuffer) + " " + waktu)
				}
			} else if Channelnya == "R998" { //TOKOPEDIA
				StockBuffer = StockTemp
				cek, _ := StockTokopedia(IdParent, int64(StockBuffer), sku)
				if cek == false {
					fmt.Println("Update stock helper " + Skunya + " gagal " + strconv.Itoa(StockBuffer) + " " + waktu)

				} else {
					fmt.Println("Update stock helper " + Skunya + " sukses " + strconv.Itoa(StockBuffer) + " " + waktu)
				}
			} else if Channelnya == "R997" { //LAZADA
				StockBuffer = StockTemp
				cek, _ := StockLazada(IdParent, IdChild1, int64(StockBuffer), sku)
				if cek == false {
					fmt.Println("Update stock helper " + Skunya + " gagal " + strconv.Itoa(StockBuffer) + " " + waktu)

				} else {
					fmt.Println("Update stock helper " + Skunya + " sukses " + strconv.Itoa(StockBuffer) + " " + waktu)
				}
			} else {
				fmt.Println("Channel SKU " + Skunya + " Tidak Ditemukan")
			}

			istok++
		}

		//end update stock API
	}
}

func InsertStockWms(sku string, user string, stock string) {
	//cek insert/update wms_sku_stock
	ObjSkuStock := stockRepository.CariSkuStockWMS(sku)
	var ObjSkuStockIns models.TableSkuStock
	var ObjSkuStockUpd models.TableSkuStock

	if ObjSkuStock.UuidSkuStock == "" {
		ObjSkuStockIns = models.TableSkuStock{
			UuidSkuStock: uuid.New().String(),
			SkuNo:        sku,
			Stock:        stock,
			CreatedBy:    user,
			CreatedDate:  time.Now(),
		}
	} else {

		ObjSkuStockUpd = models.TableSkuStock{
			UuidSkuStock: ObjSkuStock.UuidSkuStock,
			SkuNo:        ObjSkuStock.SkuNo,
			Stock:        stock,
			UpdatedBy:    user,
			UpdatedDate:  time.Now(),
		}
	}
	errs := stockRepository.SaveSkuStock(ObjSkuStockIns, ObjSkuStockUpd)
	if errs != nil {
		fmt.Println("INSERT/UPDATE DATA SKU " + sku + " STOCK GAGAL")
	}
}

// SHOPEE
type ResultStockShopee struct {
	Result ResultShopee `json:"result"`
}
type ResultShopee struct {
	Message string `json:"message"`
	Error   string `json:"error"`
	//Response  ResponseDetailShopee `json:"response"`
	RequestId string `json:"request_id"`
}

func StockShopee(parent string, child string, stock int64, sku string) (bool, ResultStockShopee) {
	status := false
	returnchannel := ResultStockShopee{}
	child1 := child
	if child1 == "" {
		child1 = "00"
	}
	urlshopee := os.Getenv("URL_API_SHOPEE") + "StockUpdateModelItem/" + parent + "/" + child1 + "/" + strconv.Itoa(int(stock)) + "/" + sku
	fmt.Println(urlshopee)
	req, err := http.NewRequest("GET", urlshopee, bytes.NewBuffer(nil))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		for key, val := range via[0].Header {
			req.Header[key] = val
		}
		return err
	}

	resp, err := client.Do(req)

	if err != nil {
		//fmt.Println("gagal get order detail shopee")

	} else {
		//fmt.Println("sukses get order detail")
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		//datas := parseJson(data)
		//datasx := utils.GetByteToInterface(data)
		//fmt.Println("======================")

		err := json.Unmarshal(data, &returnchannel)
		if err != nil {
			//return returnchannel
		}

		//fmt.Println(returnchannel)
		if returnchannel.Result.Message == "" {
			status = true
		}

	}

	return status, returnchannel
}

// ZALORA
type ResultStock struct {
	Message string `json:"Message"`
}

func StockZalora(sku string, stock int64) (bool, ResultStock) {
	status := false
	returnchannel := ResultStock{}

	url := os.Getenv("URL_WMS_ZALORA") + "UpdateStockProduct/" + sku + "/" + strconv.Itoa(int(stock))
	fmt.Println(url)
	req, err := http.NewRequest("GET", url, bytes.NewBuffer(nil))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		for key, val := range via[0].Header {
			req.Header[key] = val
		}
		return err
	}

	resp, err := client.Do(req)

	if err != nil {
		//fmt.Println("gagal get order detail shopee")

	} else {
		//fmt.Println("sukses get order detail")
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		//datas := parseJson(data)
		//datasx := utils.GetByteToInterface(data)
		//fmt.Println("======================")

		err := json.Unmarshal(data, &returnchannel)
		if err != nil {
			//return returnchannel
		}
		//fmt.Println(string(data))
		fmt.Println("Message " + returnchannel.Message)
		//fmt.Println(returnchannel)
		if returnchannel.Message == "" {
			status = true
		}

	}

	return status, returnchannel
}
func StockZaloraV2(productid string, stock int64, sku string) (bool, ResultStock) {
	status := false
	returnchannel := ResultStock{}

	url := os.Getenv("URL_WMS_ZALORA") + "UpdateStockProductV2/" + productid + "/" + strconv.Itoa(int(stock)) + "/" + sku
	fmt.Println(url)
	req, err := http.NewRequest("GET", url, bytes.NewBuffer(nil))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		for key, val := range via[0].Header {
			req.Header[key] = val
		}
		return err
	}

	resp, err := client.Do(req)

	if err != nil {
		//fmt.Println("gagal get order detail shopee")

	} else {
		//fmt.Println("sukses get order detail")
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		//datas := parseJson(data)
		//datasx := utils.GetByteToInterface(data)
		//fmt.Println("======================")

		err := json.Unmarshal(data, &returnchannel)
		if err != nil {
			//return returnchannel
		}
		//fmt.Println(string(data))
		fmt.Println("Message " + returnchannel.Message)
		//fmt.Println(returnchannel)
		if returnchannel.Message == "" {
			status = true
		}

	}

	return status, returnchannel
}

// TIKTOK
func StockTiktok(parent string, child string, stock int64) (bool, ResultStock) {
	status := false
	returnchannel := ResultStock{}

	child1 := child
	// if child1 == "" {
	// 	child1 = "00"
	// }
	url := os.Getenv("URL_WMS_TIKTOK") + "UpdateStockTiktok/" + parent + "/" + child1 + "/" + strconv.Itoa(int(stock))

	fmt.Println(url)
	req, err := http.NewRequest("GET", url, bytes.NewBuffer(nil))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		for key, val := range via[0].Header {
			req.Header[key] = val
		}
		return err
	}

	resp, err := client.Do(req)

	if err != nil {
		//fmt.Println("gagal get order detail shopee")

	} else {
		//fmt.Println("sukses get order detail")
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		//datas := parseJson(data)
		//datasx := utils.GetByteToInterface(data)
		//fmt.Println("======================")

		err := json.Unmarshal(data, &returnchannel)
		if err != nil {
			//return returnchannel
		}
		//fmt.Println(string(data))
		fmt.Println("Message " + returnchannel.Message)
		//fmt.Println(returnchannel)
		if returnchannel.Message == "" {
			status = true
		}

	}

	return status, returnchannel
}

// JDID
func StockJdId(parent string, stock int64) (bool, ResultStock) {
	status := false
	returnchannel := ResultStock{}

	url := os.Getenv("URL_WMS_JDID") + "UpdateStock/" + parent + "/" + strconv.Itoa(int(stock))

	fmt.Println(url)
	req, err := http.NewRequest("GET", url, bytes.NewBuffer(nil))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		for key, val := range via[0].Header {
			req.Header[key] = val
		}
		return err
	}

	resp, err := client.Do(req)

	if err != nil {

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		err := json.Unmarshal(data, &returnchannel)
		if err != nil {
			//return returnchannel
		}

		//fmt.Println("Message " + returnchannel.Message)

		if returnchannel.Message == "" {
			status = true
		}

	}

	return status, returnchannel
}

// BLIBLI
func StockBlibli(parent string, stock int64) (bool, ResultStock) {
	status := false
	returnchannel := ResultStock{}

	url := os.Getenv("URL_WMS_BLIBLI") + "StockUpdate/" + parent + "/" + strconv.Itoa(int(stock))

	fmt.Println(url)
	req, err := http.NewRequest("GET", url, bytes.NewBuffer(nil))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		for key, val := range via[0].Header {
			req.Header[key] = val
		}
		return err
	}

	resp, err := client.Do(req)

	if err != nil {

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		err := json.Unmarshal(data, &returnchannel)
		if err != nil {
			//return returnchannel
		}

		//fmt.Println("Message " + returnchannel.Message)

		if returnchannel.Message == "" {
			status = true
		}

	}

	return status, returnchannel
}

// BUKALAPAK
func StockBukalapak(parent string, child string, stock int64) (bool, ResultStock) {
	status := false
	returnchannel := ResultStock{}

	url := os.Getenv("URL_WMS_BUKALAPAK") + "UpdateStock/" + parent + "/" + child + "/" + strconv.Itoa(int(stock))

	fmt.Println(url)
	req, err := http.NewRequest("GET", url, bytes.NewBuffer(nil))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		for key, val := range via[0].Header {
			req.Header[key] = val
		}
		return err
	}

	resp, err := client.Do(req)

	if err != nil {

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		err := json.Unmarshal(data, &returnchannel)
		if err != nil {
			//return returnchannel
		}

		//fmt.Println("Message " + returnchannel.Message)

		if returnchannel.Message == "" {
			status = true
		}

	}

	return status, returnchannel
}

// TOKOPEDIA
func StockTokopedia(parent string, stock int64, sku string) (bool, ResultStock) {
	status := false
	returnchannel := ResultStock{}

	url := os.Getenv("URL_WMS_TOKPED") + "UpdateStock/" + parent + "/" + strconv.Itoa(int(stock)) + "/" + sku

	fmt.Println(url)
	req, err := http.NewRequest("GET", url, bytes.NewBuffer(nil))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		for key, val := range via[0].Header {
			req.Header[key] = val
		}
		return err
	}

	resp, err := client.Do(req)

	if err != nil {

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		err := json.Unmarshal(data, &returnchannel)
		if err != nil {
			//return returnchannel
		}

		//fmt.Println("Message " + returnchannel.Message)

		if returnchannel.Message == "" {
			status = true
		}

	}

	return status, returnchannel
}

// LAZADA
func StockLazada(parent string, child string, stock int64, sku string) (bool, ResultStock) {
	status := false
	returnchannel := ResultStock{}

	url := os.Getenv("URL_WMS_LAZADA") + "UpdateStock/" + parent + "/" + child + "/" + strconv.Itoa(int(stock)) + "/" + sku

	fmt.Println(url)
	req, err := http.NewRequest("GET", url, bytes.NewBuffer(nil))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		for key, val := range via[0].Header {
			req.Header[key] = val
		}
		return err
	}

	resp, err := client.Do(req)

	if err != nil {

	} else {
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		err := json.Unmarshal(data, &returnchannel)
		if err != nil {
			//return returnchannel
		}

		//fmt.Println("Message " + returnchannel.Message)

		if returnchannel.Message == "" {
			status = true
		}

	}

	return status, returnchannel
}
