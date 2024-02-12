package runner

type ModbusData struct {
	Address  int `json:"address"`
	Function int `json:"function"`
	Length   int `json:"length"`
	Value    int `json:"value"`
}

type Data struct {
	SlaveId    int          `json:"slaveId"`
	ModbusData []ModbusData `json:"modbusData"`
}

type MessageData struct {
	MessageId int  `json:"messageId"`
	LoraRssi  int  `json:"loraRssi"`
	Ts        int  `json:"ts"`
	Data      Data `json:"data"`
}

type MessageSent struct {
	MessageId int            `json:"messageId"`
	LoraRssi  int            `json:"loraRssi"`
	Ts        int            `json:"ts"`
	Data      map[string]int `json:"data"`
}

func checkFieldType(code int) string {
	field := map[int]string{
		22: "TDS",
	}

	if field[code] == "" {
		return "flow"
	}
	return field[code]
}
