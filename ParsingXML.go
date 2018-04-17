package main

import ("fmt"
		"encoding/xml")


var washPostXML = []byte(`
	<sitemapindex>
	<sitemap>
	   <loc>http://www.washingtonpost.com/news-politics-sitemap.xml</loc>
	</sitemap>
	<sitemap>
	   <loc>http://www.washingtonpost.com/news-blogs-politics-sitemap.xml</loc>
	</sitemap>
	<sitemap>
	   <loc>http://www.washingtonpost.com/news-opinions-sitemap.xml</loc>
	</sitemap>
 </sitemapindex>
`)


type SitemapIndex struct{
	Location []Location `xml:"sitemap"`
}
// [5 5]type == array
// []type == slice

type Location struct{
	Loc string `xml:"loc"`
}

func (l Location) String() string{
	return fmt.Sprintf(l.Loc)
}

func main(){
	bytes := washPostXML
	var s SitemapIndex
	xml.Unmarshal(bytes, &s)
	fmt.Println(s.Location)
}