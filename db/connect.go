package db
import (
	"detect-cycle/model"
	"fmt"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

func Connection()*pg.DB{//This indicates that the function returns a pointer to a pg.DB object,

	d_b:=pg.Connect(&pg.Options{
		Addr:"localhost:5432",
		User:"postgres",
		Password:"1234",
		Database:"postgres",
	})
	if d_b==nil{
		fmt.Println("Connection not established")
		return d_b
	}
	fmt.Println("Connection established")
	return d_b
}
func Schema(db *pg.DB)error{
	md:=[]interface{}{
		(*model.EmpManager)(nil),
	}
	for _,m:=range md{
		//db.Model specifying the model to be used for creating table

		err:=db.Model(m).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,//this ensures that the table is only created if it doesnt already exist
			
		})
		if err!=nil{
			fmt.Printf("Error while creating table")
			return err
		}
	}
	return nil
}