package mock

import (
	"fmt"
	"testing"

	"test-privy/helper"
	"test-privy/helper/constant"
	cake "test-privy/module/admin/cake"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestListCakes(t *testing.T) {
	cakeRow := make([]cake.Cake, 10)
	for i := len(cakeRow) - 1; i >= 0; i-- {
		date, rate := helper.GetDateNow(), float32(i)+0.1
		r := cake.Cake{
			Title:       fmt.Sprintf("Cake %d", 1),
			Description: fmt.Sprintf("Ini adalah cake %d", i),
			Image:       fmt.Sprintf("https://image/%d", i),
			Rating:      &rate,
			CreatedAt:   date,
			UpdatedAt:   &date,
		}
		cakeRow[i] = r
	}
	expectedRes := &constant.DataPaging{
		Total:      10,
		PageActive: 1,
		Showing:    "1 to 10 of 10",
		PageList:   []int{1},
		Rows:       cakeRow,
		Count:      10,
	}
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	req := constant.Params{
		Page:      1,
		MaxRows:   10,
		OrderBy:   "rating",
		OrderType: "desc",
		Search:    []constant.Search{},
	}
	t.Run("success with data", func(t *testing.T) {
		mockModel := NewMockCakeModel(ctl)
		gomock.InOrder(
			mockModel.EXPECT().CakeList(req).Return(expectedRes, nil),
		)

		model := new(cake.ListCakes)
		res, err := model.CakeList(&req)
		assert.Equal(t, expectedRes, res)
		assert.Nil(t, err)

	})
}
