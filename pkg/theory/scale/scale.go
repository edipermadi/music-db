package scale

import (
	"math"
	"slices"

	"github.com/edipermadi/music-db/pkg/theory/degree"
	"github.com/edipermadi/music-db/pkg/theory/pitch"
)

// Type is a type for scale
type Type int

// Scale name enumerations
const (
	Invalid         Type = iota
	Minoric         Type = iota
	Thaptic         Type = iota
	Lothic          Type = iota
	Phratic         Type = iota
	Aerathic        Type = iota
	Epathic         Type = iota
	Mynic           Type = iota
	Rothic          Type = iota
	Eporic          Type = iota
	Zyphic          Type = iota
	Epogic          Type = iota
	Lanic           Type = iota
	Pyrric          Type = iota
	Aeoloric        Type = iota
	Gonic           Type = iota
	Dalic           Type = iota
	Dygic           Type = iota
	Daric           Type = iota
	Lonic           Type = iota
	Phradic         Type = iota
	Bolic           Type = iota
	Saric           Type = iota
	Zoptic          Type = iota
	Aeraphic        Type = iota
	Byptic          Type = iota
	Aeolic          Type = iota
	Koptic          Type = iota
	Mixolyric       Type = iota
	Lydic           Type = iota
	Stathic         Type = iota
	Dadic           Type = iota
	Phrynic         Type = iota
	Epathitonic     Type = iota
	Mynitonic       Type = iota
	Rocritonic      Type = iota
	Pentatonic      Type = iota
	Thaptitonic     Type = iota
	Magitonic       Type = iota
	Daditonic       Type = iota
	Aeolyphritonic  Type = iota
	Gycritonic      Type = iota
	Pyritonic       Type = iota
	Gathitonic      Type = iota
	Ionitonic       Type = iota
	Phrynitonic     Type = iota
	Stathitonic     Type = iota
	Thalitonic      Type = iota
	Zolitonic       Type = iota
	Epogitonic      Type = iota
	Lanitonic       Type = iota
	Paptitonic      Type = iota
	Ionacritonic    Type = iota
	Phraditonic     Type = iota
	Aeoloritonic    Type = iota
	Gonitonic       Type = iota
	Dalitonic       Type = iota
	Dygitonic       Type = iota
	Aeracritonic    Type = iota
	Byptitonic      Type = iota
	Daritonic       Type = iota
	Lonitonic       Type = iota
	Ionycritonic    Type = iota
	Lothitonic      Type = iota
	Phratonic       Type = iota
	Aerathitonic    Type = iota
	Saritonic       Type = iota
	Zoptitonic      Type = iota
	Dolitonic       Type = iota
	Poritonic       Type = iota
	Aerylitonic     Type = iota
	Zagitonic       Type = iota
	Lagitonic       Type = iota
	Molitonic       Type = iota
	Staptitonic     Type = iota
	Mothitonic      Type = iota
	Aeritonic       Type = iota
	Ragitonic       Type = iota
	Ionaditonic     Type = iota
	Bocritonic      Type = iota
	Gythitonic      Type = iota
	Pagitonic       Type = iota
	Aeolythitonic   Type = iota
	Zacritonic      Type = iota
	Laritonic       Type = iota
	Thacritonic     Type = iota
	Styditonic      Type = iota
	Loritonic       Type = iota
	Aeolyritonic    Type = iota
	Goritonic       Type = iota
	Aeoloditonic    Type = iota
	Doptitonic      Type = iota
	Aeraphitonic    Type = iota
	Zathitonic      Type = iota
	Raditonic       Type = iota
	Stonitonic      Type = iota
	Syptitonic      Type = iota
	Ionythitonic    Type = iota
	Aeolanitonic    Type = iota
	Danitonic       Type = iota
	Ionaritonic     Type = iota
	Dynitonic       Type = iota
	Zyditonic       Type = iota
	Aeolacritonic   Type = iota
	Zythitonic      Type = iota
	Dyritonic       Type = iota
	Koptitonic      Type = iota
	Thocritonic     Type = iota
	Lycritonic      Type = iota
	Daptitonic      Type = iota
	Kygitonic       Type = iota
	Mocritonic      Type = iota
	Zynitonic       Type = iota
	Epygitonic      Type = iota
	Zaptitonic      Type = iota
	Kagitonic       Type = iota
	Zogitonic       Type = iota
	Epyritonic      Type = iota
	Zothitonic      Type = iota
	Phrolitonic     Type = iota
	Ionagitonic     Type = iota
	Aeolapritonic   Type = iota
	Kyritonic       Type = iota
	Ionyptitonic    Type = iota
	Gyritonic       Type = iota
	Zalitonic       Type = iota
	Stolitonic      Type = iota
	Bylitonic       Type = iota
	Thoditonic      Type = iota
	Dogitonic       Type = iota
	Phralitonic     Type = iota
	Garitonic       Type = iota
	Soptitonic      Type = iota
	Kataritonic     Type = iota
	Sylitonic       Type = iota
	Thonitonic      Type = iota
	Phropitonic     Type = iota
	Staditonic      Type = iota
	Lyditonic       Type = iota
	Mythitonic      Type = iota
	Sogitonic       Type = iota
	Gothitonic      Type = iota
	Rothitonic      Type = iota
	Zylitonic       Type = iota
	Zoditonic       Type = iota
	Zaritonic       Type = iota
	Phrythitonic    Type = iota
	Rolitonic       Type = iota
	Ranitonic       Type = iota
	Laditonic       Type = iota
	Poditonic       Type = iota
	Ionothitonic    Type = iota
	Kanitonic       Type = iota
	Ryphitonic      Type = iota
	Gylitonic       Type = iota
	Aeolycritonic   Type = iota
	Pynitonic       Type = iota
	Zanitonic       Type = iota
	Phronitonic     Type = iota
	Banitonic       Type = iota
	Aeronitonic     Type = iota
	Golitonic       Type = iota
	Dyptitonic      Type = iota
	Aerynitonic     Type = iota
	Palitonic       Type = iota
	Stothitonic     Type = iota
	Aerophitonic    Type = iota
	Katagitonic     Type = iota
	Ionoditonic     Type = iota
	Bogitonic       Type = iota
	Mogitonic       Type = iota
	Docritonic      Type = iota
	Epaditonic      Type = iota
	Mixitonic       Type = iota
	Phrothitonic    Type = iota
	Katycritonic    Type = iota
	Ionalitonic     Type = iota
	Loptitonic      Type = iota
	Thyritonic      Type = iota
	Thoptitonic     Type = iota
	Bycritonic      Type = iota
	Pathitonic      Type = iota
	Myditonic       Type = iota
	Bolitonic       Type = iota
	Bothitonic      Type = iota
	Kataditonic     Type = iota
	Koditonic       Type = iota
	Tholitonic      Type = iota
	Epathimic       Type = iota
	Mynimic         Type = iota
	Rocrimic        Type = iota
	Eporimic        Type = iota
	Thaptimic       Type = iota
	Lothimic        Type = iota
	Dyrimic         Type = iota
	Koptimic        Type = iota
	Thocrimic       Type = iota
	Aeolanimic      Type = iota
	Danimic         Type = iota
	Ionarimic       Type = iota
	Daptimic        Type = iota
	Kygimic         Type = iota
	Mocrimic        Type = iota
	Zynimic         Type = iota
	Aeolimic        Type = iota
	Zythimic        Type = iota
	Epygimic        Type = iota
	Zaptimic        Type = iota
	Kagimic         Type = iota
	Zogimic         Type = iota
	Epyrimic        Type = iota
	Lycrimic        Type = iota
	Bylimic         Type = iota
	Zothimic        Type = iota
	Phrolimic       Type = iota
	Ionagimic       Type = iota
	Aeolaphimic     Type = iota
	Kycrimic        Type = iota
	Garimic         Type = iota
	Soptimic        Type = iota
	Ionyptimic      Type = iota
	Gyrimic         Type = iota
	Zalimic         Type = iota
	Stolimic        Type = iota
	Thonimic        Type = iota
	Stadimic        Type = iota
	Thodimic        Type = iota
	Mythimic        Type = iota
	Sogimic         Type = iota
	Gogimic         Type = iota
	Rothimic        Type = iota
	Katarimic       Type = iota
	Sylimic         Type = iota
	Mixolimic       Type = iota
	Dadimic         Type = iota
	Aeolyphimic     Type = iota
	Gycrimic        Type = iota
	Pyrimic         Type = iota
	Lydimic         Type = iota
	Ionacrimic      Type = iota
	Gathimic        Type = iota
	Ionynimic       Type = iota
	Phrynimic       Type = iota
	Stathimic       Type = iota
	Thatimic        Type = iota
	Dalimic         Type = iota
	Dygimic         Type = iota
	Zolimic         Type = iota
	Epogimic        Type = iota
	Lanimic         Type = iota
	Paptimic        Type = iota
	Darmic          Type = iota
	Lonimic         Type = iota
	Ionycrimic      Type = iota
	Phradimic       Type = iota
	Aeolorimic      Type = iota
	Gonimic         Type = iota
	Phracrimic      Type = iota
	Aerathimic      Type = iota
	Sarimic         Type = iota
	Zoptimic        Type = iota
	Zeracrimic      Type = iota
	Byptimic        Type = iota
	Starimic        Type = iota
	Phrathimic      Type = iota
	Saptimic        Type = iota
	Aerodimic       Type = iota
	Macrimic        Type = iota
	Rogimic         Type = iota
	Bygimic         Type = iota
	Thycrimic       Type = iota
	Aeoladimic      Type = iota
	Dylimic         Type = iota
	Eponimic        Type = iota
	Katygimic       Type = iota
	Stalimic        Type = iota
	Stoptimic       Type = iota
	Zygimic         Type = iota
	Kataptimic      Type = iota
	Aeolaptimic     Type = iota
	Pothimic        Type = iota
	Rycrimic        Type = iota
	Ronimic         Type = iota
	Stycrimic       Type = iota
	Katorimic       Type = iota
	Epythimic       Type = iota
	Kaptimic        Type = iota
	Katythimic      Type = iota
	Madimic         Type = iota
	Aerygimic       Type = iota
	Pylimic         Type = iota
	Ionathimic      Type = iota
	Morimic         Type = iota
	Aerycrimic      Type = iota
	Ganimic         Type = iota
	Eparimic        Type = iota
	Lyrimic         Type = iota
	Phraptimic      Type = iota
	Bacrimic        Type = iota
	Phralimic       Type = iota
	Phrogimic       Type = iota
	Rathimic        Type = iota
	Katocrimic      Type = iota
	Phryptimic      Type = iota
	Katynimic       Type = iota
	Solimic         Type = iota
	Ionolimic       Type = iota
	Ionophimic      Type = iota
	Aeologimic      Type = iota
	Zadimic         Type = iota
	Sygimic         Type = iota
	Thogimic        Type = iota
	Rythimic        Type = iota
	Donimic         Type = iota
	Aeoloptimic     Type = iota
	Panimic         Type = iota
	Lodimic         Type = iota
	Laptimic        Type = iota
	Lygimic         Type = iota
	Logimic         Type = iota
	Lalimic         Type = iota
	Sothimic        Type = iota
	Phrocrimic      Type = iota
	Modimic         Type = iota
	Barimic         Type = iota
	Poptimic        Type = iota
	Sagimic         Type = iota
	Aelothimic      Type = iota
	Socrimic        Type = iota
	Syrimic         Type = iota
	Stodimic        Type = iota
	Ionocrimic      Type = iota
	Zycrimic        Type = iota
	Ionygimic       Type = iota
	Katathimic      Type = iota
	Bolimic         Type = iota
	Bothimic        Type = iota
	Katadimic       Type = iota
	Kodimic         Type = iota
	Tholimic        Type = iota
	Ralimic         Type = iota
	Kanimic         Type = iota
	Zylimic         Type = iota
	Zodimic         Type = iota
	Zarimic         Type = iota
	Phrythimic      Type = iota
	Rorimic         Type = iota
	Pynimic         Type = iota
	Zanimic         Type = iota
	Ranimic         Type = iota
	Ladimic         Type = iota
	Podimic         Type = iota
	Ionothimic      Type = iota
	Kytrimic        Type = iota
	Golimic         Type = iota
	Dyptimic        Type = iota
	Ryrimic         Type = iota
	Gylimic         Type = iota
	Aeolycrimic     Type = iota
	Palimic         Type = iota
	Stothimic       Type = iota
	Aeronimic       Type = iota
	Katagimic       Type = iota
	Phronimic       Type = iota
	Banimic         Type = iota
	Ionodimic       Type = iota
	Bogimic         Type = iota
	Mogimic         Type = iota
	Docrimic        Type = iota
	Epadimic        Type = iota
	Aerynimic       Type = iota
	Mydimic         Type = iota
	Thyptimic       Type = iota
	Phrothimic      Type = iota
	Katycrimic      Type = iota
	Ionalimic       Type = iota
	Loptimic        Type = iota
	Zagimic         Type = iota
	Lagimic         Type = iota
	Thyrimic        Type = iota
	Thothimic       Type = iota
	Bycrimic        Type = iota
	Pathimic        Type = iota
	Mothimic        Type = iota
	Aeranimic       Type = iota
	Ragimic         Type = iota
	Dolimic         Type = iota
	Porimic         Type = iota
	Aerylimic       Type = iota
	Bocrimic        Type = iota
	Gythimic        Type = iota
	Pagimic         Type = iota
	Aeolythimic     Type = iota
	Molimic         Type = iota
	Staptimic       Type = iota
	Zacrimic        Type = iota
	Larimic         Type = iota
	Thacrimic       Type = iota
	Stydimic        Type = iota
	Lorimic         Type = iota
	Ionadimic       Type = iota
	Ionythimic      Type = iota
	Aerythimic      Type = iota
	Dynimic         Type = iota
	Zydimic         Type = iota
	Zathimic        Type = iota
	Radimic         Type = iota
	Stonimic        Type = iota
	Syptimic        Type = iota
	Ponimic         Type = iota
	Kadimic         Type = iota
	Gynimic         Type = iota
	Thydimic        Type = iota
	Polimic         Type = iota
	Thanimic        Type = iota
	Lathimic        Type = iota
	Aeralimic       Type = iota
	Kynimic         Type = iota
	Stynimic        Type = iota
	Epytimic        Type = iota
	Katoptimic      Type = iota
	Galimic         Type = iota
	Kathimic        Type = iota
	Lylimic         Type = iota
	Epalimic        Type = iota
	Epacrimic       Type = iota
	Sathimic        Type = iota
	Katanimic       Type = iota
	Katyrimic       Type = iota
	Rynimic         Type = iota
	Pogimic         Type = iota
	Aeraptimic      Type = iota
	Epylimic        Type = iota
	Manimic         Type = iota
	Marimic         Type = iota
	Locrimic        Type = iota
	Rylimic         Type = iota
	Epatimic        Type = iota
	Byrimic         Type = iota
	Kocrimic        Type = iota
	Korimic         Type = iota
	Lynimic         Type = iota
	Malimic         Type = iota
	Synimic         Type = iota
	Phragimic       Type = iota
	Mycrimic        Type = iota
	Ionorimic       Type = iota
	Phrydimic       Type = iota
	Zyptimic        Type = iota
	Katothimic      Type = iota
	Phrylimic       Type = iota
	Aerothimic      Type = iota
	Stagimic        Type = iota
	Dorimic         Type = iota
	Phrycrimic      Type = iota
	Kyptimic        Type = iota
	Ionylimic       Type = iota
	Epynimic        Type = iota
	Ionogimic       Type = iota
	Kydimic         Type = iota
	Gaptimic        Type = iota
	Tharimic        Type = iota
	Ionaphimic      Type = iota
	Thoptimic       Type = iota
	Bagimic         Type = iota
	Kyrimic         Type = iota
	Sonimic         Type = iota
	Aeolonimic      Type = iota
	Rygimic         Type = iota
	Thagimic        Type = iota
	Kolimic         Type = iota
	Dycrimic        Type = iota
	Epycrimic       Type = iota
	Gocrimic        Type = iota
	Katolimic       Type = iota
	Dagimic         Type = iota
	Aeolydimic      Type = iota
	Parimic         Type = iota
	Ionaptimic      Type = iota
	Thylimic        Type = iota
	Lolimic         Type = iota
	Thalimic        Type = iota
	Stygimic        Type = iota
	Aeolygimic      Type = iota
	Aerogimic       Type = iota
	Dacrimic        Type = iota
	Baptimic        Type = iota
	Stythimic       Type = iota
	Kothimic        Type = iota
	Pygimic         Type = iota
	Rodimic         Type = iota
	Sorimic         Type = iota
	Monimic         Type = iota
	Aeragimic       Type = iota
	Epothimic       Type = iota
	Salimic         Type = iota
	Lyptimic        Type = iota
	Katonimic       Type = iota
	Gygimic         Type = iota
	Aeradimic       Type = iota
	Zyrimic         Type = iota
	Stylimic        Type = iota
	Lythimic        Type = iota
	Dodimic         Type = iota
	Katalimic       Type = iota
	Boptimic        Type = iota
	Stogimic        Type = iota
	Thynimic        Type = iota
	Aeolathimic     Type = iota
	Bythimic        Type = iota
	Padimic         Type = iota
	Dathimic        Type = iota
	Epagimic        Type = iota
	Raptimic        Type = iota
	Epolimic        Type = iota
	Sythimic        Type = iota
	Sydimic         Type = iota
	Gacrimic        Type = iota
	Borimic         Type = iota
	Sycrimic        Type = iota
	Gadimic         Type = iota
	Aeolocrimic     Type = iota
	Phrygimic       Type = iota
	WholeTone       Type = iota
	Lydian          Type = iota
	Mixolydian      Type = iota
	Aeolian         Type = iota
	Locrian         Type = iota
	Ionian          Type = iota
	Dorian          Type = iota
	Phrygian        Type = iota
	Ionythian       Type = iota
	Aeolyrian       Type = iota
	Gorian          Type = iota
	Aeolodian       Type = iota
	Doptian         Type = iota
	Aeraphian       Type = iota
	Zacrian         Type = iota
	Ionarian        Type = iota
	Dynian          Type = iota
	Zydian          Type = iota
	Zathian         Type = iota
	Radian          Type = iota
	Stonian         Type = iota
	Syptian         Type = iota
	Aeolacrian      Type = iota
	Zythian         Type = iota
	Dyrian          Type = iota
	Koptian         Type = iota
	Thocrian        Type = iota
	Aeolanian       Type = iota
	Danian          Type = iota
	Zogian          Type = iota
	Epyrian         Type = iota
	Lycrian         Type = iota
	Daptian         Type = iota
	Kygian          Type = iota
	Mocrian         Type = iota
	Zynian          Type = iota
	Phrolian        Type = iota
	Ionagian        Type = iota
	Aeodian         Type = iota
	Kycrian         Type = iota
	Epygian         Type = iota
	Zaptian         Type = iota
	Kagian          Type = iota
	Soptian         Type = iota
	Ionyptian       Type = iota
	Gyrian          Type = iota
	Zalian          Type = iota
	Stolian         Type = iota
	Bylian          Type = iota
	Zothian         Type = iota
	Thonian         Type = iota
	Phrorian        Type = iota
	Stadian         Type = iota
	Thodian         Type = iota
	Dogian          Type = iota
	Mixopyrian      Type = iota
	Garian          Type = iota
	Epathian        Type = iota
	Mythian         Type = iota
	Sogian          Type = iota
	Gogian          Type = iota
	Rothian         Type = iota
	Katarian        Type = iota
	Stylian         Type = iota
	Stathian        Type = iota
	Mixonyphian     Type = iota
	Magian          Type = iota
	Dadian          Type = iota
	Aeolylian       Type = iota
	Gycrian         Type = iota
	Pyrian          Type = iota
	Epogian         Type = iota
	Lanian          Type = iota
	Paptian         Type = iota
	Ionacrian       Type = iota
	Gathian         Type = iota
	Ionyphian       Type = iota
	Phrynian        Type = iota
	Ionycrian       Type = iota
	Phradian        Type = iota
	Aeolorian       Type = iota
	Gonian          Type = iota
	Dalian          Type = iota
	Dygian          Type = iota
	Zolian          Type = iota
	Aerathian       Type = iota
	Sarian          Type = iota
	Zoptian         Type = iota
	Aeracrian       Type = iota
	Byptian         Type = iota
	Darian          Type = iota
	Lonian          Type = iota
	Aeopian         Type = iota
	Rygian          Type = iota
	Epynian         Type = iota
	Ionogian        Type = iota
	Kydian          Type = iota
	Gaptian         Type = iota
	Tharian         Type = iota
	Epycrian        Type = iota
	Gocrian         Type = iota
	Katolian        Type = iota
	Thoptian        Type = iota
	Bagian          Type = iota
	Kyrian          Type = iota
	Sonian          Type = iota
	Parian          Type = iota
	Ionaptian       Type = iota
	Thylian         Type = iota
	Lolian          Type = iota
	Thagian         Type = iota
	Kolian          Type = iota
	Dycrian         Type = iota
	Stygian         Type = iota
	Aeolygian       Type = iota
	Aerogian        Type = iota
	Dacrian         Type = iota
	Baptian         Type = iota
	Dagian          Type = iota
	Aeolydian       Type = iota
	Stythian        Type = iota
	Kothian         Type = iota
	Pygian          Type = iota
	Rodian          Type = iota
	Sorian          Type = iota
	Monian          Type = iota
	Thalian         Type = iota
	Zorian          Type = iota
	Aeragian        Type = iota
	Epothian        Type = iota
	Salian          Type = iota
	Lyptian         Type = iota
	Katonian        Type = iota
	Gyphian         Type = iota
	Thacrian        Type = iota
	Dodian          Type = iota
	Aeolyptian      Type = iota
	Aeolonian       Type = iota
	Aeradian        Type = iota
	Aeolagian       Type = iota
	Zyrian          Type = iota
	Aeolathian      Type = iota
	Bythian         Type = iota
	Padian          Type = iota
	Rolian          Type = iota
	Pydian          Type = iota
	Thygian         Type = iota
	Katalian        Type = iota
	Saptian         Type = iota
	Aerodian        Type = iota
	Macrian         Type = iota
	Rogian          Type = iota
	Boptian         Type = iota
	Stogian         Type = iota
	Thynian         Type = iota
	Thycrian        Type = iota
	Aeoladian       Type = iota
	Dylian          Type = iota
	Eponian         Type = iota
	Katygian        Type = iota
	Starian         Type = iota
	Phrathian       Type = iota
	Stalian         Type = iota
	Stoptian        Type = iota
	Zygian          Type = iota
	Kataptian       Type = iota
	Aeolaptian      Type = iota
	Pothian         Type = iota
	Bygian          Type = iota
	Morian          Type = iota
	Rycrian         Type = iota
	Ronian          Type = iota
	Stycrian        Type = iota
	Katorian        Type = iota
	Epythian        Type = iota
	Kaptian         Type = iota
	Phraptian       Type = iota
	Bacrian         Type = iota
	Katythian       Type = iota
	Madian          Type = iota
	Aerygian        Type = iota
	Pylian          Type = iota
	Ionathian       Type = iota
	Katocrian       Type = iota
	Phryptian       Type = iota
	Katynian        Type = iota
	Aerycrian       Type = iota
	Ganian          Type = iota
	Eparian         Type = iota
	Lyrian          Type = iota
	Ionopian        Type = iota
	Aeologian       Type = iota
	Zadian          Type = iota
	Sygian          Type = iota
	Phralian        Type = iota
	Phrogian        Type = iota
	Rathian         Type = iota
	Rythian         Type = iota
	Donian          Type = iota
	Aeoloptian      Type = iota
	Panian          Type = iota
	Lodian          Type = iota
	Solian          Type = iota
	Ionolian        Type = iota
	Laptian         Type = iota
	Lygian          Type = iota
	Logian          Type = iota
	Lalian          Type = iota
	Sothian         Type = iota
	Phrocrian       Type = iota
	Thogian         Type = iota
	Katathian       Type = iota
	Modian          Type = iota
	Barian          Type = iota
	Mixolocrian     Type = iota
	Sagian          Type = iota
	Aeolothian      Type = iota
	Socrian         Type = iota
	Tholian         Type = iota
	Ralian          Type = iota
	Syrian          Type = iota
	Stodian         Type = iota
	Ionocrian       Type = iota
	Zycrian         Type = iota
	Ionygian        Type = iota
	Zarian          Type = iota
	Phrythian       Type = iota
	Rorian          Type = iota
	Bolian          Type = iota
	Bothian         Type = iota
	Katadian        Type = iota
	Kodian          Type = iota
	Ranian          Type = iota
	Ladian          Type = iota
	Podian          Type = iota
	Ionothian       Type = iota
	Kanian          Type = iota
	Zylian          Type = iota
	Zodian          Type = iota
	Golian          Type = iota
	Dyptian         Type = iota
	Ryphian         Type = iota
	Gylian          Type = iota
	Aeolycrian      Type = iota
	Pynian          Type = iota
	Zanian          Type = iota
	Palian          Type = iota
	Stothian        Type = iota
	Aerorian        Type = iota
	Katagian        Type = iota
	Phronian        Type = iota
	Banian          Type = iota
	Aeronian        Type = iota
	Loptian         Type = iota
	Ionodian        Type = iota
	Bogian          Type = iota
	Mogian          Type = iota
	Docrian         Type = iota
	Epadian         Type = iota
	Aerynian        Type = iota
	Bycrian         Type = iota
	Pathian         Type = iota
	Mydian          Type = iota
	Thyptian        Type = iota
	Phrothian       Type = iota
	Katycrian       Type = iota
	Ionalian        Type = iota
	Dolian          Type = iota
	Porian          Type = iota
	Aerylian        Type = iota
	Zagian          Type = iota
	Lagian          Type = iota
	Tyrian          Type = iota
	Mixonorian      Type = iota
	Pagian          Type = iota
	Aeolythian      Type = iota
	Molian          Type = iota
	Staptian        Type = iota
	Mothian         Type = iota
	Aeranian        Type = iota
	Ragian          Type = iota
	Larian          Type = iota
	Lythian         Type = iota
	Stydian         Type = iota
	Lorian          Type = iota
	Ionadian        Type = iota
	Bocrian         Type = iota
	Mixolythian     Type = iota
	Thadian         Type = iota
	Sanian          Type = iota
	Ionydian        Type = iota
	Epydian         Type = iota
	Katydian        Type = iota
	Mathian         Type = iota
	Aeryptian       Type = iota
	Pythian         Type = iota
	Katylian        Type = iota
	Bydian          Type = iota
	Bynian          Type = iota
	Galian          Type = iota
	Zonian          Type = iota
	Myrian          Type = iota
	Katogian        Type = iota
	Stacrian        Type = iota
	Styrian         Type = iota
	Ionyrian        Type = iota
	Phrodian        Type = iota
	Pycrian         Type = iota
	Gyptian         Type = iota
	Katacrian       Type = iota
	Sodian          Type = iota
	Bathian         Type = iota
	Mylian          Type = iota
	Godian          Type = iota
	Thorian         Type = iota
	Zocrian         Type = iota
	Stanian         Type = iota
	Epanian         Type = iota
	Konian          Type = iota
	Stocrian        Type = iota
	Kalian          Type = iota
	Phroptian       Type = iota
	Dydian          Type = iota
	Katyptian       Type = iota
	Epodian         Type = iota
	Mygian          Type = iota
	Pacrian         Type = iota
	Aerocrian       Type = iota
	Aeolarian       Type = iota
	Kythian         Type = iota
	Bonian          Type = iota
	Badian          Type = iota
	Katodian        Type = iota
	Sadian          Type = iota
	Dothian         Type = iota
	Moptian         Type = iota
	Aeryrian        Type = iota
	Epagian         Type = iota
	Raptian         Type = iota
	Epolian         Type = iota
	Sythian         Type = iota
	Sydian          Type = iota
	Epocrian        Type = iota
	Kylian          Type = iota
	Gacrian         Type = iota
	Borian          Type = iota
	Sycrian         Type = iota
	Gadian          Type = iota
	Aeolocrian      Type = iota
	Mixodorian      Type = iota
	Dathian         Type = iota
	Katoptian       Type = iota
	Ponian          Type = iota
	Kadian          Type = iota
	Gynian          Type = iota
	Thyphian        Type = iota
	Polian          Type = iota
	Thanian         Type = iota
	Epacrian        Type = iota
	Sathian         Type = iota
	Lathian         Type = iota
	Aeralian        Type = iota
	Kynian          Type = iota
	Stynian         Type = iota
	Epyphian        Type = iota
	Pogian          Type = iota
	Aeraptian       Type = iota
	Epylian         Type = iota
	Gamian          Type = iota
	Kathian         Type = iota
	Lylian          Type = iota
	Epalian         Type = iota
	Eporian         Type = iota
	Rylian          Type = iota
	Epaptian        Type = iota
	Byrian          Type = iota
	Katanian        Type = iota
	Katyrian        Type = iota
	Rynian          Type = iota
	Korian          Type = iota
	Lynian          Type = iota
	Malian          Type = iota
	Synian          Type = iota
	Phragian        Type = iota
	Manian          Type = iota
	Marian          Type = iota
	Mycrian         Type = iota
	Ionorian        Type = iota
	Phrydian        Type = iota
	Zyptian         Type = iota
	Katothian       Type = iota
	Phrylian        Type = iota
	Kocrian         Type = iota
	Ionanian        Type = iota
	Aerothian       Type = iota
	Stagian         Type = iota
	Lothian         Type = iota
	Phrycrian       Type = iota
	Kyptian         Type = iota
	Ionylian        Type = iota
	Gydian          Type = iota
	Kogian          Type = iota
	Rarian          Type = iota
	Aerolian        Type = iota
	Karian          Type = iota
	Myptian         Type = iota
	Rydian          Type = iota
	Aeolynian       Type = iota
	Aeroptian       Type = iota
	Phryrian        Type = iota
	Gothian         Type = iota
	Storian         Type = iota
	Pyptian         Type = iota
	Thydian         Type = iota
	Aerycryllic     Type = iota
	Gadyllic        Type = iota
	Solyllic        Type = iota
	Zylyllic        Type = iota
	Mixodyllic      Type = iota
	Soryllic        Type = iota
	Godyllic        Type = iota
	Epiphyllic      Type = iota
	Pynyllic        Type = iota
	Bocryllic       Type = iota
	Kogyllic        Type = iota
	Raryllic        Type = iota
	Zycryllic       Type = iota
	Mycryllic       Type = iota
	Laptyllic       Type = iota
	Pylyllic        Type = iota
	Pothyllic       Type = iota
	Phronyllic      Type = iota
	Stynyllic       Type = iota
	Rathyllic       Type = iota
	Aeryptyllic     Type = iota
	Zydyllic        Type = iota
	Katolyllic      Type = iota
	Rythyllic       Type = iota
	Locryllic       Type = iota
	Bylyllic        Type = iota
	Sogyllic        Type = iota
	Ionycryllic     Type = iota
	Koptyllic       Type = iota
	Epyryllic       Type = iota
	Soptyllic       Type = iota
	Aeolylyllic     Type = iota
	Aeracryllic     Type = iota
	Epygyllic       Type = iota
	Thonyllic       Type = iota
	Lanyllic        Type = iota
	Phrynyllic      Type = iota
	Lycryllic       Type = iota
	Ionyptyllic     Type = iota
	Epathyllic      Type = iota
	Dydyllic        Type = iota
	Thogyllic       Type = iota
	Rygyllic        Type = iota
	Bycryllic       Type = iota
	Zacryllic       Type = iota
	Panyllic        Type = iota
	Dyryllic        Type = iota
	Zathyllic       Type = iota
	Dagyllic        Type = iota
	Katalyllic      Type = iota
	Katoryllic      Type = iota
	Dodyllic        Type = iota
	Zogyllic        Type = iota
	Madyllic        Type = iota
	Dycryllic       Type = iota
	Aeologyllic     Type = iota
	Sydyllic        Type = iota
	Katogyllic      Type = iota
	Zygyllic        Type = iota
	Aeralyllic      Type = iota
	Bacryllic       Type = iota
	Aerygyllic      Type = iota
	Dathyllic       Type = iota
	Boptyllic       Type = iota
	Bagyllic        Type = iota
	Mathyllic       Type = iota
	Styptyllic      Type = iota
	Zolyllic        Type = iota
	Rocryllic       Type = iota
	Zyryllic        Type = iota
	Sagyllic        Type = iota
	Epinyllic       Type = iota
	Katagyllic      Type = iota
	Ragyllic        Type = iota
	Gothyllic       Type = iota
	Lythyllic       Type = iota
	Ionocryllic     Type = iota
	Gocryllic       Type = iota
	Epiryllic       Type = iota
	Aeradyllic      Type = iota
	Staptyllic      Type = iota
	Danyllic        Type = iota
	Goptyllic       Type = iota
	Epocryllic      Type = iota
	Ionoptyllic     Type = iota
	Aeoloryllic     Type = iota
	Thydyllic       Type = iota
	Gycryllic       Type = iota
	Lyryllic        Type = iota
	Mogyllic        Type = iota
	Katodyllic      Type = iota
	Moptyllic       Type = iota
	Dolyllic        Type = iota
	Moryllic        Type = iota
	Bydyllic        Type = iota
	Pocryllic       Type = iota
	Phracryllic     Type = iota
	Gyryllic        Type = iota
	Phrygyllic      Type = iota
	Dogyllic        Type = iota
	Thagyllic       Type = iota
	Thoptyllic      Type = iota
	Phraptyllic     Type = iota
	Gylyllic        Type = iota
	Phralyllic      Type = iota
	Dygyllic        Type = iota
	Ronyllic        Type = iota
	Epogyllic       Type = iota
	Aeoladyllic     Type = iota
	Kocryllic       Type = iota
	Lodyllic        Type = iota
	Bynyllic        Type = iota
	Kydyllic        Type = iota
	Bygyllic        Type = iota
	Phryptyllic     Type = iota
	Ionayllic       Type = iota
	Phroryllic      Type = iota
	Thyphyllic      Type = iota
	Poptyllic       Type = iota
	Mixonyllic      Type = iota
	Paptyllic       Type = iota
	Storyllic       Type = iota
	Phrycryllic     Type = iota
	Palyllic        Type = iota
	Phranyllic      Type = iota
	Stydyllic       Type = iota
	Zadyllic        Type = iota
	Zalyllic        Type = iota
	Zocryllic       Type = iota
	Katocryllic     Type = iota
	Aerathyllic     Type = iota
	Stoptyllic      Type = iota
	Lydyllic        Type = iota
	Radyllic        Type = iota
	Stagyllic       Type = iota
	Ionoryllic      Type = iota
	Phrodyllic      Type = iota
	Aeragyllic      Type = iota
	Banyllic        Type = iota
	Epothyllic      Type = iota
	Zoryllic        Type = iota
	Phrolyllic      Type = iota
	Kolyllic        Type = iota
	Thodyllic       Type = iota
	Socryllic       Type = iota
	Aeolyllic       Type = iota
	Zythyllic       Type = iota
	Aeoryllic       Type = iota
	Mixolydyllic    Type = iota
	Mixonyphyllic   Type = iota
	Aeolanyllic     Type = iota
	Thocryllic      Type = iota
	Kygyllic        Type = iota
	Ionagyllic      Type = iota
	Gogyllic        Type = iota
	Phradyllic      Type = iota
	Ioniptyllic     Type = iota
	Kycryllic       Type = iota
	Aeolaptyllic    Type = iota
	Rodyllic        Type = iota
	Ionathyllic     Type = iota
	Pythyllic       Type = iota
	Zonyllic        Type = iota
	Ryryllic        Type = iota
	Aeolothyllic    Type = iota
	Ionyryllic      Type = iota
	Rydyllic        Type = iota
	Gonyllic        Type = iota
	Rolyllic        Type = iota
	Katydyllic      Type = iota
	Zyptyllic       Type = iota
	Modyllic        Type = iota
	Maptyllic       Type = iota
	Aeraptyllic     Type = iota
	Katadyllic      Type = iota
	Magyllic        Type = iota
	Phrylyllic      Type = iota
	Epigyllic       Type = iota
	Molyllic        Type = iota
	Ponyllic        Type = iota
	Thyptyllic      Type = iota
	Ionogyllic      Type = iota
	Aeolaryllic     Type = iota
	Katygyllic      Type = iota
	Ganyllic        Type = iota
	Kyptyllic       Type = iota
	Salyllic        Type = iota
	Sanyllic        Type = iota
	Doptyllic       Type = iota
	Ionilyllic      Type = iota
	Manyllic        Type = iota
	Polyllic        Type = iota
	Stanyllic       Type = iota
	Mixotharyllic   Type = iota
	Eporyllic       Type = iota
	Aerynyllic      Type = iota
	Lonyllic        Type = iota
	Sathyllic       Type = iota
	Layllic         Type = iota
	Saryllic        Type = iota
	Thacryllic      Type = iota
	Aeolynyllic     Type = iota
	Thadyllic       Type = iota
	Lynyllic        Type = iota
	Aeolathyllic    Type = iota
	Aeolocryllic    Type = iota
	Phroptyllic     Type = iota
	Kodyllic        Type = iota
	Epaptyllic      Type = iota
	Ionoyllic       Type = iota
	Gyptyllic       Type = iota
	Aerythyllic     Type = iota
	Zagyllic        Type = iota
	Epacryllic      Type = iota
	Thorcryllic     Type = iota
	Loptyllic       Type = iota
	Katylyllic      Type = iota
	Malyllic        Type = iota
	Mydyllic        Type = iota
	Thycryllic      Type = iota
	Gythyllic       Type = iota
	Pyryllic        Type = iota
	Rycryllic       Type = iota
	Phrathyllic     Type = iota
	Badyllic        Type = iota
	Phrocryllic     Type = iota
	Staryllic       Type = iota
	Zothyllic       Type = iota
	Tharyllic       Type = iota
	Sylyllic        Type = iota
	Lothyllic       Type = iota
	Daryllic        Type = iota
	Monyllic        Type = iota
	Styryllic       Type = iota
	Aeolacryllic    Type = iota
	Raptyllic       Type = iota
	Kataryllic      Type = iota
	Aerocryllic     Type = iota
	Zanyllic        Type = iota
	Aeolonyllic     Type = iota
	Aeonyllic       Type = iota
	Kyryllic        Type = iota
	Sythyllic       Type = iota
	Katycryllic     Type = iota
	Stogyllic       Type = iota
	Ionidyllic      Type = iota
	Stonyllic       Type = iota
	Stalyllic       Type = iota
	Poryllic        Type = iota
	Mocryllic       Type = iota
	Aeolyryllic     Type = iota
	Baryllic        Type = iota
	Dalyllic        Type = iota
	Ionyphyllic     Type = iota
	Zaptyllic       Type = iota
	Garyllic        Type = iota
	Gathyllic       Type = iota
	Mixopyryllic    Type = iota
	Ionacryllic     Type = iota
	Stylyllic       Type = iota
	Stycryllic      Type = iota
	Ionothyllic     Type = iota
	Mythyllic       Type = iota
	Aerylyllic      Type = iota
	Bonyllic        Type = iota
	Tholyllic       Type = iota
	Katyryllic      Type = iota
	Sadyllic        Type = iota
	Stolyllic       Type = iota
	Logyllic        Type = iota
	Dacryllic       Type = iota
	Thynyllic       Type = iota
	Gydyllic        Type = iota
	Eparyllic       Type = iota
	Dynyllic        Type = iota
	Ionyllic        Type = iota
	Zaryllic        Type = iota
	Dythyllic       Type = iota
	Ionaryllic      Type = iota
	Laryllic        Type = iota
	Kataptyllic     Type = iota
	Sonyllic        Type = iota
	Pathyllic       Type = iota
	Loryllic        Type = iota
	Aeronyllic      Type = iota
	Pycryllic       Type = iota
	Mygyllic        Type = iota
	Lylyllic        Type = iota
	Daptyllic       Type = iota
	Ioninyllic      Type = iota
	Epaphyllic      Type = iota
	Lolyllic        Type = iota
	Stacryllic      Type = iota
	Doryllic        Type = iota
	Kadyllic        Type = iota
	Rynyllic        Type = iota
	Aerogyllic      Type = iota
	Rothyllic       Type = iota
	Kagyllic        Type = iota
	Stathyllic      Type = iota
	Thyryllic       Type = iota
	Gygyllic        Type = iota
	Sodyllic        Type = iota
	Goryllic        Type = iota
	Bothyllic       Type = iota
	Gynyllic        Type = iota
	Ionaptyllic     Type = iota
	Phryryllic      Type = iota
	Racryllic       Type = iota
	Epicryllic      Type = iota
	Stygyllic       Type = iota
	Syryllic        Type = iota
	Stythyllic      Type = iota
	Aerothyllic     Type = iota
	Mixoryllic      Type = iota
	Thanyllic       Type = iota
	Roryllic        Type = iota
	Epotyllic       Type = iota
	Epidyllic       Type = iota
	Kaptyllic       Type = iota
	MajorDiminished Type = iota
	MinorDiminished Type = iota
	Aerycrygic      Type = iota
	Gadygic         Type = iota
	Solygic         Type = iota
	Zylygic         Type = iota
	Garygic         Type = iota
	Sorygic         Type = iota
	Godygic         Type = iota
	Epithygic       Type = iota
	Ionoptygic      Type = iota
	Kalygic         Type = iota
	Ionodygic       Type = iota
	Bythygic        Type = iota
	Epygic          Type = iota
	Marygic         Type = iota
	Gaptygic        Type = iota
	Aeroptygic      Type = iota
	Mylygic         Type = iota
	Galygic         Type = iota
	Mixolydygic     Type = iota
	Ionycrygic      Type = iota
	Zoptygic        Type = iota
	Phrygygic       Type = iota
	Locrygic        Type = iota
	Gonygic         Type = iota
	Aeracrygic      Type = iota
	Aerathygic      Type = iota
	Dorygic         Type = iota
	Dycrygic        Type = iota
	Aeolygic        Type = iota
	Dydygic         Type = iota
	Tholygic        Type = iota
	Rynygic         Type = iota
	Bycrygic        Type = iota
	Zacrygic        Type = iota
	Panygic         Type = iota
	Dyrygic         Type = iota
	Loptygic        Type = iota
	Katylygic       Type = iota
	Phradygic       Type = iota
	Mixodygic       Type = iota
	Katalygic       Type = iota
	Katorygic       Type = iota
	Dogygic         Type = iota
	Zodygic         Type = iota
	Madygic         Type = iota
	Bagygic         Type = iota
	Mathygic        Type = iota
	Styptygic       Type = iota
	Zolygic         Type = iota
	Sydygic         Type = iota
	Katygic         Type = iota
	Zyphygic        Type = iota
	Aeralygic       Type = iota
	Ryptygic        Type = iota
	Apinygic        Type = iota
	Katagygic       Type = iota
	Radygic         Type = iota
	Gothygic        Type = iota
	Lythygic        Type = iota
	Bacrygic        Type = iota
	Aerygic         Type = iota
	Dathygic        Type = iota
	Boptygic        Type = iota
	Epyrygic        Type = iota
	Aeradygic       Type = iota
	Staptygic       Type = iota
	Danygic         Type = iota
	Goptygic        Type = iota
	Epocrygic       Type = iota
	Rocrygic        Type = iota
	Zyrygic         Type = iota
	Sadygic         Type = iota
	Aeolorygic      Type = iota
	Thydygic        Type = iota
	Gycrygic        Type = iota
	Lyrygic         Type = iota
	Modygic         Type = iota
	Katodygic       Type = iota
	Moptygic        Type = iota
	Ionocrygic      Type = iota
	Gocrygic        Type = iota
	Manygic         Type = iota
	Polygic         Type = iota
	Stanygic        Type = iota
	Thaptygic       Type = iota
	Eporygic        Type = iota
	Aerynygic       Type = iota
	Thyptygic       Type = iota
	Ionogygic       Type = iota
	Aeolarygic      Type = iota
	Sathygic        Type = iota
	Ladygic         Type = iota
	Sarygic         Type = iota
	Thacrygic       Type = iota
	Aeolynygic      Type = iota
	Thadygic        Type = iota
	Lynygic         Type = iota
	Doptygic        Type = iota
	Ionilygic       Type = iota
	Phrygic         Type = iota
	Aeranygic       Type = iota
	Dothygic        Type = iota
	Lydygic         Type = iota
	Stadygic        Type = iota
	Byptygic        Type = iota
	Stodygic        Type = iota
	Zynygic         Type = iota
	Lonygic         Type = iota
	Zothygic        Type = iota
	Aeolathygic     Type = iota
	Aeolocrygic     Type = iota
	Phroptygic      Type = iota
	Kodygic         Type = iota
	Eparygic        Type = iota
	Ionygic         Type = iota
	Gyptygic        Type = iota
	Aerythygic      Type = iota
	Aeolacrygic     Type = iota
	Raptygic        Type = iota
	Gythygic        Type = iota
	Pyrygic         Type = iota
	Rycrygic        Type = iota
	Phrathygic      Type = iota
	Badygic         Type = iota
	Phrocrygic      Type = iota
	Starygic        Type = iota
	Kyrygic         Type = iota
	Sythygic        Type = iota
	Katycrygic      Type = iota
	Tharygic        Type = iota
	Sylygic         Type = iota
	Lothygic        Type = iota
	Darygic         Type = iota
	Monygic         Type = iota
	Styrygic        Type = iota
	Porygic         Type = iota
	Mocrygic        Type = iota
	Aeolyrygic      Type = iota
	Barygic         Type = iota
	Katarygic       Type = iota
	Aerocrygic      Type = iota
	Zanygic         Type = iota
	Aeolonygic      Type = iota
	Aeolanygic      Type = iota
	Kaptygic        Type = iota
	Sacrygic        Type = iota
	Padygic         Type = iota
	Epilygic        Type = iota
	Kynygic         Type = iota
	Stophygic       Type = iota
	Ionidygic       Type = iota
	Stonygic        Type = iota
	Stalygic        Type = iota
	Koptygic        Type = iota
	Raphygic        Type = iota
	Zycrygic        Type = iota
	Mycrygic        Type = iota
	Laptygic        Type = iota
	Pylygic         Type = iota
	Rodygic         Type = iota
	Epolygic        Type = iota
	Epidygic        Type = iota
	Phronygic       Type = iota
	Stynygic        Type = iota
	Zydygic         Type = iota
	Aerycryllian    Type = iota
	Gadyllian       Type = iota
	Solyllian       Type = iota
	Zyphyllian      Type = iota
	Garyllian       Type = iota
	Soryllian       Type = iota
	Godyllian       Type = iota
	Epityllian      Type = iota
	Ionyllian       Type = iota
	Aeoryllian      Type = iota
	Katoryllian     Type = iota
	Dodyllian       Type = iota
	Zogyllian       Type = iota
	Madyllian       Type = iota
	Dycryllian      Type = iota
	Aeogyllian      Type = iota
	Dydyllian       Type = iota
	Thogyllian      Type = iota
	Rygyllian       Type = iota
	Bathyllian      Type = iota
	Sydyllian       Type = iota
	Katogyllian     Type = iota
	Mixodyllian     Type = iota
	Aeradyllian     Type = iota
	Ryptyllian      Type = iota
	Loptyllian      Type = iota
	Kataphyllian    Type = iota
	Phradyllian     Type = iota
	Dagyllian       Type = iota
	Katyllian       Type = iota
	Gothyllian      Type = iota
	Lythyllian      Type = iota
	Bacryllian      Type = iota
	Aerygyllian     Type = iota
	Dathyllian      Type = iota
	Boptyllian      Type = iota
	Bagyllian       Type = iota
	Mathyllian      Type = iota
	Styptyllian     Type = iota
	Zolyllian       Type = iota
	Staptyllian     Type = iota
	Danyllian       Type = iota
	Goptyllian      Type = iota
	Epocryllian     Type = iota
	Rocryllian      Type = iota
	Zyryllian       Type = iota
	Sagyllian       Type = iota
	Epinyllian      Type = iota
	Katagyllian     Type = iota
	Ragyllian       Type = iota
	Thydyllian      Type = iota
	Epiryllian      Type = iota
	Lyryllian       Type = iota
	Mogyllian       Type = iota
	Katodyllian     Type = iota
	Aerycratic      Type = iota
	Monatic         Type = iota
	Solatic         Type = iota
	Zylatic         Type = iota
	Mixolatic       Type = iota
	Soratic         Type = iota
	Godatic         Type = iota
	Eptatic         Type = iota
	Ionatic         Type = iota
	Aeolatic        Type = iota
	Thydatic        Type = iota
	Chromatic       Type = iota
)

