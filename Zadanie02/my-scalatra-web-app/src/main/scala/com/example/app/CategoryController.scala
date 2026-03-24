package com.example.app

import org.scalatra._
import org.scalatra.json._
import org.json4s._
import scala.collection.mutable.ListBuffer

case class Category(id: Int, name: String)

class CategoryController extends ScalatraServlet with JacksonJsonSupport {

  protected implicit lazy val jsonFormats: Formats = DefaultFormats

  before() {
    contentType = formats("json")
  }
  
  val categories = ListBuffer(
    Category(1, "Owoce"),
    Category(2, "Warzywa"),
  )

  get("/") {
    categories
  }

  get("/:id") {
    val id = params("id").toInt
    categories.find(_.id == id) match {
      case Some(p) => p
      case None => halt(404,"Nie znaleziono kategorii o tym ID")
    }
  }

  put("/:id") {
    val id = params("id").toInt
    val updatedCategory = parsedBody.extract[Category]

    val index = categories.indexWhere(_.id == id)
    if (index >= 0) {
      categories.update(index, updatedCategory)
      updatedCategory
    } else {
        halt(404, "Nie znaleziono kategorii do aktualizacji")
    }
  }

  post("/") {
    val newCategory = parsedBody.extract[Category]
    categories += newCategory
    newCategory
  }

  delete("/:id") {
    val id = params("id").toInt
    categories.filterInPlace(_.id != id)
    halt(204)
  }
}