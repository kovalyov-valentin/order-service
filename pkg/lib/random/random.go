package random

import (
	"encoding/json"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/kovalyov-valentin/orders-service/internal/models"
	"github.com/sirupsen/logrus"
)

func RandOrder() []byte {
	uid := uuid.New()
	trackNumber := RandString(10)
	price := GetRandInt(100, 1e6)
	sale := GetRandInt(0, 100)
	order := models.Order{
		OrderUID:    uid.String(),
		TrackNumber: trackNumber,
		Entry:       RandString(4),
		Delivery: models.Delivery{
			Name:    RandString(20),
			Phone:   "+79" + RandNumberToString(9),
			Zip:     RandString(7),
			City:    RandString(30),
			Address: RandString(30),
			Region:  RandString(30),
			Email:   RandString(10) + "@gmail.com",
		},
		Payment: models.Payment{
			Transaction:  uid.String(),
			RequestID:    RandNumberToString(10),
			Currency:     RandNumberToString(3),
			Provider:     RandNumberToString(3),
			Amount:       GetRandInt(1, 1e6),
			PaymentDt:    GetRandInt(1e8, 1e9-1),
			Bank:         RandString(10),
			DeliveryCost: GetRandInt(100, 1e5),
			GoodsTotal:   GetRandInt(1, 1e4),
			CustomFee:    GetRandInt(0, 20),
		},
		Items: []models.Item{
			{
				ChrtID:      GetRandInt(1e7, 1e9),
				TrackNumber: trackNumber,
				Price:       price,
				Rid:         RandString(20),
				Name:        RandString(20),
				Sale:        sale,
				Size:        RandString(20),
				TotalPrice:  price * (100 - sale) / 100,
				NmID:        GetRandInt(1e6, 1e7-1),
				Brand: RandString(
					GetRandInt(5, 10)) + " " +
					RandString(GetRandInt(2, 8)),
				Status: GetRandInt(100, 500),
			},
		},
		Locale:            RandString(20),
		InternalSignature: RandString(20),
		CustomerID:        RandString(20),
		DeliveryService:   RandString(20),
		ShardKey:          RandString(20),
		SmID:              12323232,
		DateCreated:       time.Now(),
		OofShard:          RandString(20),
	}

	bytes, err := json.Marshal(&order)
	if err != nil {
		logrus.Println("error serialization")
		return nil
	}
	return bytes
}

func RandString(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(b)
}

func RandNumberToString(n int) string {
	var letterRunes = []rune("0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(b)
}

func GetRandInt(min int, max int) int {
	return rand.Intn(max-min) + min
}