// AllScales return all scales
func AllScales() []Type {
	return []Type{
		// 3 notes, reference https://allthescales.org/scales.php?n=3
		Minoric,

		// 4 notes, reference https://allthescales.org/scales.php?n=4
		Thaptic, Lothic, Phratic, Aerathic,
		Epathic, Mynic, Rothic, Eporic,
		Zyphic, Epogic, Lanic, Pyrric,
		Aeoloric, Gonic, Dalic, Dygic,
		Daric, Lonic, Phradic, Bolic,
		Saric, Zoptic, Aeraphic, Byptic,
		Aeolic, Koptic, Mixolyric, Lydic,
		Stathic, Dadic,
		Phrynic,

		// 5 notes, reference https://allthescales.org/scales.php?n=5
		Epathitonic, Mynitonic, Rocritonic, Pentatonic, Thaptitonic,
		Magitonic, Daditonic, Aeolyphritonic, Gycritonic, Pyritonic,
		Gathitonic, Ionitonic, Phrynitonic, Stathitonic, Thalitonic,
		Zolitonic, Epogitonic, Lanitonic, Paptitonic, Ionacritonic,
		Phraditonic, Aeoloritonic, Gonitonic, Dalitonic, Dygitonic,
		Aeracritonic, Byptitonic, Daritonic, Lonitonic, Ionycritonic,
		Lothitonic, Phratonic, Aerathitonic, Saritonic, Zoptitonic,
		Dolitonic, Poritonic, Aerylitonic, Zagitonic, Lagitonic,
		Molitonic, Staptitonic, Mothitonic, Aeritonic, Ragitonic,
		Ionaditonic, Bocritonic, Gythitonic, Pagitonic, Aeolythitonic,
		Zacritonic, Laritonic, Thacritonic, Styditonic, Loritonic,
		Aeolyritonic, Goritonic, Aeoloditonic, Doptitonic, Aeraphitonic,
		Zathitonic, Raditonic, Stonitonic, Syptitonic, Ionythitonic,
		Aeolanitonic, Danitonic, Ionaritonic, Dynitonic, Zyditonic,
		Aeolacritonic, Zythitonic, Dyritonic, Koptitonic, Thocritonic,
		Lycritonic, Daptitonic, Kygitonic, Mocritonic, Zynitonic,
		Epygitonic, Zaptitonic, Kagitonic, Zogitonic, Epyritonic,
		Zothitonic, Phrolitonic, Ionagitonic, Aeolapritonic, Kyritonic,
		Ionyptitonic, Gyritonic, Zalitonic, Stolitonic, Bylitonic,
		Thoditonic, Dogitonic, Phralitonic, Garitonic, Soptitonic,
		Kataritonic, Sylitonic, Thonitonic, Phropitonic, Staditonic,
		Lyditonic, Mythitonic, Sogitonic, Gothitonic, Rothitonic,
		Zylitonic, Zoditonic, Zaritonic, Phrythitonic, Rolitonic,
		Ranitonic, Laditonic, Poditonic, Ionothitonic, Kanitonic,
		Ryphitonic, Gylitonic, Aeolycritonic, Pynitonic, Zanitonic,
		Phronitonic, Banitonic, Aeronitonic, Golitonic, Dyptitonic,
		Aerynitonic, Palitonic, Stothitonic, Aerophitonic, Katagitonic,
		Ionoditonic, Bogitonic, Mogitonic, Docritonic, Epaditonic,
		Mixitonic, Phrothitonic, Katycritonic, Ionalitonic, Loptitonic,
		Thyritonic, Thoptitonic, Bycritonic, Pathitonic, Myditonic,
		Bolitonic, Bothitonic, Kataditonic, Koditonic, Tholitonic,

		// 6 notes, reference https://allthescales.org/scales.php?n=6
		Epathimic, Mynimic, Rocrimic, Eporimic, Thaptimic, Lothimic,
		Dyrimic, Koptimic, Thocrimic, Aeolanimic, Danimic, Ionarimic,
		Daptimic, Kygimic, Mocrimic, Zynimic, Aeolimic, Zythimic,
		Epygimic, Zaptimic, Kagimic, Zogimic, Epyrimic, Lycrimic,
		Bylimic, Zothimic, Phrolimic, Ionagimic, Aeolaphimic, Kycrimic,
		Garimic, Soptimic, Ionyptimic, Gyrimic, Zalimic, Stolimic,
		Thonimic, Stadimic, Thodimic,
		Mythimic, Sogimic, Gogimic, Rothimic, Katarimic, Sylimic,
		Mixolimic, Dadimic, Aeolyphimic, Gycrimic, Pyrimic, Lydimic,
		Ionacrimic, Gathimic, Ionynimic, Phrynimic, Stathimic, Thatimic,
		Dalimic, Dygimic, Zolimic, Epogimic, Lanimic, Paptimic,
		Darmic, Lonimic, Ionycrimic, Phradimic, Aeolorimic, Gonimic,
		Phracrimic, Aerathimic, Sarimic, Zoptimic, Zeracrimic, Byptimic,
		Starimic, Phrathimic, Saptimic, Aerodimic, Macrimic, Rogimic,
		Bygimic, Thycrimic, Aeoladimic, Dylimic, Eponimic, Katygimic,
		Stalimic, Stoptimic, Zygimic, Kataptimic, Aeolaptimic, Pothimic,
		Rycrimic, Ronimic, Stycrimic, Katorimic, Epythimic, Kaptimic,
		Katythimic, Madimic, Aerygimic, Pylimic, Ionathimic, Morimic,
		Aerycrimic, Ganimic, Eparimic, Lyrimic, Phraptimic, Bacrimic,
		Phralimic, Phrogimic, Rathimic, Katocrimic, Phryptimic, Katynimic,
		Solimic, Ionolimic, Ionophimic, Aeologimic, Zadimic, Sygimic,
		Thogimic, Rythimic, Donimic, Aeoloptimic, Panimic, Lodimic,
		Laptimic, Lygimic, Logimic, Lalimic, Sothimic, Phrocrimic,
		Modimic, Barimic, Poptimic, Sagimic, Aelothimic, Socrimic,
		Syrimic, Stodimic, Ionocrimic, Zycrimic, Ionygimic, Katathimic,
		Bolimic, Bothimic, Katadimic, Kodimic, Tholimic, Ralimic,
		Kanimic, Zylimic, Zodimic, Zarimic, Phrythimic, Rorimic,
		Pynimic, Zanimic, Ranimic, Ladimic, Podimic, Ionothimic,
		Kytrimic, Golimic, Dyptimic, Ryrimic, Gylimic, Aeolycrimic,
		Palimic, Stothimic, Aeronimic, Katagimic, Phronimic, Banimic,
		Ionodimic, Bogimic, Mogimic, Docrimic, Epadimic, Aerynimic,
		Mydimic, Thyptimic, Phrothimic, Katycrimic, Ionalimic, Loptimic,
		Zagimic, Lagimic, Thyrimic, Thothimic, Bycrimic, Pathimic,
		Mothimic, Aeranimic, Ragimic, Dolimic, Porimic, Aerylimic,
		Bocrimic, Gythimic, Pagimic, Aeolythimic, Molimic, Staptimic,
		Zacrimic, Larimic, Thacrimic, Stydimic, Lorimic, Ionadimic,
		Ionythimic, Aerythimic,
		Dynimic, Zydimic, Zathimic, Radimic, Stonimic, Syptimic,
		Ponimic, Kadimic, Gynimic, Thydimic, Polimic, Thanimic,
		Lathimic, Aeralimic, Kynimic, Stynimic, Epytimic, Katoptimic,
		Galimic, Kathimic, Lylimic, Epalimic, Epacrimic, Sathimic,
		Katanimic, Katyrimic, Rynimic, Pogimic, Aeraptimic, Epylimic,
		Manimic, Marimic, Locrimic, Rylimic, Epatimic, Byrimic,
		Kocrimic, Korimic, Lynimic, Malimic, Synimic, Phragimic,
		Mycrimic, Ionorimic, Phrydimic, Zyptimic, Katothimic, Phrylimic,
		Aerothimic, Stagimic, Dorimic, Phrycrimic, Kyptimic, Ionylimic,
		Epynimic, Ionogimic, Kydimic, Gaptimic, Tharimic, Ionaphimic,
		Thoptimic, Bagimic, Kyrimic, Sonimic, Aeolonimic, Rygimic,
		Thagimic, Kolimic, Dycrimic, Epycrimic, Gocrimic, Katolimic,
		Dagimic, Aeolydimic, Parimic, Ionaptimic, Thylimic, Lolimic,
		Thalimic, Stygimic, Aeolygimic, Aerogimic, Dacrimic, Baptimic,
		Stythimic, Kothimic, Pygimic, Rodimic, Sorimic, Monimic,
		Aeragimic, Epothimic, Salimic, Lyptimic, Katonimic, Gygimic,
		Aeradimic, Zyrimic, Stylimic,
		Lythimic, Dodimic, Katalimic,
		Boptimic, Stogimic, Thynimic, Aeolathimic, Bythimic, Padimic,
		Dathimic, Epagimic, Raptimic, Epolimic, Sythimic, Sydimic,
		Gacrimic, Borimic, Sycrimic, Gadimic, Aeolocrimic, Phrygimic,
		WholeTone,

		// 7 notes, reference https://allthescales.org/scales.php?n=7
		Lydian, Mixolydian, Aeolian, Locrian, Ionian, Dorian, Phrygian,
		Ionythian, Aeolyrian, Gorian, Aeolodian, Doptian, Aeraphian, Zacrian,
		Ionarian, Dynian, Zydian, Zathian, Radian, Stonian, Syptian,
		Aeolacrian, Zythian, Dyrian, Koptian, Thocrian, Aeolanian, Danian,
		Zogian, Epyrian, Lycrian, Daptian, Kygian, Mocrian, Zynian,
		Phrolian, Ionagian, Aeodian, Kycrian, Epygian, Zaptian, Kagian,
		Soptian, Ionyptian, Gyrian, Zalian, Stolian, Bylian, Zothian,
		Thonian, Phrorian, Stadian, Thodian, Dogian, Mixopyrian, Garian,
		Epathian, Mythian, Sogian, Gogian, Rothian, Katarian, Stylian,
		Stathian, Mixonyphian, Magian, Dadian, Aeolylian, Gycrian, Pyrian,
		Epogian, Lanian, Paptian, Ionacrian, Gathian, Ionyphian, Phrynian,
		Ionycrian, Phradian, Aeolorian, Gonian, Dalian, Dygian, Zolian,
		Aerathian, Sarian, Zoptian, Aeracrian, Byptian, Darian, Lonian,
		Aeopian, Rygian, Epynian, Ionogian, Kydian, Gaptian, Tharian,
		Epycrian, Gocrian, Katolian, Thoptian, Bagian, Kyrian, Sonian,
		Parian, Ionaptian, Thylian, Lolian, Thagian, Kolian, Dycrian,
		Stygian, Aeolygian, Aerogian, Dacrian, Baptian, Dagian, Aeolydian,
		Stythian, Kothian, Pygian, Rodian, Sorian, Monian, Thalian,
		Zorian, Aeragian, Epothian, Salian, Lyptian, Katonian, Gyphian,
		Thacrian, Dodian, Aeolyptian, Aeolonian, Aeradian, Aeolagian, Zyrian,
		Aeolathian, Bythian, Padian, Rolian, Pydian, Thygian, Katalian,
		Saptian, Aerodian, Macrian, Rogian, Boptian, Stogian, Thynian,
		Thycrian, Aeoladian, Dylian, Eponian, Katygian, Starian, Phrathian,
		Stalian, Stoptian, Zygian, Kataptian, Aeolaptian, Pothian, Bygian,
		Morian, Rycrian, Ronian, Stycrian, Katorian, Epythian, Kaptian,
		Phraptian, Bacrian, Katythian, Madian, Aerygian, Pylian, Ionathian,
		Katocrian, Phryptian, Katynian, Aerycrian, Ganian, Eparian, Lyrian,
		Ionopian, Aeologian, Zadian, Sygian, Phralian, Phrogian, Rathian,
		Rythian, Donian, Aeoloptian, Panian, Lodian, Solian, Ionolian,
		Laptian, Lygian, Logian, Lalian, Sothian, Phrocrian, Thogian,
		Katathian, Modian, Barian, Mixolocrian, Sagian, Aeolothian, Socrian,
		Tholian, Ralian, Syrian, Stodian, Ionocrian, Zycrian, Ionygian,
		Zarian, Phrythian, Rorian, Bolian, Bothian, Katadian, Kodian,
		Ranian, Ladian, Podian, Ionothian, Kanian, Zylian, Zodian,
		Golian, Dyptian, Ryphian, Gylian, Aeolycrian, Pynian, Zanian,
		Palian, Stothian, Aerorian, Katagian, Phronian, Banian, Aeronian,
		Loptian, Ionodian, Bogian, Mogian, Docrian, Epadian, Aerynian,
		Bycrian, Pathian, Mydian, Thyptian, Phrothian, Katycrian, Ionalian,
		Dolian, Porian, Aerylian, Zagian, Lagian, Tyrian, Mixonorian,
		Pagian, Aeolythian, Molian, Staptian, Mothian, Aeranian, Ragian,
		Larian, Lythian, Stydian, Lorian, Ionadian, Bocrian, Mixolythian,
		Thadian, Sanian, Ionydian, Epydian, Katydian, Mathian, Aeryptian,
		Pythian, Katylian, Bydian, Bynian, Galian, Zonian, Myrian,
		Katogian, Stacrian, Styrian, Ionyrian, Phrodian, Pycrian, Gyptian,
		Katacrian, Sodian, Bathian, Mylian, Godian, Thorian, Zocrian,
		Stanian, Epanian, Konian, Stocrian, Kalian, Phroptian, Dydian,
		Katyptian, Epodian, Mygian, Pacrian, Aerocrian, Aeolarian, Kythian,
		Bonian, Badian, Katodian, Sadian, Dothian, Moptian, Aeryrian,
		Epagian, Raptian, Epolian, Sythian, Sydian, Epocrian, Kylian,
		Gacrian, Borian, Sycrian, Gadian, Aeolocrian, Mixodorian, Dathian,
		Katoptian, Ponian, Kadian, Gynian, Thyphian, Polian, Thanian,
		Epacrian, Sathian, Lathian, Aeralian, Kynian, Stynian, Epyphian,
		Pogian, Aeraptian, Epylian, Gamian, Kathian, Lylian, Epalian,
		Eporian, Rylian, Epaptian, Byrian, Katanian, Katyrian, Rynian,
		Korian, Lynian, Malian, Synian, Phragian, Manian, Marian,
		Mycrian, Ionorian, Phrydian, Zyptian, Katothian, Phrylian, Kocrian,
		Ionanian, Aerothian, Stagian, Lothian, Phrycrian, Kyptian, Ionylian,
		Gydian, Kogian, Rarian, Aerolian, Karian, Myptian, Rydian,
		Aeolynian, Aeroptian, Phryrian, Gothian, Storian, Pyptian, Thydian,

		// 8 notes, reference https://allthescales.org/scales.php?n=8
		Aerycryllic, Gadyllic, Solyllic, Zylyllic, Mixodyllic, Soryllic, Godyllic, Epiphyllic,
		Pynyllic, Bocryllic, Kogyllic, Raryllic, Zycryllic, Mycryllic, Laptyllic, Pylyllic,
		Pothyllic, Phronyllic, Stynyllic, Rathyllic, Aeryptyllic, Zydyllic, Katolyllic, Rythyllic,
		Locryllic, Bylyllic, Sogyllic, Ionycryllic, Koptyllic, Epyryllic, Soptyllic, Aeolylyllic,
		Aeracryllic, Epygyllic, Thonyllic, Lanyllic, Phrynyllic, Lycryllic, Ionyptyllic, Epathyllic,
		Dydyllic, Thogyllic, Rygyllic, Bycryllic, Zacryllic, Panyllic, Dyryllic, Zathyllic,
		Dagyllic, Katalyllic, Katoryllic, Dodyllic, Zogyllic, Madyllic, Dycryllic, Aeologyllic,
		Sydyllic, Katogyllic, Zygyllic, Aeralyllic,
		Bacryllic, Aerygyllic, Dathyllic, Boptyllic, Bagyllic, Mathyllic, Styptyllic, Zolyllic,
		Rocryllic, Zyryllic, Sagyllic, Epinyllic, Katagyllic, Ragyllic, Gothyllic, Lythyllic,
		Ionocryllic, Gocryllic, Epiryllic, Aeradyllic, Staptyllic, Danyllic, Goptyllic, Epocryllic,
		Ionoptyllic, Aeoloryllic, Thydyllic, Gycryllic, Lyryllic, Mogyllic, Katodyllic, Moptyllic,
		Dolyllic, Moryllic, Bydyllic, Pocryllic, Phracryllic, Gyryllic, Phrygyllic, Dogyllic,
		Thagyllic, Thoptyllic, Phraptyllic, Gylyllic, Phralyllic, Dygyllic, Ronyllic, Epogyllic,
		Aeoladyllic, Kocryllic, Lodyllic, Bynyllic, Kydyllic, Bygyllic, Phryptyllic, Ionayllic,
		Phroryllic, Thyphyllic, Poptyllic, Mixonyllic, Paptyllic, Storyllic, Phrycryllic, Palyllic,
		Phranyllic, Stydyllic, Zadyllic, Zalyllic, Zocryllic, Katocryllic, Aerathyllic, Stoptyllic,
		Lydyllic, Radyllic, Stagyllic, Ionoryllic, Phrodyllic, Aeragyllic, Banyllic, Epothyllic,
		Zoryllic, Phrolyllic, Kolyllic, Thodyllic, Socryllic, Aeolyllic, Zythyllic, Aeoryllic,
		Mixolydyllic, Mixonyphyllic, Aeolanyllic, Thocryllic, Kygyllic, Ionagyllic, Gogyllic, Phradyllic,
		Ioniptyllic, Kycryllic, Aeolaptyllic, Rodyllic, Ionathyllic, Pythyllic, Zonyllic, Ryryllic,
		Aeolothyllic, Ionyryllic, Rydyllic, Gonyllic, Rolyllic, Katydyllic, Zyptyllic, Modyllic,
		Maptyllic, Aeraptyllic, Katadyllic, Magyllic, Phrylyllic, Epigyllic, Molyllic, Ponyllic,
		Thyptyllic, Ionogyllic, Aeolaryllic, Katygyllic, Ganyllic, Kyptyllic, Salyllic, Sanyllic,
		Doptyllic, Ionilyllic, Manyllic, Polyllic, Stanyllic, Mixotharyllic, Eporyllic, Aerynyllic,
		Lonyllic, Sathyllic, Layllic, Saryllic, Thacryllic, Aeolynyllic, Thadyllic, Lynyllic,
		Aeolathyllic, Aeolocryllic, Phroptyllic, Kodyllic, Epaptyllic, Ionoyllic, Gyptyllic, Aerythyllic,
		Zagyllic, Epacryllic, Thorcryllic, Loptyllic, Katylyllic, Malyllic, Mydyllic, Thycryllic,
		Gythyllic, Pyryllic, Rycryllic, Phrathyllic, Badyllic, Phrocryllic, Staryllic, Zothyllic,
		Tharyllic, Sylyllic, Lothyllic, Daryllic, Monyllic, Styryllic, Aeolacryllic, Raptyllic,
		Kataryllic, Aerocryllic, Zanyllic, Aeolonyllic, Aeonyllic, Kyryllic, Sythyllic, Katycryllic,
		Stogyllic, Ionidyllic, Stonyllic, Stalyllic, Poryllic, Mocryllic, Aeolyryllic, Baryllic,
		Dalyllic, Ionyphyllic, Zaptyllic, Garyllic, Gathyllic, Mixopyryllic, Ionacryllic, Stylyllic,
		Stycryllic, Ionothyllic, Mythyllic, Aerylyllic, Bonyllic, Tholyllic, Katyryllic, Sadyllic,
		Stolyllic, Logyllic, Dacryllic, Thynyllic, Gydyllic, Eparyllic, Dynyllic, Ionyllic,
		Zaryllic, Dythyllic, Ionaryllic, Laryllic, Kataptyllic, Sonyllic, Pathyllic, Loryllic,
		Aeronyllic, Pycryllic, Mygyllic, Lylyllic, Daptyllic, Ioninyllic, Epaphyllic, Lolyllic,
		Stacryllic, Doryllic, Kadyllic, Rynyllic, Aerogyllic, Rothyllic, Kagyllic, Stathyllic,
		Thyryllic, Gygyllic, Sodyllic, Goryllic, Bothyllic, Gynyllic, Ionaptyllic, Phryryllic,
		Racryllic, Epicryllic, Stygyllic, Syryllic, Stythyllic, Aerothyllic, Mixoryllic, Thanyllic,
		Roryllic, Epotyllic, Epidyllic, Kaptyllic,
		MajorDiminished, MinorDiminished,

		// 9 notes, reference https://allthescales.org/scales.php?n=9
		Aerycrygic, Gadygic, Solygic, Zylygic, Garygic, Sorygic, Godygic, Epithygic, Ionoptygic,
		Kalygic, Ionodygic, Bythygic, Epygic, Marygic, Gaptygic, Aeroptygic, Mylygic, Galygic,
		Mixolydygic, Ionycrygic, Zoptygic, Phrygygic, Locrygic, Gonygic, Aeracrygic, Aerathygic, Dorygic,
		Dycrygic, Aeolygic, Dydygic, Tholygic, Rynygic, Bycrygic, Zacrygic, Panygic, Dyrygic,
		Loptygic, Katylygic, Phradygic, Mixodygic, Katalygic, Katorygic, Dogygic, Zodygic, Madygic,
		Bagygic, Mathygic, Styptygic, Zolygic, Sydygic, Katygic, Zyphygic, Aeralygic, Ryptygic,
		Apinygic, Katagygic, Radygic, Gothygic, Lythygic, Bacrygic, Aerygic, Dathygic, Boptygic,
		Epyrygic, Aeradygic, Staptygic, Danygic, Goptygic, Epocrygic, Rocrygic, Zyrygic, Sadygic,
		Aeolorygic, Thydygic, Gycrygic, Lyrygic, Modygic, Katodygic, Moptygic, Ionocrygic, Gocrygic,
		Manygic, Polygic, Stanygic, Thaptygic, Eporygic, Aerynygic, Thyptygic, Ionogygic, Aeolarygic,
		Sathygic, Ladygic, Sarygic, Thacrygic, Aeolynygic, Thadygic, Lynygic, Doptygic, Ionilygic,
		Phrygic, Aeranygic, Dothygic, Lydygic, Stadygic, Byptygic, Stodygic, Zynygic, Lonygic,
		Zothygic, Aeolathygic, Aeolocrygic, Phroptygic, Kodygic, Eparygic, Ionygic, Gyptygic, Aerythygic,
		Aeolacrygic, Raptygic, Gythygic, Pyrygic, Rycrygic, Phrathygic, Badygic, Phrocrygic, Starygic,
		Kyrygic, Sythygic, Katycrygic, Tharygic, Sylygic, Lothygic, Darygic, Monygic, Styrygic,
		Porygic, Mocrygic, Aeolyrygic, Barygic, Katarygic, Aerocrygic, Zanygic, Aeolonygic, Aeolanygic,
		Kaptygic, Sacrygic, Padygic, Epilygic, Kynygic, Stophygic, Ionidygic, Stonygic, Stalygic,
		Koptygic, Raphygic, Zycrygic, Mycrygic, Laptygic, Pylygic, Rodygic, Epolygic, Epidygic,
		Phronygic, Stynygic, Zydygic,

		// 10 notes, reference https://allthescales.org/scales.php?n=10
		Aerycryllian, Gadyllian, Solyllian, Zyphyllian, Garyllian, Soryllian, Godyllian, Epityllian, Ionyllian, Aeoryllian,
		Katoryllian, Dodyllian, Zogyllian, Madyllian, Dycryllian, Aeogyllian, Dydyllian, Thogyllian, Rygyllian, Bathyllian,
		Sydyllian, Katogyllian, Mixodyllian, Aeradyllian, Ryptyllian, Loptyllian, Kataphyllian, Phradyllian, Dagyllian, Katyllian,
		Gothyllian, Lythyllian, Bacryllian, Aerygyllian, Dathyllian, Boptyllian, Bagyllian, Mathyllian, Styptyllian, Zolyllian,
		Staptyllian, Danyllian, Goptyllian, Epocryllian, Rocryllian, Zyryllian, Sagyllian, Epinyllian, Katagyllian, Ragyllian,
		Thydyllian, Epiryllian, Lyryllian, Mogyllian, Katodyllian,

		// 11 notes, reference https://allthescales.org/scales.php?n=11
		Aerycratic, Monatic, Solatic, Zylatic, Mixolatic, Soratic, Godatic, Eptatic, Ionatic, Aeolatic, Thydatic,

		// 12 notes, reference https://allthescales.org/scales.php?n=12
		Chromatic,
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

// PitchClass returns scale pitch class where 0 is CNatural all the way to 11 (BNatural)
func (s Type) PitchClass() []int {
	class := make([]int, 0)
	for _, v := range s.Pitches(pitch.CNatural) {
		class = append(class, int(v-1))
	}
	return class
}

// PitchFlags returns a slice of positional pitch status, from CNatural to BNatural
func (s Type) PitchFlags() []bool {
	flags := make([]bool, 12)
	for _, v := range s.PitchClass() {
		flags[v] = true
	}

	return flags
}

// IntervalPattern returns scale interval pattern in semitones
func (s Type) IntervalPattern() []int {
	class := append(s.PitchClass(), 12)

	var previous int
	pattern := make([]int, 0)
	for i, v := range class {
		if i > 0 {
			pattern = append(pattern, v-previous)
		}
		previous = v
	}
	return pattern
}

// Perfection stores perfection profile
type Perfection struct {
	Perfection   int
	Imperfection int
}

// Perfection return perfection profile
func (s Type) Perfection() Perfection {
	pitches := s.Pitches(pitch.CNatural)
	pitchMap := makePitchMap(pitches)

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

// Cardinality returns scale cardinality
func (s Type) Cardinality() int {
	var cardinality int
	for _, v := range pitch.AllPitches() {
		if s.Number()&v.Number() == v.Number() {
			cardinality++
		}
	}

	return cardinality
}

// RotationalSymmetric returns true when each note has its tritone
func (s Type) RotationalSymmetric() bool {
	return s.RotationalSymmetryLevel() > 0
}

// Palindromic returns true when scale is reflective symmetric at it's root
func (s Type) Palindromic() bool {
	return s.reflectiveSymmetricAt(0)
}

// ReflectiveSymmetric returns true when scale is reflective symmetric
func (s Type) ReflectiveSymmetric() bool {
	for i := 0; i < 12; i++ {
		if s.reflectiveSymmetricAt(i) {
			return true
		}
	}

	return false
}

// ReflectiveSymmetryAxes returns axes where the scale is reflective symmetric
func (s Type) ReflectiveSymmetryAxes() []int {
	axes := make([]int, 0)
	for i := 0; i < 12; i++ {
		if s.reflectiveSymmetricAt(i) {
			axes = append(axes, i)
		}
	}

	return axes
}

// ReflectiveSymmetry returns true when scale is reflective symmetry
func (s Type) reflectiveSymmetricAt(axis int) bool {
	pitchFlags := s.PitchFlags()

	r := axis % 12
	rotated := append(pitchFlags[r:], pitchFlags[:r]...)
	return rotated[5] == rotated[7] &&
		rotated[4] == rotated[8] &&
		rotated[3] == rotated[9] &&
		rotated[2] == rotated[10] &&
		rotated[1] == rotated[11]
}

// RotationalSymmetryLevel returns number of semitone to the next symmetry
func (s Type) RotationalSymmetryLevel() int {
	slice := pitch.Slice(s.Pitches(pitch.CNatural))
	for i := 1; i < 12; i++ {
		if slice.Equal(slice.Transpose(i)) {
			return i
		}
	}

	return 0
}

// CenterOfGravity returns scale
func (s Type) CenterOfGravity(tonic pitch.Type) (float64, float64) {
	type coordinate struct {
		X float64
		Y float64
	}

	valueMap := map[int]coordinate{
		0:  {0, 1},
		1:  {0.5, 0.8660},
		2:  {0.8660, 0.5},
		3:  {1.0, 0},
		4:  {0.8660, -0.5},
		5:  {0.5, -0.8660},
		6:  {0, -1.0},
		7:  {-0.5, -0.8660},
		8:  {-0.8660, -0.5},
		9:  {-1.0, 0},
		10: {-0.8660, 0.5},
		11: {-0.5, 0.8660},
	}

	x, y := 0.0, 0.0
	for _, v := range s.PitchClass() {
		x += valueMap[v].X
		y += valueMap[v].Y
	}

	x = math.Round(x*10000) / 10000
	y = math.Round(y*10000) / 10000

	return x, y
}

// Balanced returns true when scale is balanced
func (s Type) Balanced() bool {
	x, y := s.CenterOfGravity(pitch.CNatural)
	return x == float64(0) && y == float64(0)
}

// FifthGeneratorRoot returns a pitch and degree where the scale can be generated by fetching the next fifth, assuming scale stars with C
func (s Type) FifthGeneratorRoot() pitch.WithDegree {
	return s.FifthGeneratorRootWithTonic(pitch.CNatural)
}

// FifthGeneratorRootWithTonic returns a pitch and degree where the scale can be generated by fetching the next fifth with custom tonic
func (s Type) FifthGeneratorRootWithTonic(tonic pitch.Type) pitch.WithDegree {
	allPitches := s.Pitches(tonic)
	orphanPitches := make([]pitch.WithDegree, 0)
	centerPitches := make([]pitch.WithDegree, 0)
	leftPitches := make([]pitch.WithDegree, 0)
	rightPitches := make([]pitch.WithDegree, 0)

	for i, v := range allPitches {
		prevFound := slices.Contains(allPitches, v.PreviousFifth())
		nextFound := slices.Contains(allPitches, v.NextFifth())
		withDegree := pitch.WithDegree{Pitch: v, Degree: degree.FromInt(i + 1)}
		switch {
		case !prevFound && !nextFound:
			orphanPitches = append(orphanPitches, withDegree)
		case prevFound && nextFound:
			centerPitches = append(centerPitches, withDegree)
		case prevFound:
			rightPitches = append(rightPitches, withDegree)
		default:
			leftPitches = append(leftPitches, withDegree)
		}
	}

	if len(leftPitches) == 1 {
		return leftPitches[0]
	}

	return pitch.WithDegree{}
}

func makePitchMap(pitches []pitch.Type) map[pitch.Type]struct{} {
	pitchMap := make(map[pitch.Type]struct{})
	for _, v := range pitches {
		pitchMap[v] = struct{}{}
	}

	return pitchMap
}
