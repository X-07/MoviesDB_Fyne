package importing

type countryXML struct {
	Lines []lineXML `xml:"line"`
}

type genreXML struct {
	Lines []lineXML `xml:"line"`
}

type actorsXML struct {
	Lines []line2XML `xml:"line"`
}

type audioXML struct {
	Lines []line2XML `xml:"line"`
}

type subtXML struct {
	Lines []lineXML `xml:"line"`
}

type lineXML struct {
	Col string `xml:"col"`
}

type line2XML struct {
	Col []string `xml:"col"`
}
