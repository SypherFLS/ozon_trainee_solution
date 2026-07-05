package postgres

type Link struct {
	ID       uint   `gorm:"primaryKey"`
	Original string `gorm:"size:2048;not null;uniqueIndex"`
	Short    string `gorm:"size:10;not null;uniqueIndex"`
}
