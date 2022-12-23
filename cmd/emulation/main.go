package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	time.Sleep(8 * time.Second)

	data := `000
************0000
J300000 
298000
MIR
Успешно
200000
100000000000`

	cheque := `
         Спасибо за покупку!
              Чек 945
               Оплата
              ОДОБРЕНО
СУММА:                    890.00 RUR
Комиссия:                   0.00 RUR
AID: A0000000032333    Visa Electron
TVR: 0000000123                 TSI:
CID: 80             C744EF5B0DE356E4
Карта: VISA
             Cless EMV
        400000******0000:00

Ссылка:100000000000      ТID:J300000
Мерчант:                          mo
Код авторизации:              200000
Код ответа (хост):                00
Код ответа (эмитент):             00
Дата:              19/03/21 14:50:02
Проверен на устройстве
Подпись клиента не нужна

      ________________________
         (Подпись кассира)

           Сохраните чек
     Банк
====================================`

	c, _ := os.Create("cheq.out")
	_, _ = c.WriteString(cheque)

	d, _ := os.Create("output.out")
	_, _ = d.WriteString(data)

	fmt.Println(data)
}
