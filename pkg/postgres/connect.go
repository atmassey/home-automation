package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
)

type Config struct {
	DriverName        string
	DSN               string
	Conn              *sql.DB
	DefaultStringSize uint
}

type PSQLConnection struct {
	Host     string
	User     string
	Password string
	DB_Name  string
	RawDb    *sql.DB
	DB       *gorm.DB
	Config   Config
}

func BuildConnectionString(host, user, password, dbname string) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=America/Chicago", host, user, password, dbname)
}

func (p *PSQLConnection) Close() error {
	return p.RawDb.Close()
}
func (p *PSQLConnection) Connect() error {
	var err error
	conn_str := BuildConnectionString(p.Host, p.User, p.Password, p.DB_Name)
	p.DB, err = gorm.Open(postgres.Open(conn_str), &gorm.Config{})
	if err != nil {
		log.Panicf("Problem connection to database %v", err)
	}
	p.RawDb, err = p.DB.DB()
	if err != nil {
		log.Panicf("Problem getting raw db connection.. %v", err)
	}

	//start go routine to keep connection alive
	go func() {
		for {
			time.Sleep(time.Minute * 1)
			err := p.RawDb.Ping()
			if err != nil {
				log.Fatalf("DB Ping failure. Attempting to reconnect..")
				// make sure the connection is closed.
				err = p.Close()
				if err != nil {
					log.Fatalf("Problem closing connection %v", err)
				}
				// then reconnect
				err = p.Connect()
				if err != nil {
					log.Fatalf("Error reconnecting to database: %v", err)
				}
				// then we can return from this function since the new connect
				// call  should start a new ping watchdog
				return
			}
		}
	}()
	p.Config.Conn = p.RawDb
	return err
}

func (p *PSQLConnection) GetTableDefinition(path string, tablename string, modelname string) {
	g := gen.NewGenerator(gen.Config{OutPath: path, Mode: gen.WithoutContext})
	g.UseDB(p.DB)
	// Need to take a model name as a argument to the GetTableDefinition function because there can no be
	// underscores in the model name.  The model name is used to generate the file name and the struct name.
	g.ApplyBasic(g.GenerateModelAs(tablename, modelname))
	g.Execute()
}
