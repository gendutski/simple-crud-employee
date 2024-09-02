package usecase_test

import (
	"errors"
	"simple-crud-employee/internal/entity"
	"simple-crud-employee/internal/usecase"
	"simple-crud-employee/internal/usecase/mocks"
	"strings"
	"testing"

	"github.com/go-sql-driver/mysql"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func initUsecase(ctrl *gomock.Controller) (*usecase.EmployeeUsecase, *mocks.MockEmployeeRepo) {
	repo := mocks.NewMockEmployeeRepo(ctrl)
	uc := usecase.InitEmployeeUsecase(repo)
	return uc, repo
}

func Test_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	svc, repo := initUsecase(ctrl)

	t.Run("success", func(t *testing.T) {
		payload := entity.Employee{
			EmployeeID: "1001",
			FullName:   "Firman",
			Address:    "Jakarta Timur",
		}
		repo.EXPECT().Create(&payload).Return(nil).Times(1)

		err := svc.Create(&payload)
		assert.Nil(t, err)
	})

	t.Run("payload exceeded max char", func(t *testing.T) {
		payload := entity.Employee{
			EmployeeID: "10011",
			FullName:   strings.Repeat("x", 200),
		}

		err := svc.Create(&payload)
		assert.NotNil(t, err)
		echoErr, ok := err.(*echo.HTTPError)
		assert.True(t, ok)
		assert.Equal(t, 400, echoErr.Code)
		assert.Equal(t, map[string][]*entity.ValidatorMessage{
			"employeeID": {{Tag: "max", Param: "4"}},
			"fullName":   {{Tag: "max", Param: "100"}},
			"address":    {{Tag: "required"}},
		}, echoErr.Message)
	})

	t.Run("empty required payload", func(t *testing.T) {
		payload := entity.Employee{}

		err := svc.Create(&payload)
		assert.NotNil(t, err)
		echoErr, ok := err.(*echo.HTTPError)
		assert.True(t, ok)
		assert.Equal(t, 400, echoErr.Code)
		assert.Equal(t, map[string][]*entity.ValidatorMessage{
			"employeeID": {{Tag: "required"}},
			"fullName":   {{Tag: "required"}},
			"address":    {{Tag: "required"}},
		}, echoErr.Message)
	})

	t.Run("duplicate", func(t *testing.T) {
		payload := entity.Employee{
			EmployeeID: "1001",
			FullName:   "Firman",
			Address:    "Jakarta Timur",
		}
		repo.EXPECT().Create(&payload).Return(&mysql.MySQLError{Number: 1062, Message: "duplicate"}).Times(1)

		err := svc.Create(&payload)
		assert.NotNil(t, err)
		echoErr, ok := err.(*echo.HTTPError)
		assert.True(t, ok)
		assert.Equal(t, 409, echoErr.Code)
	})

	t.Run("internal error", func(t *testing.T) {
		payload := entity.Employee{
			EmployeeID: "1001",
			FullName:   "Firman",
			Address:    "Jakarta Timur",
		}
		repo.EXPECT().Create(&payload).Return(errors.New("error")).Times(1)

		err := svc.Create(&payload)
		assert.NotNil(t, err)
		echoErr, ok := err.(*echo.HTTPError)
		assert.True(t, ok)
		assert.Equal(t, 500, echoErr.Code)
	})
}

