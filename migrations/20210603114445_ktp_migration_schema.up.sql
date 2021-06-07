CREATE TABLE IF NOT EXISTS ktp(
    id bigint(20) NOT NULL AUTO_INCREMENT,
    nama longtext,
    nik varchar(15) NOT NULL UNIQUE,
    jenis_kelamin longtext,
    tanggal_lahir datetime(3),
    alamat longtext,
    agama longtext,
    created_at datetime(3),
    updated_at datetime(3),
    PRIMARY KEY (id)
)