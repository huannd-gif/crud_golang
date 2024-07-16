package adapters

import (
	"api_crud/app/query"
	"math"
)

func calculatePaging(pageNum, pageSize, total int) query.Paging {
	return query.Paging{
		TotalCount: int(total),
		PageSize:   pageSize,
		PageNum:    pageNum,
		PageTotal:  int(math.Ceil(float64(total) / float64(pageSize))),
	}
}
