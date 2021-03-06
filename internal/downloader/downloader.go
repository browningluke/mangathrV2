package downloader

import (
	"fmt"
	"github.com/browningluke/mangathrV2/internal/logging"
	"github.com/browningluke/mangathrV2/internal/metadata"
	"github.com/browningluke/mangathrV2/internal/sources/structs"
	"github.com/schollz/progressbar/v3"
	"os"
	"time"
)

type Downloader struct {
	config *Config
	agent  metadata.Agent

	updateMode bool

	enforceChapterDuration bool
	chapterDuration        int64
}

type Job struct {
	Chapter structs.Chapter
	Bar     *progressbar.ProgressBar
}

func NewDownloader(config *Config,
	updateMode bool,
	enforceChapterDuration bool) *Downloader {
	return &Downloader{
		config:                 config,
		updateMode:             updateMode,
		enforceChapterDuration: enforceChapterDuration,
	}
}

func (d *Downloader) MetadataAgent() *metadata.Agent {
	d.agent = metadata.NewAgent(d.config.Metadata.Agent)
	return &d.agent
}

func (d *Downloader) SetChapterDuration(duration int64) {
	d.chapterDuration = duration
}

func (d *Downloader) SetTemplate(template string) {
	if template != "" {
		d.config.Output.FilenameTemplate = template
	}
}

/*
	-- Chapter Downloading --
*/

func (d *Downloader) CanDownload(path, filename string) *logging.ScraperError {
	chapterPath := d.GetChapterPath(path, filename)

	if d.config.DryRun {
		return &logging.ScraperError{
			Error:   fmt.Errorf("called with dryRun set to true"),
			Message: "DRY RUN",
			Code:    0,
		}
	}

	if _, err := os.Stat(chapterPath); err == nil {
		return &logging.ScraperError{
			Error:   fmt.Errorf("file exists at path %s", chapterPath),
			Message: "chapter already exists",
			Code:    0,
		}
	}

	return nil
}

// Download chapter. Assumes CanDownload() has been called and has returned true
func (d *Downloader) Download(path, filename string, pages []Page, bar *progressbar.ProgressBar) error {

	// Ensure chapter time is correct
	if d.enforceChapterDuration {
		timeStart := time.Now().UnixMilli()
		defer d.waitChapterDuration(timeStart)
	} else {
		// TODO: differentiate between Download & Update delay
		dur, err := time.ParseDuration(d.config.Delay.Chapter)
		if err != nil {
			return err
		}
		time.Sleep(dur)
	}

	chapterPath := d.GetChapterPath(path, filename)

	if d.config.Output.Zip {
		return d.downloadZip(pages, chapterPath, bar)
	} else {
		return d.downloadDir(pages, chapterPath, bar)
	}
}
