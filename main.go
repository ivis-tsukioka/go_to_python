package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

func main() {
	c := make(chan string)
	go add(1, 2, c)
	go add(3, 4, c)
	result01, result02 := <-c, <-c
	print(result01)
	print(result02)

	jsonfunc("dai", "tsukioka")
	jsonfunc("dai2", "tsukioka2")
}

func add(a, b int, c chan string) {
	pycode := fmt.Sprintf("import app; app.add(%d, %d)", a, b)
	out, _ := exec.Command("python", "-c", pycode).CombinedOutput()
	c <- string(out)
}

func jsonfunc(v1, v2 string) {
	stcData := JsonStr{Key1: v1, Key2: v2}
	jsonData, _ := json.Marshal(stcData)
	fmt.Println(string(jsonData))
	pycode := fmt.Sprintf("import app; app.printJSONString(%s)", jsonData)
	out, err := exec.Command("python", "-c", pycode).CombinedOutput()
	if err != nil {
		fmt.Println(fmt.Sprintf("Err %v", err))
		fmt.Println(string(out))
		return
	}
	fmt.Println(string(out))
}

type JsonStr struct {
	Key1 string `json:"key1"`
	Key2 string `json:"key2"`
}
