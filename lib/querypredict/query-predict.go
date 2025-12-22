package querypredict

type QuerySpecColumn struct {
	Name string `yaml:"name"`
	Type string `yaml:"type"`
	From string `yaml:"from"`
}

type QuerySpec struct {
	Name  string `yaml:"name"`
	Query string `yaml:"query"`
	Force bool   `yaml:"force"`
}

type QueryDocument struct {
	Queries []QuerySpec `yaml:"queries"`
}
