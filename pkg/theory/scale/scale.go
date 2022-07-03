package scale

import "github.com/edipermadi/music-db/pkg/theory/pitch"

// Type is a type for scale
type Type int

// Scale name enumerations
const (
	Invalid        Type = iota
	Minoric        Type = iota
	Thaptic        Type = iota
	Lothic         Type = iota
	Phratic        Type = iota
	Aerathic       Type = iota
	Epathic        Type = iota
	Mynic          Type = iota
	Rothic         Type = iota
	Eporic         Type = iota
	Zyphic         Type = iota
	Epogic         Type = iota
	Lanic          Type = iota
	Pyrric         Type = iota
	Aeoloric       Type = iota
	Gonic          Type = iota
	Dalic          Type = iota
	Dygic          Type = iota
	Daric          Type = iota
	Lonic          Type = iota
	Phradic        Type = iota
	Bolic          Type = iota
	Saric          Type = iota
	Zoptic         Type = iota
	Aeraphic       Type = iota
	Byptic         Type = iota
	Aeolic         Type = iota
	Koptic         Type = iota
	Mixolyric      Type = iota
	Lydic          Type = iota
	Stathic        Type = iota
	Dadic          Type = iota
	Phrynic        Type = iota
	Epathitonic    Type = iota
	Mynitonic      Type = iota
	Rocritonic     Type = iota
	Pentatonic     Type = iota
	Thaptitonic    Type = iota
	Magitonic      Type = iota
	Daditonic      Type = iota
	Aeolyphritonic Type = iota
	Gycritonic     Type = iota
	Pyritonic      Type = iota
	Gathitonic     Type = iota
	Ionitonic      Type = iota
	Phrynitonic    Type = iota
	Stathitonic    Type = iota
	Thalitonic     Type = iota
	Zolitonic      Type = iota
	Epogitonic     Type = iota
	Lanitonic      Type = iota
	Paptitonic     Type = iota
	Ionacritonic   Type = iota
	Phraditonic    Type = iota
	Aeoloritonic   Type = iota
	Gonitonic      Type = iota
	Dalitonic      Type = iota
	Dygitonic      Type = iota
	Aeracritonic   Type = iota
	Byptitonic     Type = iota
	Daritonic      Type = iota
	Lonitonic      Type = iota
	Ionycritonic   Type = iota
	Lothitonic     Type = iota
	Phratonic      Type = iota
	Aerathitonic   Type = iota
	Saritonic      Type = iota
	Zoptitonic     Type = iota
	Dolitonic      Type = iota
	Poritonic      Type = iota
	Aerylitonic    Type = iota
	Zagitonic      Type = iota
	Lagitonic      Type = iota
	Molitonic      Type = iota
	Staptitonic    Type = iota
	Mothitonic     Type = iota
	Aeritonic      Type = iota
	Ragitonic      Type = iota
	Ionaditonic    Type = iota
	Bocritonic     Type = iota
	Gythitonic     Type = iota
	Pagitonic      Type = iota
	Aeolythitonic  Type = iota
	Zacritonic     Type = iota
	Laritonic      Type = iota
	Thacritonic    Type = iota
	Styditonic     Type = iota
	Loritonic      Type = iota
	Aeolyritonic   Type = iota
	Goritonic      Type = iota
	Aeoloditonic   Type = iota
	Doptitonic     Type = iota
	Aeraphitonic   Type = iota
	Zathitonic     Type = iota
	Raditonic      Type = iota
	Stonitonic     Type = iota
	Syptitonic     Type = iota
	Ionythitonic   Type = iota
	Aeolanitonic   Type = iota
	Danitonic      Type = iota
	Ionaritonic    Type = iota
	Dynitonic      Type = iota
	Zyditonic      Type = iota
	Aeolacritonic  Type = iota
	Zythitonic     Type = iota
	Dyritonic      Type = iota
	Koptitonic     Type = iota
	Thocritonic    Type = iota
	Lycritonic     Type = iota
	Daptitonic     Type = iota
	Kygitonic      Type = iota
	Mocritonic     Type = iota
	Zynitonic      Type = iota
	Epygitonic     Type = iota
	Zaptitonic     Type = iota
	Kagitonic      Type = iota
	Zogitonic      Type = iota
	Epyritonic     Type = iota
	Zothitonic     Type = iota
	Phrolitonic    Type = iota
	Ionagitonic    Type = iota
	Aeolapritonic  Type = iota
	Kyritonic      Type = iota
	Ionyptitonic   Type = iota
	Gyritonic      Type = iota
	Zalitonic      Type = iota
	Stolitonic     Type = iota
	Bylitonic      Type = iota
	Thoditonic     Type = iota
	Dogitonic      Type = iota
	Phralitonic    Type = iota
	Garitonic      Type = iota
	Soptitonic     Type = iota
	Kataritonic    Type = iota
	Sylitonic      Type = iota
	Thonitonic     Type = iota
	Phropitonic    Type = iota
	Staditonic     Type = iota
	Lyditonic      Type = iota
	Mythitonic     Type = iota
	Sogitonic      Type = iota
	Gothitonic     Type = iota
	Rothitonic     Type = iota
	Zylitonic      Type = iota
	Zoditonic      Type = iota
	Zaritonic      Type = iota
	Phrythitonic   Type = iota
	Rolitonic      Type = iota
	Ranitonic      Type = iota
	Laditonic      Type = iota
	Poditonic      Type = iota
	Ionothitonic   Type = iota
	Kanitonic      Type = iota
	Ryphitonic     Type = iota
	Gylitonic      Type = iota
	Aeolycritonic  Type = iota
	Pynitonic      Type = iota
	Zanitonic      Type = iota
	Phronitonic    Type = iota
	Banitonic      Type = iota
	Aeronitonic    Type = iota
	Golitonic      Type = iota
	Dyptitonic     Type = iota
	Aerynitonic    Type = iota
	Palitonic      Type = iota
	Stothitonic    Type = iota
	Aerophitonic   Type = iota
	Katagitonic    Type = iota
	Ionoditonic    Type = iota
	Bogitonic      Type = iota
	Mogitonic      Type = iota
	Docritonic     Type = iota
	Epaditonic     Type = iota
	Mixitonic      Type = iota
	Phrothitonic   Type = iota
	Katycritonic   Type = iota
	Ionalitonic    Type = iota
	Loptitonic     Type = iota
	Thyritonic     Type = iota
	Thoptitonic    Type = iota
	Bycritonic     Type = iota
	Pathitonic     Type = iota
	Myditonic      Type = iota
	Bolitonic      Type = iota
	Bothitonic     Type = iota
	Kataditonic    Type = iota
	Koditonic      Type = iota
	Tholitonic     Type = iota
	Epathimic      Type = iota
	Mynimic        Type = iota
	Rocrimic       Type = iota
	Eporimic       Type = iota
	Thaptimic      Type = iota
	Lothimic       Type = iota
	Dyrimic        Type = iota
	Koptimic       Type = iota
	Thocrimic      Type = iota
	Aeolanimic     Type = iota
	Danimic        Type = iota
	Ionarimic      Type = iota
	Daptimic       Type = iota
	Kygimic        Type = iota
	Mocrimic       Type = iota
	Zynimic        Type = iota
	Aeolimic       Type = iota
	Zythimic       Type = iota
	Epygimic       Type = iota
	Zaptimic       Type = iota
	Kagimic        Type = iota
	Zogimic        Type = iota
	Epyrimic       Type = iota
	Lycrimic       Type = iota
	Bylimic        Type = iota
	Zothimic       Type = iota
	Phrolimic      Type = iota
	Ionagimic      Type = iota
	Aeolaphimic    Type = iota
	Kycrimic       Type = iota
	Garimic        Type = iota
	Soptimic       Type = iota
	Ionyptimic     Type = iota
	Gyrimic        Type = iota
	Zalimic        Type = iota
	Stolimic       Type = iota
	Thonimic       Type = iota
	Stadimic       Type = iota
	Thodimic       Type = iota
	Mythimic       Type = iota
	Sogimic        Type = iota
	Gogimic        Type = iota
	Rothimic       Type = iota
	Katarimic      Type = iota
	Sylimic        Type = iota
	Mixolimic      Type = iota
	Dadimic        Type = iota
	Aeolyphimic    Type = iota
	Gycrimic       Type = iota
	Pyrimic        Type = iota
	Lydimic        Type = iota
	Ionacrimic     Type = iota
	Gathimic       Type = iota
	Ionynimic      Type = iota
	Phrynimic      Type = iota
	Stathimic      Type = iota
	Thatimic       Type = iota
	Dalimic        Type = iota
	Dygimic        Type = iota
	Zolimic        Type = iota
	Epogimic       Type = iota
	Lanimic        Type = iota
	Paptimic       Type = iota
	Darmic         Type = iota
	Lonimic        Type = iota
	Ionycrimic     Type = iota
	Phradimic      Type = iota
	Aeolorimic     Type = iota
	Gonimic        Type = iota
	Phracrimic     Type = iota
	Aerathimic     Type = iota
	Sarimic        Type = iota
	Zoptimic       Type = iota
	Zeracrimic     Type = iota
	Byptimic       Type = iota
	Starimic       Type = iota
	Phrathimic     Type = iota
	Saptimic       Type = iota
	Aerodimic      Type = iota
	Macrimic       Type = iota
	Rogimic        Type = iota
	Bygimic        Type = iota
	Thycrimic      Type = iota
	Aeoladimic     Type = iota
	Dylimic        Type = iota
	Eponimic       Type = iota
	Katygimic      Type = iota
	Stalimic       Type = iota
	Stoptimic      Type = iota
	Zygimic        Type = iota
	Kataptimic     Type = iota
	Aeolaptimic    Type = iota
	Pothimic       Type = iota
	Rycrimic       Type = iota
	Ronimic        Type = iota
	Stycrimic      Type = iota
	Katorimic      Type = iota
	Epythimic      Type = iota
	Kaptimic       Type = iota
	Katythimic     Type = iota
	Madimic        Type = iota
	Aerygimic      Type = iota
	Pylimic        Type = iota
	Ionathimic     Type = iota
	Morimic        Type = iota
	Aerycrimic     Type = iota
	Ganimic        Type = iota
	Eparimic       Type = iota
	Lyrimic        Type = iota
	Phraptimic     Type = iota
	Bacrimic       Type = iota
	Phralimic      Type = iota
	Phrogimic      Type = iota
	Rathimic       Type = iota
	Katocrimic     Type = iota
	Phryptimic     Type = iota
	Katynimic      Type = iota
	Solimic        Type = iota
	Ionolimic      Type = iota
	Ionophimic     Type = iota
	Aeologimic     Type = iota
	Zadimic        Type = iota
	Sygimic        Type = iota
	Thogimic       Type = iota
	Rythimic       Type = iota
	Donimic        Type = iota
	Aeoloptimic    Type = iota
	Panimic        Type = iota
	Lodimic        Type = iota
	Laptimic       Type = iota
	Lygimic        Type = iota
	Logimic        Type = iota
	Lalimic        Type = iota
	Sothimic       Type = iota
	Phrocrimic     Type = iota
	Modimic        Type = iota
	Barimic        Type = iota
	Poptimic       Type = iota
	Sagimic        Type = iota
	Aelothimic     Type = iota
	Socrimic       Type = iota
	Syrimic        Type = iota
	Stodimic       Type = iota
	Ionocrimic     Type = iota
	Zycrimic       Type = iota
	Ionygimic      Type = iota
	Katathimic     Type = iota
	Bolimic        Type = iota
	Bothimic       Type = iota
	Katadimic      Type = iota
	Kodimic        Type = iota
	Tholimic       Type = iota
	Ralimic        Type = iota
	Kanimic        Type = iota
	Zylimic        Type = iota
	Zodimic        Type = iota
	Zarimic        Type = iota
	Phrythimic     Type = iota
	Rorimic        Type = iota
	Pynimic        Type = iota
	Zanimic        Type = iota
	Ranimic        Type = iota
	Ladimic        Type = iota
	Podimic        Type = iota
	Ionothimic     Type = iota
	Kytrimic       Type = iota
	Golimic        Type = iota
	Dyptimic       Type = iota
	Ryrimic        Type = iota
	Gylimic        Type = iota
	Aeolycrimic    Type = iota
	Palimic        Type = iota
	Stothimic      Type = iota
	Aeronimic      Type = iota
	Katagimic      Type = iota
	Phronimic      Type = iota
	Banimic        Type = iota
	Ionodimic      Type = iota
	Bogimic        Type = iota
	Mogimic        Type = iota
	Docrimic       Type = iota
	Epadimic       Type = iota
	Aerynimic      Type = iota
	Mydimic        Type = iota
	Thyptimic      Type = iota
	Phrothimic     Type = iota
	Katycrimic     Type = iota
	Ionalimic      Type = iota
	Loptimic       Type = iota
	Zagimic        Type = iota
	Lagimic        Type = iota
	Thyrimic       Type = iota
	Thothimic      Type = iota
	Bycrimic       Type = iota
	Pathimic       Type = iota
	Mothimic       Type = iota
	Aeranimic      Type = iota
	Ragimic        Type = iota
	Dolimic        Type = iota
	Porimic        Type = iota
	Aerylimic      Type = iota
	Bocrimic       Type = iota
	Gythimic       Type = iota
	Pagimic        Type = iota
	Aeolythimic    Type = iota
	Molimic        Type = iota
	Staptimic      Type = iota
	Zacrimic       Type = iota
	Larimic        Type = iota
	Thacrimic      Type = iota
	Stydimic       Type = iota
	Lorimic        Type = iota
	Ionadimic      Type = iota
	Ionythimic     Type = iota
	Aerythimic     Type = iota
	Dynimic        Type = iota
	Zydimic        Type = iota
	Zathimic       Type = iota
	Radimic        Type = iota
	Stonimic       Type = iota
	Syptimic       Type = iota
	Ponimic        Type = iota
	Kadimic        Type = iota
	Gynimic        Type = iota
	Thydimic       Type = iota
	Polimic        Type = iota
	Thanimic       Type = iota
	Lathimic       Type = iota
	Aeralimic      Type = iota
	Kynimic        Type = iota
	Stynimic       Type = iota
	Epytimic       Type = iota
	Katoptimic     Type = iota
	Galimic        Type = iota
	Kathimic       Type = iota
	Lylimic        Type = iota
	Epalimic       Type = iota
	Epacrimic      Type = iota
	Sathimic       Type = iota
	Katanimic      Type = iota
	Katyrimic      Type = iota
	Rynimic        Type = iota
	Pogimic        Type = iota
	Aeraptimic     Type = iota
	Epylimic       Type = iota
	Manimic        Type = iota
	Marimic        Type = iota
	Locrimic       Type = iota
	Rylimic        Type = iota
	Epatimic       Type = iota
	Byrimic        Type = iota
	Kocrimic       Type = iota
	Korimic        Type = iota
	Lynimic        Type = iota
	Malimic        Type = iota
	Synimic        Type = iota
	Phragimic      Type = iota
	Mycrimic       Type = iota
	Ionorimic      Type = iota
	Phrydimic      Type = iota
	Zyptimic       Type = iota
	Katothimic     Type = iota
	Phrylimic      Type = iota
	Aerothimic     Type = iota
	Stagimic       Type = iota
	Dorimic        Type = iota
	Phrycrimic     Type = iota
	Kyptimic       Type = iota
	Ionylimic      Type = iota
	Epynimic       Type = iota
	Ionogimic      Type = iota
	Kydimic        Type = iota
	Gaptimic       Type = iota
	Tharimic       Type = iota
	Ionaphimic     Type = iota
	Thoptimic      Type = iota
	Bagimic        Type = iota
	Kyrimic        Type = iota
	Sonimic        Type = iota
	Aeolonimic     Type = iota
	Rygimic        Type = iota
	Thagimic       Type = iota
	Kolimic        Type = iota
	Dycrimic       Type = iota
	Epycrimic      Type = iota
	Gocrimic       Type = iota
	Katolimic      Type = iota
	Dagimic        Type = iota
	Aeolydimic     Type = iota
	Parimic        Type = iota
	Ionaptimic     Type = iota
	Thylimic       Type = iota
	Lolimic        Type = iota
	Thalimic       Type = iota
	Stygimic       Type = iota
	Aeolygimic     Type = iota
	Aerogimic      Type = iota
	Dacrimic       Type = iota
	Baptimic       Type = iota
	Stythimic      Type = iota
	Kothimic       Type = iota
	Pygimic        Type = iota
	Rodimic        Type = iota
	Sorimic        Type = iota
	Monimic        Type = iota
	Aeragimic      Type = iota
	Epothimic      Type = iota
	Salimic        Type = iota
	Lyptimic       Type = iota
	Katonimic      Type = iota
	Gygimic        Type = iota
	Aeradimic      Type = iota
	Zyrimic        Type = iota
	Stylimic       Type = iota
	Lythimic       Type = iota
	Dodimic        Type = iota
	Katalimic      Type = iota
	Boptimic       Type = iota
	Stogimic       Type = iota
	Thynimic       Type = iota
	Aeolathimic    Type = iota
	Bythimic       Type = iota
	Padimic        Type = iota
	Dathimic       Type = iota
	Epagimic       Type = iota
	Raptimic       Type = iota
	Epolimic       Type = iota
	Sythimic       Type = iota
	Sydimic        Type = iota
	Gacrimic       Type = iota
	Borimic        Type = iota
	Sycrimic       Type = iota
	Gadimic        Type = iota
	Aeolocrimic    Type = iota
	Phrygimic      Type = iota
	WholeTone      Type = iota
)

