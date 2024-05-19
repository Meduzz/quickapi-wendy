package quickapiwendy

import (
	"github.com/Meduzz/helper/block"
	"github.com/Meduzz/helper/nuts"
	"github.com/Meduzz/quickapi"
	"github.com/Meduzz/wendy"
	wendyrpc "github.com/Meduzz/wendy-rpc"
	"gorm.io/gorm"
)

// Run - exposes the provided entities over nats. Queue is used to queuesubscribe if provided. Prefix to prefix topic bound to if provided.
func Run(db *gorm.DB, queue, prefix string, entities ...quickapi.Entity) error {
	nats, err := nuts.Connect()

	if err != nil {
		return err
	}

	migrations := make([]any, 0)
	modules := make([]*wendy.Module, 0)

	for _, e := range entities {
		if e.Name() != "" {
			migrations = append(migrations, e.Create())
			modules = append(modules, For(db, e))
		}
	}

	err = db.AutoMigrate(migrations...)

	if err != nil {
		return err
	}

	err = wendyrpc.ServeModules(nats, queue, prefix, modules...)

	if err != nil {
		return err
	}

	return block.Block(func() error {
		return nats.Drain()
	})
}

// For - turn provided entity into a wendy.Module.
func For(db *gorm.DB, entity quickapi.Entity) *wendy.Module {
	m := wendy.NewModule(entity.Name())
	s := NewStorage(db, entity)
	h := NewHandler(s)

	m.WithHandler("create", h.Create)
	m.WithHandler("read", h.Read)
	m.WithHandler("update", h.Update)
	m.WithHandler("delete", h.Delete)
	m.WithHandler("search", h.Search)
	m.WithHandler("patch", h.Patch)

	return m
}
