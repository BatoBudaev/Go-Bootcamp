package recipes

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"strings"
)

type DBReader interface {
	ReadData(filePath string) (*Recipes, error)
	PrintData(data *Recipes)
}

type JSONReader struct{}

func (r *JSONReader) ReadData(filePath string) (*Recipes, error) {
	jsonData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var data Recipes
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *JSONReader) PrintData(data *Recipes) {
	xmlOutput, err := xml.MarshalIndent(data, "", "    ")
	if err != nil {
		fmt.Println("Ошибка:", err)
		os.Exit(1)
	}

	fmt.Println(string(xmlOutput))
}

type XMLReader struct{}

func (r *XMLReader) ReadData(filePath string) (*Recipes, error) {
	xmlData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var data Recipes
	err = xml.Unmarshal(xmlData, &data)
	if err != nil {

		return nil, err
	}

	return &data, nil
}

func (r *XMLReader) PrintData(data *Recipes) {
	jsonOutput, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		fmt.Println("Ошибка:", err)
		os.Exit(1)
	}

	fmt.Println(string(jsonOutput))
}

func ReadFile(filePath string) (DBReader, *Recipes) {
	var reader DBReader

	if strings.HasSuffix(filePath, ".json") {
		reader = &JSONReader{}
	} else if strings.HasSuffix(filePath, ".xml") {
		reader = &XMLReader{}
	} else {
		fmt.Println("Неверный формат файла")
		os.Exit(1)
	}

	data, err := reader.ReadData(filePath)
	if err != nil {
		fmt.Println("Ошибка чтения:", err)
		os.Exit(1)
	}

	return reader, data
}
