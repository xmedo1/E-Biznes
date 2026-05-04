import pytest
from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.chrome.service import Service
from webdriver_manager.chrome import ChromeDriverManager
import time

class Test:

    @pytest.fixture(scope="class")
    def driver(self):
        options = webdriver.ChromeOptions()
        options.add_argument("--start-maximized")
        options.add_argument("--incognito")
        options.add_experimental_option("excludeSwitches", ["enable-automation"])
        driver = webdriver.Chrome(service=Service(ChromeDriverManager().install()), options=options)
        yield driver
        driver.quit()

    def test_01_page_load(self, driver):
        driver.get("https://www.saucedemo.com/")
        assert "Swag Labs" in driver.title
        assert driver.current_url == "https://www.saucedemo.com/"
        assert driver.find_element(By.CLASS_NAME, "login_logo").is_displayed()

    def test_02_username_field_visible(self, driver):
        user_input = driver.find_element(By.ID, "user-name")
        assert user_input.is_displayed()
        assert user_input.is_enabled()
        assert user_input.get_attribute("placeholder") == "Username"

    def test_03_password_field_visible(self, driver):
        pass_input = driver.find_element(By.ID, "password")
        assert pass_input.is_displayed()
        assert pass_input.get_attribute("type") == "password"
        assert pass_input.get_attribute("placeholder") == "Password"

    def test_04_invalid_login(self, driver):
        driver.find_element(By.ID, "user-name").send_keys("error_user")
        driver.find_element(By.ID, "password").send_keys("wrong_password")
        driver.find_element(By.ID, "login-button").click()
        error_msg = driver.find_element(By.CSS_SELECTOR, "[data-test='error']")
        assert error_msg.is_displayed()
        assert "Username and password do not match" in error_msg.text
        assert driver.find_element(By.CLASS_NAME, "error-button").is_displayed()

    def test_05_valid_login(self, driver):
        driver.refresh()
        driver.find_element(By.ID, "user-name").send_keys("standard_user")
        driver.find_element(By.ID, "password").send_keys("secret_sauce")
        driver.find_element(By.ID, "login-button").click()
        assert "inventory.html" in driver.current_url
        assert driver.find_element(By.CLASS_NAME, "title").text == "Products"
        assert driver.find_element(By.CLASS_NAME, "shopping_cart_link").is_displayed()

    def test_06_burger_menu_opens(self, driver):
        driver.find_element(By.ID, "react-burger-menu-btn").click()
        time.sleep(0.5)
        assert driver.find_element(By.ID, "logout_sidebar_link").is_displayed()
        assert driver.find_element(By.ID, "about_sidebar_link").is_displayed()
        assert driver.find_element(By.ID, "reset_sidebar_link").is_displayed()

    def test_07_burger_menu_closes(self, driver):
        driver.find_element(By.ID, "react-burger-cross-btn").click()
        time.sleep(0.5)
        assert "inventory.html" in driver.current_url
        assert driver.find_element(By.CLASS_NAME, "inventory_list").is_displayed()

    def test_08_inventory_items_present(self, driver):
        items = driver.find_elements(By.CLASS_NAME, "inventory_item")
        assert len(items) == 6
        assert items[0].find_element(By.CLASS_NAME, "inventory_item_name").is_displayed()
        assert "$" in items[0].find_element(By.CLASS_NAME, "inventory_item_price").text

    def test_09_sort_dropdown_visible(self, driver):
        dropdown = driver.find_element(By.CLASS_NAME, "product_sort_container")
        assert dropdown.is_displayed()
        options = dropdown.find_elements(By.TAG_NAME, "option")
        assert len(options) == 4
        assert "az" in dropdown.get_attribute("value")

    def test_10_add_backpack_to_cart(self, driver):
        driver.find_element(By.ID, "add-to-cart-sauce-labs-backpack").click()
        badge = driver.find_element(By.CLASS_NAME, "shopping_cart_badge")
        assert badge.is_displayed()
        assert badge.text == "1"
        assert driver.find_element(By.ID, "remove-sauce-labs-backpack").is_displayed()

    def test_11_remove_backpack_from_cart(self, driver):
        driver.find_element(By.ID, "remove-sauce-labs-backpack").click()
        assert len(driver.find_elements(By.CLASS_NAME, "shopping_cart_badge")) == 0
        assert driver.find_element(By.ID, "add-to-cart-sauce-labs-backpack").is_displayed()

    def test_12_add_bike_light_to_cart(self, driver):
        driver.find_element(By.ID, "add-to-cart-sauce-labs-bike-light").click()
        assert driver.find_element(By.CLASS_NAME, "shopping_cart_badge").text == "1"
        assert "Remove" in driver.find_element(By.ID, "remove-sauce-labs-bike-light").text

    def test_13_open_cart_page(self, driver):
        driver.find_element(By.CLASS_NAME, "shopping_cart_link").click()
        assert "cart.html" in driver.current_url
        assert driver.find_element(By.CLASS_NAME, "title").text == "Your Cart"
        assert len(driver.find_elements(By.CLASS_NAME, "cart_item")) == 1

    def test_14_continue_shopping(self, driver):
        driver.find_element(By.ID, "continue-shopping").click()
        assert "inventory.html" in driver.current_url
        assert driver.find_element(By.ID, "add-to-cart-sauce-labs-backpack").is_displayed()

    def test_15_go_to_product_details(self, driver):
        driver.find_element(By.ID, "item_4_title_link").click()
        assert "inventory-item.html" in driver.current_url
        assert driver.find_element(By.CLASS_NAME, "inventory_details_name").text == "Sauce Labs Backpack"
        assert driver.find_element(By.ID, "back-to-products").is_displayed()

    def test_16_back_to_products(self, driver):
        driver.find_element(By.ID, "back-to-products").click()
        assert "inventory.html" in driver.current_url
        assert driver.find_element(By.CLASS_NAME, "title").text == "Products"

    def test_17_go_to_checkout(self, driver):
        driver.find_element(By.CLASS_NAME, "shopping_cart_link").click()
        driver.find_element(By.ID, "checkout").click()
        assert "checkout-step-one.html" in driver.current_url
        assert driver.find_element(By.ID, "first-name").is_displayed()
        assert driver.find_element(By.ID, "continue").is_displayed()

    def test_18_cancel_checkout(self, driver):
        driver.find_element(By.ID, "cancel").click()
        assert "cart.html" in driver.current_url
        assert driver.find_element(By.ID, "checkout").is_displayed()

    def test_19_footer_visible(self, driver):
        assert driver.find_element(By.CLASS_NAME, "footer").is_displayed()

    def test_20_logout(self, driver):
        driver.execute_script("window.scrollTo(0, 0);")
        driver.find_element(By.ID, "react-burger-menu-btn").click()
        time.sleep(1)
        logout_link = driver.find_element(By.ID, "logout_sidebar_link")
        driver.execute_script("arguments[0].click();", logout_link)
        assert driver.current_url == "https://www.saucedemo.com/"
        assert driver.find_element(By.ID, "login-button").is_displayed()