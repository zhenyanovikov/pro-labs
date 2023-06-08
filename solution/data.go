package solution

import (
	"encoding/json"
	"log"
	"math/rand"
	"os"
)

const fileName = "data.txt"

type Data struct {
	N int

	MT Matrix
	MX Matrix
	MD Matrix
	ME Matrix
	B  Vector
	D  Vector
	A  float64
}

func GetData(n int) *Data {
	// check if file exists
	if _, err := os.Stat(fileName); err == nil {
		// file exists
		return ReadDataFromFile(n)
	}

	data := generateData(n)

	// save data to file
	SaveDataToFile(data)

	return data
}

func SaveDataToFile(data *Data) {
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err = encoder.Encode(data); err != nil {
		log.Fatal(err)
	}
}

func ReadDataFromFile(n int) *Data {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var data Data

	if err = json.NewDecoder(file).Decode(&data); err != nil {
		log.Fatal(err)
		return nil
	}

	if data.N != n {
		newData := generateData(n)
		SaveDataToFile(newData)
		return newData
	}

	return &data
}

func generateData(n int) *Data {
	data := &Data{
		N:  n,
		MT: createMatrix(n),
		MX: createMatrix(n),
		MD: createMatrix(n),
		ME: createMatrix(n),
		B:  createVector(n),
		D:  createVector(n),
		A:  rand.Float64(),
	}

	return data
}
