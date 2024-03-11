package main

import (
	"detect-cycle/db"
	"detect-cycle/model"
	"fmt"

	"github.com/go-pg/pg/v10"
)
func hasCycle(graph map[int][]int,empId int,vis,curr map[int]bool)bool{
	vis[empId]=true;
	curr[empId]=true;
	for _,manId:=range graph[empId]{
		if !vis[manId]{
			if hasCycle(graph,manId,vis,curr){
				return true;
			}
		}else if curr[manId]{
			return true;
		}
	}
	curr[empId]=false//backtracking
	return false;
}
func Check(db *pg.DB, empId int, mangId int) bool{
	//retreive all existing relationships from the database
	
	var rows []model.EmpManager
	if err:=db.Model(&rows).Select();err!=nil{
		fmt.Printf("Error in querying database:%v\n",err)
        return false
	}
	//creating graph from the data fetched
	graph:=make(map[int][]int)
	for _,row:=range rows{
		graph[row.Eid]=append(graph[row.Eid],row.Mid)
	}
	graph[empId]=append(graph[empId], mangId)
	//check for cycle
	vis:=make(map[int]bool)
	curr:=make(map[int]bool)
	for node:=range graph{
		if !vis[node]{
			if hasCycle(graph,node,vis,curr){
				return true//cycle detected

			}
		}
	}
    //no cycle detected
	return false



}
func main() {
	pg_db := db.Connection()
	if pg_db == nil {
		fmt.Println("Failed to connect to the database")
		return
	}
	err := db.Schema(pg_db)
	if err != nil {
		fmt.Println("Error creating database schema", err)
		return

	}
	fmt.Println("Database schema created successfully")
	var empId, mangId int
	fmt.Println("Enter the employeed id:")
	fmt.Scanln(&empId)
	fmt.Println("Enter the manager id:")
	fmt.Scanln(&mangId)

	if Check(pg_db, empId, mangId){
		fmt.Println("Cycle detected")
	}else{
		fmt.Println("Cycle not present!!!!")
	}
}