// AllScales return all scales
func AllScales() []Type {
	return []Type{
		// 3 notes
		Minoric,

		// 4 notes
		Thaptic,
		Lothic,
		Phratic,
		Aerathic,
		Epathic,
		Mynic,
		Rothic,
		Eporic,
		Zyphic,
		Epogic,
		Lanic,
		Pyrric,
		Aeoloric,
		Gonic,
		Dalic,
		Dygic,
		Daric,
		Lonic,
		Phradic,
		Bolic,
		Saric,
		Zoptic,
		Aeraphic,
		Byptic,
		Aeolic,
		Koptic,
		Mixolyric,
		Lydic,
		Stathic,
		Dadic,
		Phrynic,

		// 5 notes
		Epathitonic,
		Mynitonic,
		Rocritonic,
		Pentatonic,
		Thaptitonic,
		Magitonic,
		Daditonic,
		Aeolyphritonic,
		Gycritonic,
		Pyritonic,
		Gathitonic,
		Ionitonic,
		Phrynitonic,
		Stathitonic,
		Thalitonic,
		Zolitonic,
		Epogitonic,
		Lanitonic,
		Paptitonic,
		Ionacritonic,
		Phraditonic,
		Aeoloritonic,
		Gonitonic,
		Dalitonic,
		Dygitonic,
		Aeracritonic,
		Byptitonic,
		Daritonic,
		Lonitonic,
		Ionycritonic,
		Lothitonic,
		Phratonic,
		Aerathitonic,
		Saritonic,
		Zoptitonic,
		Dolitonic,
		Poritonic,
		Aerylitonic,
		Zagitonic,
		Lagitonic,
		Molitonic,
		Staptitonic,
		Mothitonic,
		Aeritonic,
		Ragitonic,
		Ionaditonic,
		Bocritonic,
		Gythitonic,
		Pagitonic,
		Aeolythitonic,
		Zacritonic,
		Laritonic,
		Thacritonic,
		Styditonic,
		Loritonic,
		Aeolyritonic,
		Goritonic,
		Aeoloditonic,
		Doptitonic,
		Aeraphitonic,
		Zathitonic,
		Raditonic,
		Stonitonic,
		Syptitonic,
		Ionythitonic,
		Aeolanitonic,
		Danitonic,
		Ionaritonic,
		Dynitonic,
		Zyditonic,
		Aeolacritonic,
		Zythitonic,
		Dyritonic,
		Koptitonic,
		Thocritonic,
		Lycritonic,
		Daptitonic,
		Kygitonic,
		Mocritonic,
		Zynitonic,
		Epygitonic,
		Zaptitonic,
		Kagitonic,
		Zogitonic,
		Epyritonic,
		Zothitonic,
		Phrolitonic,
		Ionagitonic,
		Aeolapritonic,
		Kyritonic,
		Ionyptitonic,
		Gyritonic,
		Zalitonic,
		Stolitonic,
		Bylitonic,
		Thoditonic,
		Dogitonic,
		Phralitonic,
		Garitonic,
		Soptitonic,
		Kataritonic,
		Sylitonic,
		Thonitonic,
		Phropitonic,
		Staditonic,
		Lyditonic,
		Mythitonic,
		Sogitonic,
		Gothitonic,
		Rothitonic,
		Zylitonic,
		Zoditonic,
		Zaritonic,
		Phrythitonic,
		Rolitonic,
		Ranitonic,
		Laditonic,
		Poditonic,
		Ionothitonic,
		Kanitonic,
		Ryphitonic,
		Gylitonic,
		Aeolycritonic,
		Pynitonic,
		Zanitonic,
		Phronitonic,
		Banitonic,
		Aeronitonic,
		Golitonic,
		Dyptitonic,
		Aerynitonic,
		Palitonic,
		Stothitonic,
		Aerophitonic,
		Katagitonic,
		Ionoditonic,
		Bogitonic,
		Mogitonic,
		Docritonic,
		Epaditonic,
		Mixitonic,
		Phrothitonic,
		Katycritonic,
		Ionalitonic,
		Loptitonic,
		Thyritonic,
		Thoptitonic,
		Bycritonic,
		Pathitonic,
		Myditonic,
		Bolitonic,
		Bothitonic,
		Kataditonic,
		Koditonic,
		Tholitonic,

		// 6 notes
		Epathimic,
		Mynimic,
		Rocrimic,
		Eporimic,
		Thaptimic,
		Lothimic,
		Dyrimic,
		Koptimic,
		Thocrimic,
		Aeolanimic,
		Danimic,
		Ionarimic,
		Daptimic,
		Kygimic,
		Mocrimic,
		Zynimic,
		Aeolimic,
		Zythimic,
		Epygimic,
		Zaptimic,
		Kagimic,
		Zogimic,
		Epyrimic,
		Lycrimic,
		Bylimic,
		Zothimic,
		Phrolimic,
		Ionagimic,
		Aeolaphimic,
		Kycrimic,
		Garimic,
		Soptimic,
		Ionyptimic,
		Gyrimic,
		Zalimic,
		Stolimic,
		Thonimic,
		Stadimic,
		Thodimic,
		Mythimic,
		Sogimic,
		Gogimic,
		Rothimic,
		Katarimic,
		Sylimic,
		Mixolimic,
		Dadimic,
		Aeolyphimic,
		Gycrimic,
		Pyrimic,
		Lydimic,
		Ionacrimic,
		Gathimic,
		Ionynimic,
		Phrynimic,
		Stathimic,
		Thatimic,
		Dalimic,
		Dygimic,
		Zolimic,
		Epogimic,
		Lanimic,
		Paptimic,
		Darmic,
		Lonimic,
		Ionycrimic,
		Phradimic,
		Aeolorimic,
		Gonimic,
		Phracrimic,
		Aerathimic,
		Sarimic,
		Zoptimic,
		Zeracrimic,
		Byptimic,
		Starimic,
		Phrathimic,
		Saptimic,
		Aerodimic,
		Macrimic,
		Rogimic,
		Bygimic,
		Thycrimic,
		Aeoladimic,
		Dylimic,
		Eponimic,
		Katygimic,
		Stalimic,
		Stoptimic,
		Zygimic,
		Kataptimic,
		Aeolaptimic,
		Pothimic,
		Rycrimic,
		Ronimic,
		Stycrimic,
		Katorimic,
		Epythimic,
		Kaptimic,
		Katythimic,
		Madimic,
		Aerygimic,
		Pylimic,
		Ionathimic,
		Morimic,
		Aerycrimic,
		Ganimic,
		Eparimic,
		Lyrimic,
		Phraptimic,
		Bacrimic,
		Phralimic,
		Phrogimic,
		Rathimic,
		Katocrimic,
		Phryptimic,
		Katynimic,
		Solimic,
		Ionolimic,
		Ionophimic,
		Aeologimic,
		Zadimic,
		Sygimic,
		Thogimic,
		Rythimic,
		Donimic,
		Aeoloptimic,
		Panimic,
		Lodimic,
		Laptimic,
		Lygimic,
		Logimic,
		Lalimic,
		Sothimic,
		Phrocrimic,
		Modimic,
		Barimic,
		Poptimic,
		Sagimic,
		Aelothimic,
		Socrimic,
		Syrimic,
		Stodimic,
		Ionocrimic,
		Zycrimic,
		Ionygimic,
		Katathimic,
		Bolimic,
		Bothimic,
		Katadimic,
		Kodimic,
		Tholimic,
		Ralimic,
		Kanimic,
		Zylimic,
		Zodimic,
		Zarimic,
		Phrythimic,
		Rorimic,
		Pynimic,
		Zanimic,
		Ranimic,
		Ladimic,
		Podimic,
		Ionothimic,
		Kytrimic,
		Golimic,
		Dyptimic,
		Ryrimic,
		Gylimic,
		Aeolycrimic,
		Palimic,
		Stothimic,
		Aeronimic,
		Katagimic,
		Phronimic,
		Banimic,
		Ionodimic,
		Bogimic,
		Mogimic,
		Docrimic,
		Epadimic,
		Aerynimic,
		Mydimic,
		Thyptimic,
		Phrothimic,
		Katycrimic,
		Ionalimic,
		Loptimic,
		Zagimic,
		Lagimic,
		Thyrimic,
		Thothimic,
		Bycrimic,
		Pathimic,
		Mothimic,
		Aeranimic,
		Ragimic,
		Dolimic,
		Porimic,
		Aerylimic,
		Bocrimic,
		Gythimic,
		Pagimic,
		Aeolythimic,
		Molimic,
		Staptimic,
		Zacrimic,
		Larimic,
		Thacrimic,
		Stydimic,
		Lorimic,
		Ionadimic,
		Ionythimic,
		Aerythimic,
		Dynimic,
		Zydimic,
		Zathimic,
		Radimic,
		Stonimic,
		Syptimic,
		Ponimic,
		Kadimic,
		Gynimic,
		Thydimic,
		Polimic,
		Thanimic,
		Lathimic,
		Aeralimic,
		Kynimic,
		Stynimic,
		Epytimic,
		Katoptimic,
		Galimic,
		Kathimic,
		Lylimic,
		Epalimic,
		Epacrimic,
		Sathimic,
		Katanimic,
		Katyrimic,
		Rynimic,
		Pogimic,
		Aeraptimic,
		Epylimic,
		Manimic,
		Marimic,
		Locrimic,
		Rylimic,
		Epatimic,
		Byrimic,
		Kocrimic,
		Korimic,
		Lynimic,
		Malimic,
		Synimic,
		Phragimic,
		Mycrimic,
		Ionorimic,
		Phrydimic,
		Zyptimic,
		Katothimic,
		Phrylimic,
		Aerothimic,
		Stagimic,
		Dorimic,
		Phrycrimic,
		Kyptimic,
		Ionylimic,
		Epynimic,
		Ionogimic,
		Kydimic,
		Gaptimic,
		Tharimic,
		Ionaphimic,
		Thoptimic,
		Bagimic,
		Kyrimic,
		Sonimic,
		Aeolonimic,
		Rygimic,
		Thagimic,
		Kolimic,
		Dycrimic,
		Epycrimic,
		Gocrimic,
		Katolimic,
		Dagimic,
		Aeolydimic,
		Parimic,
		Ionaptimic,
		Thylimic,
		Lolimic,
		Thalimic,
		Stygimic,
		Aeolygimic,
		Aerogimic,
		Dacrimic,
		Baptimic,
		Stythimic,
		Kothimic,
		Pygimic,
		Rodimic,
		Sorimic,
		Monimic,
		Aeragimic,
		Epothimic,
		Salimic,
		Lyptimic,
		Katonimic,
		Gygimic,
		Aeradimic,
		Zyrimic,
		Stylimic,
		Lythimic,
		Dodimic,
		Katalimic,
		Boptimic,
		Stogimic,
		Thynimic,
		Aeolathimic,
		Bythimic,
		Padimic,
		Dathimic,
		Epagimic,
		Raptimic,
		Epolimic,
		Sythimic,
		Sydimic,
		Gacrimic,
		Borimic,
		Sycrimic,
		Gadimic,
		Aeolocrimic,
		Phrygimic,
		WholeTone,
	}
}

// Pitches returns set of pitches of the scale with given tonic
func (s Type) Pitches(tonic pitch.Type) []pitch.Type {
	amount := int(tonic - pitch.CNatural)
	pitches := make([]pitch.Type, 0)
	for _, v := range pitch.AllPitches() {
		if s.Number()&v.Number() == v.Number() {
			pitches = append(pitches, v.Transpose(amount))
		}
	}

	return pitches
}

// Perfection stores perfection profile
type Perfection struct {
	Perfection   int
	Imperfection int
}

func (s Type) Perfection() Perfection {
	pitchMap := make(map[pitch.Type]struct{})
	pitches := s.Pitches(pitch.CNatural)

	for _, v := range pitches {
		pitchMap[v] = struct{}{}
	}

	var result Perfection
	for _, v := range pitches {
		_, found := pitchMap[v.NextFifth()]
		if found {
			result.Perfection++
		} else {
			result.Imperfection++
		}
	}

	return result
}
