package models

type Petugas struct {
	Id         int          `gorm:"primaryKey;autoIncrement;" json:"id"`
	Nama       string       `json:"nama"`
	Hp         string       `json:"hp"`
	Peminjaman []Peminjaman `gorm:"foreignKey:Id_Petugas;" json:"peminjaman"`
}

type Pengunjung struct {
	Id         int          `gorm:"primaryKaey;autoIncrement;" json:"id"`
	Nama       string       `json:"nama"`
	Hp         string       `json:"hp"`
	Alamat     string       `json:"alamat"`
	Peminjaman []Peminjaman `gorm:"foreignKey:Id_Pengunjung;" json:"peminjaman"`
}

type Buku struct {
	Id         int          `gorm:"primaryKaey;autoIncrement;" json:"id"`
	Judul      string       `json:"judul"`
	Pengarang  string       `json:"pengarang"`
	Penerbit   string       `json:"penerbit"`
	Jenis_Buku string       `json:"jenis_buku"`
	Peminjaman []Peminjaman `gorm:"foreignKey:Id_Buku;" json:"peminjaman"`
}

type Peminjaman struct {
	Id            int    `gorm:"primaryKaey;autoIncrement;" json:"id"`
	Id_Petugas    int    `json:"id_petugas"`
	Id_Pengunjung int    `json:"id_pengunjung"`
	Id_Buku       int    `json:"id_buku"`
	Tgl_Pinjam    string `json:"tgl_pinjam"`
	Tgl_Kembali   string `json:"tgl_kembali"`
}
