package quickapiwendy

import (
	"encoding/json"

	"github.com/Meduzz/quickapi"
	"github.com/Meduzz/quickapi-wendy/api"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type (
	storage struct {
		entity quickapi.Entity
		db     *gorm.DB
	}
)

func newStorage(db *gorm.DB, entity quickapi.Entity) *storage {
	return &storage{entity, db}
}

func (s *storage) Create(c *api.Create) (any, error) {
	e := s.entity.Create()
	err := json.Unmarshal(c.Entity, e)

	if err != nil {
		return nil, err
	}

	err = s.db.Create(e).Error

	if err != nil {
		return nil, err
	}

	return e, nil
}

func (s *storage) Read(r *api.Read) (any, error) {
	e := s.entity.Create()
	query := s.db

	scopes := createScopes(r.Filters, s.entity.Filters())

	if scopes != nil {
		query = query.Scopes(scopes...)
	}

	err := query.First(e, r.ID).Error

	if err != nil {
		return nil, err
	}

	return e, nil
}

func (s *storage) Update(u *api.Update) (any, error) {
	e := s.entity.Create()
	err := json.Unmarshal(u.Entity, e)

	if err != nil {
		return nil, err
	}

	query := s.db.Session(&gorm.Session{FullSaveAssociations: true})
	scopes := createScopes(u.Filters, s.entity.Filters())

	if scopes != nil {
		query = query.Scopes(scopes...)
	}

	err = query.Save(e).Error

	if err != nil {
		return nil, err
	}

	return e, nil
}

func (s *storage) Delete(d *api.Delete) error {
	e := s.entity.Create()
	query := s.db.Select(clause.Associations)
	scopes := createScopes(d.Filters, s.entity.Filters())

	if scopes != nil {
		query = query.Scopes(scopes...)
	}

	err := query.Delete(e, d.ID).Error

	if err != nil {
		return err
	}

	return nil
}

func (s *storage) Search(c *api.Search) (any, error) {
	data := s.entity.CreateArray()

	query := s.db.
		Offset(c.Skip).
		Limit(c.Take)

	if len(c.Where) > 0 {
		query = query.Where(c.Where)
	}

	scopes := createScopes(c.Filters, s.entity.Filters())

	if scopes != nil {
		query = query.Scopes(scopes...)
	}

	err := query.Find(&data).Error

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *storage) Patch(p *api.Patch) (any, error) {
	e := s.entity.Create()

	err := s.db.Model(e).
		Where("id = ?", p.ID).
		Updates(p.Data).Error

	if err != nil {
		return nil, err
	}

	query := s.db
	scopes := createScopes(p.Filters, s.entity.Filters())

	if scopes != nil {
		query = query.Scopes(scopes...)
	}

	err = query.Find(e, p.ID).Error

	if err != nil {
		return nil, err
	}

	return e, nil
}

func createScopes(data map[string]map[string]string, filters []*quickapi.NamedFilter) []func(*gorm.DB) *gorm.DB {
	if len(filters) == 0 {
		return nil
	}

	scopes := make([]func(*gorm.DB) *gorm.DB, 0)

	for _, filter := range filters {
		data, ok := data[filter.Name]

		if ok {
			scopes = append(scopes, filter.Scope(data))
		}
	}

	return scopes
}
