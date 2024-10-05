package app

import "gorm.io/gorm"

// Struct representing KeyValue object
type KeyValue struct {
	gorm.Model
	Key      string `gorm:"size:255;uniqueIndex"`
	ValueInt uint64 `gorm:"type:int"`
	ValueStr string `gorm:"type:string"`
}

func getBotCount() uint64 {
	kvi := &KeyValue{Key: "botCount"}
	err := db.FirstOrCreate(kvi, kvi).Error
	if err != nil {
		loge(err)
	}

	return kvi.ValueInt
}

func increaseBotCount() {
	kvi := &KeyValue{Key: "botCount"}
	err := db.FirstOrCreate(kvi, kvi).Error
	if err != nil {
		loge(err)
	}

	kvi.ValueInt++

	err = db.Save(kvi).Error
	if err != nil {
		loge(err)
	}
}

func initKeyValue() {
	kvi := &KeyValue{Key: "botCount"}
	err := db.FirstOrCreate(kvi, kvi).Error
	if err != nil {
		loge(err)
	}

	if kvi.ValueInt == 0 {
		kvi.ValueInt = 1
		err := db.Save(kvi).Error
		if err != nil {
			loge(err)
		}
	}
}
