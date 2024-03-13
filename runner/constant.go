package runner

import "strconv"

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

func checkFieldType(slaveId int, address int) string {
	slave := strconv.Itoa(slaveId)
	addrs := strconv.Itoa(address)

	if slave == "2" {
		code := slave + addrs
		field := map[string]string{
			"20":  "conductivity",
			"22":  "temperature",
			"28":  "tds",
			"210": "salinity",
			"212": "conductivity_calibration",
			"235": "conductivity_signal_1",
			"236": "conductivity_signal_2",
			"237": "temperature_signal_AD",
		}
		return field[code]
	} else {
		return "volume"
	}
}
