CREATE TABLE IF NOT EXISTS music (
    music_id INTEGER NOT NULL,
    music_name text UNIQUE NOT NULL,
    file_path text UNIQUE NOT NULL,
    file_ext varchar(5) NOT NULL,
    CONSTRAINT music_pk PRIMARY KEY (music_id)
)