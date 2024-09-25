package datafile 

import (
	"bufio"
	"os"
	"strconv"
)

func GetFloats(fileName string) ([3]float64, error){
	var numbers [3]float64
	file, err := os.Open(fileName)
	if err != nil{ //если при открытии файла произошла ошибка
		return numbers, err 
	} 
	i:=0 //хранение индекса, которому должно выполняться присвоения
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){//читаем строку из файла 
		numbers[i], err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil{//если при приобразовании значения - произошла ошибка, возвращаем её, иначе переходим к следующему индексу
			return numbers, err
		}
		i++
	}
	err = file.Close()
	if err != nil {
		return numbers, err
	}
	if scanner.Err() != nil {
		return numbers, scanner.Err()
	}
	return numbers, nil //программа возвращает массив значений и пустую ошибку 
}

