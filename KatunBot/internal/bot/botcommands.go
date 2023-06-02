package bot

import(
	"time"
	"log"
	"fmt"
	"strings"
	"katun/lib/months"
)



func (b *Bot) regVkat(userId int64, userTag string, date string, messageId int, chatid int64){

	if date == "" {
		b.sendMsg(`После команды необходимо указать дату "/vkat 20ХХ-ММ-ДД"`, chatid, messageId)
		return
	}

	parsedDate, err:= time.Parse(time.DateOnly, date)
	if err!=nil{
		log.Println(err)
		b.sendMsg(`Неверная дата, отправь дату в формате "20ХХ-мм-дд."`, chatid, messageId)
		return
	}

	if parsedDate.Before(time.Now()){
		b.sendMsg(`Это прошедшая дата, введи другую.`, chatid, messageId)
		return
	}

	_, err1:=b.Database.AddVkat(userId, userTag, parsedDate)
	if err1!=nil{
		log.Println(err1)
		if strings.Contains(fmt.Sprint(err1), "(SQLSTATE 23505)"){
			b.sendMsg("Ты уже вкатываешься!", chatid, messageId)
		}
		return
	}

	text:= fmt.Sprintf("Желаю удачи с вкатом, дедлай: %v", date)
	b.sendMsg(text, chatid, messageId)
}





func (b *Bot) deadline(userId int64, messageId int, chatid int64){

	_,finish, err:=b.Database.Time(userId)
	if err!= nil{
		log.Println(err)
		if strings.Contains(fmt.Sprint(err), "sql: no rows in result set"){
			b.sendMsg("Ты не вкатываешься! Воспульзуйся командой /vkat чобы начать вкат.", chatid, messageId)
		}
		return
	}

	parsed,err:= time.Parse(time.RFC3339, finish )
	if err!=nil{
		log.Println(err)
		return
	}

	txt:= fmt.Sprintf("Дедлайн вката: %v", monthRus.FormatDate(parsed))
	b.sendMsg(txt, chatid, messageId)
}

func (b *Bot) daysUntilVkat(userId int64, messageId int, chatid int64){
	
	_, finish, err:= b.Database.Time(userId)
	if err != nil{
		if strings.Contains(fmt.Sprint(err), "sql: no rows in result set"){
			b.sendMsg("Ты не вкатываешься! Воспульзуйся командой /vkat чобы начать вкат.",chatid, messageId )
		}
		log.Println(err)
		return
	}

	parsed,err:= time.Parse(time.RFC3339, finish )
	if err!=nil{
		log.Println(err)
		return
	}




	txt:= fmt.Sprintf("До вката осталось: %v", monthRus.DaysUntil(parsed))

	b.sendMsg(txt, chatid, messageId)


}

func (b Bot) vkatuns(messageId int, chatid int64){

	vkatunsArr, err:= b.Database.VkatMembers()
	if err!=nil{
		log.Println(err)
		return
	}

	var txtMsg string

	for _,el:= range vkatunsArr{

		parsedFinish, err1:= time.Parse(time.RFC3339 ,el.Finish)
		if err1!=nil{
			log.Println(err1)
			return
		}


		txt:= fmt.Sprintf("Вкатун: @%v, До вката осталось %v \n\n", el.UserTag, monthRus.DaysUntil(parsedFinish))
		txtMsg+= txt
	}

	b.sendMsg(txtMsg,chatid, messageId)
}

func (b *Bot) DeleteVkat(userId int64, prove string, messageId int, chatid int64){
	if prove != "удалить" {
		txt:= `Для удаления вката напишите компанду "/delete удалить"`
		b.sendMsg(txt, chatid,messageId)
		return
	}

	err:= b.Database.DeleteVkat(userId)
	if err!= nil{
		if strings.Contains(fmt.Sprint(err), "sql: no rows in result set"){
			b.sendMsg("Ты не вкатываешься! Воспульзуйся командой /vkat чобы начать вкат.",chatid, messageId )
		}
		log.Println(err)
		return
	}

	b.sendMsg("Вкат удален.",chatid, messageId)
}

func (b *Bot) updateVkat(userId int64, date string, messageId int, chatid int64){

	if date == "" {
		b.sendMsg(`После команды необходимо указать дату "/updatevkat 20ХХ-ММ-ДД"`, chatid, messageId)
		return
	}

	parsedDate, err:= time.Parse(time.DateOnly, date)
	if err!=nil{
		log.Println(err)
		b.sendMsg(`Неверная дата, отправь дату в формате "20ХХ-мм-дд."`, chatid, messageId)
		return
	}

	if parsedDate.Before(time.Now()){
		b.sendMsg(`Это прошедшая дата, введи другую.`, chatid, messageId)
		return
	}


	err= b.Database.UpdateVkat(userId,parsedDate)
	if err!=nil{
		if strings.Contains(fmt.Sprint(err), "sql: no rows in result set"){
			b.sendMsg("Ты не вкатываешься или указал ту же дату!",chatid, messageId )
		}
		log.Println(err)
		return
	}

	txt:= fmt.Sprintf("Дата вката обновлена, новый дедлайн: %v", date)

	b.sendMsg(txt, chatid, messageId)
}