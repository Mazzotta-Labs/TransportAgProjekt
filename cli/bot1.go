package cli

import (
	"TransportAgProjekt/model"
	"TransportAgProjekt/model/entity"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
)

func BotMessage() {
	bot, err := tgbotapi.NewBotAPI("5010157425:AAHCefPFUxTbLgrDr43GiqlJoBAUnoH4E8o")
	if err != nil {
		println(err)
		//log.Panic(err)
	}

	bot.Debug = true

	result := model.FindAllOrder()
	printBotOrderList(result)
	txt := printBotOrderList(model.FindAllOrder())
	println(txt)
	msg := tgbotapi.NewMessage(-712071829, txt)

	bot.Send(msg)

}

func printBotOrderList(ordersToPrint []entity.Order) string {
	var txt string
	txt = "***** Auftrag Print *****\n"
	for i, order := range ordersToPrint {
		txt = txt + "Auftrag " + toStr(i+1) + "\n" +
			"Auftrags Nr.: " + toStr(order.OrderId) + ",\n" +
			"Datum: " + order.OrderDate + ",\n" +
			"Name: " + order.CustomerPrename + ",\n" +
			"Vorname: " + order.CustomerName + ",\n" +
			"PLZ: " + order.Plz + ",\n" +
			"Ort: " + order.Town + ",\n" +
			"Strasse: " + order.Street + ",\n" +
			"Land: " + order.Country + ",\n" +
			"Fahrer Name: " + order.Name + ",\n" +
			"Fahrer Vorame :" + order.Prename + ",\n" +
			"Fahrzeug Nr. :" + order.Number + ",\n" + "\n"
	}
	return txt
}

func toStr(info int) string {
	aStr := strconv.Itoa(info)
	return aStr
}
