package com.example.app

import org.scalatra._
import org.scalatra.json._
import org.json4s._
import scala.collection.mutable.ListBuffer

class CartController extends ScalatraServlet with JacksonJsonSupport {

  protected implicit lazy val jsonFormats: Formats = DefaultFormats

  before() {
    contentType = formats("json")
  }
  
  val cartItems = ListBuffer[Product]()

  get("/") {
    cartItems
  }

  post("/") {
    val product = parsedBody.extract[Product]
    cartItems += product
    product
  }

  delete("/:id") {
    val id = params("id").toInt
    cartItems.filterInPlace(_.id != id)
    halt(204)
  }

  delete("/") {
    cartItems.clear()
    halt(204)
  }
}