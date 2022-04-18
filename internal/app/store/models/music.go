package models

type Music struct {
	ID        uint
	AlbumID   uint
	MusicType uint
	Name      string
	Info      string
	Image     string
	FilePath  string
	FileExt   string
}
