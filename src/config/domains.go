package config

type Domain struct {
	Title string
	URL   string
	Tag   string
}

var domains = []Domain{
	{
		Title: "Progy",
		URL:   "http://progy.com.br",
		Tag:   ".text-3xl",
	},
	{
		Title: "TechCrunch",
		URL:   "https://techcrunch.com/2023/08/21/b2b-inventory-marketplace-ghost-30m/",
		Tag:   "article.article-container.article--post",
	},
}

func GetDomains() []Domain {
	return domains
}
