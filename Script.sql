CREATE TABLE music (
	ID SERIAL PRIMARY KEY,
	Name VARCHAR(50),
	Album VARCHAR(50),
	AlbumArt VARCHAR(255),
	Singer VARCHAR(50),
	PublishDate INT,
	CreateAt DATE,
	UpdateAt DATE
);