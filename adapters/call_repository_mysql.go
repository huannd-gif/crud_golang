package adapters

import (
	"api_crud/app/query"
	"api_crud/domain"
	"context"
	"encoding/json"
	"errors"
	"gorm.io/gorm"
	"time"
)

type Call struct {
	Id          int        `gorm:"primaryKey,autoincrement;index;not null"`
	PhoneNumber string     `gorm:"type:varchar(20);not null"`
	Result      string     `gorm:"type:varchar(20);not null"`
	CreateAt    *time.Time `gorm:"autoCreateTime;milli;not null"`
	UpdateAt    *time.Time `gorm:"autoUpdateTime;milli"`
	CallAt      *time.Time `gorm:"autoCreateTime;milli;not null"`
	EndAt       *time.Time `gorm:"milli"`
	CallPress   *time.Time `gorm:"milli"`
	ReceiverAt  *time.Time `gorm:"milli"`
	Metadata    string     `gorm:"type:json;null"`
}

type CallMySQLRepository struct {
	db *gorm.DB
}

func (cdb CallMySQLRepository) callModelToQuery(c *Call) *query.Call {
	return &query.Call{
		Id:          c.Id,
		PhoneNumber: c.PhoneNumber,
		Result:      c.Result,
		CreateAt:    c.CreateAt,
		UpdateAt:    c.UpdateAt,
		CallAt:      c.CallAt,
		EndAt:       c.EndAt,
		CallPress:   c.CallPress,
		ReceiverAt:  c.ReceiverAt,
		Metadata:    c.Metadata,
	}
}

func (cdb CallMySQLRepository) callDomainToModel(c *domain.Call) *Call {
	metadata, err := json.Marshal(c.GetMetadata())
	if err != nil {
		metadata = nil
	}

	return &Call{
		PhoneNumber: c.GetPhoneNumber(),
		Result:      c.GetResult(),
		CallAt:      c.GetCallAt(),
		EndAt:       c.GetEndAt(),
		CallPress:   c.GetCallPress(),
		ReceiverAt:  c.GetReceiverAt(),
		Metadata:    string(metadata),
	}
}

func (cdb CallMySQLRepository) GetCalls(ctx context.Context, re query.CallRequest) (query.ListCallsPaginated, error) {
	callModels := make([]*Call, 0, 10)

	baseQuery := cdb.db.Model(Call{}).Session(&gorm.Session{})
	if re.PhoneNumber != "" {
		baseQuery = baseQuery.Where("phone_number LIKE ?", "%"+re.PhoneNumber+"%")
	}

	if re.MetadataDisplayField != "" {
		baseQuery = baseQuery.Select("id, phone_number, result, create_at, update_at, call_at, end_at, call_press, receiver_at, CONCAT('{" + re.MetadataDisplayField + ": ', JSON_UNQUOTE(JSON_EXTRACT(metadata, '$." + re.MetadataDisplayField + "')), '}') AS metadata")

		//baseQuery = baseQuery.Select("id, phone_number, result, create_at, update_at, call_at, end_at, call_press, receiver_at,JSON_EXTRACT(metadata, '$." + re.MetadataDisplayField + "') AS metadata")
	}

	total := int64(0)
	err := baseQuery.Count(&total).Error
	if err != nil {
		return query.ListCallsPaginated{}, errors.New(err.Error())
	}
	paging := calculatePaging(re.PageNum, re.PageSize, int(total))

	err = baseQuery.Order("create_at desc").Offset((re.PageNum - 1) * re.PageSize).Limit(re.PageSize).Find(&callModels).Error

	if err != nil {
		return query.ListCallsPaginated{}, errors.New(err.Error())
	}

	res := make([]*query.Call, 0, 10)
	for _, callModel := range callModels {
		res = append(res, cdb.callModelToQuery(callModel))
	}

	return query.ListCallsPaginated{
		Calls:    res,
		Metadata: paging,
	}, nil
}

func (cdb CallMySQLRepository) AddCall(ctx context.Context, c *domain.Call) error {
	callSaveDB := cdb.callDomainToModel(c)
	tx := cdb.db.Begin()
	err := tx.Model(Call{}).Create(&callSaveDB).Error
	c.SetId(callSaveDB.Id)
	c.SetCreateAt(callSaveDB.CreateAt)
	c.SetUpdateAt(callSaveDB.UpdateAt)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (cdb CallMySQLRepository) GetCallById(ctx context.Context, id int) (*Call, error) {
	call := Call{}
	err := cdb.db.Model(Call{}).Where("id = ?", id).First(&call).Error
	if err != nil {
		return nil, err
	}
	return &call, nil
}

func (cdb CallMySQLRepository) UpdateCall(ctx context.Context, c *domain.Call) error {
	callSaveDB := cdb.callDomainToModel(c)
	tx := cdb.db.Begin()
	err := tx.Model(Call{}).Where("id = ?", c.GetId()).Updates(callSaveDB).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (cdb CallMySQLRepository) DeleteCall(ctx context.Context, id *int) error {
	tx := cdb.db.Begin()
	err := tx.Model(Call{}).Where("id = ?", id).Delete(nil).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func NewCallMysqlRepository(db *gorm.DB, migrator *GORMMigrator) *CallMySQLRepository {
	if migrator != nil {
		migrator.addAutoMigrate(Call{})
	}
	return &CallMySQLRepository{
		db: db,
	}
}
