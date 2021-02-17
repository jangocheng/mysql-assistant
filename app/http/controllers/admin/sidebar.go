package admin

import (
	"fmt"
	"owen2020/app/models"
	"owen2020/app/models/dao"
	"owen2020/app/resp/out"

	"github.com/gin-gonic/gin"
)

type menuList []models.Menu

type menuItem struct {
	models.Menu
	Child []menuItem `json:"_child"`
}

//GetSideBarTree 获取无限级菜单
func GetSideBarTree(c *gin.Context) {
	fmt.Println("get sidebar tree")
	list := dao.GetMenuList()
	ml := menuList(list)
	//ret := processToTree(list, 0, 0)
	ret := ml.processToTree(0, 0)
	out.NewSuccess(ret).JSONOK(c)
}

func (m *menuList) processToTree(pid int, level int) []menuItem {
	var menuTree []menuItem
	if level == 10 {
		return menuTree
	}

	list := m.findChildren(pid)
	if len(list) == 0 {
		return menuTree
	}

	for _, v := range list {
		child := m.processToTree(v.MenuId, level+1)
		menuTree = append(menuTree, menuItem{v, child})
	}

	return menuTree
}

func (m *menuList) findChildren(pid int) []models.Menu {
	child := []models.Menu{}

	for _, v := range *m {
		if v.Pid == pid {
			child = append(child, v)
		}
	}
	return child
}

//func processToTree(menuList []models.Menu, pid int, level int) []menuTreeType {
//	var menuTree []menuTreeType
//	if level == 10 {
//		return menuTree
//	}
//
//	list := findChildren(menuList, pid)
//	if len(list) == 0 {
//		return menuTree
//	}
//
//	for _, v := range list {
//		child := processToTree(menuList, v.MenuId, level+1)
//		menuTree = append(menuTree, menuTreeType{v, child})
//	}
//
//	return menuTree
//}
//
//func findChildren(menuList []models.Menu, pid int) []models.Menu {
//	child := []models.Menu{}
//
//	for _, v := range menuList {
//		if v.Pid == pid {
//			child = append(child, v)
//		}
//	}
//	return child
//}
