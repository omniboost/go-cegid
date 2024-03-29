package cegid

// fixed:"{startPos},{endPos},[{alignment},[{padChar}]]"
type ThirdPartyAccount struct {
	FIXE               string `fixed:"1,3"`
	IDENTIFIANT        string `fixed:"4,6"`
	CODE               string `fixed:"7,23"`
	LIBELLE            string `fixed:"24,58"`
	NATURE             string `fixed:"59,61"`
	LETTRABLE          string `fixed:"62,62"`
	COLLECTIF          string `fixed:"63,79"`
	EAN                string `fixed:"80,96"`
	TABLE1             string `fixed:"97,113"`
	TABLE2             string `fixed:"114,130"`
	TABLE3             string `fixed:"131,147"`
	TABLE4             string `fixed:"148,164"`
	TABLE5             string `fixed:"165,181"`
	TABLE6             string `fixed:"182,198"`
	TABLE7             string `fixed:"199,215"`
	TABLE8             string `fixed:"216,232"`
	TABLE9             string `fixed:"233,249"`
	TABLE10            string `fixed:"250,266"`
	ADRESSE1           string `fixed:"267,301"`
	ADRESSE2           string `fixed:"302,336"`
	ADRESSE3           string `fixed:"337,371"`
	CODEPOSTAL         string `fixed:"372,380"`
	VILLE              string `fixed:"381,415"`
	DOMICILIATION      string `fixed:"416,439"`
	ETABLISSEMENT      string `fixed:"440,444"`
	GUICHET            string `fixed:"445,449"`
	COMPTE             string `fixed:"450,460"`
	CLE                string `fixed:"461,462"`
	PAYS               string `fixed:"463,465"`
	LIBELLEABREGE      string `fixed:"466,482"`
	LANGUE             string `fixed:"483,485"`
	MULTIDEVISE        string `fixed:"486,486"`
	DEVISEDUTIERS      string `fixed:"487,489"`
	TELEPHONE          string `fixed:"490,514"`
	FAX                string `fixed:"515,539"`
	REGIMETVA          string `fixed:"540,542"`
	MODEREGLEMENT      string `fixed:"543,545"`
	COMMENTAIRE        string `fixed:"546,580"`
	NIF                string `fixed:"581,597"`
	SIRET              string `fixed:"598,614"`
	APE                string `fixed:"615,619"`
	PRENOM             string `fixed:"620,654"`
	CONTACTSERVICE     string `fixed:"655,689"`
	CONTACTFONCTION    string `fixed:"690,724"`
	CONTACTTELEPHONE   string `fixed:"725,749"`
	CONTACTFAX         string `fixed:"750,774"`
	CONTACTTELEX       string `fixed:"775,799"`
	CONTACTRVA         string `fixed:"800,849"`
	CONTACTCIVILITE    string `fixed:"850,852"`
	CONTACTPRINCIPAL   string `fixed:"853,853"`
	FORMEJURIDIQUE     string `fixed:"854,856"`
	RIBPRINCIPAL       string `fixed:"857,857"`
	TVAENCAISSEMENT    string `fixed:"858,860"`
	PAYEUR             string `fixed:"861,877"`
	ISPAYEUR           string `fixed:"878,878"`
	AVOIRRBT           string `fixed:"879,879"`
	RELANCEREGLEMENT   string `fixed:"880,882"`
	RELANCETRAITE      string `fixed:"883,885"`
	CONFIDENTIEL       string `fixed:"886,886"`
	CORRESP1           string `fixed:"887,903"`
	CORRESP2           string `fixed:"904,920"`
	ESCOMPTE           string `fixed:"921,940,right"`
	REMISE             string `fixed:"941,960,right"`
	FACTURE            string `fixed:"961,977"`
	JURIDIQUE          string `fixed:"978,980"`
	CREDITDEMANDE      string `fixed:"981,1000,right"`
	CREDITACCORDE      string `fixed:"1001,1020,right"`
	CREDITPLAFOND      string `fixed:"1021,1040,right"`
	FERME              string `fixed:"1041,1041"`
	FACTUREHT          string `fixed:"1042,1042"`
	SOCIETEGROUPEBE    string `fixed:"1043,1059"`
	RELANCETRAITEBE    string `fixed:"1060,1062"`
	RELANCEREGLEMENTBE string `fixed:"1063,1065"`
}
