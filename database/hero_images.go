package database

import (
	"github.com/CollCaz/UniSite/database/gen/model"
	t "github.com/CollCaz/UniSite/database/gen/table"
	s "github.com/go-jet/jet/v2/sqlite"
)

type HeroImage struct {
	ImageUrl string `validate:"required,url"`
	Title    string
	SubTitle string `validate:"required_with=Title"`
}

type HeroImages []HeroImage

type GetHeroImagesArgs struct {
	Amount int
}

func (d *DataService) GetHeroImages(args GetHeroImagesArgs) (HeroImages, error) {
	stmt := s.
		SELECT(t.HeroImages.ImageURL, t.HeroImages.Title, t.HeroImages.SubTitle).
		FROM(t.HeroImages).
		ORDER_BY(t.HeroImages.UpdatedAt.DESC()).
		LIMIT(int64(args.Amount))

	dest := []model.HeroImages{}

	err := stmt.Query(d.db, &dest)
	if err != nil {
		d.logger.Error(err.Error())
		return HeroImages{}, err
	}

	res := HeroImages{}
	for _, img := range dest {
		res = append(res, HeroImage{
			Title:    img.Title,
			ImageUrl: img.ImageURL,
			SubTitle: img.SubTitle,
		})
	}

	return res, nil

}

func (d *DataService) InsertHeroImage(args HeroImage) (HeroImage, error) {
	stmt := t.HeroImages.
		INSERT(
			t.HeroImages.Title,
			t.HeroImages.SubTitle,
			t.HeroImages.ImageURL,
		).VALUES(
		args.Title,
		args.SubTitle,
		args.ImageUrl,
	).RETURNING(
		t.HeroImages.Title,
		t.HeroImages.SubTitle,
		t.HeroImages.ImageURL,
	)

	dest := model.HeroImages{}

	err := stmt.Query(d.db, &dest)
	if err != nil {
		d.logger.Error(err.Error())
		return HeroImage{}, err
	}

	res := HeroImage{
		Title:    dest.Title,
		SubTitle: dest.SubTitle,
		ImageUrl: dest.ImageURL,
	}

	return res, nil
}
