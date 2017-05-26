package data

import (
	"../util"
	"../model"
	"fmt"
	"encoding/json"
	"github.com/satori/go.uuid"
	"time"
)

func FindTGT(ticket string) *model.TGT {
	key := util.TICKET_GRANTING_TICKET_PREFIX + ticket
	val, err := Cli.Get(key).Result()
	if err != nil {
		fmt.Println(err)
	}
	var tgt = new(model.TGT)
	jsonErr := json.Unmarshal([]byte(val), tgt)
	if err != nil {
		fmt.Println(jsonErr)
		return nil
	}
	return tgt
}

func FindST(ticket string) *model.ServiceTicket{
	key := util.SERVICE_TICKET_PREFIX + ticket
	val, err := Cli.Get(key).Result()
	if err != nil {
		fmt.Println(err)
	}
	var st = new(model.ServiceTicket)
	jsonErr := json.Unmarshal([]byte(val), st)
	if err != nil {
		fmt.Println(jsonErr)
		return nil
	}
	return st
}

func AddSTToTGT(tgt *model.TGT, st *model.ServiceTicket) {
	sts := tgt.St
	newSts := append(sts, st.St)
	tgt.St = newSts

	jsonB, err := json.Marshal(tgt)
	if err != nil {
		fmt.Println(err)
		return
	}

	key := util.TICKET_GRANTING_TICKET_PREFIX + tgt.Tgt

	rErr := Cli.Set(key, string(jsonB), time.Millisecond*util.TICKET_GRANTING_TICKET_TIME_TO_LIVE).Err()
	if rErr != nil {
		fmt.Println(err)
	}
}

func GrantServiceTicket(tgt string, service string) *model.ServiceTicket {
	sTicket := util.ST_PREFIX + uuid.NewV4().String()
	var st = new(model.ServiceTicket)
	st.Service = service
	st.St = sTicket
	st.Tgt = tgt

	jsonB, err := json.Marshal(st)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	key := util.SERVICE_TICKET_PREFIX + sTicket
	rErr := Cli.Set(key, string(jsonB), time.Millisecond*util.SERVICE_TICKET_TIME_TO_LIVE).Err()
	if rErr != nil {
		fmt.Println(err)
		return nil
	}
	return st
}

func GrantTicketGrantingTicket(username string, st string) *model.TGT {
	tTicket := util.TGT_PREFIX + uuid.NewV4().String()
	var tgt = new(model.TGT)
	tgt.Tgt = tTicket
	if st != "" {
		sts := []string{st}
		tgt.St = sts
	}

	tgt.Username = username

	jsonB, err := json.Marshal(tgt)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	key := util.TICKET_GRANTING_TICKET_PREFIX + tTicket
	rErr := Cli.Set(key, string(jsonB), time.Millisecond*util.TICKET_GRANTING_TICKET_TIME_TO_LIVE).Err()
	if rErr != nil {
		fmt.Println(err)
		return nil
	}
	return tgt
}
