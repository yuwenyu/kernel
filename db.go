/**
 * Copyright 2019 YuwenYu.  All rights reserved.
**/

package kernel

import (
	"fmt"
	"log"
	"strconv"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var dbEngines map[int]db

type DB interface {
	Engine() *xorm.Engine
}

type db struct {
	id     int
	engine *xorm.Engine
	mx     sync.Mutex
}

var _ DB = &db{}

func NewDB(src int) *db {
	if src <= 0 {
		src = 1
	}

	var object *db
	if v, ok := dbEngines[src]; ok {
		object = &v
	} else {
		object = &db{id: src}
	}

	return object
}

func (odbc *db) Engine() *xorm.Engine {
	if odbc.engine == nil {
		odbc.instanceMaster()
	}

	return odbc.engine
}

func (odbc *db) instanceMaster() *db {
	odbc.mx.Lock()
	defer odbc.mx.Unlock()

	if odbc.engine != nil {
		return odbc
	}

	if len(dbEngines) == 0 {
		dbEngines = make(map[int]db)
	} else {
		if v, ok := dbEngines[odbc.id]; ok {
			if odbc.engine == nil {
				odbc.engine = v.engine
			}

			return odbc
		}
	}

	ds := odbc.initDataSource()
	driverSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		ds.username, ds.password, ds.host, ds.port, ds.table)
	engine, err := xorm.NewEngine(ds.dn, driverSource)

	if err != nil {
		log.Fatalf("db.DbInstanceMaster,", err)
		return nil
	}

	engine.SetMaxOpenConns(ds.maxOpen)
	engine.SetMaxIdleConns(ds.maxIdle)

	engine.ShowSQL(ds.showedSQL)
	engine.SetTZDatabase(SysTimeLocation)

	if ds.cachedSQL {
		cached := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)
		engine.SetDefaultCacher(cached)
	}

	odbc.engine = engine

	dbEngines[odbc.id] = *odbc

	return odbc
}

type dataSource struct {
	dn        string
	host      string
	port      int
	table     string
	username  string
	password  string
	maxOpen   int
	maxIdle   int
	showedSQL bool
	cachedSQL bool
}

func (odbc *db) initDataSource() *dataSource {
	var section string = MapConfLists[ConfDB][0] + StrUL + strconv.Itoa(odbc.id)
	var c INI = NewIni().LoadByFN(ConfDB)

	dn 		:= c.K(section, MapConfParam[MapConfLists[ConfDB][0]][0]).String()
	host 	:= c.K(section, MapConfParam[MapConfLists[ConfDB][0]][1]).String()
	table 	:= c.K(section, MapConfParam[MapConfLists[ConfDB][0]][3]).String()
	username:= c.K(section, MapConfParam[MapConfLists[ConfDB][0]][4]).String()
	password:= c.K(section, MapConfParam[MapConfLists[ConfDB][0]][5]).String()

	port, errPort	:= c.K(section, MapConfParam[MapConfLists[ConfDB][0]][2]).Int()
	if errPort != nil {port = KDbPort}

	maxOpen, errOpen:= c.K(section, MapConfParam[MapConfLists[ConfDB][0]][6]).Int()
	if errOpen != nil {maxOpen = KDbMaxOpen}

	maxIdle, errIdle:= c.K(section, MapConfParam[MapConfLists[ConfDB][0]][7]).Int()
	if errIdle != nil {maxIdle = KDbMaxIdle}

	showedSQL, errShowedSQL := c.K(section, MapConfParam[MapConfLists[ConfDB][0]][8]).Bool()
	if errShowedSQL != nil {showedSQL = KDbShowedSQL}

	cachedSQL, errCachedSQL := c.K(section, MapConfParam[MapConfLists[ConfDB][0]][9]).Bool()
	if errCachedSQL != nil {cachedSQL = KDbCachedSQL}

	return &dataSource{
		dn:        dn,
		host:      host,
		port:      port,
		table:     table,
		username:  username,
		password:  password,
		maxOpen:   maxOpen,
		maxIdle:   maxIdle,
		showedSQL: showedSQL,
		cachedSQL: cachedSQL,
	}
}
