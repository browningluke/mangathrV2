package mangadex

import (
	"fmt"
	"github.com/browningluke/mangathrV2/internal/downloader"
	"github.com/browningluke/mangathrV2/internal/logging"
	"github.com/browningluke/mangathrV2/internal/sources/structs"
	"github.com/browningluke/mangathrV2/internal/utils"
	"unicode/utf8"
)

func calculateDuration(numChapters int) int64 {
	// 60 seconds / CHAPTERSPERMIN = x = seconds per chapter
	// x * 1000 = milliseconds per chapter

	duration := int64((60 / CHAPTERSPERMIN) * 1000)
	if numChapters < CHAPTERSPERMIN {
		// Not going to exceed limit during batch, duration doesn't matter
		duration = int64(500)
	}
	return duration
}

func buildDownloadQueue(selectedChapters []structs.Chapter) (jobs []downloader.Job, maxRuneCount int) {
	var downloadQueue []downloader.Job
	maxRC := 0 // Used for padding (e.g. Chapter 10 vs Chapter 10.5)
	for _, chapter := range selectedChapters {
		downloadQueue = append(downloadQueue, downloader.Job{Chapter: chapter})

		// Check if string length is max in list
		if runeCount := utf8.RuneCountInString(chapter.Metadata.Num); runeCount > maxRC {
			maxRC = runeCount
		}
	}
	return downloadQueue, maxRC
}

func (m *Scraper) runDownloadJob(job downloader.Job, dl *downloader.Downloader,
	path string, maxRuneCount int) *logging.ScraperError {

	// Get chapter pages
	pages, err := m.getChapterPages(job.Chapter.ID)
	if err != nil {
		return err
	}

	// Initialize progress bar
	progress := utils.CreateProgressBar(len(pages), maxRuneCount, job.Chapter.Metadata.Num)

	// Get chapter filename
	dl.SetTemplate(m.config.FilenameTemplate)
	filename := dl.GetNameFromTemplate(job)

	// Set MetadataAgent values
	(*dl.MetadataAgent()).
		SetFromStruct(job.Chapter.Metadata).
		SetPageCount(len(pages))

	// Check if download is possible
	err = dl.CanDownload(path, filename)
	if err != nil {
		return err
	}

	downloadErr := dl.Download(path, filename, pages, progress)
	if downloadErr != nil {
		if err := dl.Cleanup(path, filename); err != nil {
			logging.Errorln(err)
			fmt.Printf("An error occurred when deleting failed chapter: %s", filename)
		}
		return &logging.ScraperError{
			Error:   downloadErr,
			Message: "An error occurred while downloading a page",
			Code:    0,
		}
	}

	fmt.Println("") // Create a new bar for each chapter
	return nil
}

// Download selected chapters. Handles errors itself. Returns array of chapters that succeeded
func (m *Scraper) Download(dl *downloader.Downloader, downloadType string) []structs.Chapter {
	logging.Debugln("Downloading...")

	dl.SetChapterDuration(calculateDuration(len(m.selectedChapters)))

	// downloadType is one of ["download", "update"]
	path := dl.CreateDirectory(m.manga.title, downloadType)

	downloadQueue, maxRuneCount := buildDownloadQueue(m.selectedChapters)

	// Execute download queue, potential to add workerpool here later
	var succeededChapters []structs.Chapter
	for _, job := range downloadQueue {
		err := m.runDownloadJob(job, dl, path, maxRuneCount)

		// Print error to screen, abandon chapter, and continue
		if err != nil {
			logging.Errorln(err.Error)
			fmt.Printf("Chapter %s skipping... reason: %s\n", job.Chapter.Metadata.Num, err.Message)
			continue
		}

		succeededChapters = append(succeededChapters, job.Chapter)
	}

	return succeededChapters
}
