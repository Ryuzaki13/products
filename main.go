package main

import (
	"github.com/gin-gonic/gin"
)

type Product struct {
	ID    int
	Name  string
	Desc  string
	Image string
	Price float64
}

var products []Product = []Product{
	Product{
		ID:    1,
		Name:  "Авокадо",
		Desc:  "The avocado, a tree likely originating from south-central Mexico, is classified as a member of the flowering plant family Lauraceae. The fruit of the plant, also called an avocado, is botanically a large berry containing a single large seed.",
		Image: "1.jpeg",
		Price: 199,
	},
	Product{
		ID:    2,
		Name:  "Апельсин",
		Desc:  "Плод апельсинового дерева — это гибрид мандарина и помело. Цитрус имеет округлую форму, двухслойную кожуру и сочную мякоть. Переводится название фрукта как «яблоко из Китая», откуда его и завезли в Европу. Первые письменные упоминания об апельсине имеют возраст более 2200 лет. Лечебные свойства апельсинов были известны людям ещё до нашей эры.",
		Image: "2.png",
		Price: 259,
	},
	Product{
		ID:    3,
		Name:  "Яблоко зеленое",
		Desc:  "Зеленое яблоко является целый кладовой витаминов, среди которых витамины группы В, витамины А, С, Е, К, РР, Р, фолиевая кислота, инозит. ... Кожура зеленого яблока содержит фитонциды. Хорошо реагирует кожа лица на массаж просто долькой зеленого яблока. Кожа сразу «оживает».",
		Image: "3.jpeg",
		Price: 199,
	},
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("./*html")
	router.Static("/assets", "./assets")
	router.GET("/", func(context *gin.Context) {
		context.HTML(200, "index.html", nil)
	})
	router.GET("/product", func(context *gin.Context) {
		context.JSON(200, products)
	})
	_ = router.Run(":8080")
}
