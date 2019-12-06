package day3

//Siswa ...
type Siswa struct {
	Name       string
	NameEncode string
	score      int
}

// Chiper ...
type Chiper interface {
	Encode() string
	Decode() string
}

//Encode ...
func (s *Siswa) Encode() {
	s.NameEncode = ""
	for index := 0; index < len(s.Name); index++ {
		ascii := 97 + ((s.Name[index] - 97 + 13) % 26)
		s.NameEncode += string(ascii)
	}
}

//Decode ...
func (s *Siswa) Decode() {
	s.Name = ""
	for index := 0; index < len(s.NameEncode); index++ {
		ascii := 97 + ((s.NameEncode[index] - 97 + 13) % 26)
		s.Name += string(ascii)
	}
}

//SetName ...
// func (s *Siswa) SetName(Name string) {
// 	s.Name = Name
// }

//SetNameEncode ...
// func (s *Siswa) SetNameEncode(NameEncode string) {
// 	s.NameEncode = NameEncode
// }
