package repository

import (
	"encoding/csv"
	"os"
	"strconv"

	"github.com/Le-MaliX/ACADEMY-GO-Q42021/domain/model"
)

func openCsv() ([][]string, error) {
	file, err := os.Open("./infrastructure/datastore/monsters.csv")
	if err != nil {
		return nil, err
	}

	defer file.Close()

	f := csv.NewReader(file)
	f.Comma = ';'
	lines, err := f.ReadAll()
	if err != nil {
		return nil, err
	}
	return lines, nil
}

func GetMonstersData() ([]model.Monster, error) {
	csvLines, err := openCsv()
	if err != nil {
		return nil, err
	}

	var monsters []model.Monster

	for _, line := range csvLines {
		if line[0] == "ID" {
			continue
		}
		id, err := strconv.Atoi(line[0])
		if err != nil {
			return nil, err
		}

		hp, err := strconv.Atoi(line[19])
		if err != nil {
			return nil, err
		}

		monster := model.Monster{
			Id:              id,
			Name:            line[1],
			ChallengeRating: line[2],
			HPDice:          line[18],
			HP:              hp,
		}
		monsters = append(monsters, monster)
	}

	return monsters[1:], nil
}
