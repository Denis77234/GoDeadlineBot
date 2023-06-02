package postgres

import (
	"database/sql"
	"log"
	"time"
	"katun/internal/database/member"

	_ "github.com/jackc/pgx/stdlib"
)



type Postgres struct{
	Pdb *sql.DB
}


func New(config Config) (*Postgres, error) {

	db, err:= sql.Open("pgx", config.ConnString())
	if err!=nil{
		return &Postgres{}, err
	}

	return &Postgres{Pdb: db}, nil
}

func (p *Postgres) AddVkat(userId int64, userTag string, time time.Time) (int, error){
	var insertedId int
	err:= p.Pdb.QueryRow("INSERT INTO vkat (user_id, start_time, finish_time, user_tag) VALUES ($1, NOW(), $2, $3) RETURNING vkat_id", userId, time, userTag).Scan(&insertedId)
	if err!=nil{
		return 0, err
	}
	return insertedId, nil
}

func (p *Postgres) Time(user int64) (string, string, error){
	var start string
	var finish string

	err:= p.Pdb.QueryRow("SELECT start_time, finish_time FROM vkat WHERE user_id = $1", user).Scan(&start, &finish)
	if err!= nil{
		return "", "", err
	}
	return start, finish, nil
}

func (p *Postgres) VkatMembers() ([]member.Member, error){
	memberArr:= make([]member.Member,0,10)

	rows,err:= p.Pdb.Query("SELECT user_id, start_time, finish_time, user_tag FROM vkat")
	if err!=nil{
		log.Println(err)
		return memberArr, err
	}

	defer rows.Close()

	for rows.Next() {
		var member member.Member

		rows.Scan(&member.UserId, &member.Start, &member.Finish, &member.UserTag)

		memberArr= append(memberArr, member)
	}


	return memberArr, nil
}

func (p *Postgres) DeleteVkat(userid int64) (error){
	var id int
	err:= p.Pdb.QueryRow("DELETE FROM vkat WHERE user_id = $1 RETURNING vkat_id", userid).Scan(&id)
	if err!=nil{
		log.Println(err)
		return err
	}
	return nil
}

func (p *Postgres) UpdateVkat(userid int64, time time.Time) (error) {
	var id int
	err:= p.Pdb.QueryRow("UPDATE vkat SET finish_time = $1 WHERE user_id= $2",time ,userid).Scan(&id)
	if err!=nil{
		log.Println(err)
		return err
	}
	return nil
}