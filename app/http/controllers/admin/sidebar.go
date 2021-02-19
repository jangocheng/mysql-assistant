package admin

import (
	"github.com/gin-gonic/gin"
	"owen2020/app/models"
	"owen2020/app/models/dao"
	"owen2020/app/resp/out"
	"strings"
)

// 左侧菜单树
type menuList []models.Menu

type menuItem struct {
	models.Menu
	Child []menuItem `json:"_child"`
}

// 左侧菜单树 -end

// 菜单选择列表
type menuSelectItem struct {
	models.Menu
	Level     int    `json:"level"`
	TitleShow string `json:"title_show"`
}

// 菜单选择列表 -end

//GetSideBarTree 获取无限级菜单
func GetSideBarTree(c *gin.Context) {
	list := dao.GetMenuList()
	ml := menuList(list)
	//ret := processToTree(list, 0, 0)
	ret := ml.processToTree(0, 0)
	out.NewSuccess(ret).JSONOK(c)
}

func (m *menuList) processToTree(pid int, level int) []menuItem {
	var mTree []menuItem
	if level == 10 {
		return mTree
	}

	list := m.findChildren(pid)
	if len(list) == 0 {
		return mTree
	}

	for _, v := range list {
		child := m.processToTree(v.MenuId, level+1)
		mTree = append(mTree, menuItem{v, child})
	}

	return mTree
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

// 关于子结构体转父结构体的说明
//https://stackoverflow.com/questions/46989247/how-to-convert-parent-type-to-child-type
func (m menuItem) ConvertToMenu() models.Menu {
	menu := models.Menu{}

	menu.MenuId = m.MenuId
	menu.Title = m.Title
	menu.Pid = m.Pid
	menu.Sort = m.Sort
	menu.Hide = m.Hide
	menu.Pathname = m.Pathname
	menu.Iconfont = m.Iconfont
	menu.CreatedAt = m.CreatedAt
	menu.UpdatedAt = m.UpdatedAt
	menu.IsDeleted = m.IsDeleted
	menu.DeletedAt = m.DeletedAt

	return menu
}

// -- 菜单选择列表 --
func TreeToSelect(msP *[]menuSelectItem, mt []menuItem, level int) {
	for _, v := range mt {
		prefixStr := strings.Repeat("&nbsp;", level*2)
		prefixStr += "└"
		showTitle := prefixStr + v.Menu.Title
		item := menuSelectItem{v.ConvertToMenu(), level, showTitle}
		*msP = append(*msP, item)
		if len(v.Child) > 0 {
			TreeToSelect(msP, v.Child, level+1)
		}
	}
}

//
//func menuTreeToSelect(tree []models.Menu, level int) {
//	for _, v := range tree {
//		prefixStr := strings.Repeat("&nbsp;", level*2)
//		prefixStr += "└"
//		item := menuSelectItem{v, level, prefixStr + v.Title}
//
//	}
//}

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
