package main

import (
	"goblog/database"
	"goblog/routes"
	"goblog/utils"
)

func dfs(candidates []int, idx, target, sum int, results []int) {
	if idx >= len(candidates) {
		return
	}
	if sum == target {

	}
	results = append(results, targ)
}

func combinationSum2(candidates []int, target int) [][]int {

}

func main() {
	utils.InitConfig()
	database.InitDatabase()
	routes.InitRouter()
}
