package table

import "github.com/lvxiaorun/logger"

type Room struct {
	Id        int64
	CreatedAt int
	UpdatedAt int
	Describe  string
}

func (t *Room) TableName() string {
	return "room"
}

func (t *Room) List() ([]*Room, error) {
	var rooms []*Room
	err := DB.Model(t).Scan(&rooms).Error
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	return rooms, nil
}

func (t *Room) Show() (*Room, error) {
	var room Room
	err := DB.First(&room, t.Id).Error
	if err != nil {
		return nil, err
	}
	return &room, nil
}
