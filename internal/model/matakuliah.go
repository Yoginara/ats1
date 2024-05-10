package model

type Matakuliah struct {
	ID               string `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama_Mata_kuliah string `bson:"nama_matakuliah,omitempty" json:"nama_matakuliah,omitempty"`
	Deskripsi        string `bson:"deskripsi,omitempty" json:"deskripsi,omitempty"`
	Jumlah_SKS       string `bson:"jumlah_sks,omitempty" json:"jumlah_sks,omitempty"`
	Semester         string `bson:"semester,omitempty" json:"semester,omitempty"`
}
