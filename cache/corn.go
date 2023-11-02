package cache

import (
	"regexp"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/robfig/cron"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"server-template/model"
)

var Cron *cron.Cron = cron.New()

var DLCReg = `element id="([a-zA-Z0-9]{64}i0)"`

var DLCToAssets = map[string][]model.Asset{}

var Postgres *gorm.DB
var startId int64 = 0

func InitCronPostgres(db *gorm.DB) {
	Postgres = db
}

func cronForBRC420AssetsDLC() {
	var assets []model.Asset
	Postgres.Where("id > ? and type = '' and tag = '' and category = '' and collection = ''", startId).Find(&assets)

	for _, asset := range assets {
		oss_url := "https://inscriptions.oss-ap-southeast-1.aliyuncs.com/inscriptions/" + asset.InscriptionId

		body := ""
		for {
			resp, err := resty.New().R().
				Get(oss_url)
			if err != nil {
				log.WithFields(log.Fields{
					"id":             asset.Id,
					"inscription_id": asset.InscriptionId,
				}).Errorf("Failed to fetch inscription content: %v, sleep 60s", err)
				time.Sleep(60 * time.Second)
				continue
			}

			body = string(resp.Body())

			if strings.Contains(body, "<Code>NoSuchKey</Code>") {
				log.WithFields(log.Fields{
					"id":             asset.Id,
					"inscription_id": asset.InscriptionId,
				}).Error("No Such Key, sleep 60s.")
				time.Sleep(60 * time.Second)
				continue
			}
			break
		}

		subs := regexp.MustCompile(DLCReg).FindAllStringSubmatch(body, -1)
		brc420AssetsInscriptionIdList := []string{}
		for _, sub := range subs {
			brc420AssetsInscriptionIdList = append(brc420AssetsInscriptionIdList, sub[1])
		}

		var tempAssets []model.Asset
		Postgres.Where("inscription_id in ?", brc420AssetsInscriptionIdList).Find(&tempAssets)
		DLCToAssets[asset.InscriptionId] = tempAssets
	}

	if len(assets) != 0 {
		startId = assets[len(assets)-1].Id
	}
}

func InitCron() {
	err := Cron.AddFunc("0 * * * * *", cronForBRC420AssetsDLC)

	if err != nil {
		panic(err)
	}

	Cron.Start()
}
