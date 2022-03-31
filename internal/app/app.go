package app

import (
	"archive/zip"
	"context"
	"fmt"
	"io"
	"io/fs"
	"path"
	"strconv"
	"time"

	"github.com/edanko/gen2dxf/pkg/draw"
	"github.com/edanko/gen2dxf/pkg/store"

	"github.com/edanko/gen"
	"golang.org/x/sync/errgroup"
)

type App struct {
	Params *Params
	Writer io.Writer
}

type Params struct {
	Workers     int    `mapstructure:"workers"`
	CustomOrder string `mapstructure:"custom-order"`

	MinLength      int  `mapstructure:"min-length"`
	SpecialChars   bool `mapstructure:"special-chars"`
	CapitalLetters bool `mapstructure:"capital-letters"`
}

func New(writer io.Writer) *App {
	return &App{
		Params: &Params{},
		Writer: writer,
	}
}

type result struct {
	key  string
	part *gen.PartData
}

func (a *App) Convert(ctx context.Context, source fs.FS) error {
	g, ctx := errgroup.WithContext(ctx)
	files := make(chan string)

	g.Go(func() error {
		defer close(files)
		return fs.WalkDir(source, ".", func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if d.IsDir() {
				return nil
			}
			select {
			case files <- path:
			case <-ctx.Done():
				return ctx.Err()
			}
			return nil
		})
	})

	c := make(chan result)
	for i := 0; i < a.Params.Workers; i++ {
		g.Go(func() error {
			for file := range files {
				f, err := source.Open(file)
				if err != nil {
					return err
				}
				defer f.Close()

				partMap := gen.ReadPlate(f)
				if g == nil {
					return nil
				}
				for key, part := range partMap {
					select {
					case c <- result{key, part}:
					case <-ctx.Done():
						return ctx.Err()
					}
				}
			}
			return nil
		})
	}
	go func() {
		g.Wait()
		close(c)
	}()

	parts := store.New()
	for r := range c {
		parts.Store(r.key, r.part)
	}

	if err := g.Wait(); err != nil {
		return err
	}

	zipWriter := zip.NewWriter(a.Writer)

	for _, k := range parts.Keys() {
		p, ok := parts.Load(k)
		if !ok {
			continue
		}

		if a.Params.CustomOrder != "" {
			p.ShipNo = a.Params.CustomOrder
		}

		dir := time.Now().Format("06.01.02") + " " + p.BlockNo + " dxf"
		subdir := strconv.FormatFloat(p.Thickness, 'g', -1, 64)
		dir = path.Join(dir, subdir)
		file := fmt.Sprintf("%s-%s-%s.dxf", p.ShipNo, p.BlockNo, p.PosNo)
		out := path.Join(dir, file)

		w, err := zipWriter.Create(out) // "csv/test.csv")
		if err != nil {
			return err
		}

		drawer := draw.NewDrawer(w)
		err = drawer.PartToDXF(p)
		if err != nil {
			return err
		}

	}

	return zipWriter.Close()
}

func (a *App) validateParams() error {
	params := a.Params
	if params.MinLength < 8 {
		return fmt.Errorf("min-length (%d) must not be less than 8", params.MinLength)
	}
	return nil
}