func Test_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	svc, repo := initUsecase(ctrl)

	t.Run("success", func(t *testing.T) {
		employeeID := "1001"
		employee := &entity.Employee{
			EmployeeID: "1001",
			FullName:   "Firman",
			Address:    "Jakarta Timur",
		}
		payload := entity.Employee{
			EmployeeID: "1002",
			FullName:   "Firman Darmawan",
			Address:    "Cipinang Melayu, Jakarta Timur",
		}

		repo.EXPECT().GetDetail(employeeID).Return(employee, nil).Times(1)
		repo.EXPECT().Update(payload, employee).Return(nil).Times(1)

		err := svc.Update(employeeID, payload)
		assert.Nil(t, err)
	})

	t.Run("payload exceeded max char", func(t *testing.T) {
		employeeID := "1001"
		payload := entity.Employee{
			EmployeeID: "10022",
			FullName:   strings.Repeat("f", 200),
		}

		err := svc.Update(employeeID, payload)
		assert.NotNil(t, err)
		echoErr, ok := err.(*echo.HTTPError)
		assert.True(t, ok)
		assert.Equal(t, 400, echoErr.Code)
		assert.Equal(t, map[string][]*entity.ValidatorMessage{
			"employeeID": {{Tag: "max", Param: "4"}},
			"fullName":   {{Tag: "max", Param: "100"}},
			"address":    {{Tag: "required"}},
		}, echoErr.Message)
	})

	t.Run("duplicate", func(t *testing.T) {
		employeeID := "1001"
		employee := &entity.Employee{
			EmployeeID: "1001",
			FullName:   "Firman",
			Address:    "Jakarta Timur",
		}
		payload := entity.Employee{
			EmployeeID: "1002",
			FullName:   "Firman Darmawan",
			Address:    "Cipinang Melayu, Jakarta Timur",
		}

		repo.EXPECT().GetDetail(employeeID).Return(employee, nil).Times(1)
		repo.EXPECT().Update(payload, employee).Return(&mysql.MySQLError{Number: 1062, Message: "duplicate"}).Times(1)

		err := svc.Update(employeeID, payload)
		assert.NotNil(t, err)
		echoErr, ok := err.(*echo.HTTPError)
		assert.True(t, ok)
		assert.Equal(t, 409, echoErr.Code)
	})

	t.Run("employee not found", func(t *testing.T) {
		employeeID := "1001"
		payload := entity.Employee{
			EmployeeID: "1002",
			FullName:   "Firman Darmawan",
			Address:    "Cipinang Melayu, Jakarta Timur",
		}

		repo.EXPECT().GetDetail(employeeID).Return(nil, gorm.ErrRecordNotFound).Times(1)

		err := svc.Update(employeeID, payload)
		assert.NotNil(t, err)
		echoErr, ok := err.(*echo.HTTPError)
		assert.True(t, ok)
		assert.Equal(t, 404, echoErr.Code)
	})

	t.Run("internal error", func(t *testing.T) {
		employeeID := "1001"
		employee := &entity.Employee{
			EmployeeID: "1001",
			FullName:   "Firman",
			Address:    "Jakarta Timur",
		}
		payload := entity.Employee{
			EmployeeID: "1002",
			FullName:   "Firman Darmawan",
			Address:    "Cipinang Melayu, Jakarta Timur",
		}

		repo.EXPECT().GetDetail(employeeID).Return(employee, nil).Times(1)
		repo.EXPECT().Update(payload, employee).Return(errors.New("error")).Times(1)

		err := svc.Update(employeeID, payload)
		assert.NotNil(t, err)
		echoErr, ok := err.(*echo.HTTPError)
		assert.True(t, ok)
		assert.Equal(t, 500, echoErr.Code)
	})
}

func Test_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	svc, repo := initUsecase(ctrl)

	t.Run("success", func(t *testing.T) {
		employeeID := "1001"
		employee := &entity.Employee{
			EmployeeID: "1001",
			FullName:   "Firman",
			Address:    "Jakarta Timur",
		}

		repo.EXPECT().GetDetail(employeeID).Return(employee, nil).Times(1)
		repo.EXPECT().Delete(employee).Return(nil).Times(1)

		err := svc.Delete(employeeID)
		assert.Nil(t, err)
	})

	t.Run("not found employee", func(t *testing.T) {
		employeeID := "1001"

		repo.EXPECT().GetDetail(employeeID).Return(nil, gorm.ErrRecordNotFound).Times(1)

		err := svc.Delete(employeeID)
		assert.Nil(t, err)
	})

	t.Run("error find employee", func(t *testing.T) {
		employeeID := "1001"

		repo.EXPECT().GetDetail(employeeID).Return(nil, errors.New("error")).Times(1)

		err := svc.Delete(employeeID)
		assert.NotNil(t, err)
		echoErr, ok := err.(*echo.HTTPError)
		assert.True(t, ok)
		assert.Equal(t, 500, echoErr.Code)
	})
}

func Test_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	svc, repo := initUsecase(ctrl)

	t.Run("success", func(t *testing.T) {
		req := &entity.QueryRequest{}
		res := &entity.EmployeeListResponse{
			Employees: []*entity.Employee{{
				EmployeeID: "1001",
				FullName:   "Firman",
				Address:    "Jakarta Timur",
			}},
		}

		repo.EXPECT().GetList(req).Return(res, nil).Times(1)

		resp, err := svc.Get(req)
		assert.Nil(t, err)
		assert.Equal(t, res, resp)
	})
}
