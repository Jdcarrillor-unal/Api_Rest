package models

import "time"

type Data struct {
	Id           int       `json:"id"`
	Temperatura  string    `json:"temperatura"`
	Humedad      string    `json:"humedad"`
	Boton        bool      `json:"boton"`
	Created_date time.Time `json:"time"`
}

func GetData(sql string, conditional interface{}) *Data {
	data := &Data{}
	rows, err := Query(sql, conditional)
	if err != nil {
		return data
	}
	for rows.Next() {
		rows.Scan(&data.Id, &data.Temperatura, &data.Humedad, &data.Boton, &data.Created_date)
	}
	return data
}

func GetTempById(id string) *Data {
	sql := "SELECT id,temperatura,humedad,boton,created_date FROM DATA WHERE id=?"
	return GetData(sql, id)
}
