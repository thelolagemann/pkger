package stuffing

import (
	"io"

	"github.com/markbates/pkger/here"
	"github.com/markbates/pkger/pkging"
	"github.com/markbates/pkger/pkging/mem"
	"github.com/markbates/pkger/pkging/stdos"
)

func Stuff(w io.Writer, cur here.Info, paths []pkging.Path) error {
	disk, err := stdos.New()
	if err != nil {
		return err
	}

	pkg, err := mem.New(cur)
	if err != nil {
		return err
	}

	for _, pt := range paths {
		err = func() error {
			f, err := disk.Open(pt.String())
			if err != nil {
				return err
			}
			defer f.Close()

			fi, err := f.Stat()
			if err != nil {
				return err
			}

			if err := pkg.Add(fi, f); err != nil {
				return err
			}

			return nil
			// WithInfo(ng, og)
		}()
		if err != nil {
			return err
		}
	}

	b, err := pkg.MarshalEmbed()
	if err != nil {
		return err
	}
	_, err = w.Write(b)
	return err
}