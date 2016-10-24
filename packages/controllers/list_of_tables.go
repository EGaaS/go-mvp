// Copyright 2016 The go-daylight Authors
// This file is part of the go-daylight library.
//
// The go-daylight library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-daylight library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-daylight library. If not, see <http://www.gnu.org/licenses/>.

package controllers

import (
	"github.com/EGaaS/go-mvp/packages/utils"
)

type listOfTablesPage struct {
	Alert        string
	SignData     string
	ShowSignData bool
	CountSignArr []int
	Lang         map[string]string
	WalletId     int64
	CitizenId    int64
	TxType       string
	TxTypeId     int64
	TimeNow      int64
	Global       string
	Tables       []map[string]string
}

func (c *Controller) ListOfTables() (string, error) {

	var err error

	global := c.r.FormValue("global")
	prefix := c.StateIdStr
	if global == "1" {
		prefix = "global"
	}
	tables, err := c.GetAll(`SELECT * FROM "`+prefix+`_tables"`, -1)
	if err != nil {
		return "", utils.ErrInfo(err)
	}
	for i, data := range tables {
		count, err := c.Single(`SELECT count(id) FROM "` + data["name"] + `"`).Int64()
		if err != nil {
			return "", utils.ErrInfo(err)
		}
		tables[i]["count"] = utils.Int64ToStr(count)
	}

	TemplateStr, err := makeTemplate("list_of_tables", "listOfTables", &listOfTablesPage{
		Alert:        c.Alert,
		Lang:         c.Lang,
		ShowSignData: c.ShowSignData,
		Global:       global,
		SignData:     "",
		WalletId:     c.SessWalletId,
		CitizenId:    c.SessCitizenId,
		CountSignArr: c.CountSignArr,
		Tables:       tables})
	if err != nil {
		return "", utils.ErrInfo(err)
	}
	return TemplateStr, nil
}
