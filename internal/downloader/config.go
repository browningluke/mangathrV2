package downloader

type Config struct {
	DryRun            bool `yaml:"dryRun"`
	SimultaneousPages int  `yaml:"simultaneousPages"`
	PageRetries       int  `yaml:"pageRetries"`
	Delay             struct {
		Page          string
		Chapter       string
		UpdateChapter string `yaml:"updateChapter"`
	}
	Output struct {
		Path             string
		UpdatePath       string `yaml:"updatePath"`
		Zip              bool
		FilenameTemplate string `yaml:"filenameTemplate"`
	}
	Metadata struct {
		Agent    string
		Location string
	}
}
