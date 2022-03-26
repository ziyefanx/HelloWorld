package service

import (
	"awesomeProject1/model"
	"fmt"
)

type StudentHandler interface {
	SatisfyGraduate(majorCredit, artCredit, PECredit int) bool
}

type NormalStudent struct{}

func (s NormalStudent) SatisfyGraduate(majorCredit, artCredit, PECredit int) bool {
	if majorCredit >= 10 && artCredit >= 3 && PECredit >= 5 {
		return true
	}
	return false
}

type ArtStudent struct{}

func (s ArtStudent) SatisfyGraduate(majorCredit, artCredit, PECredit int) bool {
	if majorCredit >= 5 && artCredit >= 10 && PECredit >= 3 {
		return true
	}
	return false
}

// 工厂模式
func GetStudentHandler(stuType model.StudentType) (StudentHandler, error) {
	switch stuType {
	case model.StudentTypeNormal:
		return NormalStudent{}, nil
	case model.StudentTypeArt:
		return ArtStudent{}, nil
	}
	return nil, fmt.Errorf("unsupport student type")
}
