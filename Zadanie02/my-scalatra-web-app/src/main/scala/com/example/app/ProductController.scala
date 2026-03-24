package com.example.app

import org.scalatra._
import org.scalatra.json._
import org.json4s._
import scala.collection.mutable.ListBuffer

case class Product(id: Int, name: String, price: Double)

class ProductController extends ScalatraServlet with JacksonJsonSupport {

  protected implicit lazy val jsonFormats: Formats = DefaultFormats

  before() {
    contentType = formats("json")
  }
  
  val products = ListBuffer(
    Product(1,"Banan", 1.5),
    Product(2,"Jablko", 1.2)
  )

  get("/") {
    products
  }

  get("/:id") {
    val id = params("id").toInt
    products.find(_.id == id) match {
      case Some(p) => p
      case None => halt(404,"Nie znaleziono produktu o tym ID")
    }
  }

  put("/:id") {
    val id = params("id").toInt
    val updatedProduct = parsedBody.extract[Product]
    
    val index = products.indexWhere(_.id == id)
    if (index >= 0) {
      products.update(index, updatedProduct)
      updatedProduct
    } else {
      halt(404, "Nie znaleziono produktu o tym ID do modyfikacji")
    }
  }

  post("/") {
    val newProduct = parsedBody.extract[Product]
    products += newProduct
    newProduct
  }

  delete("/:id") {
    val id = params("id").toInt
    products.filterInPlace(_.id != id)
    halt(204)
  }
}
