package cegid

// fixed:"{startPos},{endPos},[{alignment},[{padChar}]]"
type Entry struct {
	JOURNAL               string `fixed:"1,3"`
	DATECOMPTABLE         string `fixed:"4,11"`
	TYPE_PIECE            string `fixed:"12,13"`
	GENERAL               string `fixed:"14,30"`
	TYPE_CPTE             string `fixed:"31,31"`
	AUXILIAIRE_OU_SECTION string `fixed:"32,48"`
	REFINTERNE            string `fixed:"49,83"`
	LIBELLE               string `fixed:"84,118"`
	MODEPAIE              string `fixed:"119,121"`
	ECHEANCE              string `fixed:"122,129"`
	SENS                  string `fixed:"130,130"`
	MONTANT1              string `fixed:"131,150,right"`
	TYPE_ECRITURE         string `fixed:"151,151"`
	NUMEROPIECE           string `fixed:"152,159"`
	DEVISE                string `fixed:"160,162"`
	TAUXDEV               string `fixed:"163,172"`
	CODEMONTANT           string `fixed:"173,175"`
	MONTANT2              string `fixed:"176,195,right"`
	MONTANT3              string `fixed:"196,215,right"`
	ETABLISSEMENT         string `fixed:"216,218"`
	AXE                   string `fixed:"219,220"`
	NUMECHE               string `fixed:"221,222"`
}
