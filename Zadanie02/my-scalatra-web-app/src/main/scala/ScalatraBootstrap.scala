import com.example.app._
import org.scalatra._
import jakarta.servlet.ServletContext

class ScalatraBootstrap extends LifeCycle {
  override def init(context: ServletContext): Unit = {
    context.mount(new ProductController, "/products/*")
    context.mount(new CategoryController, "/categories/*")
    context.mount(new CartController, "/cart/*")
  }
}
