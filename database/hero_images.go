package database

import (
	"github.com/CollCaz/UniSite/database/gen/unicontentdb/public/model"
	t "github.com/CollCaz/UniSite/database/gen/unicontentdb/public/table"
	s "github.com/go-jet/jet/v2/postgres"
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

type joinedHeroImageModel struct {
	model.HeroImages
	model.Image
}

func (d *DataService) scanHeroImages(stmt s.Statement) (HeroImages, error) {
	dest := []joinedHeroImageModel{}
	err := stmt.Query(d.db, &dest)
	if err != nil {
		return HeroImages{}, err
	}

	res := HeroImages{}
	for _, img := range dest {
		res = append(res, HeroImage{
			Title:    img.HeroImages.Title,
			ImageUrl: img.ImageURL,
			SubTitle: img.SubTitle,
		})
	}

	return res, nil
}

func (d *DataService) scanHeroImage(stmt s.Statement) (HeroImage, error) {
	dest := joinedHeroImageModel{}
	err := stmt.Query(d.db, &dest)
	if err != nil {
		return HeroImage{}, err
	}

	heroImage := HeroImage{
		Title:    dest.HeroImages.Title,
		ImageUrl: dest.ImageURL,
		SubTitle: dest.SubTitle,
	}

	return heroImage, nil
}

func (d *DataService) GetHeroImages(args GetHeroImagesArgs) (HeroImages, error) {
	stmt := s.
		SELECT(t.HeroImages.Title, t.HeroImages.SubTitle, t.Image.ImageURL).
		FROM(
			t.HeroImages,
			t.HeroImages.INNER_JOIN(t.Image, t.HeroImages.ImageID.EQ(t.Image.ID)),
		).
		ORDER_BY(t.HeroImages.UpdatedAt.DESC()).
		LIMIT(int64(args.Amount))

	return d.scanHeroImages(stmt)

}

func (d *DataService) InsertHeroImage(args HeroImage) (HeroImage, error) {
	cte := t.Image.
		INSERT(
			t.Image.Title,
			t.Image.ImageURL,
		).VALUES(
		args.Title,
		args.ImageUrl,
	).RETURNING(
		t.Image.ID,
	)

	stmt := t.HeroImages.
		INSERT(
			t.HeroImages.Title,
			t.HeroImages.SubTitle,
			t.HeroImages.ImageID,
		).VALUES(
		args.Title,
		args.SubTitle,
		cte,
	).RETURNING(
		t.HeroImages.Title,
		t.HeroImages.SubTitle,
		t.HeroImages.ImageID,
	)

	return d.scanHeroImage(stmt)
}
