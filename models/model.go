package models

type UniqueModel interface {
	SetUniqueId()
}

type PartialUpdater interface {
	CreatePartialUpdateMap()
}
