package commands

import (
	"fmt"
	"github.com/araddon/dateparse"
	"github.com/photoprism/photoprism/internal/photoprism"
	"github.com/urfave/cli"
	"log"
)

var ExportCommand = cli.Command{
	Name:   "export",
	Usage:  "Exports photos as JPEG",
	Flags:  exportFlags,
	Action: exportAction,
}

var exportFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "name, n",
		Usage: "Sub-directory name",
	},
	cli.StringFlag{
		Name:  "after, a",
		Usage: "Start date e.g. 2017/04/15",
	},
	cli.StringFlag{
		Name:  "before, b",
		Usage: "End date e.g. 2018/05/02",
	},
	cli.IntFlag{
		Name:  "size, s",
		Usage: "Image size in pixels",
		Value: 2560,
	},
}

func exportAction(context *cli.Context) error {
	conf := photoprism.NewConfig(context)

	if err := conf.CreateDirectories(); err != nil {
		log.Fatal(err)
	}

	before := context.String("before")
	after := context.String("after")

	if before == "" || after == "" {
		fmt.Println("You need to provide before and after dates for export, e.g.\n\nphotoprism export --after 2018/04/10 --before '2018/04/15 23:00:00'")

		return nil
	}

	afterDate, _ := dateparse.ParseAny(after)
	beforeDate, _ := dateparse.ParseAny(before)
	afterDateFormatted := afterDate.Format("20060102")
	beforeDateFormatted := beforeDate.Format("20060102")

	name := context.String("name")

	if name == "" {
		if afterDateFormatted == beforeDateFormatted {
			name = beforeDateFormatted
		} else {
			name = fmt.Sprintf("%s - %s", afterDateFormatted, beforeDateFormatted)
		}
	}

	exportPath := fmt.Sprintf("%s/%s", conf.ExportPath, name)
	size := context.Int("size")
	originals := photoprism.FindOriginalsByDate(conf.OriginalsPath, afterDate, beforeDate)

	fmt.Printf("Exporting photos to %s...\n", exportPath)

	photoprism.ExportPhotosFromOriginals(originals, conf.ThumbnailsPath, exportPath, size)

	fmt.Println("Done.")

	return nil
}
