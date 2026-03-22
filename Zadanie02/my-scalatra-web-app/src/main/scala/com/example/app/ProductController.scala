package com.example.app

import org.scalatra._

case class Product(id: Int, name: String, price: Double)

class ProductController extends ScalatraServlet {
  
  var products = List(
    Product(1,"Banan", 1.5),
    Product(2,"Jablko", 1.2)
  )

  get("/") {
    "Produkty: " + products.map(_.name).mkString(", ")
  }

  get("/:id") {
    val id = params("id").toInt
    products.find(_.id == id) match {
      case Some(p) => s"Produkt: ${p.name}, Cena: ${p.price}"
      case None => NotFound("Nie znaleziono produktu o tym ID")
    }
  }
}
