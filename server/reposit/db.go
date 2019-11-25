package reposit

import (
	"misc/grpc-rest-mock/grpc/model"
	"strconv"
)

func GetData(from string, to string) ([]*model.StudentDetails, error) {
	rows := make([]*model.StudentDetails, 0, 100)
	var i int32
	for i = 1; i <= 2; i++ {
		str := strconv.Itoa(int(i))
		rows = append(rows, &model.StudentDetails{
			Id:          i,
			Name:        "Test" + str,
			RollNo:      "Roll-" + str,
			Age:         i + 20,
			ExamCleared: true,
		})
	}
	return rows, nil
}
