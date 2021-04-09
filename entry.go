package cegid

// fixed:"{startPos},{endPos},[{alignment},[{padChar}]]"
type Entry struct {
	JOURNAL             string `fixed:"1,3"`
	DATECOMPTABLE       string `fix:"4,11"`
	TYPEPIECE           string `fix:"12,13"`
	GENERAL             string `fix:"14,30"`
	TYPECPTE            string `fix:"31,31"`
	AUXILIAIREOUSECTION string `fix:"32,48"`
	REFINTERNE          string `fix:"49,83"`
	LIBELLE             string `fix:"84,118"`
	MODEPAIE            string `fix:"119,121"`
	ECHEANCE            string `fix:"122,129"`
	SENS                string `fix:"130,130"`
	MONTANT1            string `fix:"131,150"`
	TYPE_ECRITURE       string `fix:"151,151"`
	NUMEROPIECE         string `fix:"152,159"`
	DEVISE              string `fix:"160,162"`
	TAUXDEV             string `fix:"163,172"`
	CODEMONTANT         string `fix:"173,175"`
	MONTANT2            string `fix:"176,195"`
	MONTANT3            string `fix:"196,215"`
	ETABLISSEMENT       string `fix:"216,218"`
	AXE                 string `fix:"219,220"`
	NUMECHE             string `fix:"221,222"`
}
