package blog

import "newsCaptors/models"

type IndexController struct {
	baseController
}

func (this *IndexController) Index(){
	captorModel := new(models.CaptorInfo)
	var dataList []*models.CaptorInfo
	ok := captorModel.FindAll(&dataList)
	if !ok {
		this.Ctx.WriteString("no data")
		return
	}

	this.Data["count"] = len(dataList)
	this.Data["dataList"] = dataList
	this.setHeadMetas()
	this.display("index")
}