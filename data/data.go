package data

type ClassMate struct {
	no        int8
	nama      string
	alamat    string
	pekerjaan string
	alasan    string
}

var classMates = []ClassMate{
	{
		no:        1,
		nama:      "Juni Dio Kasandra",
		alamat:    "Indonesia",
		pekerjaan: "Mahasiswa",
		alasan:    "Bootcamp",
	},
	{
		no:        2,
		nama:      "Rehan",
		alamat:    "Indonesia",
		pekerjaan: "Mahasiswa",
		alasan:    "Bootcamp",
	},
	{
		no:        3,
		nama:      "Riski",
		alamat:    "Indonesia",
		pekerjaan: "Karyawan Swasta",
		alasan:    "Bootcamp",
	},
	{
		no:        4,
		nama:      "Farhan",
		alamat:    "Jerman",
		pekerjaan: "Sales Specialist",
		alasan:    "Bootcamp",
	},
	{
		no:        5,
		nama:      "Burhan",
		alamat:    "Jerman",
		pekerjaan: "Ahli waris CEO PT. Abadi Selamanya",
		alasan:    "Bootcamp",
	},
}

func (cm *ClassMate) GetNo() int8 {
	return cm.no
}

func (cm *ClassMate) GetNama() string {
	return cm.nama
}

func (cm *ClassMate) GetAlamat() string {
	return cm.alamat
}

func (cm *ClassMate) GetPekerjaan() string {
	return cm.pekerjaan
}

func (cm *ClassMate) GetAlasan() string {
	return cm.alasan
}

func GetTotalClassMate() int {
	return len(classMates)
}

func GetClassMate(noAbsen int) ClassMate {
	return classMates[noAbsen-1]
}
