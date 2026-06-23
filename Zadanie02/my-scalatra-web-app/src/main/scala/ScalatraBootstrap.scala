import com.example.app._
import org.scalatra._
import jakarta.servlet.ServletContext

class ScalatraBootstrap extends LifeCycle {
  override def init(context: ServletContext): Unit = {
    context.setInitParameter(CorsSupport.AllowedOriginsKey, "http://localhost:3000,http://localhost:4000")
    context.setInitParameter(CorsSupport.AllowedMethodsKey, "GET,POST,PUT,DELETE,OPTIONS")
    context.setInitParameter(CorsSupport.AllowCredentialsKey, "true")

    context.mount(new ProductController, "/products/*")
    context.mount(new CategoryController, "/categories/*")
    context.mount(new CartController, "/cart/*")
  }
}